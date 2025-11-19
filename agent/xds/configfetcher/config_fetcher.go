// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/xds/configfetcher/config_fetcher.go
Config Fetcher module for agent layer

## Tags
agent, configuration, envoy, service-mesh

## Exports
ConfigFetcher

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/configfetcher/config_fetcher.go> a code:Module ;
    code:name "agent/xds/configfetcher/config_fetcher.go" ;
    code:description "Config Fetcher module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :ConfigFetcher ;
    code:tags "agent", "configuration", "envoy", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package configfetcher

// ConfigFetcher is the interface the agent needs to expose
// for the xDS server to fetch agent config, currently only one field is fetched
type ConfigFetcher interface {
	AdvertiseAddrLAN() string
}
