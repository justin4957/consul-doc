// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/authmethod/testing.go
Testing module for agent layer

## Tags
agent, authentication, security

## Exports
RequireIdentityMatch

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/authmethod/testing.go> a code:Module ;
    code:name "agent/consul/authmethod/testing.go" ;
    code:description "Testing module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :RequireIdentityMatch ;
    code:tags "agent", "authentication", "security" .
<!-- End LinkedDoc RDF -->
*/
package authmethod

import (
	"sort"

	"github.com/hashicorp/go-bexpr"
	"github.com/mitchellh/go-testing-interface"
	"github.com/stretchr/testify/require"
)

// RequireIdentityMatch tests to see if the given Identity matches the provided
// projected vars and filters for testing purpose.
func RequireIdentityMatch(t testing.T, id *Identity, projectedVars map[string]string, filters ...string) {
	t.Helper()

	gotNames := id.ProjectedVarNames()

	require.Equal(t, projectedVars, id.ProjectedVars)

	expectNames := make([]string, 0, len(projectedVars))
	for k := range projectedVars {
		expectNames = append(expectNames, k)
	}
	sort.Strings(expectNames)
	sort.Strings(gotNames)

	require.Equal(t, expectNames, gotNames)
	require.Nil(t, id.EnterpriseMeta)

	for _, filter := range filters {
		eval, err := bexpr.CreateEvaluatorForType(filter, nil, id.SelectableFields)
		if err != nil {
			t.Fatalf("filter %q got err: %v", filter, err)
		}

		result, err := eval.Evaluate(id.SelectableFields)
		if err != nil {
			t.Fatalf("filter %q got err: %v", filter, err)
		}

		if !result {
			t.Fatalf("filter %q did not match", filter)
		}
	}
}
