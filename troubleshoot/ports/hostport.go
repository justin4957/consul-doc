// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: troubleshoot/ports/hostport.go
Hostport module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<troubleshoot/ports/hostport.go> a code:Module ;
    code:name "troubleshoot/ports/hostport.go" ;
    code:description "Hostport module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package ports

type hostPort struct {
	host string
	port string
}
