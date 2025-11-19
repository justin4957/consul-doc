// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/connect/ca/provider_consul_config.go
Provider Consul Config module for agent layer

## Linked Modules
- [agent/connect](../agent/connect)
- [agent/structs](../agent/structs)

## Tags
agent, configuration, mtls, service-mesh

## Exports
ParseConsulCAConfig

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/ca/provider_consul_config.go> a code:Module ;
    code:name "agent/connect/ca/provider_consul_config.go" ;
    code:description "Provider Consul Config module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/connect" ;
        code:path "../agent/connect"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :ParseConsulCAConfig ;
    code:tags "agent", "configuration", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package ca

import (
	"fmt"
	"time"

	"github.com/go-viper/mapstructure/v2"
	"github.com/hashicorp/consul/agent/connect"
	"github.com/hashicorp/consul/agent/structs"
)

func ParseConsulCAConfig(raw map[string]interface{}) (*structs.ConsulCAProviderConfig, error) {
	config := defaultConsulCAProviderConfig()
	decodeConf := &mapstructure.DecoderConfig{
		DecodeHook:       structs.ParseDurationFunc(),
		Result:           &config,
		WeaklyTypedInput: true,
	}

	decoder, err := mapstructure.NewDecoder(decodeConf)
	if err != nil {
		return nil, err
	}

	if err := decoder.Decode(raw); err != nil {
		return nil, fmt.Errorf("error decoding config: %s", err)
	}

	if config.PrivateKey == "" && config.RootCert != "" {
		return nil, fmt.Errorf("must provide a private key when providing a root cert")
	}

	if err := config.CommonCAProviderConfig.Validate(); err != nil {
		return nil, err
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func defaultConsulCAProviderConfig() structs.ConsulCAProviderConfig {
	return structs.ConsulCAProviderConfig{
		CommonCAProviderConfig: defaultCommonConfig(),
	}
}
func defaultCommonConfig() structs.CommonCAProviderConfig {
	return structs.CommonCAProviderConfig{
		LeafCertTTL:         3 * 24 * time.Hour,
		IntermediateCertTTL: 24 * 365 * time.Hour,
		PrivateKeyType:      connect.DefaultPrivateKeyType,
		PrivateKeyBits:      connect.DefaultPrivateKeyBits,
		RootCertTTL:         10 * 24 * 365 * time.Hour,
	}
}
