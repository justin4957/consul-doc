// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/leafcert/signer_netrpc.go
Signer Netrpc module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, communication, networking

## Exports
NetRPC, NewNetRPCCertSigner

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/leafcert/signer_netrpc.go> a code:Module ;
    code:name "agent/leafcert/signer_netrpc.go" ;
    code:description "Signer Netrpc module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :NetRPC, :NewNetRPCCertSigner ;
    code:tags "agent", "communication", "networking" .
<!-- End LinkedDoc RDF -->
*/
package leafcert

import (
	"context"

	"github.com/hashicorp/consul/agent/structs"
)

// NetRPC is an interface that an NetRPC client must implement. This is a helper
// interface that is implemented by the agent delegate so that Type
// implementations can request NetRPC access.
type NetRPC interface {
	RPC(ctx context.Context, method string, args any, reply any) error
}

// NewNetRPCCertSigner returns a CertSigner that uses net-rpc to sign certs.
func NewNetRPCCertSigner(netRPC NetRPC) CertSigner {
	return &netRPCCertSigner{netRPC: netRPC}
}

type netRPCCertSigner struct {
	// NetRPC is an RPC client for remote cert signing requests.
	netRPC NetRPC
}

var _ CertSigner = (*netRPCCertSigner)(nil)

func (s *netRPCCertSigner) SignCert(ctx context.Context, args *structs.CASignRequest) (*structs.IssuedCert, error) {
	var reply structs.IssuedCert
	err := s.netRPC.RPC(ctx, "ConnectCA.Sign", args, &reply)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
