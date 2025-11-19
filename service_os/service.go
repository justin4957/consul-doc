// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: service_os/service.go
Service module for internal layer

## Tags
internal

## Exports
Shutdown

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<service_os/service.go> a code:Module ;
    code:name "service_os/service.go" ;
    code:description "Service module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :Shutdown ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package service_os

var chanGraceExit = make(chan int)

func Shutdown_Channel() <-chan int {
	return chanGraceExit
}
