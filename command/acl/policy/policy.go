// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/acl/policy/policy.go
Policy module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
access-control, authorization, cli, security, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/acl/policy/policy.go> a code:Module ;
    code:name "command/acl/policy/policy.go" ;
    code:description "Policy module for cli layer" ;
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
package policy

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

const synopsis = "Manage Consul's ACL policies"
const help = `
Usage: consul acl policy <subcommand> [options] [args]

  This command has subcommands for managing Consul's ACL policies.
  Here are some simple examples, and more detailed examples are available
  in the subcommands or the documentation.

  Create a new ACL policy:

      $ consul acl policy create -name "new-policy" \
                                 -description "This is an example policy" \
                                 -datacenter "dc1" \
                                 -datacenter "dc2" \
                                 -rules @rules.hcl
  List all policies:

      $ consul acl policy list

  Update a policy:

      $ consul acl policy update -name "other-policy" -datacenter "dc1"

  Read a policy:

    $ consul acl policy read -id 0479e93e-091c-4475-9b06-79a004765c24

  Delete a policy

    $ consul acl policy delete -name "my-policy"

  For more examples, ask for subcommand help or view the documentation.
`
