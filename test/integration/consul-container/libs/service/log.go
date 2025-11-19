// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/libs/service/log.go
Log module for internal layer

## Tags
internal

## Exports
LogConsumer

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/libs/service/log.go> a code:Module ;
    code:name "test/integration/consul-container/libs/service/log.go" ;
    code:description "Log module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :LogConsumer ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package service

import (
	"fmt"
	"os"

	"github.com/testcontainers/testcontainers-go"
)

type LogConsumer struct {
	Prefix string
}

var _ testcontainers.LogConsumer = (*LogConsumer)(nil)

func (c *LogConsumer) Accept(log testcontainers.Log) {
	switch log.LogType {
	case "STDOUT":
		fmt.Fprint(os.Stdout, c.Prefix+" ~~ "+string(log.Content))
	case "STDERR":
		fmt.Fprint(os.Stderr, c.Prefix+" ~~ "+string(log.Content))
	}
}
