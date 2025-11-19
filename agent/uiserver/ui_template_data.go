// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/uiserver/ui_template_data.go
Ui Template Data module for agent layer

## Linked Modules
- [agent/config](../agent/config)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/uiserver/ui_template_data.go> a code:Module ;
    code:name "agent/uiserver/ui_template_data.go" ;
    code:description "Ui Template Data module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/config" ;
        code:path "../agent/config"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package uiserver

import (
	"encoding/json"

	"github.com/hashicorp/consul/agent/config"
)

// uiTemplateDataFromConfig returns the base set of variables that should be
// injected into the UI's Env based on the given runtime UI config.
func uiTemplateDataFromConfig(cfg *config.RuntimeConfig) (map[string]interface{}, error) {

	uiCfg := map[string]interface{}{
		"metrics_provider": cfg.UIConfig.MetricsProvider,
		// We explicitly MUST NOT pass the metrics_proxy object since it might
		// contain add_headers with secrets that the UI shouldn't know e.g. API
		// tokens for the backend. The provider should either require the proxy to
		// be configured and then use that or hit the backend directly from the
		// browser.
		"metrics_proxy_enabled":   cfg.UIConfig.MetricsProxy.BaseURL != "",
		"dashboard_url_templates": cfg.UIConfig.DashboardURLTemplates,
		"hcp_enabled":             cfg.UIConfig.HCPEnabled,
	}

	// Only set this if there is some actual JSON or we'll cause a JSON
	// marshalling error later during serving which ends up being silent.
	if cfg.UIConfig.MetricsProviderOptionsJSON != "" {
		uiCfg["metrics_provider_options"] = json.RawMessage(cfg.UIConfig.MetricsProviderOptionsJSON)
	}

	d := map[string]interface{}{
		"ContentPath":       cfg.UIConfig.ContentPath,
		"ACLsEnabled":       cfg.ACLsEnabled,
		"HCPEnabled":        cfg.UIConfig.HCPEnabled,
		"UIConfig":          uiCfg,
		"LocalDatacenter":   cfg.Datacenter,
		"PrimaryDatacenter": cfg.PrimaryDatacenter,
		"PeeringEnabled":    cfg.PeeringEnabled,
	}

	// Also inject additional provider scripts if needed, otherwise strip the
	// comment.
	if len(cfg.UIConfig.MetricsProviderFiles) > 0 {
		d["ExtraScripts"] = []string{
			cfg.UIConfig.ContentPath + compiledProviderJSPath,
		}
	}

	return d, nil
}
