// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: acl/testing.go
Testing module for acl layer

## Tags
access-control, acl, authorization, security

## Exports
RequirePermissionDeniedError, RequirePermissionDeniedMessage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<acl/testing.go> a code:Module ;
    code:name "acl/testing.go" ;
    code:description "Testing module for acl layer" ;
    code:language "go" ;
    code:layer "acl" ;
    code:exports :RequirePermissionDeniedError, :RequirePermissionDeniedMessage ;
    code:tags "access-control", "acl", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package acl

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func RequirePermissionDeniedError(t testing.TB, err error, authz Authorizer, _ *AuthorizerContext, resource Resource, accessLevel AccessLevel, resourceID string) {
	t.Helper()
	if err == nil {
		t.Fatal("An error is expected but got nil.")
	}
	if v, ok := err.(PermissionDeniedError); ok {
		require.Equal(t, v.Resource, resource)
		require.Equal(t, v.AccessLevel, accessLevel)
		require.Equal(t, v.ResourceID.Name, resourceID)
	} else {
		t.Fatalf("Expected a permission denied error got %T %vp", err, err)
	}
}

func RequirePermissionDeniedMessage(t testing.TB, msg string, authz interface{}, _ *AuthorizerContext, resource Resource, accessLevel AccessLevel, resourceID string) {
	require.NotEmpty(t, msg, "expected non-empty error message")

	baseRegex := ` lacks permission '(\S*):(\S*)' on \"([^\"]*)\"(?: in partition \"([^\"]*)\" in namespace \"([^\"]*)\")?\s*$`

	var resourceIDFound string
	if authz == nil {
		expr := "^Permission denied" + `: provided token` + baseRegex
		re, _ := regexp.Compile(expr)
		matched := re.FindStringSubmatch(msg)

		require.NotNil(t, matched, fmt.Sprintf("RE %q didn't match %q", expr, msg))
		require.Equal(t, string(resource), matched[1], "resource")
		require.Equal(t, accessLevel.String(), matched[2], "access level")
		resourceIDFound = matched[3]
	} else {
		expr := "^Permission denied" + `: token with AccessorID '(\S*)'` + baseRegex
		re, _ := regexp.Compile(expr)
		matched := re.FindStringSubmatch(msg)

		require.NotNil(t, matched, fmt.Sprintf("RE %q didn't match %q", expr, msg))
		require.Equal(t, extractAccessorID(authz), matched[1], "auth")
		require.Equal(t, string(resource), matched[2], "resource")
		require.Equal(t, accessLevel.String(), matched[3], "access level")
		resourceIDFound = matched[4]
	}
	// AuthorizerContext information should be checked here
	require.Contains(t, resourceIDFound, resourceID, "resource id")
}
