// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/check.go
Check module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)
- [types](../types)

## Tags
agent, health-checks, monitoring

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/check.go> a code:Module ;
    code:name "agent/check.go" ;
    code:description "Check module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "types" ;
        code:path "../types"
    ] ;
    code:tags "agent", "health-checks", "monitoring" .
<!-- End LinkedDoc RDF -->
*/
package agent

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/types"
)

// persistedCheck is used to serialize a check and write it to disk
// so that it may be restored later on.
type persistedCheck struct {
	Check   *structs.HealthCheck
	ChkType *structs.CheckType
	Token   string
	Source  string
}

// persistedCheckState is used to persist the current state of a given
// check. This is different from the check definition, and includes an
// expiration timestamp which is used to determine staleness on later
// agent restarts.
type persistedCheckState struct {
	CheckID types.CheckID
	Output  string
	Status  string
	Expires int64
	acl.EnterpriseMeta
}
