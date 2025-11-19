// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build windows

/*
# Module: service_os/service_windows.go
Service Windows module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<service_os/service_windows.go> a code:Module ;
    code:name "service_os/service_windows.go" ;
    code:description "Service Windows module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package service_os

import (
	wsvc "golang.org/x/sys/windows/svc"
)

type serviceWindows struct{}

func init() {
	interactive, err := wsvc.IsAnInteractiveSession()
	if err != nil {
		panic(err)
	}
	if interactive {
		return
	}
	go func() {
		_ = wsvc.Run("", serviceWindows{})
	}()
}

func (serviceWindows) Execute(args []string, r <-chan wsvc.ChangeRequest, s chan<- wsvc.Status) (svcSpecificEC bool, exitCode uint32) {
	const accCommands = wsvc.AcceptStop | wsvc.AcceptShutdown
	s <- wsvc.Status{State: wsvc.StartPending}

	s <- wsvc.Status{State: wsvc.Running, Accepts: accCommands}
	for {
		c := <-r
		switch c.Cmd {
		case wsvc.Interrogate:
			s <- c.CurrentStatus
		case wsvc.Stop, wsvc.Shutdown:
			chanGraceExit <- 1
			s <- wsvc.Status{State: wsvc.StopPending}
			return false, 0
		}
	}

	return false, 0
}
