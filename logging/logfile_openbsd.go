// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build openbsd

/*
# Module: logging/logfile_openbsd.go
Logfile Openbsd module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<logging/logfile_openbsd.go> a code:Module ;
    code:name "logging/logfile_openbsd.go" ;
    code:description "Logfile Openbsd module for internal layer" ;
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
	createTime := stat_t.Ctim
	return time.Unix(int64(createTime.Sec), int64(createTime.Nsec)) //nolint:unconvert
}
