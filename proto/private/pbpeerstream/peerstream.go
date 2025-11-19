// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: proto/private/pbpeerstream/peerstream.go
Peerstream module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbpeerstream/peerstream.go> a code:Module ;
    code:name "proto/private/pbpeerstream/peerstream.go" ;
    code:description "Peerstream module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package pbpeerstream

func (x Operation) GoString() string {
	return x.String()
}
