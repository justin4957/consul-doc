// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/kv/impexp/kvimpexp.go
Kvimpexp module for cli layer

## Linked Modules
- [api](../api)

## Tags
cli, user-interface

## Exports
Entry, ToEntry

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/kv/impexp/kvimpexp.go> a code:Module ;
    code:name "command/kv/impexp/kvimpexp.go" ;
    code:description "Kvimpexp module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:exports :Entry, :ToEntry ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package impexp

import (
	"encoding/base64"

	"github.com/hashicorp/consul/api"
)

type Entry struct {
	Key       string `json:"key"`
	Flags     uint64 `json:"flags"`
	Value     string `json:"value"`
	Namespace string `json:"namespace,omitempty"`
	Partition string `json:"partition,omitempty"`
}

func ToEntry(pair *api.KVPair) *Entry {
	return &Entry{
		Key:       pair.Key,
		Flags:     pair.Flags,
		Value:     base64.StdEncoding.EncodeToString(pair.Value),
		Namespace: pair.Namespace,
		Partition: pair.Partition,
	}
}
