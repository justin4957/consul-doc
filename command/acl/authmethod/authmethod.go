// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/acl/authmethod/authmethod.go
Authmethod module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
access-control, authentication, authorization, cli, security, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/acl/authmethod/authmethod.go> a code:Module ;
    code:name "command/acl/authmethod/authmethod.go" ;
    code:description "Authmethod module for cli layer" ;
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
package authmethod

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

const synopsis = "Manage Consul's ACL auth methods"
const help = `
Usage: consul acl auth-method <subcommand> [options] [args]

  This command has subcommands for managing Consul's ACL auth methods.
  Here are some simple examples, and more detailed examples are available in
  the subcommands or the documentation.

  Create a new auth method:

    $ consul acl auth-method create -type "kubernetes" \
                            -name "my-k8s" \
                            -description "This is an example kube auth method" \
                            -kubernetes-host "https://apiserver.example.com:8443" \
                            -kubernetes-ca-cert @/path/to/kube.ca.crt \
                            -kubernetes-service-account-jwt "JWT_CONTENTS"

  List all auth methods:

    $ consul acl auth-method list

  Update all editable fields of the auth method:

    $ consul acl auth-method update -name "my-k8s" \
                            -description "new description" \
                            -kubernetes-host "https://new-apiserver.example.com:8443" \
                            -kubernetes-ca-cert @/path/to/new-kube.ca.crt \
                            -kubernetes-service-account-jwt "NEW_JWT_CONTENTS"

  Read an auth method:

    $ consul acl auth-method read -name my-k8s

  Delete an auth method:

    $ consul acl auth-method delete -name my-k8s

  For more examples, ask for subcommand help or view the documentation.
`
