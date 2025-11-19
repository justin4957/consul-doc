// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: api/exported_services.go
Exported Services module for api layer

## Tags
api, client

## Exports
ResolvedConsumers, ResolvedExportedService

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<api/exported_services.go> a code:Module ;
    code:name "api/exported_services.go" ;
    code:description "Exported Services module for api layer" ;
    code:language "go" ;
    code:layer "api" ;
    code:exports :ResolvedConsumers, :ResolvedExportedService ;
    code:tags "api", "client" .
<!-- End LinkedDoc RDF -->
*/
package api

type ResolvedExportedService struct {
	// Service is the name of the service which is exported.
	Service string

	// Partition of the service
	Partition string `json:",omitempty"`

	// Namespace of the service
	Namespace string `json:",omitempty"`

	// Consumers is a list of downstream consumers of the service.
	Consumers ResolvedConsumers
}

type ResolvedConsumers struct {
	Peers      []string `json:",omitempty"`
	Partitions []string `json:",omitempty"`
}

func (c *Client) ExportedServices(q *QueryOptions) ([]ResolvedExportedService, *QueryMeta, error) {

	r := c.newRequest("GET", "/v1/exported-services")
	r.setQueryOptions(q)
	rtt, resp, err := c.doRequest(r)
	if err != nil {
		return nil, nil, err
	}
	defer closeResponseBody(resp)
	if err := requireOK(resp); err != nil {
		return nil, nil, err
	}

	qm := &QueryMeta{}
	parseQueryMeta(resp, qm)
	qm.RequestTime = rtt

	var expSvcs []ResolvedExportedService

	if err := decodeBody(resp, &expSvcs); err != nil {
		return nil, nil, err
	}

	return expSvcs, qm, nil
}
