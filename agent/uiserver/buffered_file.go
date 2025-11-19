// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/uiserver/buffered_file.go
Buffered File module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/uiserver/buffered_file.go> a code:Module ;
    code:name "agent/uiserver/buffered_file.go" ;
    code:description "Buffered File module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package uiserver

import (
	"bytes"
	"io/fs"
	"os"
	"time"
)

// bufferedFile implements fs.File and allows us to modify a file from disk by
// writing out the new version into a buffer and then serving file reads from
// that.
type bufferedFile struct {
	buf  *bytes.Buffer
	info fs.FileInfo
}

var _ fs.FileInfo = (*bufferedFile)(nil)

func newBufferedFile(buf *bytes.Buffer, info fs.FileInfo) *bufferedFile {
	return &bufferedFile{
		buf:  buf,
		info: info,
	}
}

func (b *bufferedFile) Stat() (fs.FileInfo, error) {
	return b, nil
}

func (b *bufferedFile) Read(bytes []byte) (int, error) {
	return b.buf.Read(bytes)
}

func (b *bufferedFile) Close() error {
	return nil
}

func (b *bufferedFile) Name() string {
	return b.info.Name()
}

func (b *bufferedFile) Size() int64 {
	return int64(b.buf.Len())
}

func (b *bufferedFile) Mode() os.FileMode {
	return b.info.Mode()
}

func (b *bufferedFile) ModTime() time.Time {
	return b.info.ModTime()
}

func (b *bufferedFile) IsDir() bool {
	return false
}

func (b *bufferedFile) Sys() any {
	return nil
}
