// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/version/version.go
Version module for cli layer

## Linked Modules
- [agent/consul](../agent/consul)
- [command/flags](../command/flags)
- [version](../version)

## Tags
cli, user-interface

## Exports
New, RPCVersionInfo, VersionInfo

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/version/version.go> a code:Module ;
    code:name "command/version/version.go" ;
    code:description "Version module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "agent/consul" ;
        code:path "../agent/consul"
    ], [
        code:name "command/flags" ;
        code:path "../command/flags"
    ], [
        code:name "version" ;
        code:path "../version"
    ] ;
    code:exports :New, :RPCVersionInfo, :VersionInfo ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package version

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/mitchellh/cli"

	"github.com/hashicorp/consul/agent/consul"
	"github.com/hashicorp/consul/command/flags"
	"github.com/hashicorp/consul/version"
)

func New(ui cli.Ui) *cmd {
	c := &cmd{UI: ui}
	c.init()
	return c
}

type cmd struct {
	UI     cli.Ui
	flags  *flag.FlagSet
	format string
	help   string
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	c.flags.StringVar(
		&c.format,
		"format",
		PrettyFormat,
		fmt.Sprintf("Output format {%s}", strings.Join(GetSupportedFormats(), "|")))
	c.help = flags.Usage(help, c.flags)

}

type RPCVersionInfo struct {
	Default int
	Min     int
	Max     int
}

type VersionInfo struct {
	HumanVersion string `json:"-"`
	Version      string
	Revision     string
	Prerelease   string
	FIPS         string `json:",omitempty"`
	BuildDate    time.Time
	RPC          RPCVersionInfo
}

func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); err != nil {
		return 1
	}

	formatter, err := NewFormatter(c.format)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	// We parse this here because consul version is used in our 'smoke' tests and we want to fail early
	buildDate, err := time.Parse(time.RFC3339, version.BuildDate)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	out, err := formatter.Format(&VersionInfo{
		HumanVersion: version.GetHumanVersion(),
		Version:      version.Version,
		Revision:     version.GitCommit,
		Prerelease:   version.VersionPrerelease,
		BuildDate:    buildDate,
		FIPS:         version.GetFIPSInfo(),
		RPC: RPCVersionInfo{
			Default: consul.DefaultRPCProtocol,
			Min:     int(consul.ProtocolVersionMin),
			Max:     consul.ProtocolVersionMax,
		},
	})
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}
	c.UI.Output(out)
	return 0
}

func (c *cmd) Synopsis() string {
	return "Prints the Consul version"
}

func (c *cmd) Help() string {
	return flags.Usage(c.help, nil)
}

const synopsis = "Output Consul version information"
const help = `
Usage: consul version [options]
`
