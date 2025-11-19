// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resource/resourcetest/fs.go
Fs module for internal layer

## Linked Modules
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
ParseResourcesFromFilesystem

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/resourcetest/fs.go> a code:Module ;
    code:name "internal/resource/resourcetest/fs.go" ;
    code:description "Fs module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :ParseResourcesFromFilesystem ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package resourcetest

import (
	"fmt"
	"io/fs"

	"github.com/hashicorp/consul/proto-public/pbresource"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

// ParseResourcesFromFilesystem will walk the filesystem at the given path
// and parse all files as protobuf/JSON resources.
func ParseResourcesFromFilesystem(t T, files fs.FS, path string) []*pbresource.Resource {
	t.Helper()

	var resources []*pbresource.Resource
	err := fs.WalkDir(files, path, func(fpath string, dent fs.DirEntry, _ error) error {
		if dent.IsDir() {
			return nil
		}

		data, err := fs.ReadFile(files, fpath)
		if err != nil {
			return err
		}

		var res pbresource.Resource
		err = protojson.Unmarshal(data, &res)
		if err != nil {
			return fmt.Errorf("error decoding data from %s: %w", fpath, err)
		}

		resources = append(resources, &res)
		return nil
	})

	require.NoError(t, err)
	return resources
}
