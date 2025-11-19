// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/hcp/discover/discover.go
Discover module for agent layer

## Linked Modules
- [agent/hcp/client](../agent/hcp/client)
- [agent/hcp/config](../agent/hcp/config)

## Tags
agent

## Exports
Provider

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/hcp/discover/discover.go> a code:Module ;
    code:name "agent/hcp/discover/discover.go" ;
    code:description "Discover module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/hcp/client" ;
        code:path "../agent/hcp/client"
    ], [
        code:name "agent/hcp/config" ;
        code:path "../agent/hcp/config"
    ] ;
    code:exports :Provider ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package discover

import (
	"context"
	"fmt"
	"log"
	"time"

	hcpclient "github.com/hashicorp/consul/agent/hcp/client"
	"github.com/hashicorp/consul/agent/hcp/config"
)

type Provider struct {
}

var (
	defaultTimeout = 5 * time.Second
)

type providerConfig struct {
	config.CloudConfig

	timeout time.Duration
}

func (p *Provider) Addrs(args map[string]string, l *log.Logger) ([]string, error) {
	cfg, err := parseArgs(args)
	if err != nil {
		return nil, err
	}

	client, err := hcpclient.NewClient(cfg.CloudConfig)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.timeout)
	defer cancel()
	servers, err := client.DiscoverServers(ctx)
	if err != nil {
		return nil, err
	}

	return servers, nil
}

func (p *Provider) Help() string {
	return ""
}

func parseArgs(args map[string]string) (cfg providerConfig, err error) {
	cfg.timeout = defaultTimeout

	if id, ok := args["resource_id"]; ok {
		cfg.ResourceID = id
	} else {
		err = fmt.Errorf("'resource_id' was not found and is required")
	}

	if cid, ok := args["client_id"]; ok {
		cfg.ClientID = cid
	}

	if csec, ok := args["client_secret"]; ok {
		cfg.ClientSecret = csec
	}

	if timeoutRaw, ok := args["timeout"]; ok {
		timeout, err := time.ParseDuration(timeoutRaw)
		if err != nil {
			return cfg, err
		}
		cfg.timeout = timeout
	}
	return
}
