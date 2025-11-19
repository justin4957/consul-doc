// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/acl/templatedpolicy/templated_policy.go
Templated Policy module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
access-control, authorization, cli, security, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/acl/templatedpolicy/templated_policy.go> a code:Module ;
    code:name "command/acl/templatedpolicy/templated_policy.go" ;
    code:description "Templated Policy module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/flags" ;
        code:path "../command/flags"
    ] ;
    code:exports :New ;
    code:tags "access-control", "authorization", "cli", "security", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package templatedpolicy

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

const synopsis = "Manage Consul's ACL templated policies"
const help = `
Usage: consul acl templated-policy <subcommand> [options] [args]

  This command has subcommands for managing Consul ACL templated policies.
  Here are some simple examples, and more detailed examples are available
  in the subcommands or the documentation.

  List all templated policies:

      $ consul acl templated-policy list

  Preview the policy rendered by the ACL templated policy:

      $ consul acl templated-policy preview -name "builtin/service" -var "name:api"

  Read a templated policy with name:

      $ consul acl templated-policy read -name "builtin/service"

  For more examples, ask for subcommand help or view the documentation.
`
