// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/go-sso/oidcauth/internal/strutil/util.go
Util module for internal layer

## Tags
authentication, internal, security

## Exports
StrListContains

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/go-sso/oidcauth/internal/strutil/util.go> a code:Module ;
    code:name "internal/go-sso/oidcauth/internal/strutil/util.go" ;
    code:description "Util module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :StrListContains ;
    code:tags "authentication", "internal", "security" .
<!-- End LinkedDoc RDF -->
*/
package strutil

// StrListContains looks for a string in a list of strings.
func StrListContains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}
