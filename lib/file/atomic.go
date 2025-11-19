// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: lib/file/atomic.go
Atomic module for internal layer

## Tags
internal

## Exports
WriteAtomic, WriteAtomicWithPerms

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<lib/file/atomic.go> a code:Module ;
    code:name "lib/file/atomic.go" ;
    code:description "Atomic module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :WriteAtomic, :WriteAtomicWithPerms ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package file

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hashicorp/go-uuid"
)

// WriteAtomic writes the given contents to a temporary file in the same
// directory, does an fsync and then renames the file to its real path
func WriteAtomic(path string, contents []byte) error {
	return WriteAtomicWithPerms(path, contents, 0700, 0600)
}

func WriteAtomicWithPerms(path string, contents []byte, dirPerms, filePerms os.FileMode) error {

	uuid, err := uuid.GenerateUUID()
	if err != nil {
		return err
	}
	tempPath := fmt.Sprintf("%s-%s.tmp", path, uuid)

	if err := os.MkdirAll(filepath.Dir(path), dirPerms); err != nil {
		return err
	}
	fh, err := os.OpenFile(tempPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, filePerms)
	if err != nil {
		return err
	}
	if _, err := fh.Write(contents); err != nil {
		fh.Close()
		os.Remove(tempPath)
		return err
	}
	if err := fh.Sync(); err != nil {
		fh.Close()
		os.Remove(tempPath)
		return err
	}
	if err := fh.Close(); err != nil {
		os.Remove(tempPath)
		return err
	}
	if err := os.Rename(tempPath, path); err != nil {
		os.Remove(tempPath)
		return err
	}
	return nil
}
