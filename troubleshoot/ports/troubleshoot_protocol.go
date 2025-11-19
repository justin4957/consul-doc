// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: troubleshoot/ports/troubleshoot_protocol.go
Troubleshoot Protocol module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<troubleshoot/ports/troubleshoot_protocol.go> a code:Module ;
    code:name "troubleshoot/ports/troubleshoot_protocol.go" ;
    code:description "Troubleshoot Protocol module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package ports

type troubleShootProtocol interface {
	dialPort(hostPort *hostPort) error
}
