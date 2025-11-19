// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/logging.go
Logging module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/logging.go> a code:Module ;
    code:name "agent/consul/logging.go" ;
    code:description "Logging module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"sync"

	"github.com/hashicorp/go-hclog"
)

type loggerStore struct {
	root  hclog.Logger
	l     sync.Mutex
	cache map[string]hclog.Logger
}

func newLoggerStore(root hclog.Logger) *loggerStore {
	return &loggerStore{
		root:  root,
		cache: make(map[string]hclog.Logger),
	}
}

func (ls *loggerStore) Named(name string) hclog.Logger {
	ls.l.Lock()
	defer ls.l.Unlock()
	l, ok := ls.cache[name]
	if !ok {
		l = ls.root.Named(name)
		ls.cache[name] = l
	}
	return l
}
