// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/libs/utils/utils.go
Utils module for internal layer

## Tags
internal

## Exports
BoolToPointer, IntToPointer, JQFilter, RandName, StringToPointer

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/libs/utils/utils.go> a code:Module ;
    code:name "test/integration/consul-container/libs/utils/utils.go" ;
    code:description "Utils module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :BoolToPointer, :IntToPointer, :JQFilter, :RandName, :StringToPointer ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/itchyny/gojq"
	"github.com/teris-io/shortid"
)

func RandName(name string) string {
	shortID, err := shortid.New(1, shortid.DefaultABC, 6666)
	if err != nil {
		return ""
	}
	id, err := shortID.Generate()
	if err != nil {
		return ""
	}
	id = strings.ToLower(id)
	return name + "-" + id
}

// JQFilter uses the provided "jq" filter to parse json.
// Matching results are returned as a slice of strings.
func JQFilter(config, filter string) ([]string, error) {
	result := []string{}
	query, err := gojq.Parse(filter)
	if err != nil {
		return nil, err
	}

	var m interface{}
	err = json.Unmarshal([]byte(config), &m)
	if err != nil {
		return nil, err
	}

	iter := query.Run(m)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return nil, err
		}
		s := fmt.Sprintf("%v", v)
		result = append(result, s)
	}
	return result, nil
}

func IntToPointer(i int) *int {
	return &i
}

func BoolToPointer(b bool) *bool {
	return &b
}

func StringToPointer(s string) *string {
	return &s
}
