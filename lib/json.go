// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: lib/json.go
Json module for internal layer

## Tags
internal

## Exports
DecodeJSON, UnmarshalJSON

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<lib/json.go> a code:Module ;
    code:name "lib/json.go" ;
    code:description "Json module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :DecodeJSON, :UnmarshalJSON ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package lib

import (
	"bytes"
	"encoding/json"
	"io"
)

// DecodeJSON is a convenience function to create a JSON decoder
// set it up to disallow unknown fields and then decode into the
// given value
func DecodeJSON(data io.Reader, out interface{}) error {
	if data == nil {
		return io.EOF
	}

	decoder := json.NewDecoder(data)
	decoder.DisallowUnknownFields()
	return decoder.Decode(&out)
}

// UnmarshalJSON is a convenience function around calling
// DecodeJSON. It will mainly be useful in many of our
// UnmarshalJSON methods for structs.
func UnmarshalJSON(data []byte, out interface{}) error {
	return DecodeJSON(bytes.NewReader(data), out)
}
