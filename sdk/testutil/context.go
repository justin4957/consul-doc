// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: sdk/testutil/context.go
Context module for internal layer

## Tags
internal

## Exports
CleanerT, TestContext

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sdk/testutil/context.go> a code:Module ;
    code:name "sdk/testutil/context.go" ;
    code:description "Context module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :CleanerT, :TestContext ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package testutil

import (
	"context"
)

type CleanerT interface {
	Helper()
	Cleanup(func())
}

func TestContext(t CleanerT) context.Context {
	t.Helper()
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	return ctx
}
