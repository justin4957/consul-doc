// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/flags/flag_map_value.go
Flag Map Value module for cli layer

## Tags
cli, user-interface

## Exports
FlagMapValue

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/flags/flag_map_value.go> a code:Module ;
    code:name "command/flags/flag_map_value.go" ;
    code:description "Flag Map Value module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:exports :FlagMapValue ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package flags

import (
	"flag"
	"fmt"
	"strings"
)

// Ensure implements
var _ flag.Value = (*FlagMapValue)(nil)

// FlagMapValue is a flag implementation used to provide key=value semantics
// multiple times.
type FlagMapValue map[string]string

func (h *FlagMapValue) String() string {
	return fmt.Sprintf("%v", *h)
}

func (h *FlagMapValue) Set(value string) error {
	idx := strings.Index(value, "=")
	if idx == -1 {
		return fmt.Errorf("Missing \"=\" value in argument: %s", value)
	}

	key, value := value[0:idx], value[idx+1:]

	if *h == nil {
		*h = make(map[string]string)
	}

	headers := *h
	headers[key] = value
	*h = headers

	return nil
}

// Merge will overlay this value if it has been set.
func (h *FlagMapValue) Merge(onto map[string]string) {
	if h == nil || onto == nil {
		return
	}
	for k, v := range *h {
		if _, ok := onto[k]; !ok {
			onto[k] = v
		}
	}
}
