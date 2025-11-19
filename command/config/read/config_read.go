// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/config/read/config_read.go
Config Read module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
cli, configuration, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/config/read/config_read.go> a code:Module ;
    code:name "command/config/read/config_read.go" ;
    code:description "Config Read module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/flags" ;
        code:path "../command/flags"
    ] ;
    code:exports :New ;
    code:tags "cli", "configuration", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package read

import (
	"encoding/json"
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

	kind string
	name string
}

func (c *cmd) init() {
	// TODO(rb): needs a way to print the metadata so you know the modify index to use for 'config write -cas'

	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	c.flags.StringVar(&c.kind, "kind", "", "The kind of configuration to read.")
	c.flags.StringVar(&c.name, "name", "", "The name of configuration to read.")
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

	if c.kind == "" {
		c.UI.Error("Must specify the -kind parameter")
		return 1
	}

	if c.name == "" {
		c.UI.Error("Must specify the -name parameter")
		return 1
	}

	client, err := c.http.APIClient()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error connect to Consul agent: %s", err))
		return 1
	}

	entry, _, err := client.ConfigEntries().Get(c.kind, c.name, nil)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error reading config entry %s/%s: %v", c.kind, c.name, err))
		return 1
	}

	b, err := json.MarshalIndent(entry, "", "    ")
	if err != nil {
		c.UI.Error("Failed to encode output data")
		return 1
	}

	c.UI.Info(string(b))
	return 0
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return flags.Usage(c.help, nil)
}

const (
	synopsis = "Read a centralized config entry"
	help     = `
Usage: consul config read [options] -kind <config kind> -name <config name>

  Reads the config entry specified by the given kind and name and outputs its
  JSON representation.

  Example:

    $ consul config read -kind proxy-defaults -name global
`
)
