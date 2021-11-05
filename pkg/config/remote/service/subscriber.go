// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	"github.com/DataDog/datadog-agent/pkg/api/security"
	"github.com/DataDog/datadog-agent/pkg/config"
	"github.com/DataDog/datadog-agent/pkg/config/remote/service/tuf"
	"github.com/DataDog/datadog-agent/pkg/proto/pbgo"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

const errorRetryInterval = 3 * time.Second

// SubscriberCallback defines the function called when a new configuration was fetched
type SubscriberCallback func(config *pbgo.ConfigResponse) error

// Subscriber describes a product's configuration subscriber
type Subscriber struct {
	product     pbgo.Product
	refreshRate time.Duration
	lastUpdate  time.Time
	lastVersion uint64
	callback    SubscriberCallback
}

type coreAgentStream struct {
	stream      pbgo.AgentSecure_GetConfigUpdatesClient
	streamCtx   context.Context
	cancel      context.CancelFunc
	agentClient pbgo.AgentSecureClient
	mx          sync.Mutex
}

func newCoreAgentStream(streamCtx context.Context) (*coreAgentStream, error) {
	cas := &coreAgentStream{
		streamCtx: streamCtx,
	}
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	ctx, cancel := context.WithCancel(context.Background())
	cas.cancel = cancel

	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf(":%v", config.Datadog.GetInt("cmd_port")),
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		cancel()
		return nil, err
	}

	cas.agentClient = pbgo.NewAgentSecureClient(conn)

	cas.connect()

	return cas, nil
}

func (cas *coreAgentStream) connect() {
	cas.mx.Lock()
	defer cas.mx.Unlock()
	stream := cas.getStream()
	cas.stream = stream
}

func (cas *coreAgentStream) getStream() pbgo.AgentSecure_GetConfigUpdatesClient {
	var stream pbgo.AgentSecure_GetConfigUpdatesClient
	var err error
	for {
		stream, err = cas.agentClient.GetConfigUpdates(cas.streamCtx)
		if err != nil {
			log.Errorf("Failed to establish channel to core-agent, retrying in %s...", errorRetryInterval)
			time.Sleep(errorRetryInterval)
			continue
		} else {
			log.Debugf("Successfully established channel to core-agent")
			break
		}
	}
	return stream
}

func (cas *coreAgentStream) sendTracerInfos(tracerInfos chan *pbgo.TracerInfo, currentConfigSnapshotVersion uint64, product pbgo.Product, streamCtx context.Context) {
	for {
		select {
		case tracerInfo := <-tracerInfos:
			request := pbgo.SubscribeConfigRequest{
				CurrentConfigSnapshotVersion: currentConfigSnapshotVersion,
				Product:                      product,
				TracerInfo:                   tracerInfo,
			}
			log.Trace("Sending subscribe config requests with tracer infos to core-agent")
			if err := cas.stream.Send(&request); err != nil {
				log.Warnf("Error writing tracer infos to stream: %s", err)
				time.Sleep(errorRetryInterval)
				cas.connect()
				continue
			}
		case <-streamCtx.Done():
			return
		}
	}
}

// NewSubscriber returns a new subscriber with the specified refresh rate and a callback
func NewSubscriber(product pbgo.Product, refreshRate time.Duration, callback SubscriberCallback) *Subscriber {
	return &Subscriber{
		product:     product,
		refreshRate: refreshRate,
		callback:    callback,
	}
}

// NewGRPCSubscriber returns a new gRPC stream based subscriber.
func NewGRPCSubscriber(product pbgo.Product, callback SubscriberCallback) (context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())

	token, err := security.FetchAuthToken()
	if err != nil {
		cancel()
		err = fmt.Errorf("unable to fetch authentication token: %w", err)
		log.Infof("unable to establish stream, will possibly retry: %s", err)
		return nil, err
	}

	currentConfigSnapshotVersion := uint64(0)
	client := tuf.NewDirectorPartialClient()

	go func() {
		log.Debug("Waiting for configuration from remote config management")

		streamCtx, streamCancel := context.WithCancel(
			metadata.NewOutgoingContext(ctx, metadata.MD{
				"authorization": []string{fmt.Sprintf("Bearer %s", token)},
			}),
		)
		defer streamCancel()

		cas, err := newCoreAgentStream(streamCtx)
		if err != nil {
			cancel()
			return
		}

		log.Debug("Processing response from remote config management")

		for {
			request := pbgo.SubscribeConfigRequest{
				CurrentConfigSnapshotVersion: currentConfigSnapshotVersion,
				Product:                      product,
			}
			stream := cas.getStream()
			err := stream.SendMsg(request)
			if err != nil {
				log.Warnf("Error sending message to core agent: %s", err)
				continue
			}

			for {
				// Get new event from stream
				configResponse, err := stream.Recv()
				if err == io.EOF {
					continue
				} else if err != nil {
					log.Warnf("Stopped listening for configuration from remote config management: %s", err)
					time.Sleep(errorRetryInterval)
					break
				}

				if err := client.Verify(configResponse); err != nil {
					log.Errorf("Partial verify failed: %s", err)
					continue
				}

				log.Infof("Got config for product %s", product)
				if err := callback(configResponse); err == nil {
					currentConfigSnapshotVersion = configResponse.ConfigSnapshotVersion
				}
			}
		}
	}()

	return cancel, nil
}

// NewTracerGRPCSubscriber returns a new gRPC stream based subscriber.
func NewTracerGRPCSubscriber(product pbgo.Product, callback SubscriberCallback, tracerInfos chan *pbgo.TracerInfo) (context.CancelFunc, error) {
	ctx, cancel := context.WithCancel(context.Background())

	token, err := security.FetchAuthToken()
	if err != nil {
		cancel()
		err = fmt.Errorf("unable to fetch authentication token: %w", err)
		log.Infof("unable to establish stream, will possibly retry: %s", err)
		return nil, err
	}

	currentConfigSnapshotVersion := uint64(0)
	client := tuf.NewDirectorPartialClient()

	go func() {
		streamCtx, streamCancel := context.WithCancel(
			metadata.NewOutgoingContext(ctx, metadata.MD{
				"authorization": []string{fmt.Sprintf("Bearer %s", token)},
			}),
		)
		defer streamCancel()

		cas, err := newCoreAgentStream(streamCtx)
		if err != nil {
			log.Errorf("Error creating core agent stream: %s", err)
			return
		}

		go cas.sendTracerInfos(tracerInfos, currentConfigSnapshotVersion, product, streamCtx)
		readConfigs(cas, client, product, callback, currentConfigSnapshotVersion)
	}()

	return cancel, nil
}

func readConfigs(cas *coreAgentStream, client *tuf.DirectorPartialClient, product pbgo.Product, callback SubscriberCallback, currentConfigSnapshotVersion uint64) {
	for {
		log.Debug("Waiting for new config")
		configResponse, err := cas.stream.Recv()
		if err == io.EOF {
			continue
		} else if err != nil {
			log.Warnf("Stopped listening for configuration from remote config management: %s", err)
			time.Sleep(errorRetryInterval)
			cas.connect()
			continue
		}

		if err := client.Verify(configResponse); err != nil {
			log.Errorf("Partial verify failed: %s", err)
			continue
		}

		log.Infof("Got config for product %s", product)
		if err := callback(configResponse); err == nil {
			currentConfigSnapshotVersion = configResponse.ConfigSnapshotVersion
		}
	}
}
