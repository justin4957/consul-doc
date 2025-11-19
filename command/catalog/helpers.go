// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/catalog/helpers.go
Helpers module for cli layer

## Tags
cli, discovery, registry, user-interface

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/catalog/helpers.go> a code:Module ;
    code:name "command/catalog/helpers.go" ;
    code:description "Helpers module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:tags "cli", "discovery", "registry", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package catalog

import (
	"fmt"
	"sort"
	"strings"
)

// mapToKV converts a map[string]string into a human-friendly key=value list,
// sorted by name.
func mapToKV(m map[string]string, joiner string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	r := make([]string, len(keys))
	for i, k := range keys {
		r[i] = fmt.Sprintf("%s=%s", k, m[k])
	}
	return strings.Join(r, joiner)
}
