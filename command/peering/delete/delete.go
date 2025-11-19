// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/peering/delete/delete.go
Delete module for cli layer

## Linked Modules
- [api](../api)
- [command/flags](../command/flags)

## Tags
cli, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/peering/delete/delete.go> a code:Module ;
    code:name "command/peering/delete/delete.go" ;
    code:description "Delete module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ], [
        code:name "command/flags" ;
        code:path "../command/flags"
    ] ;
    code:exports :New ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package delete

import (
	"context"
	"flag"
	"fmt"

	"github.com/mitchellh/cli"

	"github.com/hashicorp/consul/api"
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
	http  *flags.HTTPFlags
	help  string

	name string
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)

	c.flags.StringVar(&c.name, "name", "", "(Required) The local name assigned to the peer cluster.")

	c.http = &flags.HTTPFlags{}
	flags.Merge(c.flags, c.http.ClientFlags())
	flags.Merge(c.flags, c.http.PartitionFlag())
	c.help = flags.Usage(help, c.flags)
}

func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); err != nil {
		return 1
	}

	if c.name == "" {
		c.UI.Error("Missing the required -name flag")
		return 1
	}

	client, err := c.http.APIClient()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error connecting to Consul agent: %s", err))
		return 1
	}

	peerings := client.Peerings()

	_, err = peerings.Delete(context.Background(), c.name, &api.WriteOptions{})
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error deleting peering for %s: %v", c.name, err))
		return 1
	}

	c.UI.Info(fmt.Sprintf("Successfully submitted peering connection, %s, for deletion", c.name))
	return 0
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return flags.Usage(c.help, nil)
}

const (
	synopsis = "Delete a peering connection"
	help     = `
Usage: consul peering delete [options] -name <peer name>

  Delete a peering connection.  Consul deletes all data imported from the peer 
  in the background. The peering connection is removed after all associated 
  data has been deleted. Operators can still read the peering connections 
  while the data is being removed. A 'DeletedAt' field will be populated with 
  the timestamp of when the peering was marked for deletion.

  Example:

    $ consul peering delete -name west-dc
`
)
