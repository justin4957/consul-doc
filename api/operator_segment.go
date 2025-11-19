// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: api/operator_segment.go
Operator Segment module for api layer

## Tags
api, client

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<api/operator_segment.go> a code:Module ;
    code:name "api/operator_segment.go" ;
    code:description "Operator Segment module for api layer" ;
    code:language "go" ;
    code:layer "api" ;
    code:tags "api", "client" .
<!-- End LinkedDoc RDF -->
*/
package api

// SegmentList returns all the available LAN segments.
func (op *Operator) SegmentList(q *QueryOptions) ([]string, *QueryMeta, error) {
	var out []string
	qm, err := op.c.query("/v1/operator/segment", &out, q)
	if err != nil {
		return nil, nil, err
	}
	return out, qm, nil
}
