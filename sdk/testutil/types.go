// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: sdk/testutil/types.go
Types module for internal layer

## Tags
data-model, internal, types

## Exports
TestingTB

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sdk/testutil/types.go> a code:Module ;
    code:name "sdk/testutil/types.go" ;
    code:description "Types module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :TestingTB ;
    code:tags "data-model", "internal", "types" .
<!-- End LinkedDoc RDF -->
*/
package testutil

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var nilInf TestingTB = nil

// Assertion that our TestingTB can be passed to
var _ require.TestingT = nilInf
var _ assert.TestingT = nilInf

// TestingTB is an interface that describes the implementation of the testing object.
// Using an interface that describes testing.TB instead of the actual implementation
// makes testutil usable in a wider variety of contexts (e.g. use with ginkgo : https://godoc.org/github.com/onsi/ginkgo#GinkgoT)
type TestingTB interface {
	Cleanup(func())
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	Setenv(key, value string)
	TempDir() string
}
