// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/controller/cache/index/errors.go
Errors module for internal layer

## Tags
internal

## Exports
MissingRequiredIndexError

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/controller/cache/index/errors.go> a code:Module ;
    code:name "internal/controller/cache/index/errors.go" ;
    code:description "Errors module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :MissingRequiredIndexError ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package index

import (
	"fmt"
)

type MissingRequiredIndexError struct {
	Name string
}

func (err MissingRequiredIndexError) Error() string {
	return fmt.Sprintf("the indexer produced no value for the required %q index", err.Name)
}
