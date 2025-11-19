// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/kv/kv.go
CLI command parent for key-value store operations (get, put, delete).

## Linked Modules
- [api/kv](../../api/kv.go) - KV API client
- [command/registry](../registry.go) - Command registry

## Tags
cli, kv, storage, commands

## Exports
New, cmd, Run

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/kv/kv.go> a code:Module ;
    code:language "go" ;
    code:name "command/kv/kv.go" ;
    code:description "CLI command parent for key-value store operations (get, put, delete)" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "api/kv" ;
        code:path "../../api/kv.go" ;
        code:relationship "KV API client"
    ], [
        code:name "command/registry" ;
        code:path "../registry.go" ;
        code:relationship "Command registry"
    ] ;
    code:exports :New, :cmd, :Run ;
    code:tags "cli", "kv", "storage", "commands" .
<!-- End LinkedDoc RDF -->
*/
package kv

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

const synopsis = "Interact with the key-value store"
const help = `
Usage: consul kv <subcommand> [options] [args]

  This command has subcommands for interacting with Consul's key-value
  store. Here are some simple examples, and more detailed examples are
  available in the subcommands or the documentation.

  Create or update the key named "redis/config/connections" with the value "5":

      $ consul kv put redis/config/connections 5

  Read this value back:

      $ consul kv get redis/config/connections

  Or get detailed key information:

      $ consul kv get -detailed redis/config/connections

  Finally, delete the key:

      $ consul kv delete redis/config/connections

  For more examples, ask for subcommand help or view the documentation.
`
