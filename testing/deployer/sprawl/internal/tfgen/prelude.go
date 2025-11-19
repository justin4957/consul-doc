// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: testing/deployer/sprawl/internal/tfgen/prelude.go
Prelude module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<testing/deployer/sprawl/internal/tfgen/prelude.go> a code:Module ;
    code:name "testing/deployer/sprawl/internal/tfgen/prelude.go" ;
    code:description "Prelude module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package tfgen

const terraformPrelude = `provider "docker" {
  host = "unix:///var/run/docker.sock"
}

terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 2.0"
    }
  }
  required_version = ">= 0.13"
}
`
