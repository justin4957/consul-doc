// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/test/envoy_extensions/testdata/wasm_test_files/wasm_add_header.go
Wasm Add Header module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/test/envoy_extensions/testdata/wasm_test_files/wasm_add_header.go> a code:Module ;
    code:name "test/integration/consul-container/test/envoy_extensions/testdata/wasm_test_files/wasm_add_header.go" ;
    code:description "Wasm Add Header module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultPluginContext
}

func (p *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpHeaders{}
}

type httpHeaders struct {
	// Embed the default http context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultHttpContext
}

func (ctx *httpHeaders) OnHttpResponseHeaders(int, bool) types.Action {
	proxywasm.LogDebug("adding header: x-test:true")

	err := proxywasm.AddHttpResponseHeader("x-test", "true")
	if err != nil {
		proxywasm.LogCriticalf("failed to add test header to response: %v", err)
	}

	return types.ActionContinue
}
