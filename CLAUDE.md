# Consul GraphFS LinkedDoc/RDF Integration

## Purpose
This repository is being annotated with GraphFS LinkedDoc+RDF comments to test and validate the linkedoc/rdf functionality from the [graphfs](../graphfs) repository. The goal is to transform the Consul codebase into a queryable knowledge graph using semantic metadata.

## GraphFS Overview
GraphFS is a semantic code filesystem toolkit that parses RDF/Turtle documentation embedded in source code and exposes it as a semantic graph queryable via SPARQL and GraphQL. It transforms implicit code relationships into explicit, navigable, and analyzable knowledge.

## LinkedDoc+RDF Comment Format

When adding GraphFS comments to code files, follow this format:

```go
/*
# Module: path/to/file.go
Brief description of the module's purpose.

## Linked Modules
- [module-name](relative/path/to/module.go) - Relationship description
- [another-module](path/to/another.go) - Relationship description

## Tags
tag1, tag2, tag3

## Exports
ExportedType, ExportedFunction, ExportedInterface

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<this> a code:Module ;
    code:name "path/to/file.go" ;
    code:description "Brief description of the module's purpose" ;
    code:layer "layer-name" ;
    code:linksTo [
        code:name "module-name" ;
        code:path "relative/path/to/module.go" ;
        code:relationship "Relationship description"
    ] ;
    code:exports :ExportedType, :ExportedFunction, :ExportedInterface ;
    code:tags "tag1", "tag2", "tag3" .
<!-- End LinkedDoc RDF -->
*/
package packagename
```

## Consul Architecture Layers

When annotating Consul code, use these architectural layers:

- **api** - Public API client library and interfaces
- **agent** - Agent runtime, HTTP endpoints, service/check registration
- **acl** - Access control and authorization
- **service-mesh** - Connect service mesh functionality, proxies, intentions
- **discovery** - Service discovery, health checking, catalog
- **consensus** - Raft consensus, leader election, state replication
- **config** - Configuration management, config entries
- **storage** - State store, snapshots, persistence
- **network** - Networking utilities, RPC, gossip (Serf)
- **security** - TLS, certificates, encryption
- **grpc** - gRPC services and middleware
- **cli** - Command-line interface
- **internal** - Internal utilities and helpers

## Common Tags

Use these tags to categorize modules:

- **security** - Security-related functionality (ACL, TLS, auth)
- **api** - API endpoints and handlers
- **storage** - Data storage and persistence
- **networking** - Network communication and protocols
- **service-mesh** - Service mesh capabilities
- **health** - Health checking
- **discovery** - Service discovery
- **consensus** - Distributed consensus
- **monitoring** - Metrics, logging, observability
- **configuration** - Configuration management
- **testing** - Test utilities and mocks

## Annotation Priorities

Focus on annotating these high-value areas first:

1. **Core agent modules** (agent/agent.go, agent/consul/*.go)
2. **ACL system** (acl/*.go)
3. **API layer** (api/*.go, agent/*_endpoint.go)
4. **Service mesh/Connect** (connect/*.go, agent/xds/*.go)
5. **State store** (agent/consul/state/*.go)
6. **Configuration** (agent/config/*.go)
7. **Health checking** (agent/checks/*.go)
8. **Internal services** (internal/*/*.go)

## Testing GraphFS Integration

After adding LinkedDoc comments, test them using:

```bash
# Initialize GraphFS (if not already done)
cd /Users/coolbeans/Development/dev/consul-doc
graphfs init

# Scan codebase and validate RDF metadata
graphfs scan --validate --stats

# Query example: Find all security-related modules
graphfs query 'SELECT ?module ?description WHERE {
  ?module <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <https://schema.codedoc.org/Module> .
  ?module <https://schema.codedoc.org/description> ?description .
  ?module <https://schema.codedoc.org/tags> "security" .
}'

# Export knowledge graph
graphfs scan --output consul-graph.json

# Generate visualizations
graphfs viz --format dot > consul-architecture.dot
```

## Guidelines

1. **Start small** - Begin with key modules before attempting comprehensive coverage
2. **Accuracy** - Ensure relationship descriptions are accurate and meaningful
3. **Consistency** - Use consistent terminology across annotations
4. **Validation** - Run `graphfs scan --validate` frequently to catch errors
5. **Documentation** - Annotations should enhance code understanding, not duplicate it
6. **Incrementalism** - Add annotations as you work on code, don't try to annotate everything at once

## Benefits for Consul Development

GraphFS annotations will enable:

- **Impact analysis** - Understand downstream effects before refactoring
- **Architecture validation** - Enforce layering rules and security boundaries
- **Code navigation** - Find dependencies and relationships quickly
- **Documentation** - Auto-generate architecture diagrams and module docs
- **AI assistance** - Provide rich context to AI coding assistants
- **Onboarding** - Help new developers understand the codebase structure

## Resources

- [GraphFS Repository](../graphfs)
- [GraphFS README](../graphfs/README.md)
- [LinkedDoc Format Specification](../graphfs/docs/linkedoc-format.md)
- [SPARQL Query Examples](../graphfs/examples/minimal-app/examples/query-examples.md)
