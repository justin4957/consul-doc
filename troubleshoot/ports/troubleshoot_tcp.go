// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: troubleshoot/ports/troubleshoot_tcp.go
Troubleshoot Tcp module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<troubleshoot/ports/troubleshoot_tcp.go> a code:Module ;
    code:name "troubleshoot/ports/troubleshoot_tcp.go" ;
    code:description "Troubleshoot Tcp module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package ports

import (
	"net"
	"time"
)

type troubleShootTcp struct {
}

func (tcp *troubleShootTcp) dialPort(hostPort *hostPort) error {
	address := net.JoinHostPort(hostPort.host, hostPort.port)

	// Attempt to establish a TCP connection with a timeout.
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}
