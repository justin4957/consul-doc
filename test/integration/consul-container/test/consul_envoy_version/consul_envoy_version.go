// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/test/consul_envoy_version/consul_envoy_version.go
Consul Envoy Version module for internal layer

## Linked Modules
- [envoyextensions/xdscommon](../envoyextensions/xdscommon)

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/test/consul_envoy_version/consul_envoy_version.go> a code:Module ;
    code:name "test/integration/consul-container/test/consul_envoy_version/consul_envoy_version.go" ;
    code:description "Consul Envoy Version module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "envoyextensions/xdscommon" ;
        code:path "../envoyextensions/xdscommon"
    ] ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/hashicorp/consul/envoyextensions/xdscommon"
)

type consulEnvoyVersions struct {
	ConsulVersion string
	EnvoyVersions []string
}

func main() {
	cev := consulEnvoyVersions{}

	// Get Consul Version
	data, err := os.ReadFile("./version/VERSION")
	if err != nil {
		panic(err)
	}
	cVersion := strings.TrimSpace(string(data))

	cev.EnvoyVersions = append(cev.EnvoyVersions, xdscommon.EnvoyVersions...)

	// ensure the versions are properly sorted latest to oldest
	sort.Sort(sort.Reverse(sort.StringSlice(cev.EnvoyVersions)))

	ceVersions := consulEnvoyVersions{
		ConsulVersion: cVersion,
		EnvoyVersions: cev.EnvoyVersions,
	}
	output, err := json.Marshal(ceVersions)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(output))
}
