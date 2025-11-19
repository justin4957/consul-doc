// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: acl/acl.go
Core ACL configuration and types for access control in Consul.

## Linked Modules
- [authorizer](./authorizer.go) - Authorization interface and enforcement
- [policy](./policy.go) - Policy definitions and parsing
- [resolver](./resolver/result.go) - ACL token resolution

## Tags
security, access-control, configuration

## Exports
Config, ExportFetcher, ExportedServices, AnonymousTokenID, WildcardName, AliasIfAnonymousToken

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<acl/acl.go> a code:Module ;
    code:language "go" ;
    code:name "acl/acl.go" ;
    code:description "Core ACL configuration and types for access control in Consul" ;
    code:layer "acl" ;
    code:linksTo [
        code:name "authorizer" ;
        code:path "./authorizer.go" ;
        code:relationship "Authorization interface and enforcement"
    ], [
        code:name "policy" ;
        code:path "./policy.go" ;
        code:relationship "Policy definitions and parsing"
    ], [
        code:name "resolver" ;
        code:path "./resolver/result.go" ;
        code:relationship "ACL token resolution"
    ] ;
    code:exports :Config, :ExportFetcher, :ExportedServices, :AnonymousTokenID, :WildcardName, :AliasIfAnonymousToken ;
    code:tags "security", "access-control", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package acl

const (
	WildcardName = "*"

	// AnonymousTokenID is the AccessorID of the anonymous token.
	// When logging or displaying to users, use acl.AliasIfAnonymousToken
	// to convert this to AnonymousTokenAlias.
	AnonymousTokenID     = "00000000-0000-0000-0000-000000000002"
	AnonymousTokenAlias  = "anonymous token"
	AnonymousTokenSecret = "anonymous"

	ReservedBuiltinPrefix = "builtin/"
)

// Config encapsulates all of the generic configuration parameters used for
// policy parsing and enforcement
type Config struct {
	// WildcardName is the string that represents a request to authorize a wildcard permission
	WildcardName string

	//by default errors, but in certain instances we want to make sure to maintain backwards compatabilty
	WarnOnDuplicateKey bool

	// embedded enterprise configuration
	EnterpriseConfig
}

type ExportFetcher interface {
	// ExportsForPartition returns the config entry defining exports for a partition
	ExportsForPartition(partition string) ExportedServices
}

type ExportedServices struct {
	// Data is a map of [namespace] -> [service] -> [list of partitions the service is exported to]
	// This includes both the names of typical service instances and their corresponding sidecar proxy
	// instance names. Meaning that if "web" is exported, "web-sidecar-proxy" instances will also be
	// shown as exported.
	Data map[string]map[string][]string
}

// GetWildcardName will retrieve the configured wildcard name or provide a default
// in the case that the config is Nil or the wildcard name is unset.
func (c *Config) GetWildcardName() string {
	if c == nil || c.WildcardName == "" {
		return WildcardName
	}
	return c.WildcardName
}

// Close will relinquish any resources this Config might be holding on to or
// managing.
func (c *Config) Close() {
	if c != nil {
		c.EnterpriseConfig.Close()
	}
}

// AliasIfAnonymousToken returns the string "anonymous token" if
// accessorID is acl.AnonymousTokenID. Used for better
// UX when logging the accessorID.
func AliasIfAnonymousToken(accessorID string) string {
	if accessorID == AnonymousTokenID {
		return AnonymousTokenAlias
	}
	return accessorID
}
