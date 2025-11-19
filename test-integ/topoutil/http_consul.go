// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test-integ/topoutil/http_consul.go
Http Consul module for internal layer

## Tags
internal

## Exports
RequestRegisterService

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test-integ/topoutil/http_consul.go> a code:Module ;
    code:name "test-integ/topoutil/http_consul.go" ;
    code:description "Http Consul module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :RequestRegisterService ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package topoutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// RequestRegisterService registers a service at the given node address
// using consul http request.
//
// The service definition must be a JSON string.
func RequestRegisterService(clusterHttpCli *http.Client, nodeAddress string, serviceDefinition string, token string) error {
	var js json.RawMessage
	if err := json.Unmarshal([]byte(serviceDefinition), &js); err != nil {
		return fmt.Errorf("failed to unmarshal service definition: %s", err)
	}

	u, err := url.Parse(nodeAddress)
	if err != nil {
		return fmt.Errorf("failed to parse node address: %s", err)
	}
	u.Path = "/v1/agent/service/register"

	req, err := http.NewRequest(http.MethodPut, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return fmt.Errorf("failed to create request: %s", err)
	}

	if token != "" {
		req.Header.Set("X-Consul-Token", token)
	}

	resp, err := clusterHttpCli.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service: %s", resp.Status)
	}

	return nil
}
