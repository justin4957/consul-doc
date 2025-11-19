// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: command/registry_ce.go
Registry Ce module for cli layer

## Linked Modules
- [command/cli](../command/cli)

## Tags
cli, user-interface

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/registry_ce.go> a code:Module ;
    code:name "command/registry_ce.go" ;
    code:description "Registry Ce module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/cli" ;
        code:path "../command/cli"
    ] ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package command

import (
	mcli "github.com/mitchellh/cli"

	"github.com/hashicorp/consul/command/cli"
)

func registerEnterpriseCommands(_ cli.Ui, _ map[string]mcli.CommandFactory) {}
