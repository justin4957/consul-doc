// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: main.go
Main module for internal layer

## Linked Modules
- [command](../command)
- [command/cli](../command/cli)
- [command/version](../command/version)
- [service_os](../service_os)

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<main.go> a code:Module ;
    code:name "main.go" ;
    code:description "Main module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "command" ;
        code:path "../command"
    ], [
        code:name "command/cli" ;
        code:path "../command/cli"
    ], [
        code:name "command/version" ;
        code:path "../command/version"
    ], [
        code:name "service_os" ;
        code:path "../service_os"
    ] ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	mcli "github.com/mitchellh/cli"

	"github.com/hashicorp/consul/command"
	"github.com/hashicorp/consul/command/cli"
	"github.com/hashicorp/consul/command/version"
	_ "github.com/hashicorp/consul/service_os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	log.SetOutput(io.Discard)

	ui := &cli.BasicUI{
		BasicUi: mcli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr},
	}
	cmds := command.RegisteredCommands(ui)
	var names []string
	for c := range cmds {
		names = append(names, c)
	}

	cli := &mcli.CLI{
		Args:         os.Args[1:],
		Commands:     cmds,
		Autocomplete: true,
		Name:         "consul",
		HelpFunc:     mcli.FilteredHelpFunc(names, mcli.BasicHelpFunc("consul")),
		HelpWriter:   os.Stdout,
		ErrorWriter:  os.Stderr,
	}

	if cli.IsVersion() {
		cmd := version.New(ui)
		return cmd.Run(nil)
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %v\n", err)
		return 1
	}

	return exitCode
}
