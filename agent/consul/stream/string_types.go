// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/stream/string_types.go
String Types module for agent layer

## Tags
agent, data-model, types

## Exports
StringSubject, StringTopic

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/stream/string_types.go> a code:Module ;
    code:name "agent/consul/stream/string_types.go" ;
    code:description "String Types module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :StringSubject, :StringTopic ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package stream

// StringSubject can be used as a Subject for Events sent to the EventPublisher
type StringSubject string

func (s StringSubject) String() string { return string(s) }

// StringTopic can be used as a Topic for Events sent to the EventPublisher
type StringTopic string

func (s StringTopic) String() string { return string(s) }
