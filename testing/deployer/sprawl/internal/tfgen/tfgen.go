// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: testing/deployer/sprawl/internal/tfgen/tfgen.go
Tfgen module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<testing/deployer/sprawl/internal/tfgen/tfgen.go> a code:Module ;
    code:name "testing/deployer/sprawl/internal/tfgen/tfgen.go" ;
    code:description "Tfgen module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package tfgen

import (
	"embed"
	"text/template"
)

//go:embed templates/container-app-dataplane.tf.tmpl
//go:embed templates/container-app-sidecar.tf.tmpl
//go:embed templates/container-app.tf.tmpl
//go:embed templates/container-consul.tf.tmpl
//go:embed templates/container-mgw.tf.tmpl
//go:embed templates/container-mgw-dataplane.tf.tmpl
//go:embed templates/container-pause.tf.tmpl
//go:embed templates/container-proxy.tf.tmpl
//go:embed templates/container-coredns.tf.tmpl
var content embed.FS

var (
	tfAppDataplaneT         = template.Must(template.ParseFS(content, "templates/container-app-dataplane.tf.tmpl"))
	tfAppSidecarT           = template.Must(template.ParseFS(content, "templates/container-app-sidecar.tf.tmpl"))
	tfAppT                  = template.Must(template.ParseFS(content, "templates/container-app.tf.tmpl"))
	tfConsulT               = template.Must(template.ParseFS(content, "templates/container-consul.tf.tmpl"))
	tfMeshGatewayT          = template.Must(template.ParseFS(content, "templates/container-mgw.tf.tmpl"))
	tfMeshGatewayDataplaneT = template.Must(template.ParseFS(content, "templates/container-mgw-dataplane.tf.tmpl"))
	tfPauseT                = template.Must(template.ParseFS(content, "templates/container-pause.tf.tmpl"))
	tfForwardProxyT         = template.Must(template.ParseFS(content, "templates/container-proxy.tf.tmpl"))
	tfCorednsT              = template.Must(template.ParseFS(content, "templates/container-coredns.tf.tmpl"))
)
