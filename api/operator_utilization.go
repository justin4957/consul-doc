// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: api/operator_utilization.go
Operator Utilization module for api layer

## Tags
api, client

## Exports
Census, UtilizationBundleRequest

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<api/operator_utilization.go> a code:Module ;
    code:name "api/operator_utilization.go" ;
    code:description "Operator Utilization module for api layer" ;
    code:language "go" ;
    code:layer "api" ;
    code:exports :Census, :UtilizationBundleRequest ;
    code:tags "api", "client" .
<!-- End LinkedDoc RDF -->
*/
package api

import (
	"io"
	"strconv"
)

type Census struct {
	client *Client
}

// UtilizationBundleRequest configures generation of a census utilization bundle.
type UtilizationBundleRequest struct {
	Message    string
	TodayOnly  bool
	SendReport bool
}

func (c *Client) Census() *Census {
	return &Census{c}
}

// Utilization generates a census utilization bundle by calling the
// /v1/operator/utilization endpoint. The returned byte slice contains the bundle
// JSON payload suitable for saving to disk or further processing.
func (c *Census) Utilization(req *UtilizationBundleRequest, q *QueryOptions) ([]byte, *QueryMeta, error) {
	r := c.client.newRequest("GET", "/v1/operator/utilization")
	if q != nil {
		r.setQueryOptions(q)
	}
	if req != nil {
		if req.Message != "" {
			r.params.Set("message", req.Message)
		}
		if req.TodayOnly {
			r.params.Set("today_only", strconv.FormatBool(true))
		}
		if req.SendReport {
			r.params.Set("send_report", strconv.FormatBool(true))
		}
	}

	rtt, resp, err := c.client.doRequest(r)
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return body, qm, nil
}
