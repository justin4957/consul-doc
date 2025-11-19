## Consul Architecture Overview

```mermaid
flowchart LR
    agent_cache_types_service_gateways_go["agent/cache-types/service_gateways.go"
    agent_proxycfg_internal_watch_watchmap_go["agent/proxycfg/internal/watch/watchmap.go"
    agent_structs_config_entry_jwt_provider_go["agent/structs/config_entry_jwt_provider.go"
    agent_proxycfg_glue_intentions_go["agent/proxycfg-glue/intentions.go"
    agent_proxycfg_sources_catalog_mock_Watcher_go["agent/proxycfg-sources/catalog/mock_Watcher.go"
    agent_structs_acl_cache_go["agent/structs/acl_cache.go"
    agent_grpc_middleware_stats_go["agent/grpc-middleware/stats.go"
    agent_consul_rate_handler_go["agent/consul/rate/handler.go"
    agent_structs_config_entry_sameness_group_go["agent/structs/config_entry_sameness_group.go"
    api_status_go["api/status.go"
    agent_xds_z_xds_packages_go["agent/xds/z_xds_packages.go"
    acl_policy_ce_go["acl/policy_ce.go"
    agent_cache_entry_go["agent/cache/entry.go"
    agent_consul_discoverychain_testing_go["agent/consul/discoverychain/testing.go"
    agent_envoyextensions_builtin_otel_access_logging_structs_go["agent/envoyextensions/builtin/otel-access-logging/structs.go"
    agent_kvs_endpoint_go["agent/kvs_endpoint.go"
    agent_signal_unix_go["agent/signal_unix.go"
    agent_connect_generate_go["agent/connect/generate.go"
    agent_grpc_internal_listener_go["agent/grpc-internal/listener.go"
    agent_consul_raft_handle_go["agent/consul/raft_handle.go"
    agent_consul_state_intention_go["agent/consul/state/intention.go"
    agent_consul_state_usage_ce_go["agent/consul/state/usage_ce.go"
    agent_mock_notify_go["agent/mock/notify.go"
    agent_proxycfg_glue_leafcerts_go["agent/proxycfg-glue/leafcerts.go"
    agent_structs_check_definition_go["agent/structs/check_definition.go"
    agent_consul_gateways_controller_gateways_go["agent/consul/gateways/controller_gateways.go"
    agent_grpc_internal_services_subscribe_logger_go["agent/grpc-internal/services/subscribe/logger.go"
    agent_proxycfg_glue_gateway_services_go["agent/proxycfg-glue/gateway_services.go"
    agent_dns_context_go["agent/dns/context.go"
    agent_proxycfg_state_go["agent/proxycfg/state.go"
    agent_xds_delta_go["agent/xds/delta.go"
    agent_cache_watch_go["agent/cache/watch.go"
    agent_consul_discovery_chain_endpoint_go["agent/consul/discovery_chain_endpoint.go"
    agent_consul_intention_endpoint_go["agent/consul/intention_endpoint.go"
    agent_structs_config_entry_gateways_go["agent/structs/config_entry_gateways.go"
    agent_xds_jwt_authn_go["agent/xds/jwt_authn.go"
    agent_xds_platform_net_fallback_go["agent/xds/platform/net_fallback.go"
    agent_cache_types_gateway_services_go["agent/cache-types/gateway_services.go"
    agent_consul_reporting_reporting_go["agent/consul/reporting/reporting.go"
    agent_consul_stream_event_snapshot_go["agent/consul/stream/event_snapshot.go"
    agent_consul_wanfed_pool_go["agent/consul/wanfed/pool.go"
    agent_proxycfg_glue_trust_bundle_go["agent/proxycfg-glue/trust_bundle.go"
    agent_hcp_deps_go["agent/hcp/deps.go"
    agent_hcp_telemetry_doc_go["agent/hcp/telemetry/doc.go"
    agent_proxycfg_glue_exported_peered_services_go["agent/proxycfg-glue/exported_peered_services.go"
    agent_consul_authmethod_testauth_testing_ce_go["agent/consul/authmethod/testauth/testing_ce.go"
    agent_hcp_bootstrap_constants_constants_go["agent/hcp/bootstrap/constants/constants.go"
    agent_proxycfg_mesh_gateway_ce_go["agent/proxycfg/mesh_gateway_ce.go"
    agent_cache_types_intention_upstreams_go["agent/cache-types/intention_upstreams.go"
    agent_consul_state_coordinate_go["agent/consul/state/coordinate.go"
    agent_proxycfg_glue_intentions_ce_go["agent/proxycfg-glue/intentions_ce.go"
    agent_connect_ca_provider_aws_go["agent/connect/ca/provider_aws.go"
    agent_connect_csr_go["agent/connect/csr.go"
    agent_consul_state_txn_go["agent/consul/state/txn.go"
    agent_hcp_telemetry_gauge_store_go["agent/hcp/telemetry/gauge_store.go"
    agent_xds_locality_policy_go["agent/xds/locality_policy.go"
    api_operator_usage_go["api/operator_usage.go"
    agent_consul_operator_usage_endpoint_go["agent/consul/operator_usage_endpoint.go"
    agent_structs_txn_go["agent/structs/txn.go"
    agent_structs_snapshot_go["agent/structs/snapshot.go"
    api_config_entry_go["api/config_entry.go"
    agent_hcp_telemetry_provider_go["agent/hcp/telemetry_provider.go"
    agent_structs_config_entry_ce_go["agent/structs/config_entry_ce.go"
    agent_blockingquery_mock_RequestOptions_go["agent/blockingquery/mock_RequestOptions.go"
    agent_grpc_external_services_peerstream_health_snapshot_go["agent/grpc-external/services/peerstream/health_snapshot.go"
    agent_structs_discovery_chain_ce_go["agent/structs/discovery_chain_ce.go"
    agent_connect_uri_mesh_gateway_ce_go["agent/connect/uri_mesh_gateway_ce.go"
    agent_metrics_testing_go["agent/metrics/testing.go"
    agent_proxycfg_glue_service_http_checks_go["agent/proxycfg-glue/service_http_checks.go"
    agent_ui_endpoint_go["agent/ui_endpoint.go"
    agent_xds_config_config_go["agent/xds/config/config.go"
    agent_proxycfg_naming_ce_go["agent/proxycfg/naming_ce.go"
    agent_watch_handler_go["agent/watch_handler.go"
    agent_config_endpoint_go["agent/config_endpoint.go"
    agent_connect_parsing_go["agent/connect/parsing.go"
    agent_consul_session_endpoint_go["agent/consul/session_endpoint.go"
    agent_cache_types_peerings_go["agent/cache-types/peerings.go"
    agent_consul_state_config_entry_sameness_group_go["agent/consul/state/config_entry_sameness_group.go"
    agent_discovery_chain_endpoint_go["agent/discovery_chain_endpoint.go"
    agent_local_state_go["agent/local/state.go"
    agent_remote_exec_go["agent/remote_exec.go"
    agent_xds_endpoints_go["agent/xds/endpoints.go"
    agent_consul_leader_connect_ca_go["agent/consul/leader_connect_ca.go"
    agent_consul_raft_rpc_go["agent/consul/raft_rpc.go"
    api_raw_go["api/raw.go"
    agent_proxycfg_testing_ce_go["agent/proxycfg/testing_ce.go"
    agent_structs_aclfilter_filter_go["agent/structs/aclfilter/filter.go"
    agent_xds_listeners_apigateway_go["agent/xds/listeners_apigateway.go"
    agent_xds_server_ce_go["agent/xds/server_ce.go"
    acl_testing_go["acl/testing.go"
    acl_MockAuthorizer_go["acl/MockAuthorizer.go"
    agent_proxycfg_sources_catalog_config_source_go["agent/proxycfg-sources/catalog/config_source.go"
    agent_consul_fsm_commands_ce_go["agent/consul/fsm/commands_ce.go"
    api_config_entry_status_go["api/config_entry_status.go"
    agent_connect_uri_agent_ce_go["agent/connect/uri_agent_ce.go"
    agent_consul_enterprise_config_ce_go["agent/consul/enterprise_config_ce.go"
    agent_blockingquery_mock_FSMServer_go["agent/blockingquery/mock_FSMServer.go"
    agent_structs_check_type_go["agent/structs/check_type.go"
    agent_uiserver_ui_template_data_go["agent/uiserver/ui_template_data.go"
    agent_xds_response_response_go["agent/xds/response/response.go"
    api_watch_funcs_go["api/watch/funcs.go"
    acl_chained_authorizer_go["acl/chained_authorizer.go"
    agent_consul_autopilot_ce_go["agent/consul/autopilot_ce.go"
    agent_consul_state_config_entry_intention_go["agent/consul/state/config_entry_intention.go"
    agent_operator_endpoint_ce_go["agent/operator_endpoint_ce.go"
    agent_proxycfg_testing_ingress_gateway_go["agent/proxycfg/testing_ingress_gateway.go"
    agent_structs_connect_go["agent/structs/connect.go"
    agent_user_event_go["agent/user_event.go"
    acl_enterprisemeta_ce_go["acl/enterprisemeta_ce.go"
    agent_config_flags_go["agent/config/flags.go"
    agent_consul_filter_go["agent/consul/filter.go"
    agent_structs_structs_deepcopy_go["agent/structs/structs.deepcopy.go"
    agent_connect_uri_service_ce_go["agent/connect/uri_service_ce.go"
    agent_consul_auth_mock_ACLCache_go["agent/consul/auth/mock_ACLCache.go"
    connect_proxy_testing_go["connect/proxy/testing.go"
    agent_connect_testing_ca_go["agent/connect/testing_ca.go"
    agent_connect_uri_mesh_gateway_go["agent/connect/uri_mesh_gateway.go"
    agent_consul_state_prepared_query_index_go["agent/consul/state/prepared_query_index.go"
    agent_dns_buffer_response_writer_go["agent/dns/buffer_response_writer.go"
    agent_structs_connect_proxy_config_ce_go["agent/structs/connect_proxy_config_ce.go"
    agent_connect_testing_spiffe_go["agent/connect/testing_spiffe.go"
    agent_hcp_scada_mock_Provider_go["agent/hcp/scada/mock_Provider.go"
    agent_structs_config_entry_file_system_certificate_go["agent/structs/config_entry_file_system_certificate.go"
    agent_utilization_endpoint_go["agent/utilization_endpoint.go"
    agent_consul_session_timers_go["agent/consul/session_timers.go"
    agent_connect_ca_endpoint_go["agent/connect_ca_endpoint.go"
    agent_consul_server_serf_go["agent/consul/server_serf.go"
    agent_event_endpoint_go["agent/event_endpoint.go"
    agent_proxycfg_glue_discovery_chain_go["agent/proxycfg-glue/discovery_chain.go"
    agent_consul_state_connect_ca_events_go["agent/consul/state/connect_ca_events.go"
    agent_consul_state_catalog_go["agent/consul/state/catalog.go"
    agent_structs_config_entry_mesh_go["agent/structs/config_entry_mesh.go"
    connect_certgen_certgen_go["connect/certgen/certgen.go"
    agent_consul_auto_config_backend_go["agent/consul/auto_config_backend.go"
    agent_structs_acl_templated_policy_ce_go["agent/structs/acl_templated_policy_ce.go"
    agent_consul_acl_authmethod_go["agent/consul/acl_authmethod.go"
    agent_structs_testing_catalog_go["agent/structs/testing_catalog.go"
    api_config_entry_rate_limit_ip_go["api/config_entry_rate_limit_ip.go"
    agent_consul_stream_string_types_go["agent/consul/stream/string_types.go"
    agent_consul_stream_subscription_go["agent/consul/stream/subscription.go"
    agent_consul_server_log_verification_go["agent/consul/server_log_verification.go"
    agent_envoyextensions_builtin_property_override_property_override_go["agent/envoyextensions/builtin/property-override/property_override.go"
    agent_proxycfg_sources_catalog_mock_ConfigManager_go["agent/proxycfg-sources/catalog/mock_ConfigManager.go"
    agent_auto_config_config_go["agent/auto-config/config.go"
    agent_connect_x509_patch_go["agent/connect/x509_patch.go"
    agent_consul_xdscapacity_capacity_go["agent/consul/xdscapacity/capacity.go"
    agent_proxycfg_glue_health_go["agent/proxycfg-glue/health.go"
    agent_hcp_client_errors_go["agent/hcp/client/errors.go"
    agent_sidecar_service_go["agent/sidecar_service.go"
    agent_consul_usagemetrics_usagemetrics_go["agent/consul/usagemetrics/usagemetrics.go"
    agent_consul_leader_intentions_ce_go["agent/consul/leader_intentions_ce.go"
    agent_leafcert_roots_go["agent/leafcert/roots.go"
    agent_pool_peeked_conn_go["agent/pool/peeked_conn.go"
    agent_structs_peering_go["agent/structs/peering.go"
    agent_auto_config_run_go["agent/auto-config/run.go"
    agent_consul_rtt_go["agent/consul/rtt.go"
    agent_denylist_go["agent/denylist.go"
    api_config_entry_intentions_go["api/config_entry_intentions.go"
    agent_leafcert_cached_roots_go["agent/leafcert/cached_roots.go"
    api_lock_go["api/lock.go"
    agent_grpc_external_services_acl_login_go["agent/grpc-external/services/acl/login.go"
    agent_consul_options_ce_go["agent/consul/options_ce.go"
    agent_hcp_bootstrap_config_loader_loader_go["agent/hcp/bootstrap/config-loader/loader.go"
    agent_operator_endpoint_go["agent/operator_endpoint.go"
    connect_proxy_config_go["connect/proxy/config.go"
    agent_connect_ca_mock_Provider_go["agent/connect/ca/mock_Provider.go"
    agent_consul_client_go["agent/consul/client.go"
    agent_consul_state_session_go["agent/consul/state/session.go"
    connect_testing_go["connect/testing.go"
    agent_grpc_external_services_acl_mock_Login_go["agent/grpc-external/services/acl/mock_Login.go"
    agent_connect_ca_provider_vault_auth_azure_go["agent/connect/ca/provider_vault_auth_azure.go"
    agent_consul_rpc_go["agent/consul/rpc.go"
    agent_grpc_external_services_resource_delete_go["agent/grpc-external/services/resource/delete.go"
    agent_http_register_go["agent/http_register.go"
    agent_grpc_external_services_serverdiscovery_server_go["agent/grpc-external/services/serverdiscovery/server.go"
    agent_grpc_external_services_peerstream_subscription_view_go["agent/grpc-external/services/peerstream/subscription_view.go"
    api_api_go["api/api.go"
    agent_consul_state_coordinate_ce_go["agent/consul/state/coordinate_ce.go"
    agent_consul_fsm_log_verification_chunking_shim_go["agent/consul/fsm/log_verification_chunking_shim.go"
    api_agent_go["api/agent.go"
    agent_envoyextensions_builtin_ext_authz_ext_authz_go["agent/envoyextensions/builtin/ext-authz/ext_authz.go"
    agent_consul_peering_backend_ce_go["agent/consul/peering_backend_ce.go"
    agent_grpc_external_services_serverdiscovery_mock_ACLResolver_go["agent/grpc-external/services/serverdiscovery/mock_ACLResolver.go"
    agent_proxycfg_sources_catalog_mock_SessionLimiter_go["agent/proxycfg-sources/catalog/mock_SessionLimiter.go"
    api_operator_area_go["api/operator_area.go"
    agent_intentions_endpoint_go["agent/intentions_endpoint.go"
    acl_static_authorizer_go["acl/static_authorizer.go"
    agent_cache_types_intention_match_go["agent/cache-types/intention_match.go"
    agent_checks_docker_unix_go["agent/checks/docker_unix.go"
    agent_consul_rate_metrics_go["agent/consul/rate/metrics.go"
    agent_proxycfg_snapshot_go["agent/proxycfg/snapshot.go"
    api_watch_plan_go["api/watch/plan.go"
    acl_resolver_result_go["acl/resolver/result.go"
    agent_grpc_middleware_rate_go["agent/grpc-middleware/rate.go"
    api_coordinate_go["api/coordinate.go"
    agent_config_config_deepcopy_go["agent/config/config.deepcopy.go"
    agent_consul_state_acl_ce_go["agent/consul/state/acl_ce.go"
    agent_grpc_external_testutils_acl_go["agent/grpc-external/testutils/acl.go"
    agent_proxycfg_proxycfg_deepcopy_go["agent/proxycfg/proxycfg.deepcopy.go"
    agent_proxycfg_testing_upstreams_ce_go["agent/proxycfg/testing_upstreams_ce.go"
    api_operator_raft_go["api/operator_raft.go"
    agent_checks_docker_go["agent/checks/docker.go"
    agent_consul_tenancy_bridge_ce_go["agent/consul/tenancy_bridge_ce.go"
    agent_rpcclient_health_health_go["agent/rpcclient/health/health.go"
    agent_consul_prepared_query_template_go["agent/consul/prepared_query/template.go"
    agent_grpc_external_services_peerstream_replication_go["agent/grpc-external/services/peerstream/replication.go"
    agent_grpc_internal_tracker_go["agent/grpc-internal/tracker.go"
    api_config_entry_jwt_provider_go["api/config_entry_jwt_provider.go"
    agent_envoyextensions_builtin_ext_authz_structs_go["agent/envoyextensions/builtin/ext-authz/structs.go"
    agent_pool_pool_go["agent/pool/pool.go"
    agent_consul_state_acl_events_go["agent/consul/state/acl_events.go"
    agent_hcp_client_telemetry_config_go["agent/hcp/client/telemetry_config.go"
    agent_proxycfg_config_snapshot_glue_go["agent/proxycfg/config_snapshot_glue.go"
    agent_consul_reporting_reportingmock_mock_StateDelegate_go["agent/consul/reporting/reportingmock/mock_StateDelegate.go"
    agent_http_ce_go["agent/http_ce.go"
    agent_metadata_server_go["agent/metadata/server.go"
    agent_structs_connect_ca_go["agent/structs/connect_ca.go"
    agent_consul_auto_encrypt_endpoint_go["agent/consul/auto_encrypt_endpoint.go"
    agent_consul_serf_filter_go["agent/consul/serf_filter.go"
    agent_exec_exec_windows_go["agent/exec/exec_windows.go"
    agent_rpc_middleware_rate_limit_mappings_go["agent/rpc/middleware/rate_limit_mappings.go"
    agent_systemd_notify_go["agent/systemd/notify.go"
    acl_policy_authorizer_ce_go["acl/policy_authorizer_ce.go"
    agent_consul_kvs_endpoint_go["agent/consul/kvs_endpoint.go"
    agent_config_file_watcher_go["agent/config/file_watcher.go"
    agent_cacheshim_cache_go["agent/cacheshim/cache.go"
    agent_consul_health_endpoint_ce_go["agent/consul/health_endpoint_ce.go"
    agent_consul_leader_log_verification_go["agent/consul/leader_log_verification.go"
    agent_structs_connect_ce_go["agent/structs/connect_ce.go"
    api_prepared_query_go["api/prepared_query.go"
    agent_consul_discoverychain_gateway_go["agent/consul/discoverychain/gateway.go"
    agent_prepared_query_endpoint_go["agent/prepared_query_endpoint.go"
    api_health_go["api/health.go"
    agent_config_limits_windows_go["agent/config/limits_windows.go"
    agent_agent_go["agent/agent.go"
    agent_ae_trigger_go["agent/ae/trigger.go"
    agent_connect_ca_common_go["agent/connect/ca/common.go"
    agent_consul_acl_endpoint_ce_go["agent/consul/acl_endpoint_ce.go"
    agent_consul_state_intention_ce_go["agent/consul/state/intention_ce.go"
    agent_consul_state_kvs_ce_go["agent/consul/state/kvs_ce.go"
    agent_structs_acl_ce_go["agent/structs/acl_ce.go"
    agent_auto_config_config_ce_go["agent/auto-config/config_ce.go"
    agent_consul_config_go["agent/consul/config.go"
    agent_consul_controller_queue_rate_go["agent/consul/controller/queue/rate.go"
    agent_envoyextensions_builtin_otel_access_logging_otel_access_logging_go["agent/envoyextensions/builtin/otel-access-logging/otel_access_logging.go"
    agent_consul_leader_connect_go["agent/consul/leader_connect.go"
    agent_dns_go["agent/dns.go"
    agent_submatview_handler_go["agent/submatview/handler.go"
    agent_structs_auto_encrypt_go["agent/structs/auto_encrypt.go"
    agent_consul_authmethod_ssoauth_sso_ce_go["agent/consul/authmethod/ssoauth/sso_ce.go"
    agent_consul_merge_go["agent/consul/merge.go"
    agent_connect_ca_provider_vault_auth_aws_go["agent/connect/ca/provider_vault_auth_aws.go"
    agent_consul_acl_replication_go["agent/consul/acl_replication.go"
    agent_connect_uri_service_go["agent/connect/uri_service.go"
    agent_proxycfg_testing_peering_go["agent/proxycfg/testing_peering.go"
    api_watch_watch_go["api/watch/watch.go"
    agent_consul_internal_endpoint_go["agent/consul/internal_endpoint.go"
    agent_grpc_external_testutils_mock_server_transport_stream_go["agent/grpc-external/testutils/mock_server_transport_stream.go"
    agent_hcp_config_config_go["agent/hcp/config/config.go"
    agent_structs_testing_connect_proxy_config_go["agent/structs/testing_connect_proxy_config.go"
    agent_grpc_external_services_dataplane_get_supported_features_go["agent/grpc-external/services/dataplane/get_supported_features.go"
    agent_local_testing_go["agent/local/testing.go"
    agent_consul_leader_metrics_go["agent/consul/leader_metrics.go"
    agent_leafcert_generate_go["agent/leafcert/generate.go"
    agent_leafcert_cert_go["agent/leafcert/cert.go"
    agent_structs_config_entry_discoverychain_go["agent/structs/config_entry_discoverychain.go"
    agent_config_default_ce_go["agent/config/default_ce.go"
    agent_configentry_doc_go["agent/configentry/doc.go"
    agent_consul_state_census_utilization_go["agent/consul/state/census_utilization.go"
    agent_notify_go["agent/notify.go"
    agent_xds_testcommon_testcommon_go["agent/xds/testcommon/testcommon.go"
    agent_acl_endpoint_go["agent/acl_endpoint.go"
    agent_cache_types_connect_ca_root_go["agent/cache-types/connect_ca_root.go"
    agent_grpc_external_services_acl_mock_TokenWriter_go["agent/grpc-external/services/acl/mock_TokenWriter.go"
    agent_consul_state_indexer_go["agent/consul/state/indexer.go"
    agent_xds_rbac_go["agent/xds/rbac.go"
    agent_xds_jwt_authn_ce_go["agent/xds/jwt_authn_ce.go"
    agent_consul_autopilotevents_ready_servers_events_go["agent/consul/autopilotevents/ready_servers_events.go"
    agent_structs_acl_go["agent/structs/acl.go"
    agent_structs_config_entry_inline_certificate_go["agent/structs/config_entry_inline_certificate.go"
    acl_resolver_danger_go["acl/resolver/danger.go"
    agent_consul_leader_registrator_v1_go["agent/consul/leader_registrator_v1.go"
    agent_consul_server_ce_go["agent/consul/server_ce.go"
    agent_consul_state_memdb_go["agent/consul/state/memdb.go"
    agent_structs_config_entry_intentions_ce_go["agent/structs/config_entry_intentions_ce.go"
    agent_structs_identity_go["agent/structs/identity.go"
    connect_proxy_listener_go["connect/proxy/listener.go"
    agent_cache_types_prepared_query_go["agent/cache-types/prepared_query.go"
    agent_consul_autopilot_go["agent/consul/autopilot.go"
    agent_grpc_external_services_resource_write_status_go["agent/grpc-external/services/resource/write_status.go"
    test_integration_connect_envoy_test_sds_server_sds_go["test/integration/connect/envoy/test-sds-server/sds.go"
    agent_consul_prepared_query_endpoint_ce_go["agent/consul/prepared_query_endpoint_ce.go"
    agent_proxycfg_naming_go["agent/proxycfg/naming.go"
    agent_proxycfg_testing_connect_proxy_go["agent/proxycfg/testing_connect_proxy.go"
    agent_structs_errors_go["agent/structs/errors.go"
    agent_exec_exec_unix_go["agent/exec/exec_unix.go"
    agent_leafcert_watch_go["agent/leafcert/watch.go"
    agent_proxycfg_terminating_gateway_go["agent/proxycfg/terminating_gateway.go"
    agent_submatview_local_materializer_go["agent/submatview/local_materializer.go"
    agent_consul_configentry_backend_go["agent/consul/configentry_backend.go"
    agent_consul_stream_event_publisher_go["agent/consul/stream/event_publisher.go"
    agent_grpc_external_services_resource_write_go["agent/grpc-external/services/resource/write.go"
    agent_hcp_telemetry_otlp_transform_go["agent/hcp/telemetry/otlp_transform.go"
    agent_log_drop_mock_Logger_go["agent/log-drop/mock_Logger.go"
    agent_proxycfg_testing_api_gateway_go["agent/proxycfg/testing_api_gateway.go"
    agent_proxycfg_testing_tproxy_go["agent/proxycfg/testing_tproxy.go"
    agent_connect_ca_provider_consul_config_go["agent/connect/ca/provider_consul_config.go"
    agent_consul_context_go["agent/consul/context.go"
    agent_consul_state_config_entry_events_go["agent/consul/state/config_entry_events.go"
    agent_grpc_internal_pipe_go["agent/grpc-internal/pipe.go"
    agent_proxycfg_sources_local_mock_ConfigManager_go["agent/proxycfg-sources/local/mock_ConfigManager.go"
    agent_consul_fsm_snapshot_go["agent/consul/fsm/snapshot.go"
    agent_xds_failover_policy_ce_go["agent/xds/failover_policy_ce.go"
    agent_consul_servercert_manager_go["agent/consul/servercert/manager.go"
    agent_grpc_external_services_resource_mutate_and_validate_go["agent/grpc-external/services/resource/mutate_and_validate.go"
    agent_testagent_go["agent/testagent.go"
    agent_consul_authmethod_kubeauth_testing_go["agent/consul/authmethod/kubeauth/testing.go"
    agent_consul_wanfed_wanfed_go["agent/consul/wanfed/wanfed.go"
    agent_hcp_scada_capabilities_go["agent/hcp/scada/capabilities.go"
    agent_hcp_manager_go["agent/hcp/manager.go"
    acl_errors_ce_go["acl/errors_ce.go"
    agent_config_deprecated_go["agent/config/deprecated.go"
    agent_consul_acl_go["agent/consul/acl.go"
    agent_rpc_middleware_recovery_go["agent/rpc/middleware/recovery.go"
    connect_proxy_proxy_go["connect/proxy/proxy.go"
    acl_policy_go["acl/policy.go"
    agent_consul_authmethod_kubeauth_k8s_ce_go["agent/consul/authmethod/kubeauth/k8s_ce.go"
    agent_consul_operator_endpoint_go["agent/consul/operator_endpoint.go"
    agent_blockingquery_blockingquery_go["agent/blockingquery/blockingquery.go"
    agent_consul_acl_authmethod_ce_go["agent/consul/acl_authmethod_ce.go"
    agent_peering_endpoint_go["agent/peering_endpoint.go"
    api_config_entry_file_system_certificate_go["api/config_entry_file_system_certificate.go"
    agent_consul_client_serf_go["agent/consul/client_serf.go"
    acl_policy_merger_ce_go["acl/policy_merger_ce.go"
    acl_errors_go["acl/errors.go"
    agent_config_runtime_go["agent/config/runtime.go"
    agent_hcp_client_metrics_client_go["agent/hcp/client/metrics_client.go"
    agent_setup_go["agent/setup.go"
    agent_grpc_internal_services_subscribe_subscribe_go["agent/grpc-internal/services/subscribe/subscribe.go"
    agent_submatview_mock_ACLResolver_go["agent/submatview/mock_ACLResolver.go"
    agent_structs_config_entry_intentions_go["agent/structs/config_entry_intentions.go"
    agent_structs_operator_go["agent/structs/operator.go"
    api_peering_go["api/peering.go"
    agent_consul_state_config_entry_exported_services_go["agent/consul/state/config_entry_exported_services.go"
    agent_grpc_internal_resolver_registry_go["agent/grpc-internal/resolver/registry.go"
    agent_hcp_testserver_main_go["agent/hcp/testserver/main.go"
    agent_configentry_discoverychain_go["agent/configentry/discoverychain.go"
    agent_xds_listeners_go["agent/xds/listeners.go"
    agent_xds_listeners_ingress_go["agent/xds/listeners_ingress.go"
    api_connect_go["api/connect.go"
    agent_consul_type_registry_go["agent/consul/type_registry.go"
    agent_uiserver_buffered_file_go["agent/uiserver/buffered_file.go"
    agent_consul_gateways_controller_gateways_ce_go["agent/consul/gateways/controller_gateways_ce.go"
    agent_proxycfg_glue_internal_service_dump_go["agent/proxycfg-glue/internal_service_dump.go"
    agent_consul_acl_replication_types_go["agent/consul/acl_replication_types.go"
    agent_snapshot_endpoint_go["agent/snapshot_endpoint.go"
    agent_cache_types_discovery_chain_go["agent/cache-types/discovery_chain.go"
    agent_grpc_external_services_resource_mock_ACLResolver_go["agent/grpc-external/services/resource/mock_ACLResolver.go"
    agent_connect_ca_provider_vault_auth_go["agent/connect/ca/provider_vault_auth.go"
    agent_consul_fsm_snapshot_ce_go["agent/consul/fsm/snapshot_ce.go"
    agent_xds_failover_policy_go["agent/xds/failover_policy.go"
    api_internal_go["api/internal.go"
    agent_connect_ca_provider_vault_auth_k8s_go["agent/connect/ca/provider_vault_auth_k8s.go"
    agent_connect_uri_server_go["agent/connect/uri_server.go"
    agent_xds_xds_go["agent/xds/xds.go"
    agent_cache_types_intention_upstreams_destination_go["agent/cache-types/intention_upstreams_destination.go"
    agent_proxycfg_glue_config_entry_go["agent/proxycfg-glue/config_entry.go"
    agent_reload_go["agent/reload.go"
    api_config_entry_discoverychain_go["api/config_entry_discoverychain.go"
    agent_consul_state_catalog_events_go["agent/consul/state/catalog_events.go"
    agent_grpc_middleware_handshake_go["agent/grpc-middleware/handshake.go"
    api_namespace_go["api/namespace.go"
    agent_consul_auth_token_writer_go["agent/consul/auth/token_writer.go"
    agent_consul_config_endpoint_go["agent/consul/config_endpoint.go"
    agent_consul_server_grpc_go["agent/consul/server_grpc.go"
    agent_grpc_middleware_testutil_testservice_simple_pb_go["agent/grpc-middleware/testutil/testservice/simple.pb.go"
    agent_leafcert_leafcert_go["agent/leafcert/leafcert.go"
    agent_proxycfg_manager_go["agent/proxycfg/manager.go"
    agent_cache_mock_Type_go["agent/cache/mock_Type.go"
    agent_grpc_middleware_testutil_testservice_simple_grpc_pb_go["agent/grpc-middleware/testutil/testservice/simple_grpc.pb.go"
    agent_structs_autopilot_ce_go["agent/structs/autopilot_ce.go"
    agent_token_store_go["agent/token/store.go"
    agent_consul_state_autopilot_go["agent/consul/state/autopilot.go"
    types_node_id_go["types/node_id.go"
    agent_checks_alias_go["agent/checks/alias.go"
    agent_consul_configentry_backend_ce_go["agent/consul/configentry_backend_ce.go"
    agent_consul_leader_peering_go["agent/consul/leader_peering.go"
    agent_federation_state_endpoint_go["agent/federation_state_endpoint.go"
    agent_hcp_client_mock_Client_go["agent/hcp/client/mock_Client.go"
    agent_checks_os_service_go["agent/checks/os_service.go"
    agent_connect_ca_provider_vault_auth_alicloud_go["agent/connect/ca/provider_vault_auth_alicloud.go"
    agent_enterprise_delegate_ce_go["agent/enterprise_delegate_ce.go"
    agent_grpc_internal_resolver_resolver_go["agent/grpc-internal/resolver/resolver.go"
    agent_structs_catalog_ce_go["agent/structs/catalog_ce.go"
    api_snapshot_go["api/snapshot.go"
    agent_auto_config_auto_config_ce_go["agent/auto-config/auto_config_ce.go"
    agent_consul_acl_server_go["agent/consul/acl_server.go"
    agent_consul_rate_handler_ce_go["agent/consul/rate/handler_ce.go"
    agent_consul_state_config_entry_ce_go["agent/consul/state/config_entry_ce.go"
    agent_acl_go["agent/acl.go"
    agent_consul_config_replication_go["agent/consul/config_replication.go"
    agent_grpc_external_services_peerstream_subscription_blocking_go["agent/grpc-external/services/peerstream/subscription_blocking.go"
    api_session_go["api/session.go"
    agent_cache_types_trust_bundle_go["agent/cache-types/trust_bundle.go"
    agent_grpc_middleware_recovery_go["agent/grpc-middleware/recovery.go"
    agent_metrics_go["agent/metrics.go"
    agent_status_endpoint_go["agent/status_endpoint.go"
    agent_connect_uri_agent_go["agent/connect/uri_agent.go"
    agent_consul_controller_queue_defer_go["agent/consul/controller/queue/defer.go"
    agent_proxycfg_connect_proxy_go["agent/proxycfg/connect_proxy.go"
    agent_structs_testing_intention_go["agent/structs/testing_intention.go"
    agent_checks_docker_windows_go["agent/checks/docker_windows.go"
    agent_consul_server_register_go["agent/consul/server_register.go"
    agent_grpc_external_services_peerstream_stream_resources_go["agent/grpc-external/services/peerstream/stream_resources.go"
    agent_cache_types_mock_RPC_go["agent/cache-types/mock_RPC.go"
    agent_hcp_testing_go["agent/hcp/testing.go"
    agent_cache_cache_go["agent/cache/cache.go"
    api_connect_ca_go["api/connect_ca.go"
    agent_configentry_config_entry_go["agent/configentry/config_entry.go"
    agent_consul_usagemetrics_usagemetrics_ce_go["agent/consul/usagemetrics/usagemetrics_ce.go"
    agent_configentry_service_config_go["agent/configentry/service_config.go"
    agent_structs_system_metadata_go["agent/structs/system_metadata.go"
    agent_cache_types_catalog_datacenters_go["agent/cache-types/catalog_datacenters.go"
    agent_connect_authz_go["agent/connect/authz.go"
    agent_consul_server_overview_go["agent/consul/server_overview.go"
    agent_grpc_external_services_resource_list_go["agent/grpc-external/services/resource/list.go"
    agent_consul_snapshot_endpoint_go["agent/consul/snapshot_endpoint.go"
    agent_proxycfg_upstreams_go["agent/proxycfg/upstreams.go"
    agent_service_manager_go["agent/service_manager.go"
    agent_consul_operator_raft_endpoint_go["agent/consul/operator_raft_endpoint.go"
    agent_consul_state_schema_ce_go["agent/consul/state/schema_ce.go"
    agent_hcp_scada_scada_go["agent/hcp/scada/scada.go"
    agent_xds_secrets_go["agent/xds/secrets.go"
    api_catalog_go["api/catalog.go"
    api_semaphore_go["api/semaphore.go"
    agent_consul_stream_event_buffer_go["agent/consul/stream/event_buffer.go"
    agent_consul_state_session_ce_go["agent/consul/state/session_ce.go"
    agent_signal_windows_go["agent/signal_windows.go"
    agent_grpc_external_services_acl_mock_Validator_go["agent/grpc-external/services/acl/mock_Validator.go"
    agent_proxycfg_testing_go["agent/proxycfg/testing.go"
    agent_consul_multilimiter_multilimiter_go["agent/consul/multilimiter/multilimiter.go"
    agent_grpc_external_services_peerstream_stream_tracker_go["agent/grpc-external/services/peerstream/stream_tracker.go"
    agent_structs_config_entry_exports_ce_go["agent/structs/config_entry_exports_ce.go"
    agent_xds_resources_go["agent/xds/resources.go"
    api_partition_go["api/partition.go"
    agent_consul_state_kvs_go["agent/consul/state/kvs.go"
    agent_consul_util_go["agent/consul/util.go"
    agent_dns_ce_go["agent/dns_ce.go"
    api_content_type_go["api/content_type.go"
    acl_validation_go["acl/validation.go"
    agent_configentry_resolve_go["agent/configentry/resolve.go"
    agent_connect_ca_provider_vault_auth_jwt_go["agent/connect/ca/provider_vault_auth_jwt.go"
    agent_consul_segment_ce_go["agent/consul/segment_ce.go"
    agent_consul_state_peering_go["agent/consul/state/peering.go"
    agent_proxycfg_glue_service_list_go["agent/proxycfg-glue/service_list.go"
    agent_config_runtime_ce_go["agent/config/runtime_ce.go"
    agent_structs_federation_state_go["agent/structs/federation_state.go"
    agent_xds_server_go["agent/xds/server.go"
    agent_auto_config_auto_encrypt_go["agent/auto-config/auto_encrypt.go"
    agent_checks_os_service_unix_go["agent/checks/os_service_unix.go"
    agent_grpc_external_services_resource_mock_Registry_go["agent/grpc-external/services/resource/mock_Registry.go"
    agent_xds_accesslogs_accesslogs_go["agent/xds/accesslogs/accesslogs.go"
    agent_config_ratelimited_file_watcher_go["agent/config/ratelimited_file_watcher.go"
    agent_grpc_external_utils_go["agent/grpc-external/utils.go"
    agent_hcp_client_http_client_go["agent/hcp/client/http_client.go"
    agent_structs_dns_go["agent/structs/dns.go"
    agent_structs_discovery_chain_go["agent/structs/discovery_chain.go"
    agent_apiserver_go["agent/apiserver.go"
    agent_consul_controller_reconciler_go["agent/consul/controller/reconciler.go"
    agent_proxycfg_testing_upstreams_go["agent/proxycfg/testing_upstreams.go"
    api_acl_go["api/acl.go"
    agent_consul_state_catalog_ce_go["agent/consul/state/catalog_ce.go"
    agent_grpc_external_forward_go["agent/grpc-external/forward.go"
    agent_proxycfg_glue_peering_list_go["agent/proxycfg-glue/peering_list.go"
    agent_uiserver_buf_index_fs_go["agent/uiserver/buf_index_fs.go"
    api_discovery_chain_go["api/discovery_chain.go"
    agent_cache_types_config_entry_go["agent/cache-types/config_entry.go"
    agent_configentry_compare_go["agent/configentry/compare.go"
    agent_proxycfg_glue_federation_state_list_mesh_gateways_go["agent/proxycfg-glue/federation_state_list_mesh_gateways.go"
    agent_structs_intention_go["agent/structs/intention.go"
    acl_policy_merger_go["acl/policy_merger.go"
    agent_grpc_external_services_peerstream_mock_ACLResolver_go["agent/grpc-external/services/peerstream/mock_ACLResolver.go"
    api_operator_audit_go["api/operator_audit.go"
    acl_acl_go["acl/acl.go"
    agent_consul_discoverychain_compile_go["agent/consul/discoverychain/compile.go"
    connect_proxy_conn_go["connect/proxy/conn.go"
    agent_cache_type_go["agent/cache/type.go"
    agent_consul_state_config_entry_go["agent/consul/state/config_entry.go"
    agent_envoyextensions_registered_extensions_go["agent/envoyextensions/registered_extensions.go"
    agent_grpc_internal_balancer_registry_go["agent/grpc-internal/balancer/registry.go"
    agent_connect_ca_provider_vault_go["agent/connect/ca/provider_vault.go"
    agent_grpc_external_options_go["agent/grpc-external/options.go"
    agent_submatview_store_go["agent/submatview/store.go"
    api_operator_segment_go["api/operator_segment.go"
    agent_cache_types_node_services_go["agent/cache-types/node_services.go"
    agent_catalog_endpoint_go["agent/catalog_endpoint.go"
    agent_grpc_external_services_connectca_watch_roots_go["agent/grpc-external/services/connectca/watch_roots.go"
    agent_session_endpoint_go["agent/session_endpoint.go"
    agent_config_merge_go["agent/config/merge.go"
    agent_envoyextensions_builtin_aws_lambda_aws_lambda_go["agent/envoyextensions/builtin/aws-lambda/aws_lambda.go"
    agent_grpc_external_limiter_limiter_go["agent/grpc-external/limiter/limiter.go"
    agent_grpc_external_services_resource_mock_Backend_go["agent/grpc-external/services/resource/mock_Backend.go"
    agent_structs_connect_proxy_config_go["agent/structs/connect_proxy_config.go"
    agent_grpc_external_services_resource_watch_go["agent/grpc-external/services/resource/watch.go"
    api_debug_go["api/debug.go"
    agent_consul_state_query_ce_go["agent/consul/state/query_ce.go"
    agent_proxycfg_sources_local_sync_go["agent/proxycfg-sources/local/sync.go"
    agent_consul_controller_queue_queue_go["agent/consul/controller/queue/queue.go"
    agent_rpc_peering_testing_go["agent/rpc/peering/testing.go"
    agent_consul_state_config_entry_intention_ce_go["agent/consul/state/config_entry_intention_ce.go"
    api_exported_services_go["api/exported_services.go"
    connect_resolver_go["connect/resolver.go"
    agent_auto_config_auto_config_go["agent/auto-config/auto_config.go"
    agent_consul_operator_autopilot_endpoint_go["agent/consul/operator_autopilot_endpoint.go"
    agent_grpc_external_services_dataplane_mock_ACLResolver_go["agent/grpc-external/services/dataplane/mock_ACLResolver.go"
    agent_proxycfg_glue_resolved_service_config_go["agent/proxycfg-glue/resolved_service_config.go"
    agent_consul_acl_server_ce_go["agent/consul/acl_server_ce.go"
    agent_consul_peering_backend_go["agent/consul/peering_backend.go"
    agent_grpc_external_services_resource_testing_testing_ce_go["agent/grpc-external/services/resource/testing/testing_ce.go"
    agent_proxycfg_data_sources_ce_go["agent/proxycfg/data_sources_ce.go"
    agent_rpcclient_configentry_configentry_go["agent/rpcclient/configentry/configentry.go"
    api_operator_autopilot_go["api/operator_autopilot.go"
    agent_consul_enterprise_client_ce_go["agent/consul/enterprise_client_ce.go"
    agent_envoyextensions_builtin_property_override_structpatcher_go["agent/envoyextensions/builtin/property-override/structpatcher.go"
    agent_keyring_go["agent/keyring.go"
    api_config_entry_mesh_go["api/config_entry_mesh.go"
    agent_config_agent_limits_go["agent/config/agent_limits.go"
    agent_proxycfg_mesh_gateway_go["agent/proxycfg/mesh_gateway.go"
    agent_structs_config_entry_go["agent/structs/config_entry.go"
    agent_hcp_telemetry_otel_exporter_go["agent/hcp/telemetry/otel_exporter.go"
    agent_submatview_materializer_go["agent/submatview/materializer.go"
    agent_rpc_peering_service_go["agent/rpc/peering/service.go"
    agent_structs_envoy_extension_go["agent/structs/envoy_extension.go"
    agent_structs_structs_deepcopy_ce_go["agent/structs/structs.deepcopy_ce.go"
    agent_hcp_bootstrap_bootstrap_go["agent/hcp/bootstrap/bootstrap.go"
    api_event_go["api/event.go"
    agent_ae_ae_go["agent/ae/ae.go"
    agent_consul_enterprise_server_ce_go["agent/consul/enterprise_server_ce.go"
    agent_grpc_external_services_peerstream_subscription_state_go["agent/grpc-external/services/peerstream/subscription_state.go"
    agent_agent_endpoint_go["agent/agent_endpoint.go"
    agent_checks_check_go["agent/checks/check.go"
    agent_consul_authmethod_testing_go["agent/consul/authmethod/testing.go"
    agent_consul_leader_federation_state_ae_go["agent/consul/leader_federation_state_ae.go"
    agent_consul_state_query_go["agent/consul/state/query.go"
    agent_token_persistence_go["agent/token/persistence.go"
    agent_cache_types_trust_bundles_go["agent/cache-types/trust_bundles.go"
    agent_consul_auth_login_go["agent/consul/auth/login.go"
    agent_consul_watch_server_local_go["agent/consul/watch/server_local.go"
    agent_proxycfg_glue_glue_go["agent/proxycfg-glue/glue.go"
    agent_agent_endpoint_ce_go["agent/agent_endpoint_ce.go"
    agent_checks_os_service_windows_go["agent/checks/os_service_windows.go"
    agent_hcp_discover_discover_go["agent/hcp/discover/discover.go"
    agent_xds_configfetcher_config_fetcher_go["agent/xds/configfetcher/config_fetcher.go"
    connect_tls_go["connect/tls.go"
    agent_consul_state_delay_ce_go["agent/consul/state/delay_ce.go"
    agent_grpc_middleware_rate_limit_mappings_gen_go["agent/grpc-middleware/rate_limit_mappings.gen.go"
    agent_xds_protocol_trace_go["agent/xds/protocol_trace.go"
    api_operator_utilization_go["api/operator_utilization.go"
    agent_consul_connect_ca_endpoint_go["agent/consul/connect_ca_endpoint.go"
    agent_consul_status_endpoint_go["agent/consul/status_endpoint.go"
    agent_setup_ce_go["agent/setup_ce.go"
    agent_structs_protobuf_compat_go["agent/structs/protobuf_compat.go"
    agent_consul_state_graveyard_ce_go["agent/consul/state/graveyard_ce.go"
    agent_grpc_external_services_resource_testing_testing_go["agent/grpc-external/services/resource/testing/testing.go"
    agent_grpc_internal_client_go["agent/grpc-internal/client.go"
    agent_xds_clusters_go["agent/xds/clusters.go"
    agent_cache_types_peered_upstreams_go["agent/cache-types/peered_upstreams.go"
    agent_consul_state_federation_state_go["agent/consul/state/federation_state.go"
    agent_grpc_external_services_peerstream_subscription_manager_go["agent/grpc-external/services/peerstream/subscription_manager.go"
    agent_proxycfg_testing_terminating_gateway_go["agent/proxycfg/testing_terminating_gateway.go"
    agent_proxycfg_api_gateway_ce_go["agent/proxycfg/api_gateway_ce.go"
    agent_retry_join_go["agent/retry_join.go"
    agent_consul_acl_client_go["agent/consul/acl_client.go"
    agent_consul_server_connect_go["agent/consul/server_connect.go"
    agent_proxycfg_glue_intention_upstreams_go["agent/proxycfg-glue/intention_upstreams.go"
    agent_rpc_middleware_interceptors_go["agent/rpc/middleware/interceptors.go"
    agent_structs_autopilot_go["agent/structs/autopilot.go"
    connect_service_go["connect/service.go"
    agent_consul_multilimiter_mock_RateLimiter_go["agent/consul/multilimiter/mock_RateLimiter.go"
    agent_consul_state_tombstone_gc_go["agent/consul/state/tombstone_gc.go"
    agent_consul_stream_event_go["agent/consul/stream/event.go"
    agent_proxycfg_glue_health_blocking_go["agent/proxycfg-glue/health_blocking.go"
    agent_consul_options_go["agent/consul/options.go"
    agent_proxycfg_api_gateway_go["agent/proxycfg/api_gateway.go"
    agent_proxycfg_glue_peered_upstreams_go["agent/proxycfg-glue/peered_upstreams.go"
    api_config_entry_exports_go["api/config_entry_exports.go"
    acl_policy_authorizer_go["acl/policy_authorizer.go"
    agent_consul_flood_go["agent/consul/flood.go"
    agent_translate_addr_go["agent/translate_addr.go"
    agent_grpc_external_services_connectca_sign_go["agent/grpc-external/services/connectca/sign.go"
    agent_structs_structs_ce_go["agent/structs/structs_ce.go"
    agent_cache_types_catalog_service_list_go["agent/cache-types/catalog_service_list.go"
    agent_connect_ca_provider_vault_auth_gcp_go["agent/connect/ca/provider_vault_auth_gcp.go"
    agent_connect_sni_go["agent/connect/sni.go"
    agent_consul_authmethod_awsauth_aws_go["agent/consul/authmethod/awsauth/aws.go"
    agent_consul_controller_controller_go["agent/consul/controller/controller.go"
    agent_health_endpoint_go["agent/health_endpoint.go"
    agent_structs_config_entry_sameness_group_ce_go["agent/structs/config_entry_sameness_group_ce.go"
    agent_auto_config_tls_go["agent/auto-config/tls.go"
    agent_consul_txn_endpoint_go["agent/consul/txn_endpoint.go"
    agent_grpc_middleware_testutil_fake_sink_go["agent/grpc-middleware/testutil/fake_sink.go"
    acl_authorizer_go["acl/authorizer.go"
    agent_auto_config_server_addr_go["agent/auto-config/server_addr.go"
    agent_cache_types_health_services_go["agent/cache-types/health_services.go"
    agent_config_config_ce_go["agent/config/config_ce.go"
    agent_consul_system_metadata_go["agent/consul/system_metadata.go"
    agent_grpc_external_services_dataplane_get_envoy_bootstrap_params_go["agent/grpc-external/services/dataplane/get_envoy_bootstrap_params.go"
    agent_hcp_mock_TelemetryProvider_go["agent/hcp/mock_TelemetryProvider.go"
    agent_structs_config_entry_jwt_provider_ce_go["agent/structs/config_entry_jwt_provider_ce.go"
    agent_uiserver_redirect_fs_go["agent/uiserver/redirect_fs.go"
    agent_consul_rate_mock_RequestLimitsHandler_go["agent/consul/rate/mock_RequestLimitsHandler.go"
    api_operator_keyring_go["api/operator_keyring.go"
    agent_consul_leader_connect_ca_ce_go["agent/consul/leader_connect_ca_ce.go"
    agent_consul_prepared_query_walk_go["agent/consul/prepared_query/walk.go"
    agent_consul_state_events_go["agent/consul/state/events.go"
    agent_structs_config_entry_discoverychain_ce_go["agent/structs/config_entry_discoverychain_ce.go"
    agent_consul_state_catalog_events_ce_go["agent/consul/state/catalog_events_ce.go"
    agent_connect_ca_provider_go["agent/connect/ca/provider.go"
    agent_connect_uri_go["agent/connect/uri.go"
    agent_consul_leader_intentions_go["agent/consul/leader_intentions.go"
    agent_leafcert_leafcert_test_helpers_go["agent/leafcert/leafcert_test_helpers.go"
    agent_netutil_network_go["agent/netutil/network.go"
    agent_rpcclient_common_go["agent/rpcclient/common.go"
    agent_envoyextensions_builtin_wasm_wasm_go["agent/envoyextensions/builtin/wasm/wasm.go"
    agent_hcp_config_mock_CloudConfig_go["agent/hcp/config/mock_CloudConfig.go"
    agent_nodeid_go["agent/nodeid.go"
    types_checks_go["types/checks.go"
    agent_rpcclient_configentry_view_go["agent/rpcclient/configentry/view.go"
    agent_xds_testing_go["agent/xds/testing.go"
    agent_configentry_merge_service_config_go["agent/configentry/merge_service_config.go"
    agent_consul_gateway_locator_go["agent/consul/gateway_locator.go"
    api_config_entry_sameness_group_go["api/config_entry_sameness_group.go"
    agent_cache_request_go["agent/cache/request.go"
    agent_grpc_external_testutils_server_go["agent/grpc-external/testutils/server.go"
    agent_consul_coordinate_endpoint_go["agent/consul/coordinate_endpoint.go"
    agent_grpc_middleware_testutil_testservice_fake_service_go["agent/grpc-middleware/testutil/testservice/fake_service.go"
    agent_xds_locality_policy_ce_go["agent/xds/locality_policy_ce.go"
    agent_cache_types_options_go["agent/cache-types/options.go"
    agent_cache_types_testing_go["agent/cache-types/testing.go"
    agent_grpc_middleware_auth_interceptor_go["agent/grpc-middleware/auth_interceptor.go"
    agent_proxycfg_proxycfg_go["agent/proxycfg/proxycfg.go"
    agent_proxycfg_sources_local_config_source_go["agent/proxycfg-sources/local/config_source.go"
    api_txn_go["api/txn.go"
    agent_consul_discoverychain_string_stack_go["agent/consul/discoverychain/string_stack.go"
    agent_grpc_external_services_connectca_mock_ACLResolver_go["agent/grpc-external/services/connectca/mock_ACLResolver.go"
    agent_cache_testing_go["agent/cache/testing.go"
    agent_consul_operator_backend_go["agent/consul/operator_backend.go"
    agent_consul_subscribe_backend_go["agent/consul/subscribe_backend.go"
    agent_rpcclient_health_view_go["agent/rpcclient/health/view.go"
    agent_structs_service_definition_go["agent/structs/service_definition.go"
    acl_acl_ce_go["acl/acl_ce.go"
    agent_catalog_endpoint_ce_go["agent/catalog_endpoint_ce.go"
    agent_consul_authmethod_authmethods_ce_go["agent/consul/authmethod/authmethods_ce.go"
    agent_consul_acl_token_exp_go["agent/consul/acl_token_exp.go"
    agent_structs_testing_service_definition_go["agent/structs/testing_service_definition.go"
    agent_consul_state_graveyard_go["agent/consul/state/graveyard.go"
    agent_grpc_external_services_serverdiscovery_watch_servers_go["agent/grpc-external/services/serverdiscovery/watch_servers.go"
    api_operator_license_go["api/operator_license.go"
    agent_consul_auth_token_writer_ce_go["agent/consul/auth/token_writer_ce.go"
    agent_consul_state_peering_ce_go["agent/consul/state/peering_ce.go"
    agent_structs_config_entry_mesh_ce_go["agent/structs/config_entry_mesh_ce.go"
    agent_consul_federation_state_replication_go["agent/consul/federation_state_replication.go"
    agent_consul_state_acl_go["agent/consul/state/acl.go"
    agent_grpc_external_querymeta_go["agent/grpc-external/querymeta.go"
    acl_authorizer_ce_go["acl/authorizer_ce.go"
    agent_config_doc_go["agent/config/doc.go"
    agent_hcp_mock_Manager_go["agent/hcp/mock_Manager.go"
    agent_grpc_external_testutils_fsm_go["agent/grpc-external/testutils/fsm.go"
    agent_cache_types_catalog_services_go["agent/cache-types/catalog_services.go"
    agent_consul_server_metadata_go["agent/consul/server_metadata.go"
    agent_blockingquery_mock_ResponseMeta_go["agent/blockingquery/mock_ResponseMeta.go"
    agent_cache_types_service_checks_go["agent/cache-types/service_checks.go"
    agent_consul_reporting_reportingmock_mock_ServerDelegate_go["agent/consul/reporting/reportingmock/mock_ServerDelegate.go"
    agent_consul_replication_go["agent/consul/replication.go"
    agent_grpc_external_services_resource_mock_TenancyBridge_go["agent/grpc-external/services/resource/mock_TenancyBridge.go"
    agent_proxycfg_data_sources_go["agent/proxycfg/data_sources.go"
    agent_consul_config_ce_go["agent/consul/config_ce.go"
    agent_consul_fsm_decode_downgrade_go["agent/consul/fsm/decode_downgrade.go"
    agent_consul_server_lookup_go["agent/consul/server_lookup.go"
    agent_leafcert_structs_go["agent/leafcert/structs.go"
    agent_xds_naming_naming_go["agent/xds/naming/naming.go"
    agent_structs_acl_templated_policy_go["agent/structs/acl_templated_policy.go"
    agent_util_go["agent/util.go"
    agent_cache_types_rpc_go["agent/cache-types/rpc.go"
    agent_consul_prepared_query_endpoint_go["agent/consul/prepared_query_endpoint.go"
    agent_consul_stream_noop_go["agent/consul/stream/noop.go"
    agent_config_flagset_go["agent/config/flagset.go"
    agent_consul_authmethod_testauth_testing_go["agent/consul/authmethod/testauth/testing.go"
    agent_grpc_external_services_configentry_server_go["agent/grpc-external/services/configentry/server.go"
    agent_structs_structs_go["agent/structs/structs.go"
    agent_connect_ca_provider_vault_auth_approle_go["agent/connect/ca/provider_vault_auth_approle.go"
    agent_consul_authmethod_kubeauth_k8s_go["agent/consul/authmethod/kubeauth/k8s.go"
    agent_consul_reporting_reporting_ce_go["agent/consul/reporting/reporting_ce.go"
    agent_consul_state_catalog_schema_deepcopy_go["agent/consul/state/catalog_schema.deepcopy.go"
    agent_structs_config_entry_exports_go["agent/structs/config_entry_exports.go"
    agent_consul_federation_state_endpoint_go["agent/consul/federation_state_endpoint.go"
    agent_pool_peek_go["agent/pool/peek.go"
    agent_xds_platform_net_linux_go["agent/xds/platform/net_linux.go"
    agent_config_segment_ce_go["agent/config/segment_ce.go"
    agent_hcp_client_client_go["agent/hcp/client/client.go"
    api_connect_intention_go["api/connect_intention.go"
    agent_cache_types_service_dump_go["agent/cache-types/service_dump.go"
    agent_consul_logging_go["agent/consul/logging.go"
    agent_consul_merge_ce_go["agent/consul/merge_ce.go"
    agent_consul_v2_config_entry_exports_shim_go["agent/consul/v2_config_entry_exports_shim.go"
    agent_hcp_telemetry_custom_metrics_go["agent/hcp/telemetry/custom_metrics.go"
    agent_consul_authmethod_ssoauth_sso_go["agent/consul/authmethod/ssoauth/sso.go"
    agent_consul_state_usage_go["agent/consul/state/usage.go"
    agent_leafcert_signer_netrpc_go["agent/leafcert/signer_netrpc.go"
    agent_log_drop_log_drop_go["agent/log-drop/log-drop.go"
    agent_uiserver_uiserver_go["agent/uiserver/uiserver.go"
    agent_connect_uri_signing_go["agent/connect/uri_signing.go"
    agent_hcp_telemetry_otel_sink_go["agent/hcp/telemetry/otel_sink.go"
    agent_connect_ca_testing_go["agent/connect/ca/testing.go"
    agent_consul_discoverychain_compile_ce_go["agent/consul/discoverychain/compile_ce.go"
    agent_http_go["agent/http.go"
    agent_connect_common_names_go["agent/connect/common_names.go"
    agent_structs_prepared_query_go["agent/structs/prepared_query.go"
    agent_consul_state_prepared_query_go["agent/consul/state/prepared_query.go"
    agent_structs_config_entry_status_go["agent/structs/config_entry_status.go"
    agent_auto_config_persist_go["agent/auto-config/persist.go"
    agent_auto_config_config_translate_go["agent/auto-config/config_translate.go"
    agent_exec_exec_go["agent/exec/exec.go"
    agent_proxycfg_sources_local_local_go["agent/proxycfg-sources/local/local.go"
    agent_consul_auto_config_endpoint_go["agent/consul/auto_config_endpoint.go"
    agent_consul_state_config_entry_sameness_group_ce_go["agent/consul/state/config_entry_sameness_group_ce.go"
    agent_consul_stats_fetcher_go["agent/consul/stats_fetcher.go"
    agent_consul_tenancy_bridge_go["agent/consul/tenancy_bridge.go"
    agent_grpc_internal_balancer_balancer_go["agent/grpc-internal/balancer/balancer.go"
    agent_pool_conn_go["agent/pool/conn.go"
    agent_acl_ce_go["agent/acl_ce.go"
    agent_checks_grpc_go["agent/checks/grpc.go"
    agent_envoyextensions_builtin_wasm_structs_go["agent/envoyextensions/builtin/wasm/structs.go"
    agent_config_limits_go["agent/config/limits.go"
    agent_consul_state_operations_ce_go["agent/consul/state/operations_ce.go"
    agent_structs_config_entry_apigw_jwt_ce_go["agent/structs/config_entry_apigw_jwt_ce.go"
    api_operator_go["api/operator.go"
    agent_cache_mock_Request_go["agent/cache/mock_Request.go"
    agent_config_config_go["agent/config/config.go"
    agent_grpc_external_services_resource_list_by_owner_go["agent/grpc-external/services/resource/list_by_owner.go"
    agent_grpc_external_services_peerstream_testing_go["agent/grpc-external/services/peerstream/testing.go"
    agent_grpc_external_services_resource_server_ce_go["agent/grpc-external/services/resource/server_ce.go"
    agent_xds_extensionruntime_runtime_config_go["agent/xds/extensionruntime/runtime_config.go"
    agent_structs_testing_go["agent/structs/testing.go"
    agent_consul_fsm_decode_ce_go["agent/consul/fsm/decode_ce.go"
    agent_proxycfg_testing_mesh_gateway_go["agent/proxycfg/testing_mesh_gateway.go"
    api_config_entry_inline_certificate_go["api/config_entry_inline_certificate.go"
    agent_connect_ca_provider_consul_go["agent/connect/ca/provider_consul.go"
    agent_consul_acl_ce_go["agent/consul/acl_ce.go"
    agent_consul_state_catalog_schema_go["agent/consul/state/catalog_schema.go"
    agent_grpc_external_services_connectca_mock_CAManager_go["agent/grpc-external/services/connectca/mock_CAManager.go"
    agent_proxycfg_ingress_gateway_go["agent/proxycfg/ingress_gateway.go"
    agent_rpc_peering_validate_go["agent/rpc/peering/validate.go"
    agent_structs_intention_ce_go["agent/structs/intention_ce.go"
    agent_leafcert_util_go["agent/leafcert/util.go"
    agent_rpc_operator_service_go["agent/rpc/operator/service.go"
    api_config_entry_gateways_go["api/config_entry_gateways.go"
    agent_cache_types_resolved_service_config_go["agent/cache-types/resolved_service_config.go"
    agent_debug_host_go["agent/debug/host.go"
    agent_envoyextensions_builtin_lua_lua_go["agent/envoyextensions/builtin/lua/lua.go"
    agent_hcp_bootstrap_testing_go["agent/hcp/bootstrap/testing.go"
    agent_structs_catalog_go["agent/structs/catalog.go"
    agent_check_go["agent/check.go"
    agent_consul_state_connect_ca_go["agent/consul/state/connect_ca.go"
    api_kv_go["api/kv.go"
    agent_cache_types_exported_peered_services_go["agent/cache-types/exported_peered_services.go"
    agent_consul_session_ttl_go["agent/consul/session_ttl.go"
    agent_consul_state_config_entry_exported_services_ce_go["agent/consul/state/config_entry_exported_services_ce.go"
    agent_grpc_external_services_acl_server_go["agent/grpc-external/services/acl/server.go"
    agent_token_store_ce_go["agent/token/store_ce.go"
    agent_config_default_go["agent/config/default.go"
    agent_consul_authmethod_authmethods_go["agent/consul/authmethod/authmethods.go"
    agent_consul_controller_doc_go["agent/consul/controller/doc.go"
    agent_consul_state_system_metadata_go["agent/consul/state/system_metadata.go"
    agent_consul_state_schema_go["agent/consul/state/schema.go"
    agent_cache_types_catalog_list_services_go["agent/cache-types/catalog_list_services.go"
    agent_cache_types_federation_state_list_gateways_go["agent/cache-types/federation_state_list_gateways.go"
    agent_consul_state_acl_schema_go["agent/consul/state/acl_schema.go"
    agent_consul_state_mock_publishFuncType_go["agent/consul/state/mock_publishFuncType.go"
    agent_grpc_external_services_connectca_server_go["agent/grpc-external/services/connectca/server.go"
    agent_grpc_external_services_resource_read_go["agent/grpc-external/services/resource/read.go"
    agent_agent_ce_go["agent/agent_ce.go"
    agent_coordinate_endpoint_go["agent/coordinate_endpoint.go"
    agent_envoyextensions_registered_extensions_ce_go["agent/envoyextensions/registered_extensions_ce.go"
    agent_consul_state_config_entry_schema_go["agent/consul/state/config_entry_schema.go"
    agent_submatview_rpc_materializer_go["agent/submatview/rpc_materializer.go"
    agent_txn_endpoint_go["agent/txn_endpoint.go"


    classDef style0 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style1 fill:#4CAF50,stroke:#333,stroke-width:2px
    classDef style2 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style3 fill:#90CAF9,stroke:#333,stroke-width:2px

    class agent_cache_types_service_gateways_go style0
    class agent_proxycfg_internal_watch_watchmap_go style0
    class agent_structs_config_entry_jwt_provider_go style0
    class agent_proxycfg_glue_intentions_go style0
    class agent_proxycfg_sources_catalog_mock_Watcher_go style0
    class agent_structs_acl_cache_go style0
    class agent_grpc_middleware_stats_go style0
    class agent_consul_rate_handler_go style0
    class agent_structs_config_entry_sameness_group_go style0
    class api_status_go style1
    class agent_xds_z_xds_packages_go style0
    class acl_policy_ce_go style2
    class agent_cache_entry_go style0
    class agent_consul_discoverychain_testing_go style0
    class agent_envoyextensions_builtin_otel_access_logging_structs_go style0
    class agent_kvs_endpoint_go style0
    class agent_signal_unix_go style0
    class agent_connect_generate_go style0
    class agent_grpc_internal_listener_go style0
    class agent_consul_raft_handle_go style0
    class agent_consul_state_intention_go style0
    class agent_consul_state_usage_ce_go style0
    class agent_mock_notify_go style0
    class agent_proxycfg_glue_leafcerts_go style0
    class agent_structs_check_definition_go style1
    class agent_consul_gateways_controller_gateways_go style0
    class agent_grpc_internal_services_subscribe_logger_go style0
    class agent_proxycfg_glue_gateway_services_go style0
    class agent_dns_context_go style0
    class agent_proxycfg_state_go style0
    class agent_xds_delta_go style0
    class agent_cache_watch_go style0
    class agent_consul_discovery_chain_endpoint_go style0
    class agent_consul_intention_endpoint_go style0
    class agent_structs_config_entry_gateways_go style0
    class agent_xds_jwt_authn_go style0
    class agent_xds_platform_net_fallback_go style0
    class agent_cache_types_gateway_services_go style0
    class agent_consul_reporting_reporting_go style0
    class agent_consul_stream_event_snapshot_go style0
    class agent_consul_wanfed_pool_go style0
    class agent_proxycfg_glue_trust_bundle_go style0
    class agent_hcp_deps_go style0
    class agent_hcp_telemetry_doc_go style0
    class agent_proxycfg_glue_exported_peered_services_go style0
    class agent_consul_authmethod_testauth_testing_ce_go style0
    class agent_hcp_bootstrap_constants_constants_go style0
    class agent_proxycfg_mesh_gateway_ce_go style0
    class agent_cache_types_intention_upstreams_go style0
    class agent_consul_state_coordinate_go style0
    class agent_proxycfg_glue_intentions_ce_go style0
    class agent_connect_ca_provider_aws_go style0
    class agent_connect_csr_go style0
    class agent_consul_state_txn_go style0
    class agent_hcp_telemetry_gauge_store_go style0
    class agent_xds_locality_policy_go style0
    class api_operator_usage_go style1
    class agent_consul_operator_usage_endpoint_go style0
    class agent_structs_txn_go style0
    class agent_structs_snapshot_go style0
    class api_config_entry_go style1
    class agent_hcp_telemetry_provider_go style0
    class agent_structs_config_entry_ce_go style0
    class agent_blockingquery_mock_RequestOptions_go style0
    class agent_grpc_external_services_peerstream_health_snapshot_go style0
    class agent_structs_discovery_chain_ce_go style0
    class agent_connect_uri_mesh_gateway_ce_go style0
    class agent_metrics_testing_go style0
    class agent_proxycfg_glue_service_http_checks_go style0
    class agent_ui_endpoint_go style0
    class agent_xds_config_config_go style0
    class agent_proxycfg_naming_ce_go style0
    class agent_watch_handler_go style0
    class agent_config_endpoint_go style0
    class agent_connect_parsing_go style0
    class agent_consul_session_endpoint_go style0
    class agent_cache_types_peerings_go style0
    class agent_consul_state_config_entry_sameness_group_go style0
    class agent_discovery_chain_endpoint_go style0
    class agent_local_state_go style0
    class agent_remote_exec_go style0
    class agent_xds_endpoints_go style0
    class agent_consul_leader_connect_ca_go style0
    class agent_consul_raft_rpc_go style0
    class api_raw_go style1
    class agent_proxycfg_testing_ce_go style0
    class agent_structs_aclfilter_filter_go style0
    class agent_xds_listeners_apigateway_go style0
    class agent_xds_server_ce_go style0
    class acl_testing_go style2
    class acl_MockAuthorizer_go style2
    class agent_proxycfg_sources_catalog_config_source_go style0
    class agent_consul_fsm_commands_ce_go style0
    class api_config_entry_status_go style1
    class agent_connect_uri_agent_ce_go style0
    class agent_consul_enterprise_config_ce_go style0
    class agent_blockingquery_mock_FSMServer_go style0
    class agent_structs_check_type_go style0
    class agent_uiserver_ui_template_data_go style0
    class agent_xds_response_response_go style0
    class api_watch_funcs_go style1
    class acl_chained_authorizer_go style2
    class agent_consul_autopilot_ce_go style0
    class agent_consul_state_config_entry_intention_go style0
    class agent_operator_endpoint_ce_go style0
    class agent_proxycfg_testing_ingress_gateway_go style0
    class agent_structs_connect_go style0
    class agent_user_event_go style0
    class acl_enterprisemeta_ce_go style2
    class agent_config_flags_go style0
    class agent_consul_filter_go style0
    class agent_structs_structs_deepcopy_go style0
    class agent_connect_uri_service_ce_go style0
    class agent_consul_auth_mock_ACLCache_go style0
    class connect_proxy_testing_go style3
    class agent_connect_testing_ca_go style0
    class agent_connect_uri_mesh_gateway_go style0
    class agent_consul_state_prepared_query_index_go style0
    class agent_dns_buffer_response_writer_go style0
    class agent_structs_connect_proxy_config_ce_go style0
    class agent_connect_testing_spiffe_go style0
    class agent_hcp_scada_mock_Provider_go style0
    class agent_structs_config_entry_file_system_certificate_go style0
    class agent_utilization_endpoint_go style0
    class agent_consul_session_timers_go style0
    class agent_connect_ca_endpoint_go style0
    class agent_consul_server_serf_go style0
    class agent_event_endpoint_go style0
    class agent_proxycfg_glue_discovery_chain_go style0
    class agent_consul_state_connect_ca_events_go style0
    class agent_consul_state_catalog_go style0
    class agent_structs_config_entry_mesh_go style0
    class connect_certgen_certgen_go style3
    class agent_consul_auto_config_backend_go style0
    class agent_structs_acl_templated_policy_ce_go style0
    class agent_consul_acl_authmethod_go style0
    class agent_structs_testing_catalog_go style0
    class api_config_entry_rate_limit_ip_go style1
    class agent_consul_stream_string_types_go style0
    class agent_consul_stream_subscription_go style0
    class agent_consul_server_log_verification_go style0
    class agent_envoyextensions_builtin_property_override_property_override_go style0
    class agent_proxycfg_sources_catalog_mock_ConfigManager_go style0
    class agent_auto_config_config_go style0
    class agent_connect_x509_patch_go style0
    class agent_consul_xdscapacity_capacity_go style0
    class agent_proxycfg_glue_health_go style0
    class agent_hcp_client_errors_go style0
    class agent_sidecar_service_go style0
    class agent_consul_usagemetrics_usagemetrics_go style0
    class agent_consul_leader_intentions_ce_go style0
    class agent_leafcert_roots_go style0
    class agent_pool_peeked_conn_go style0
    class agent_structs_peering_go style1
    class agent_auto_config_run_go style0
    class agent_consul_rtt_go style0
    class agent_denylist_go style0
    class api_config_entry_intentions_go style1
    class agent_leafcert_cached_roots_go style0
    class api_lock_go style1
    class agent_grpc_external_services_acl_login_go style0
    class agent_consul_options_ce_go style0
    class agent_hcp_bootstrap_config_loader_loader_go style0
    class agent_operator_endpoint_go style0
    class connect_proxy_config_go style3
    class agent_connect_ca_mock_Provider_go style0
    class agent_consul_client_go style0
    class agent_consul_state_session_go style0
    class connect_testing_go style3
    class agent_grpc_external_services_acl_mock_Login_go style0
    class agent_connect_ca_provider_vault_auth_azure_go style0
    class agent_consul_rpc_go style0
    class agent_grpc_external_services_resource_delete_go style0
    class agent_http_register_go style0
    class agent_grpc_external_services_serverdiscovery_server_go style0
    class agent_grpc_external_services_peerstream_subscription_view_go style0
    class api_api_go style1
    class agent_consul_state_coordinate_ce_go style0
    class agent_consul_fsm_log_verification_chunking_shim_go style0
    class api_agent_go style1
    class agent_envoyextensions_builtin_ext_authz_ext_authz_go style0
    class agent_consul_peering_backend_ce_go style0
    class agent_grpc_external_services_serverdiscovery_mock_ACLResolver_go style0
    class agent_proxycfg_sources_catalog_mock_SessionLimiter_go style0
    class api_operator_area_go style1
    class agent_intentions_endpoint_go style0
    class acl_static_authorizer_go style2
    class agent_cache_types_intention_match_go style0
    class agent_checks_docker_unix_go style0
    class agent_consul_rate_metrics_go style0
    class agent_proxycfg_snapshot_go style0
    class api_watch_plan_go style1
    class acl_resolver_result_go style2
    class agent_grpc_middleware_rate_go style0
    class api_coordinate_go style1
    class agent_config_config_deepcopy_go style0
    class agent_consul_state_acl_ce_go style0
    class agent_grpc_external_testutils_acl_go style0
    class agent_proxycfg_proxycfg_deepcopy_go style0
    class agent_proxycfg_testing_upstreams_ce_go style0
    class api_operator_raft_go style1
    class agent_checks_docker_go style0
    class agent_consul_tenancy_bridge_ce_go style0
    class agent_rpcclient_health_health_go style0
    class agent_consul_prepared_query_template_go style0
    class agent_grpc_external_services_peerstream_replication_go style0
    class agent_grpc_internal_tracker_go style0
    class api_config_entry_jwt_provider_go style1
    class agent_envoyextensions_builtin_ext_authz_structs_go style0
    class agent_pool_pool_go style0
    class agent_consul_state_acl_events_go style0
    class agent_hcp_client_telemetry_config_go style0
    class agent_proxycfg_config_snapshot_glue_go style0
    class agent_consul_reporting_reportingmock_mock_StateDelegate_go style0
    class agent_http_ce_go style0
    class agent_metadata_server_go style0
    class agent_structs_connect_ca_go style0
    class agent_consul_auto_encrypt_endpoint_go style0
    class agent_consul_serf_filter_go style0
    class agent_exec_exec_windows_go style0
    class agent_rpc_middleware_rate_limit_mappings_go style0
    class agent_systemd_notify_go style0
    class acl_policy_authorizer_ce_go style2
    class agent_consul_kvs_endpoint_go style0
    class agent_config_file_watcher_go style0
    class agent_cacheshim_cache_go style0
    class agent_consul_health_endpoint_ce_go style0
    class agent_consul_leader_log_verification_go style0
    class agent_structs_connect_ce_go style0
    class api_prepared_query_go style1
    class agent_consul_discoverychain_gateway_go style0
    class agent_prepared_query_endpoint_go style0
    class api_health_go style1
    class agent_config_limits_windows_go style0
    class agent_agent_go style0
    class agent_ae_trigger_go style0
    class agent_connect_ca_common_go style0
    class agent_consul_acl_endpoint_ce_go style0
    class agent_consul_state_intention_ce_go style0
    class agent_consul_state_kvs_ce_go style0
    class agent_structs_acl_ce_go style0
    class agent_auto_config_config_ce_go style0
    class agent_consul_config_go style0
    class agent_consul_controller_queue_rate_go style0
    class agent_envoyextensions_builtin_otel_access_logging_otel_access_logging_go style0
    class agent_consul_leader_connect_go style0
    class agent_dns_go style0
    class agent_submatview_handler_go style0
    class agent_structs_auto_encrypt_go style0
    class agent_consul_authmethod_ssoauth_sso_ce_go style0
    class agent_consul_merge_go style0
    class agent_connect_ca_provider_vault_auth_aws_go style0
    class agent_consul_acl_replication_go style0
    class agent_connect_uri_service_go style0
    class agent_proxycfg_testing_peering_go style0
    class api_watch_watch_go style1
    class agent_consul_internal_endpoint_go style0
    class agent_grpc_external_testutils_mock_server_transport_stream_go style0
    class agent_hcp_config_config_go style0
    class agent_structs_testing_connect_proxy_config_go style0
    class agent_grpc_external_services_dataplane_get_supported_features_go style0
    class agent_local_testing_go style0
    class agent_consul_leader_metrics_go style0
    class agent_leafcert_generate_go style0
    class agent_leafcert_cert_go style0
    class agent_structs_config_entry_discoverychain_go style0
    class agent_config_default_ce_go style0
    class agent_configentry_doc_go style0
    class agent_consul_state_census_utilization_go style0
    class agent_notify_go style0
    class agent_xds_testcommon_testcommon_go style0
    class agent_acl_endpoint_go style0
    class agent_cache_types_connect_ca_root_go style0
    class agent_grpc_external_services_acl_mock_TokenWriter_go style0
    class agent_consul_state_indexer_go style0
    class agent_xds_rbac_go style0
    class agent_xds_jwt_authn_ce_go style0
    class agent_consul_autopilotevents_ready_servers_events_go style0
    class agent_structs_acl_go style0
    class agent_structs_config_entry_inline_certificate_go style0
    class acl_resolver_danger_go style2
    class agent_consul_leader_registrator_v1_go style0
    class agent_consul_server_ce_go style0
    class agent_consul_state_memdb_go style0
    class agent_structs_config_entry_intentions_ce_go style0
    class agent_structs_identity_go style0
    class connect_proxy_listener_go style3
    class agent_cache_types_prepared_query_go style0
    class agent_consul_autopilot_go style0
    class agent_grpc_external_services_resource_write_status_go style0
    class test_integration_connect_envoy_test_sds_server_sds_go style3
    class agent_consul_prepared_query_endpoint_ce_go style0
    class agent_proxycfg_naming_go style0
    class agent_proxycfg_testing_connect_proxy_go style0
    class agent_structs_errors_go style0
    class agent_exec_exec_unix_go style0
    class agent_leafcert_watch_go style0
    class agent_proxycfg_terminating_gateway_go style0
    class agent_submatview_local_materializer_go style0
    class agent_consul_configentry_backend_go style0
    class agent_consul_stream_event_publisher_go style0
    class agent_grpc_external_services_resource_write_go style0
    class agent_hcp_telemetry_otlp_transform_go style0
    class agent_log_drop_mock_Logger_go style0
    class agent_proxycfg_testing_api_gateway_go style0
    class agent_proxycfg_testing_tproxy_go style0
    class agent_connect_ca_provider_consul_config_go style0
    class agent_consul_context_go style0
    class agent_consul_state_config_entry_events_go style0
    class agent_grpc_internal_pipe_go style0
    class agent_proxycfg_sources_local_mock_ConfigManager_go style0
    class agent_consul_fsm_snapshot_go style0
    class agent_xds_failover_policy_ce_go style0
    class agent_consul_servercert_manager_go style0
    class agent_grpc_external_services_resource_mutate_and_validate_go style0
    class agent_testagent_go style0
    class agent_consul_authmethod_kubeauth_testing_go style0
    class agent_consul_wanfed_wanfed_go style0
    class agent_hcp_scada_capabilities_go style0
    class agent_hcp_manager_go style0
    class acl_errors_ce_go style2
    class agent_config_deprecated_go style0
    class agent_consul_acl_go style0
    class agent_rpc_middleware_recovery_go style0
    class connect_proxy_proxy_go style3
    class acl_policy_go style2
    class agent_consul_authmethod_kubeauth_k8s_ce_go style0
    class agent_consul_operator_endpoint_go style0
    class agent_blockingquery_blockingquery_go style0
    class agent_consul_acl_authmethod_ce_go style0
    class agent_peering_endpoint_go style0
    class api_config_entry_file_system_certificate_go style1
    class agent_consul_client_serf_go style0
    class acl_policy_merger_ce_go style2
    class acl_errors_go style2
    class agent_config_runtime_go style0
    class agent_hcp_client_metrics_client_go style0
    class agent_setup_go style0
    class agent_grpc_internal_services_subscribe_subscribe_go style0
    class agent_submatview_mock_ACLResolver_go style0
    class agent_structs_config_entry_intentions_go style0
    class agent_structs_operator_go style0
    class api_peering_go style1
    class agent_consul_state_config_entry_exported_services_go style0
    class agent_grpc_internal_resolver_registry_go style0
    class agent_hcp_testserver_main_go style0
    class agent_configentry_discoverychain_go style0
    class agent_xds_listeners_go style0
    class agent_xds_listeners_ingress_go style0
    class api_connect_go style1
    class agent_consul_type_registry_go style0
    class agent_uiserver_buffered_file_go style0
    class agent_consul_gateways_controller_gateways_ce_go style0
    class agent_proxycfg_glue_internal_service_dump_go style0
    class agent_consul_acl_replication_types_go style0
    class agent_snapshot_endpoint_go style0
    class agent_cache_types_discovery_chain_go style0
    class agent_grpc_external_services_resource_mock_ACLResolver_go style0
    class agent_connect_ca_provider_vault_auth_go style0
    class agent_consul_fsm_snapshot_ce_go style0
    class agent_xds_failover_policy_go style0
    class api_internal_go style1
    class agent_connect_ca_provider_vault_auth_k8s_go style0
    class agent_connect_uri_server_go style0
    class agent_xds_xds_go style0
    class agent_cache_types_intention_upstreams_destination_go style0
    class agent_proxycfg_glue_config_entry_go style0
    class agent_reload_go style0
    class api_config_entry_discoverychain_go style1
    class agent_consul_state_catalog_events_go style0
    class agent_grpc_middleware_handshake_go style0
    class api_namespace_go style1
    class agent_consul_auth_token_writer_go style0
    class agent_consul_config_endpoint_go style0
    class agent_consul_server_grpc_go style0
    class agent_grpc_middleware_testutil_testservice_simple_pb_go style0
    class agent_leafcert_leafcert_go style0
    class agent_proxycfg_manager_go style3
    class agent_cache_mock_Type_go style0
    class agent_grpc_middleware_testutil_testservice_simple_grpc_pb_go style0
    class agent_structs_autopilot_ce_go style0
    class agent_token_store_go style0
    class agent_consul_state_autopilot_go style0
    class types_node_id_go style1
    class agent_checks_alias_go style0
    class agent_consul_configentry_backend_ce_go style0
    class agent_consul_leader_peering_go style0
    class agent_federation_state_endpoint_go style0
    class agent_hcp_client_mock_Client_go style0
    class agent_checks_os_service_go style0
    class agent_connect_ca_provider_vault_auth_alicloud_go style0
    class agent_enterprise_delegate_ce_go style0
    class agent_grpc_internal_resolver_resolver_go style0
    class agent_structs_catalog_ce_go style0
    class api_snapshot_go style1
    class agent_auto_config_auto_config_ce_go style0
    class agent_consul_acl_server_go style0
    class agent_consul_rate_handler_ce_go style0
    class agent_consul_state_config_entry_ce_go style0
    class agent_acl_go style0
    class agent_consul_config_replication_go style0
    class agent_grpc_external_services_peerstream_subscription_blocking_go style0
    class api_session_go style1
    class agent_cache_types_trust_bundle_go style0
    class agent_grpc_middleware_recovery_go style0
    class agent_metrics_go style0
    class agent_status_endpoint_go style0
    class agent_connect_uri_agent_go style0
    class agent_consul_controller_queue_defer_go style0
    class agent_proxycfg_connect_proxy_go style0
    class agent_structs_testing_intention_go style0
    class agent_checks_docker_windows_go style0
    class agent_consul_server_register_go style0
    class agent_grpc_external_services_peerstream_stream_resources_go style0
    class agent_cache_types_mock_RPC_go style0
    class agent_hcp_testing_go style0
    class agent_cache_cache_go style0
    class api_connect_ca_go style1
    class agent_configentry_config_entry_go style0
    class agent_consul_usagemetrics_usagemetrics_ce_go style0
    class agent_configentry_service_config_go style0
    class agent_structs_system_metadata_go style0
    class agent_cache_types_catalog_datacenters_go style0
    class agent_connect_authz_go style0
    class agent_consul_server_overview_go style0
    class agent_grpc_external_services_resource_list_go style0
    class agent_consul_snapshot_endpoint_go style0
    class agent_proxycfg_upstreams_go style0
    class agent_service_manager_go style0
    class agent_consul_operator_raft_endpoint_go style0
    class agent_consul_state_schema_ce_go style0
    class agent_hcp_scada_scada_go style0
    class agent_xds_secrets_go style0
    class api_catalog_go style1
    class api_semaphore_go style1
    class agent_consul_stream_event_buffer_go style0
    class agent_consul_state_session_ce_go style0
    class agent_signal_windows_go style0
    class agent_grpc_external_services_acl_mock_Validator_go style0
    class agent_proxycfg_testing_go style0
    class agent_consul_multilimiter_multilimiter_go style0
    class agent_grpc_external_services_peerstream_stream_tracker_go style0
    class agent_structs_config_entry_exports_ce_go style0
    class agent_xds_resources_go style0
    class api_partition_go style1
    class agent_consul_state_kvs_go style0
    class agent_consul_util_go style0
    class agent_dns_ce_go style0
    class api_content_type_go style1
    class acl_validation_go style2
    class agent_configentry_resolve_go style0
    class agent_connect_ca_provider_vault_auth_jwt_go style0
    class agent_consul_segment_ce_go style0
    class agent_consul_state_peering_go style0
    class agent_proxycfg_glue_service_list_go style0
    class agent_config_runtime_ce_go style0
    class agent_structs_federation_state_go style0
    class agent_xds_server_go style3
    class agent_auto_config_auto_encrypt_go style0
    class agent_checks_os_service_unix_go style0
    class agent_grpc_external_services_resource_mock_Registry_go style0
    class agent_xds_accesslogs_accesslogs_go style0
    class agent_config_ratelimited_file_watcher_go style0
    class agent_grpc_external_utils_go style0
    class agent_hcp_client_http_client_go style0
    class agent_structs_dns_go style0
    class agent_structs_discovery_chain_go style0
    class agent_apiserver_go style1
    class agent_consul_controller_reconciler_go style0
    class agent_proxycfg_testing_upstreams_go style0
    class api_acl_go style1
    class agent_consul_state_catalog_ce_go style0
    class agent_grpc_external_forward_go style0
    class agent_proxycfg_glue_peering_list_go style0
    class agent_uiserver_buf_index_fs_go style0
    class api_discovery_chain_go style1
    class agent_cache_types_config_entry_go style0
    class agent_configentry_compare_go style0
    class agent_proxycfg_glue_federation_state_list_mesh_gateways_go style0
    class agent_structs_intention_go style1
    class acl_policy_merger_go style2
    class agent_grpc_external_services_peerstream_mock_ACLResolver_go style0
    class api_operator_audit_go style1
    class acl_acl_go style2
    class agent_consul_discoverychain_compile_go style0
    class connect_proxy_conn_go style3
    class agent_cache_type_go style0
    class agent_consul_state_config_entry_go style0
    class agent_envoyextensions_registered_extensions_go style0
    class agent_grpc_internal_balancer_registry_go style0
    class agent_connect_ca_provider_vault_go style0
    class agent_grpc_external_options_go style0
    class agent_submatview_store_go style0
    class api_operator_segment_go style1
    class agent_cache_types_node_services_go style0
    class agent_catalog_endpoint_go style0
    class agent_grpc_external_services_connectca_watch_roots_go style0
    class agent_session_endpoint_go style0
    class agent_config_merge_go style0
    class agent_envoyextensions_builtin_aws_lambda_aws_lambda_go style0
    class agent_grpc_external_limiter_limiter_go style0
    class agent_grpc_external_services_resource_mock_Backend_go style0
    class agent_structs_connect_proxy_config_go style0
    class agent_grpc_external_services_resource_watch_go style0
    class api_debug_go style1
    class agent_consul_state_query_ce_go style0
    class agent_proxycfg_sources_local_sync_go style0
    class agent_consul_controller_queue_queue_go style0
    class agent_rpc_peering_testing_go style0
    class agent_consul_state_config_entry_intention_ce_go style0
    class api_exported_services_go style1
    class connect_resolver_go style3
    class agent_auto_config_auto_config_go style0
    class agent_consul_operator_autopilot_endpoint_go style0
    class agent_grpc_external_services_dataplane_mock_ACLResolver_go style0
    class agent_proxycfg_glue_resolved_service_config_go style0
    class agent_consul_acl_server_ce_go style0
    class agent_consul_peering_backend_go style0
    class agent_grpc_external_services_resource_testing_testing_ce_go style0
    class agent_proxycfg_data_sources_ce_go style0
    class agent_rpcclient_configentry_configentry_go style0
    class api_operator_autopilot_go style1
    class agent_consul_enterprise_client_ce_go style0
    class agent_envoyextensions_builtin_property_override_structpatcher_go style0
    class agent_keyring_go style0
    class api_config_entry_mesh_go style1
    class agent_config_agent_limits_go style0
    class agent_proxycfg_mesh_gateway_go style0
    class agent_structs_config_entry_go style1
    class agent_hcp_telemetry_otel_exporter_go style0
    class agent_submatview_materializer_go style0
    class agent_rpc_peering_service_go style0
    class agent_structs_envoy_extension_go style0
    class agent_structs_structs_deepcopy_ce_go style0
    class agent_hcp_bootstrap_bootstrap_go style0
    class api_event_go style1
    class agent_ae_ae_go style0
    class agent_consul_enterprise_server_ce_go style0
    class agent_grpc_external_services_peerstream_subscription_state_go style0
    class agent_agent_endpoint_go style1
    class agent_checks_check_go style0
    class agent_consul_authmethod_testing_go style0
    class agent_consul_leader_federation_state_ae_go style0
    class agent_consul_state_query_go style0
    class agent_token_persistence_go style0
    class agent_cache_types_trust_bundles_go style0
    class agent_consul_auth_login_go style0
    class agent_consul_watch_server_local_go style0
    class agent_proxycfg_glue_glue_go style0
    class agent_agent_endpoint_ce_go style0
    class agent_checks_os_service_windows_go style0
    class agent_hcp_discover_discover_go style0
    class agent_xds_configfetcher_config_fetcher_go style0
    class connect_tls_go style3
    class agent_consul_state_delay_ce_go style0
    class agent_grpc_middleware_rate_limit_mappings_gen_go style0
    class agent_xds_protocol_trace_go style0
    class api_operator_utilization_go style1
    class agent_consul_connect_ca_endpoint_go style0
    class agent_consul_status_endpoint_go style0
    class agent_setup_ce_go style0
    class agent_structs_protobuf_compat_go style0
    class agent_consul_state_graveyard_ce_go style0
    class agent_grpc_external_services_resource_testing_testing_go style0
    class agent_grpc_internal_client_go style0
    class agent_xds_clusters_go style0
    class agent_cache_types_peered_upstreams_go style0
    class agent_consul_state_federation_state_go style0
    class agent_grpc_external_services_peerstream_subscription_manager_go style0
    class agent_proxycfg_testing_terminating_gateway_go style0
    class agent_proxycfg_api_gateway_ce_go style0
    class agent_retry_join_go style0
    class agent_consul_acl_client_go style0
    class agent_consul_server_connect_go style0
    class agent_proxycfg_glue_intention_upstreams_go style0
    class agent_rpc_middleware_interceptors_go style0
    class agent_structs_autopilot_go style0
    class connect_service_go style3
    class agent_consul_multilimiter_mock_RateLimiter_go style0
    class agent_consul_state_tombstone_gc_go style0
    class agent_consul_stream_event_go style0
    class agent_proxycfg_glue_health_blocking_go style0
    class agent_consul_options_go style0
    class agent_proxycfg_api_gateway_go style0
    class agent_proxycfg_glue_peered_upstreams_go style0
    class api_config_entry_exports_go style1
    class acl_policy_authorizer_go style2
    class agent_consul_flood_go style0
    class agent_translate_addr_go style0
    class agent_grpc_external_services_connectca_sign_go style0
    class agent_structs_structs_ce_go style0
    class agent_cache_types_catalog_service_list_go style0
    class agent_connect_ca_provider_vault_auth_gcp_go style0
    class agent_connect_sni_go style0
    class agent_consul_authmethod_awsauth_aws_go style0
    class agent_consul_controller_controller_go style0
    class agent_health_endpoint_go style0
    class agent_structs_config_entry_sameness_group_ce_go style0
    class agent_auto_config_tls_go style0
    class agent_consul_txn_endpoint_go style0
    class agent_grpc_middleware_testutil_fake_sink_go style0
    class acl_authorizer_go style2
    class agent_auto_config_server_addr_go style0
    class agent_cache_types_health_services_go style0
    class agent_config_config_ce_go style0
    class agent_consul_system_metadata_go style0
    class agent_grpc_external_services_dataplane_get_envoy_bootstrap_params_go style0
    class agent_hcp_mock_TelemetryProvider_go style0
    class agent_structs_config_entry_jwt_provider_ce_go style0
    class agent_uiserver_redirect_fs_go style0
    class agent_consul_rate_mock_RequestLimitsHandler_go style0
    class api_operator_keyring_go style1
    class agent_consul_leader_connect_ca_ce_go style0
    class agent_consul_prepared_query_walk_go style0
    class agent_consul_state_events_go style0
    class agent_structs_config_entry_discoverychain_ce_go style0
    class agent_consul_state_catalog_events_ce_go style0
    class agent_connect_ca_provider_go style0
    class agent_connect_uri_go style0
    class agent_consul_leader_intentions_go style0
    class agent_leafcert_leafcert_test_helpers_go style0
    class agent_netutil_network_go style0
    class agent_rpcclient_common_go style0
    class agent_envoyextensions_builtin_wasm_wasm_go style0
    class agent_hcp_config_mock_CloudConfig_go style0
    class agent_nodeid_go style0
    class types_checks_go style1
    class agent_rpcclient_configentry_view_go style0
    class agent_xds_testing_go style0
    class agent_configentry_merge_service_config_go style0
    class agent_consul_gateway_locator_go style0
    class api_config_entry_sameness_group_go style1
    class agent_cache_request_go style0
    class agent_grpc_external_testutils_server_go style0
    class agent_consul_coordinate_endpoint_go style0
    class agent_grpc_middleware_testutil_testservice_fake_service_go style0
    class agent_xds_locality_policy_ce_go style0
    class agent_cache_types_options_go style0
    class agent_cache_types_testing_go style0
    class agent_grpc_middleware_auth_interceptor_go style0
    class agent_proxycfg_proxycfg_go style0
    class agent_proxycfg_sources_local_config_source_go style0
    class api_txn_go style1
    class agent_consul_discoverychain_string_stack_go style0
    class agent_grpc_external_services_connectca_mock_ACLResolver_go style0
    class agent_cache_testing_go style0
    class agent_consul_operator_backend_go style0
    class agent_consul_subscribe_backend_go style0
    class agent_rpcclient_health_view_go style0
    class agent_structs_service_definition_go style1
    class acl_acl_ce_go style2
    class agent_catalog_endpoint_ce_go style0
    class agent_consul_authmethod_authmethods_ce_go style0
    class agent_consul_acl_token_exp_go style0
    class agent_structs_testing_service_definition_go style0
    class agent_consul_state_graveyard_go style0
    class agent_grpc_external_services_serverdiscovery_watch_servers_go style0
    class api_operator_license_go style1
    class agent_consul_auth_token_writer_ce_go style0
    class agent_consul_state_peering_ce_go style0
    class agent_structs_config_entry_mesh_ce_go style0
    class agent_consul_federation_state_replication_go style0
    class agent_consul_state_acl_go style0
    class agent_grpc_external_querymeta_go style0
    class acl_authorizer_ce_go style2
    class agent_config_doc_go style0
    class agent_hcp_mock_Manager_go style0
    class agent_grpc_external_testutils_fsm_go style0
    class agent_cache_types_catalog_services_go style0
    class agent_consul_server_metadata_go style0
    class agent_blockingquery_mock_ResponseMeta_go style0
    class agent_cache_types_service_checks_go style0
    class agent_consul_reporting_reportingmock_mock_ServerDelegate_go style0
    class agent_consul_replication_go style0
    class agent_grpc_external_services_resource_mock_TenancyBridge_go style0
    class agent_proxycfg_data_sources_go style0
    class agent_consul_config_ce_go style0
    class agent_consul_fsm_decode_downgrade_go style0
    class agent_consul_server_lookup_go style0
    class agent_leafcert_structs_go style0
    class agent_xds_naming_naming_go style0
    class agent_structs_acl_templated_policy_go style0
    class agent_util_go style0
    class agent_cache_types_rpc_go style0
    class agent_consul_prepared_query_endpoint_go style0
    class agent_consul_stream_noop_go style0
    class agent_config_flagset_go style0
    class agent_consul_authmethod_testauth_testing_go style0
    class agent_grpc_external_services_configentry_server_go style0
    class agent_structs_structs_go style1
    class agent_connect_ca_provider_vault_auth_approle_go style0
    class agent_consul_authmethod_kubeauth_k8s_go style0
    class agent_consul_reporting_reporting_ce_go style0
    class agent_consul_state_catalog_schema_deepcopy_go style0
    class agent_structs_config_entry_exports_go style0
    class agent_consul_federation_state_endpoint_go style0
    class agent_pool_peek_go style0
    class agent_xds_platform_net_linux_go style0
    class agent_config_segment_ce_go style0
    class agent_hcp_client_client_go style0
    class api_connect_intention_go style1
    class agent_cache_types_service_dump_go style0
    class agent_consul_logging_go style0
    class agent_consul_merge_ce_go style0
    class agent_consul_v2_config_entry_exports_shim_go style0
    class agent_hcp_telemetry_custom_metrics_go style0
    class agent_consul_authmethod_ssoauth_sso_go style0
    class agent_consul_state_usage_go style0
    class agent_leafcert_signer_netrpc_go style0
    class agent_log_drop_log_drop_go style0
    class agent_uiserver_uiserver_go style0
    class agent_connect_uri_signing_go style0
    class agent_hcp_telemetry_otel_sink_go style0
    class agent_connect_ca_testing_go style0
    class agent_consul_discoverychain_compile_ce_go style0
    class agent_http_go style0
    class agent_connect_common_names_go style0
    class agent_structs_prepared_query_go style0
    class agent_consul_state_prepared_query_go style0
    class agent_structs_config_entry_status_go style0
    class agent_auto_config_persist_go style0
    class agent_auto_config_config_translate_go style0
    class agent_exec_exec_go style0
    class agent_proxycfg_sources_local_local_go style0
    class agent_consul_auto_config_endpoint_go style0
    class agent_consul_state_config_entry_sameness_group_ce_go style0
    class agent_consul_stats_fetcher_go style0
    class agent_consul_tenancy_bridge_go style0
    class agent_grpc_internal_balancer_balancer_go style0
    class agent_pool_conn_go style0
    class agent_acl_ce_go style0
    class agent_checks_grpc_go style0
    class agent_envoyextensions_builtin_wasm_structs_go style0
    class agent_config_limits_go style0
    class agent_consul_state_operations_ce_go style0
    class agent_structs_config_entry_apigw_jwt_ce_go style0
    class api_operator_go style1
    class agent_cache_mock_Request_go style0
    class agent_config_config_go style0
    class agent_grpc_external_services_resource_list_by_owner_go style0
    class agent_grpc_external_services_peerstream_testing_go style0
    class agent_grpc_external_services_resource_server_ce_go style0
    class agent_xds_extensionruntime_runtime_config_go style0
    class agent_structs_testing_go style0
    class agent_consul_fsm_decode_ce_go style0
    class agent_proxycfg_testing_mesh_gateway_go style0
    class api_config_entry_inline_certificate_go style1
    class agent_connect_ca_provider_consul_go style0
    class agent_consul_acl_ce_go style0
    class agent_consul_state_catalog_schema_go style0
    class agent_grpc_external_services_connectca_mock_CAManager_go style0
    class agent_proxycfg_ingress_gateway_go style0
    class agent_rpc_peering_validate_go style0
    class agent_structs_intention_ce_go style0
    class agent_leafcert_util_go style0
    class agent_rpc_operator_service_go style0
    class api_config_entry_gateways_go style1
    class agent_cache_types_resolved_service_config_go style0
    class agent_debug_host_go style0
    class agent_envoyextensions_builtin_lua_lua_go style0
    class agent_hcp_bootstrap_testing_go style0
    class agent_structs_catalog_go style0
    class agent_check_go style0
    class agent_consul_state_connect_ca_go style0
    class api_kv_go style1
    class agent_cache_types_exported_peered_services_go style0
    class agent_consul_session_ttl_go style0
    class agent_consul_state_config_entry_exported_services_ce_go style0
    class agent_grpc_external_services_acl_server_go style0
    class agent_token_store_ce_go style0
    class agent_config_default_go style0
    class agent_consul_authmethod_authmethods_go style0
    class agent_consul_controller_doc_go style0
    class agent_consul_state_system_metadata_go style0
    class agent_consul_state_schema_go style0
    class agent_cache_types_catalog_list_services_go style0
    class agent_cache_types_federation_state_list_gateways_go style0
    class agent_consul_state_acl_schema_go style0
    class agent_consul_state_mock_publishFuncType_go style0
    class agent_grpc_external_services_connectca_server_go style0
    class agent_grpc_external_services_resource_read_go style0
    class agent_agent_ce_go style0
    class agent_coordinate_endpoint_go style0
    class agent_envoyextensions_registered_extensions_ce_go style0
    class agent_consul_state_config_entry_schema_go style0
    class agent_submatview_rpc_materializer_go style0
    class agent_txn_endpoint_go style0

```
