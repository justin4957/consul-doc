// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/leader_log_verification.go
Leader Log Verification module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/leader_log_verification.go> a code:Module ;
    code:name "agent/consul/leader_log_verification.go" ;
    code:description "Leader Log Verification module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"context"
	"time"

	"github.com/hashicorp/consul/agent/structs"
)

func (s *Server) startLogVerification(ctx context.Context) error {
	return s.leaderRoutineManager.Start(ctx, raftLogVerifierRoutineName, s.runLogVerification)
}

func (s *Server) stopLogVerification() {
	s.leaderRoutineManager.Stop(raftLogVerifierRoutineName)
}

func (s *Server) runLogVerification(ctx context.Context) error {
	// This shouldn't be possible but bit of a safety check
	if !s.config.LogStoreConfig.Verification.Enabled ||
		s.config.LogStoreConfig.Verification.Interval == 0 {
		return nil
	}
	ticker := time.NewTicker(s.config.LogStoreConfig.Verification.Interval)
	defer ticker.Stop()

	logger := s.logger.Named("raft.logstore.verifier")
	for {
		select {
		case <-ticker.C:
			// Attempt to send a checkpoint message
			typ := structs.RaftLogVerifierCheckpoint | structs.IgnoreUnknownTypeFlag
			raw, err := s.raftApplyMsgpack(typ, nil)
			if err != nil {
				logger.Error("sending verification checkpoint failed", "err", err)
			} else {
				index, ok := raw.(uint64)
				if !ok {
					index = 0
				}
				logger.Debug("sent verification checkpoint", "index", int64(index))
			}

		case <-ctx.Done():
			return nil
		}
	}
}
