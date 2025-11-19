// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Package controller contains a re-implementation of the Kubernetes
// [controller-runtime](https://github.com/kubernetes-sigs/controller-runtime)
// with the core using Consul's event publishing pipeline rather than
// Kubernetes' client list/watch APIs.
//
// Generally this package enables defining asynchronous control loops
// meant to be run on a Consul cluster's leader that reconcile derived state
// in config entries that might be dependent on multiple sources.

/*
# Module: agent/consul/controller/doc.go
Doc module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/controller/doc.go> a code:Module ;
    code:name "agent/consul/controller/doc.go" ;
    code:description "Doc module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package controller
