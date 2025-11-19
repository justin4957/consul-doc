// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/acl/bindingrule/list/bindingrule_list.go
Bindingrule List module for cli layer

## Linked Modules
- [command/acl/bindingrule](../command/acl/bindingrule)
- [command/flags](../command/flags)

## Tags
access-control, authorization, cli, security, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/acl/bindingrule/list/bindingrule_list.go> a code:Module ;
    code:name "command/acl/bindingrule/list/bindingrule_list.go" ;
    code:description "Bindingrule List module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/acl/bindingrule" ;
        code:path "../command/acl/bindingrule"
    ], [
        code:name "command/flags" ;
        code:path "../command/flags"
    ] ;
    code:exports :New ;
    code:tags "access-control", "authorization", "cli", "security", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package bindingrulelist

import (
	"flag"
	"fmt"
	"strings"

	"github.com/hashicorp/consul/command/acl/bindingrule"
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

	authMethodName string

	showMeta bool
	format   string
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)

	c.flags.BoolVar(
		&c.showMeta,
		"meta",
		false,
		"Indicates that binding rule metadata such "+
			"as the raft indices should be shown for each entry.",
	)

	c.flags.StringVar(
		&c.authMethodName,
		"method",
		"",
		"Only show rules linked to the auth method with the given name.",
	)

	c.flags.StringVar(
		&c.format,
		"format",
		bindingrule.PrettyFormat,
		fmt.Sprintf("Output format {%s}", strings.Join(bindingrule.GetSupportedFormats(), "|")),
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

	client, err := c.http.APIClient()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error connecting to Consul agent: %s", err))
		return 1
	}

	rules, _, err := client.ACL().BindingRuleList(c.authMethodName, nil)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Failed to retrieve the binding rule list: %v", err))
		return 1
	}

	formatter, err := bindingrule.NewFormatter(c.format, c.showMeta)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	out, err := formatter.FormatBindingRuleList(rules)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	if out != "" {
		c.UI.Info(out)
	}

	return 0
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return flags.Usage(c.help, nil)
}

const (
	synopsis = "Lists ACL binding rules"
	help     = `
Usage: consul acl binding-rule list [options]

  Lists all the ACL binding rules.

  Show all:

    $ consul acl binding-rule list

  Show all for a specific auth method:

    $ consul acl binding-rule list -method="my-method"
`
)
