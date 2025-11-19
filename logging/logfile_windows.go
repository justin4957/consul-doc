// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: logging/logfile_windows.go
Logfile Windows module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<logging/logfile_windows.go> a code:Module ;
    code:name "logging/logfile_windows.go" ;
    code:description "Logfile Windows module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package logging

import (
	"os"
	"time"
)

func (l *LogFile) createTime(stat os.FileInfo) time.Time {
	// Use `ModTime` as an approximation if the exact create time is not
	// available.
	// On Windows, the file create time is not updated after the active log
	// rotates, so use `ModTime` as an approximation as well.
	return stat.ModTime()
}
