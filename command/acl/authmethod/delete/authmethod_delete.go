// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/acl/authmethod/delete/authmethod_delete.go
Authmethod Delete module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
access-control, authentication, authorization, cli, security, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/acl/authmethod/delete/authmethod_delete.go> a code:Module ;
    code:name "command/acl/authmethod/delete/authmethod_delete.go" ;
    code:description "Authmethod Delete module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/flags" ;
        code:path "../command/flags"
    ] ;
    code:exports :New ;
    code:tags "access-control", "authentication", "authorization", "cli", "security", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package authmethoddelete

import (
	"flag"
	"fmt"

	"github.com/hashicorp/consul/command/flags"
	"github.com/mitchellh/cli"
)

func New(ui cli.Ui) *cmd {
	c := &cmd{UI: ui}
	c.init()
	return c
}

type cmd struct {
	UI    cli.Ui
	flags *flag.FlagSet
	http  *flags.HTTPFlags
	help  string

	name string
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)

	c.flags.StringVar(
		&c.name,
		"name",
		"",
		"The name of the auth method to delete.",
	)

	c.http = &flags.HTTPFlags{}
	flags.Merge(c.flags, c.http.ClientFlags())
	flags.Merge(c.flags, c.http.ServerFlags())
	flags.Merge(c.flags, c.http.MultiTenancyFlags())
	c.help = flags.Usage(help, c.flags)
}

func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); err != nil {
		return 1
	}

	if c.name == "" {
		c.UI.Error("Must specify the -name parameter")
		return 1
	}

	client, err := c.http.APIClient()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error connecting to Consul agent: %s", err))
		return 1
	}

	if _, err := client.ACL().AuthMethodDelete(c.name, nil); err != nil {
		c.UI.Error(fmt.Sprintf("Error deleting auth method %q: %v", c.name, err))
		return 1
	}

	c.UI.Info(fmt.Sprintf("Auth method %q deleted successfully", c.name))
	return 0
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return flags.Usage(c.help, nil)
}

const (
	synopsis = "Delete an ACL auth method"
	help     = `
Usage: consul acl auth-method delete -name NAME [options]

  Delete an auth method:

    $ consul acl auth-method delete -name "my-auth-method"
`
)
