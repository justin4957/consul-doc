// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: command/troubleshoot/ports/troubleshoot_ports.go
Troubleshoot Ports module for cli layer

## Linked Modules
- [command/cli](../command/cli)
- [command/flags](../command/flags)
- [troubleshoot/ports](../troubleshoot/ports)

## Tags
cli, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/troubleshoot/ports/troubleshoot_ports.go> a code:Module ;
    code:name "command/troubleshoot/ports/troubleshoot_ports.go" ;
    code:description "Troubleshoot Ports module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/cli" ;
        code:path "../command/cli"
    ], [
        code:name "command/flags" ;
        code:path "../command/flags"
    ], [
        code:name "troubleshoot/ports" ;
        code:path "../troubleshoot/ports"
    ] ;
    code:exports :New ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package ports

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/troubleshoot/ports"
	"os"

	"github.com/hashicorp/consul/command/cli"
	"github.com/hashicorp/consul/command/flags"
)

func New(ui cli.Ui) *cmd {
	c := &cmd{UI: ui}
	c.init()
	return c
}

type cmd struct {
	UI    cli.Ui
	flags *flag.FlagSet
	help  string

	// flags
	host  string
	ports string
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)

	c.flags.StringVar(&c.host, "host", os.Getenv("CONSUL_HTTP_ADDR"), "The consul server host")

	c.flags.StringVar(&c.ports, "ports", "", "Custom ports to troubleshoot")

	c.help = flags.Usage(help, c.flags)
}

func (c *cmd) Run(args []string) int {

	if err := c.flags.Parse(args); err != nil {
		c.UI.Error(fmt.Sprintf("Failed to parse args: %v", err))
		return 1
	}

	if c.host == "" {
		c.UI.Error("-host is required. or set environment variable CONSUL_HTTP_ADDR")
		return 1
	}

	if c.ports == "" {
		ports.TroubleshootDefaultPorts(c.host)
	} else {
		ports.TroubleShootCustomPorts(c.host, c.ports)
	}
	return 0
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return c.help
}

const (
	synopsis = "Prints open and closed ports on the Consul server"
	help     = `
Usage: consul troubleshoot ports [options]
	Checks ports for TCP connectivity. Add the -ports flag to check specific ports or omit the -ports flag to check default ports. 
	Refer to the following reference for default ports: https://developer.hashicorp.com/consul/docs/install/ports

	consul troubleshoot ports -host localhost

	or 
	export CONSUL_HTTP_ADDR=localhost
	consul troubleshoot ports 
	
	Use the -ports flag to check non-default ports, for example:
	consul troubleshoot ports -host localhost -ports 1023,1024
	or 
	export CONSUL_HTTP_ADDR=localhost
	consul troubleshoot ports -ports 1234,8500 
`
)
