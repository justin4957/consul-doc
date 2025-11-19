// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build darwin || freebsd || netbsd

/*
# Module: logging/logfile_darwinandfreebsd.go
Logfile Darwinandfreebsd module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<logging/logfile_darwinandfreebsd.go> a code:Module ;
    code:name "logging/logfile_darwinandfreebsd.go" ;
    code:description "Logfile Darwinandfreebsd module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package logging

import (
	"os"
	"syscall"
	"time"
)

func (l *LogFile) createTime(stat os.FileInfo) time.Time {
	stat_t := stat.Sys().(*syscall.Stat_t)
	createTime := stat_t.Ctimespec
	// Sec and Nsec are int32 in 32-bit architectures.
	return time.Unix(int64(createTime.Sec), int64(createTime.Nsec)) //nolint:unconvert
}
