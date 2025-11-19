// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: logging/log_levels.go
Log Levels module for internal layer

## Tags
internal

## Exports
AllowedLogLevels, LevelFromString, ValidateLogLevel

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<logging/log_levels.go> a code:Module ;
    code:name "logging/log_levels.go" ;
    code:description "Log Levels module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :AllowedLogLevels, :LevelFromString, :ValidateLogLevel ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package logging

import (
	"strings"

	"github.com/hashicorp/go-hclog"
)

var (
	allowedLogLevels = []string{"TRACE", "DEBUG", "INFO", "WARN", "ERR", "ERROR"}
)

func AllowedLogLevels() []string {
	var c []string
	copy(c, allowedLogLevels)
	return c
}

// ValidateLogLevel verifies that a new log level is valid
func ValidateLogLevel(minLevel string) bool {
	newLevel := strings.ToUpper(minLevel)
	for _, level := range allowedLogLevels {
		if level == newLevel {
			return true
		}
	}
	return false
}

// Backwards compatibility with former ERR log level
func LevelFromString(level string) hclog.Level {
	if strings.ToUpper(level) == "ERR" {
		level = "ERROR"
	}
	return hclog.LevelFromString(level)
}
