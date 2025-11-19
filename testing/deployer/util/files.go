// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: testing/deployer/util/files.go
Files module for internal layer

## Tags
internal

## Exports
AddFileToHash, FileExists, FilesExist, HashFile

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<testing/deployer/util/files.go> a code:Module ;
    code:name "testing/deployer/util/files.go" ;
    code:description "Files module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :AddFileToHash, :FileExists, :FilesExist, :HashFile ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package util

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/crypto/blake2b"
)

func FilesExist(parent string, paths ...string) (bool, error) {
	for _, p := range paths {
		ok, err := FileExists(filepath.Join(parent, p))
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func HashFile(path string) (string, error) {
	hash, err := blake2b.New256(nil)
	if err != nil {
		return "", err
	}

	if err := AddFileToHash(path, hash); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func AddFileToHash(path string, w io.Writer) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(w, f)
	return err
}
