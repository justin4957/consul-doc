// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: sdk/testutil/retry/output.go
Output module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sdk/testutil/retry/output.go> a code:Module ;
    code:name "sdk/testutil/retry/output.go" ;
    code:description "Output module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package retry

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

func dedup(a []string) string {
	if len(a) == 0 {
		return ""
	}
	seen := map[string]struct{}{}
	var b bytes.Buffer
	for _, s := range a {
		if _, ok := seen[s]; ok {
			continue
		}
		seen[s] = struct{}{}
		b.WriteString(s)
		b.WriteRune('\n')
	}
	return b.String()
}

func decorate(s string) string {
	_, file, line, ok := runtime.Caller(3)
	if ok {
		n := strings.LastIndex(file, "/")
		if n >= 0 {
			file = file[n+1:]
		}
	} else {
		file = "???"
		line = 1
	}
	return fmt.Sprintf("%s:%d: %s", file, line, s)
}
