// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: lib/path.go
Path module for internal layer

## Tags
internal

## Exports
EnsurePath

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<lib/path.go> a code:Module ;
    code:name "lib/path.go" ;
    code:description "Path module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :EnsurePath ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package lib

import (
	"os"
	"path/filepath"
)

// EnsurePath is used to make sure a path exists
func EnsurePath(path string, dir bool) error {
	if !dir {
		path = filepath.Dir(path)
	}
	return os.MkdirAll(path, 0755)
}
