// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: test/integration/consul-container/libs/utils/version_ce.go
Version Ce module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/libs/utils/version_ce.go> a code:Module ;
    code:name "test/integration/consul-container/libs/utils/version_ce.go" ;
    code:description "Version Ce module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package utils

const (
	defaultImageName   = DefaultImageNameCE
	ImageVersionSuffix = ""
	isInEnterpriseRepo = false
)
