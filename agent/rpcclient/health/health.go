// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/rpcclient/health/health.go
Health module for agent layer

## Linked Modules
- [agent/cache](../agent/cache)
- [agent/rpcclient](../agent/rpcclient)
- [agent/structs](../agent/structs)
- [agent/submatview](../agent/submatview)
- [proto/private/pbsubscribe](../proto/private/pbsubscribe)

## Tags
agent, communication, health-checks, monitoring, networking

## Exports
Client

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/rpcclient/health/health.go> a code:Module ;
    code:name "agent/rpcclient/health/health.go" ;
    code:description "Health module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/cache" ;
        code:path "../agent/cache"
    ], [
        code:name "agent/rpcclient" ;
        code:path "../agent/rpcclient"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "agent/submatview" ;
        code:path "../agent/submatview"
    ], [
        code:name "proto/private/pbsubscribe" ;
        code:path "../proto/private/pbsubscribe"
    ] ;
    code:exports :Client ;
    code:tags "agent", "communication", "health-checks", "monitoring", "networking" .
<!-- End LinkedDoc RDF -->
*/
package health

import (
	"context"

	"github.com/hashicorp/consul/agent/rpcclient"
	"google.golang.org/grpc/connectivity"

	"github.com/hashicorp/consul/agent/cache"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/agent/submatview"
	"github.com/hashicorp/consul/proto/private/pbsubscribe"
)

// Client provides access to service health data.
type Client struct {
	rpcclient.Client
}

// IsReadyForStreaming will indicate if the underlying gRPC connection is ready.
func (c *Client) IsReadyForStreaming() bool {
	conn := c.MaterializerDeps.Conn
	if conn == nil {
		return false
	}

	return conn.GetState() == connectivity.Ready
}

func (c *Client) ServiceNodes(
	ctx context.Context,
	req structs.ServiceSpecificRequest,
) (structs.IndexedCheckServiceNodes, cache.ResultMeta, error) {
	// Note: if MergeCentralConfig is requested, default to using the RPC backend for now
	// as the streaming backend and materializer does not have support for merging yet.
	if c.useStreaming(req) && (req.UseCache || req.MinQueryIndex > 0) && !req.MergeCentralConfig {
		c.QueryOptionDefaults(&req.QueryOptions)

		result, err := c.ViewStore.Get(ctx, c.newServiceRequest(req))
		if err != nil {
			return structs.IndexedCheckServiceNodes{}, cache.ResultMeta{}, err
		}
		meta := cache.ResultMeta{Index: result.Index, Hit: result.Cached}
		return *result.Value.(*structs.IndexedCheckServiceNodes), meta, err
	}

	out, md, err := c.getServiceNodes(ctx, req)
	if err != nil {
		return out, md, err
	}

	// TODO: DNSServer emitted a metric here, do we still need it?
	if req.AllowStale && req.MaxStaleDuration > 0 && out.LastContact > req.MaxStaleDuration {
		req.AllowStale = false
		err := c.NetRPC.RPC(context.Background(), "Health.ServiceNodes", &req, &out)
		return out, cache.ResultMeta{}, err
	}

	return out, md, err
}

func (c *Client) getServiceNodes(
	ctx context.Context,
	req structs.ServiceSpecificRequest,
) (structs.IndexedCheckServiceNodes, cache.ResultMeta, error) {
	var out structs.IndexedCheckServiceNodes
	if !req.UseCache {
		err := c.NetRPC.RPC(context.Background(), "Health.ServiceNodes", &req, &out)
		return out, cache.ResultMeta{}, err
	}

	raw, md, err := c.Cache.Get(ctx, c.CacheName, &req)
	if err != nil {
		return out, md, err
	}

	value, ok := raw.(*structs.IndexedCheckServiceNodes)
	if !ok {
		panic("wrong response type for cachetype.HealthServicesName")
	}

	return *value, md, nil
}

func (c *Client) Notify(
	ctx context.Context,
	req structs.ServiceSpecificRequest,
	correlationID string,
	cb cache.Callback,
) error {
	if c.useStreaming(req) {
		sr := c.newServiceRequest(req)
		return c.ViewStore.NotifyCallback(ctx, sr, correlationID, cb)
	}

	return c.Cache.NotifyCallback(ctx, c.CacheName, &req, correlationID, cb)
}

func (c *Client) useStreaming(req structs.ServiceSpecificRequest) bool {
	return c.UseStreamingBackend &&
		!req.Ingress &&
		// Streaming is incompatible with NearestN queries (due to lack of ordering),
		// so we can only use it if the NearestN would never work (Node == "")
		// or if we explicitly say to ignore the Node field for queries (agentless xDS).
		(req.Source.Node == "" || req.Source.DisableNode) &&
		// Streaming is incompatible with SamenessGroup queries at the moment because
		// the subscribe functionality maps to queries based on the service name and tenancy information
		// it does not support the ability to subscribe to the same service in different partitions or peers
		// and materialize the results into a single view with the first healthy sameness group member.
		req.SamenessGroup == ""
}

func (c *Client) newServiceRequest(req structs.ServiceSpecificRequest) serviceRequest {
	return serviceRequest{
		ServiceSpecificRequest: req,
		deps:                   c.MaterializerDeps,
	}
}

var _ submatview.Request = (*serviceRequest)(nil)

type serviceRequest struct {
	structs.ServiceSpecificRequest
	deps rpcclient.MaterializerDeps
}

func (r serviceRequest) CacheInfo() cache.RequestInfo {
	return r.ServiceSpecificRequest.CacheInfo()
}

func (r serviceRequest) Type() string {
	return "agent.rpcclient.health.serviceRequest"
}

func (r serviceRequest) NewMaterializer() (submatview.Materializer, error) {
	view, err := NewHealthView(r.ServiceSpecificRequest)
	if err != nil {
		return nil, err
	}
	deps := submatview.Deps{
		View:    view,
		Logger:  r.deps.Logger,
		Request: NewMaterializerRequest(r.ServiceSpecificRequest),
	}

	return submatview.NewRPCMaterializer(pbsubscribe.NewStateChangeSubscriptionClient(r.deps.Conn), deps), nil
}
