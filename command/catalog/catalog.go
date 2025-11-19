// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/catalog/catalog.go
Catalog module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
cli, discovery, registry, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/catalog/catalog.go> a code:Module ;
    code:name "command/catalog/catalog.go" ;
    code:description "Catalog module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/flags" ;
        code:path "../command/flags"
    ] ;
    code:exports :New ;
    code:tags "cli", "discovery", "registry", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package catalog

import (
	"github.com/hashicorp/consul/command/flags"
	"github.com/mitchellh/cli"
)

func New() *cmd {
	return &cmd{}
}

type cmd struct{}

func (c *cmd) Run(args []string) int {
	return cli.RunResultHelp
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return flags.Usage(help, nil)
}

const synopsis = "Interact with the catalog"
const help = `
Usage: consul catalog <subcommand> [options] [args]

  This command has subcommands for interacting with Consul's catalog. The
  catalog should not be confused with the agent, although the APIs and
  responses may be similar.

  Here are some simple examples, and more detailed examples are available
  in the subcommands or the documentation.

  List all datacenters:

      $ consul catalog datacenters

  List all nodes:

      $ consul catalog nodes

  List all services:

      $ consul catalog services

  For more examples, ask for subcommand help or view the documentation.
`
