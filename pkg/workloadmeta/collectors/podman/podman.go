// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package podman

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/containers/podman/v3/libpod"
	"github.com/containers/podman/v3/libpod/define"

	"github.com/DataDog/datadog-agent/pkg/workloadmeta"
	"github.com/DataDog/datadog-agent/pkg/workloadmeta/collectors/util"
)

const (
	collectorID   = "podman"
	componentName = "workloadmeta-podman"
	expireFreq    = 10 * time.Second
)

type collector struct {
	podmanRuntime *libpod.Runtime
	store         workloadmeta.Store
	expire        *util.Expire
}

func init() {
	workloadmeta.RegisterCollector(collectorID, func() workloadmeta.Collector {
		return &collector{}
	})
}

func (c *collector) Start(ctx context.Context, store workloadmeta.Store) error {
	// TODO: check if feature is enabled

	var err error
	c.podmanRuntime, err = libpod.NewRuntime(ctx)
	if err != nil {
		return err
	}

	c.store = store

	c.expire = util.NewExpire(expireFreq)

	return nil
}

func (c *collector) Pull(ctx context.Context) error {
	containers, err := c.podmanRuntime.GetAllContainers()
	if err != nil {
		return err
	}

	var events []workloadmeta.CollectorEvent

	for _, container := range containers {
		event, err := convertToEvent(container)
		if err != nil {
			// TODO: log
			continue
		}

		events = append(events, event)
	}

	expires := c.expire.ComputeExpires()
	for _, expired := range expires {
		events = append(events, workloadmeta.CollectorEvent{
			Type:   workloadmeta.EventTypeUnset,
			Source: workloadmeta.SourcePodman,
			Entity: expired,
		})
	}

	c.store.Notify(events)

	return nil
}

func convertToEvent(container *libpod.Container) (workloadmeta.CollectorEvent, error) {
	spec := container.Spec()
	annotations := make(map[string]string)
	if spec != nil {
		annotations = spec.Annotations
	}

	envs, err := envVars(container)
	if err != nil {
		return workloadmeta.CollectorEvent{}, err
	}

	image, err := workloadmeta.NewContainerImage(container.RawImageName())
	if err != nil {
		return workloadmeta.CollectorEvent{}, err
	}

	pid, err := container.PID() // Returns 0 if the container is not running
	if err != nil {
		return workloadmeta.CollectorEvent{}, err
	}

	state, err := container.State()
	if err != nil {
		return workloadmeta.CollectorEvent{}, err
	}
	isRunning := state == define.ContainerStateRunning

	startedAt, err := container.StartedTime()
	if err != nil {
		return workloadmeta.CollectorEvent{}, err
	}

	finishedAt, err := container.FinishedTime()
	if err != nil {
		return workloadmeta.CollectorEvent{}, err
	}

	return workloadmeta.CollectorEvent{
		Type:   workloadmeta.EventTypeSet,
		Source: workloadmeta.SourcePodman,
		Entity: &workloadmeta.Container{
			EntityID: workloadmeta.EntityID{
				Kind: workloadmeta.KindContainer,
				ID:   container.ID(),
			},
			EntityMeta: workloadmeta.EntityMeta{
				Name:        container.Name(),
				Namespace:   container.Namespace(),
				Annotations: annotations,
				Labels:      container.Labels(),
			},
			EnvVars:    envs,
			Hostname:   container.Hostname(),
			Image:      image, // TODO: check if image_name tag also includes repo in containerd
			NetworkIPs: nil, // TODO
			PID:        pid,
			Ports:      nil, // TODO
			Runtime:    workloadmeta.ContainerRuntimePodman,
			State: workloadmeta.ContainerState{
				Running:    isRunning,
				StartedAt:  startedAt,
				FinishedAt: finishedAt,
			},
		},
	}, nil
}

func envVars(container *libpod.Container) (map[string]string, error) {
	res := make(map[string]string)

	if container.Spec() == nil || container.Spec().Process == nil {
		return res, nil
	}

	for _, env := range container.Spec().Process.Env {
		envSplit := strings.SplitN(env, "=", 2)

		if len(envSplit) < 2 {
			return nil, errors.New("unexpected environment variable format")
		}

		res[envSplit[0]] = envSplit[1]
	}

	return res, nil
}
