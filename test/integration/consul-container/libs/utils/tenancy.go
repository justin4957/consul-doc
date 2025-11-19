// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/libs/utils/tenancy.go
Tenancy module for internal layer

## Linked Modules
- [api](../api)

## Tags
internal

## Exports
CompatQueryOpts, DefaultToEmpty, NamespaceOrDefault, PartitionOrDefault

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/libs/utils/tenancy.go> a code:Module ;
    code:name "test/integration/consul-container/libs/utils/tenancy.go" ;
    code:description "Tenancy module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:exports :CompatQueryOpts, :DefaultToEmpty, :NamespaceOrDefault, :PartitionOrDefault ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package utils

import "github.com/hashicorp/consul/api"

func PartitionOrDefault(name string) string {
	if name == "" {
		return "default"
	}
	return name
}
func NamespaceOrDefault(name string) string {
	if name == "" {
		return "default"
	}
	return name
}

func DefaultToEmpty(name string) string {
	if name == "default" {
		return ""
	}
	return name
}

// CompatQueryOpts cleans a QueryOptions so that Partition and Namespace fields
// are compatible with CE or ENT
// TODO: not sure why we can't do this server-side
func CompatQueryOpts(opts *api.QueryOptions) *api.QueryOptions {
	opts.Partition = DefaultToEmpty(opts.Partition)
	opts.Namespace = DefaultToEmpty(opts.Namespace)
	return opts
}
