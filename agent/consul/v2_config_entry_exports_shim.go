// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/v2_config_entry_exports_shim.go
V2 Config Entry Exports Shim module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/consul/controller/queue](../agent/consul/controller/queue)
- [agent/consul/state](../agent/consul/state)
- [agent/consul/stream](../agent/consul/stream)
- [agent/structs](../agent/structs)
- [internal/controller](../internal/controller)
- [logging](../logging)
- [proto/private/pbconfigentry](../proto/private/pbconfigentry)

## Tags
agent, configuration

## Exports
NewExportedServicesShim

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/v2_config_entry_exports_shim.go> a code:Module ;
    code:name "agent/consul/v2_config_entry_exports_shim.go" ;
    code:description "V2 Config Entry Exports Shim module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/consul/controller/queue" ;
        code:path "../agent/consul/controller/queue"
    ], [
        code:name "agent/consul/state" ;
        code:path "../agent/consul/state"
    ], [
        code:name "agent/consul/stream" ;
        code:path "../agent/consul/stream"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "internal/controller" ;
        code:path "../internal/controller"
    ], [
        code:name "logging" ;
        code:path "../logging"
    ], [
        code:name "proto/private/pbconfigentry" ;
        code:path "../proto/private/pbconfigentry"
    ] ;
    code:exports :NewExportedServicesShim ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/consul/controller/queue"
	"github.com/hashicorp/consul/agent/consul/state"
	"github.com/hashicorp/consul/agent/consul/stream"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/logging"
	"github.com/hashicorp/consul/proto/private/pbconfigentry"
	"github.com/hashicorp/go-hclog"
)

type v1ServiceExportsShim struct {
	s       *Server
	eventCh chan controller.Event
}

func NewExportedServicesShim(s *Server) *v1ServiceExportsShim {
	eventCh := make(chan controller.Event)
	return &v1ServiceExportsShim{
		s:       s,
		eventCh: eventCh,
	}
}

func (s *v1ServiceExportsShim) Start(ctx context.Context) {
	logger := s.s.logger.Named(logging.V2ExportsShim)

	// TODO replace this with a proper supervisor.
	for ctx.Err() == nil {
		err := subscribeToExportedServicesEvents(ctx, logger, s.s.publisher, s.eventCh)

		if err != nil {
			logger.Warn("encountered an error while streaming exported services", "error", err)
			select {
			case <-time.After(time.Second):
			case <-ctx.Done():
				return
			}
		} else {
			return
		}
	}
}

func subscribeToExportedServicesEvents(ctx context.Context, logger hclog.Logger, publisher *stream.EventPublisher, eventCh chan controller.Event) error {
	subscription, err := publisher.Subscribe(&stream.SubscribeRequest{
		Topic:   state.EventTopicExportedServices,
		Subject: stream.SubjectWildcard,
	})
	if err != nil {
		return err
	}
	defer subscription.Unsubscribe()
	var index uint64

	for {
		event, err := subscription.Next(ctx)
		switch {
		case errors.Is(err, context.Canceled):
			return nil
		case err != nil:
			return err
		}

		if event.IsFramingEvent() {
			continue
		}

		if event.Index <= index {
			continue
		}

		index = event.Index
		e := event.Payload.ToSubscriptionEvent(event.Index)
		configEntry := e.GetConfigEntry().GetConfigEntry()

		if configEntry.GetKind() != pbconfigentry.Kind_KindExportedServices {
			logger.Error("unexpected config entry kind", "kind", configEntry.GetKind())
			continue
		}
		partition := acl.PartitionOrDefault(configEntry.GetEnterpriseMeta().GetPartition())

		eventCh <- controller.Event{
			Obj: exportedServiceItemType{partition: partition},
		}
	}
}

func (s *v1ServiceExportsShim) EventChannel() chan controller.Event {
	return s.eventCh
}

func (s *v1ServiceExportsShim) GetExportedServicesConfigEntry(_ context.Context, name string, entMeta *acl.EnterpriseMeta) (*structs.ExportedServicesConfigEntry, error) {
	_, entry, err := s.s.fsm.State().ConfigEntry(nil, structs.ExportedServices, name, entMeta)
	if err != nil {
		return nil, err
	}

	if entry == nil {
		return nil, nil
	}

	exp, ok := entry.(*structs.ExportedServicesConfigEntry)
	if !ok {
		return nil, fmt.Errorf("exported services config entry is the wrong type: expected ExportedServicesConfigEntry, actual: %T", entry)
	}

	return exp, nil
}

func (s *v1ServiceExportsShim) WriteExportedServicesConfigEntry(_ context.Context, cfg *structs.ExportedServicesConfigEntry) error {
	if err := cfg.Normalize(); err != nil {
		return err
	}

	if err := cfg.Validate(); err != nil {
		return err
	}

	req := &structs.ConfigEntryRequest{
		Op:    structs.ConfigEntryUpsert,
		Entry: cfg,
	}

	_, err := s.s.raftApply(structs.ConfigEntryRequestType, req)
	return err
}

func (s *v1ServiceExportsShim) DeleteExportedServicesConfigEntry(_ context.Context, name string, entMeta *acl.EnterpriseMeta) error {
	if entMeta == nil {
		entMeta = acl.DefaultEnterpriseMeta()
	}

	req := &structs.ConfigEntryRequest{
		Op: structs.ConfigEntryDelete,
		Entry: &structs.ExportedServicesConfigEntry{
			Name:           name,
			EnterpriseMeta: *entMeta,
		},
	}

	if err := req.Entry.Normalize(); err != nil {
		return err
	}

	_, err := s.s.raftApply(structs.ConfigEntryRequestType, req)
	return err
}

type exportedServiceItemType struct {
	partition string
}

var _ queue.ItemType = (*exportedServiceItemType)(nil)

func (e exportedServiceItemType) Key() string {
	return e.partition
}
