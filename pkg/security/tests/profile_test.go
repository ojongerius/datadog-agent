// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// +build functionaltests

package tests

import (
	"fmt"
	"syscall"
	"testing"
	"time"

	sprobe "github.com/DataDog/datadog-agent/pkg/security/probe"
	"github.com/DataDog/datadog-agent/pkg/security/rules"
	"github.com/hashicorp/go-multierror"
)

func TestProfile(t *testing.T) {
	rule := &rules.RuleDefinition{
		ID:         "test_rule",
		Expression: `link.file.path == "{{.Root}}/test-link" && link.file.destination.path == "{{.Root}}/test2-link" && link.file.uid == 98 && link.file.gid == 99 && link.file.destination.uid == 98 && link.file.destination.gid == 99`,
	}

	test, err := newTestModule(t, nil, []*rules.RuleDefinition{rule}, testOpts{
		profiles: map[string]struct {
			selector string
			rules    []*rules.RuleDefinition
		}{
			"test-profile": {
				selector: "process.file.name == \"testsuite\"",
				rules: []*rules.RuleDefinition{
					{
						ID:         "exec_ls",
						Expression: "exec.file.path == \"/usr/bin/ls\"",
					},
					{
						ID:         "read_proc",
						Expression: "open.file.path =~ \"/proc/*\"",
						// not allowed but better
						// Expression: "open.file.path =~ r\"/proc/[0-9]*(/(stat|statm|status|mountinfo|cmdline))?\"",
					},
					{
						ID:         "write_dev_null",
						Expression: "open.file.path =~ \"/dev/null\"",
					},
					{
						ID:         "exec_getenforce",
						Expression: "exec.file.path =~ \"/usr/sbin/getenforce\"",
					},
					{
						ID:         "open_test_dir",
						Expression: "open.file.path =~ \"/tmp/test-secagent-*\"",
					},
					{
						ID:         "open_etc_passwd",
						Expression: "open.file.path == \"/etc/passwd\"",
					},
					{
						ID:         "open_etc_group",
						Expression: "open.file.path == \"/etc/group\"",
					},
					{
						ID:         "exec_docker",
						Expression: "exec.file.path == \"/usr/bin/docker\"",
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer test.Close()

	t.Run("exec-cat", func(t *testing.T) {
		var errProfile *multierror.Error
		err := test.GetSignal(t, func() error {
			_, err := syscall.ForkExec("/usr/bin/cat", []string{}, nil)
			return err
		}, nil, func(event *sprobe.Event, profile *rules.Profile) {
			path, _ := event.GetFieldValue("exec.file.path")
			if path != "/usr/bin/cat" {
				errProfile = multierror.Append(errProfile, fmt.Errorf("unexpected %s event", event.GetEventType()))
			}
		})
		if err != nil {
			t.Error(err)
		}
		if errProfile != nil {
			t.Error(errProfile)
		}
	})

	test.timeout = 2 * time.Second

	t.Run("exec-ls", func(t *testing.T) {
		var event string
		err := test.GetSignal(t, func() error {
			_, err := syscall.ForkExec("/usr/bin/ls", []string{}, nil)
			return err
		}, nil, func(evt *sprobe.Event, profile *rules.Profile) {
			event = evt.String()
		})
		if err != testTimeout {
			t.Errorf("should get a timeout error with no event, got %+v", err)
			t.Logf("Got event %s", event)
		}
	})

	t.Run("no-event", func(t *testing.T) {
		var event string
		err := test.GetSignal(t, func() error {
			return nil
		}, nil, func(evt *sprobe.Event, profile *rules.Profile) {
			event = evt.String()
		})
		if err != testTimeout {
			t.Errorf("should get a timeout error with no event, got %+v", err)
			t.Logf("Got event %s", event)
		}
	})
}
