// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/rate/handler_ce.go
Handler Ce module for agent layer

## Tags
agent

## Exports
IPLimitConfig

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/rate/handler_ce.go> a code:Module ;
    code:name "agent/consul/rate/handler_ce.go" ;
    code:description "Handler Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :IPLimitConfig ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package rate

type IPLimitConfig struct{}

func (h *Handler) UpdateIPConfig(cfg IPLimitConfig) {
	// noop
}

func (h *Handler) ipGlobalLimit(op Operation) *limit {
	return nil
}

func (h *Handler) ipCategoryLimit(op Operation) *limit {
	return nil
}
