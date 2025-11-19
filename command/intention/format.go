// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/intention/format.go
Format module for cli layer

## Linked Modules
- [api](../api)

## Tags
cli, user-interface

## Exports
FormatDestination, FormatSource

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/intention/format.go> a code:Module ;
    code:name "command/intention/format.go" ;
    code:description "Format module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:exports :FormatDestination, :FormatSource ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package intention

import (
	"github.com/hashicorp/consul/api"
)

// FormatSource returns the namespace/name format for the source. This is
// different from (*api.Intention).SourceString in that the default namespace
// is not omitted.
func FormatSource(i *api.Intention) string {
	return partString(i.SourceNS, i.SourceName)
}

// FormatDestination returns the namespace/name format for the destination.
// This is different from (*api.Intention).DestinationString in that the
// default namespace is not omitted.
func FormatDestination(i *api.Intention) string {
	return partString(i.DestinationNS, i.DestinationName)
}

func partString(ns, n string) string {
	if ns == "" {
		return n
	}
	return ns + "/" + n
}
