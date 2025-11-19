// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/enterprise_delegate_ce.go
Enterprise Delegate Ce module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/enterprise_delegate_ce.go> a code:Module ;
    code:name "agent/enterprise_delegate_ce.go" ;
    code:description "Enterprise Delegate Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package agent

// enterpriseDelegate has no functions in CE
type enterpriseDelegate interface{}
