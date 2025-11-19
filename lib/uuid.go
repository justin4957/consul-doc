// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: lib/uuid.go
Uuid module for internal layer

## Tags
internal

## Exports
GenerateUUID, UUIDCheckFunc

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<lib/uuid.go> a code:Module ;
    code:name "lib/uuid.go" ;
    code:description "Uuid module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :GenerateUUID, :UUIDCheckFunc ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package lib

import (
	"github.com/hashicorp/go-uuid"
)

// UUIDCheckFunc should determine whether the given UUID is actually
// unique and allowed to be used
type UUIDCheckFunc func(string) (bool, error)

func GenerateUUID(checkFn UUIDCheckFunc) (string, error) {
	for {
		id, err := uuid.GenerateUUID()
		if err != nil {
			return "", err
		}

		if checkFn == nil {
			return id, nil
		}

		if ok, err := checkFn(id); err != nil {
			return "", err
		} else if ok {
			return id, nil
		}
	}
}
