// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !linux

/*
# Module: sdk/iptables/iptables_executor_unsupported.go
Iptables Executor Unsupported module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sdk/iptables/iptables_executor_unsupported.go> a code:Module ;
    code:name "sdk/iptables/iptables_executor_unsupported.go" ;
    code:description "Iptables Executor Unsupported module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package iptables

import "errors"

// iptablesExecutor implements IptablesProvider and errors out on any non-linux OS.
type iptablesExecutor struct {
	cfg Config
}

func (i *iptablesExecutor) AddRule(_ string, _ ...string) {}

func (i *iptablesExecutor) ApplyRules(string) error {
	return errors.New("applying traffic redirection rules with 'iptables' is not supported on this operating system; only linux OS is supported")
}

func (i *iptablesExecutor) Rules() []string {
	return nil
}

func (i *iptablesExecutor) ClearAllRules() {
}
