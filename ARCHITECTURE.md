## Consul Module Distribution

```mermaid
flowchart LR
    command_connect_ca_ca_go["command/connect/ca/ca.go"
    connect_proxy_config_go["connect/proxy/config.go"
    grpcmocks_proto_public_pbdns_mock_DNSServiceServer_go["grpcmocks/proto-public/pbdns/mock_DNSServiceServer.go"
    agent_consul_state_session_go["agent/consul/state/session.go"
    agent_hcp_client_errors_go["agent/hcp/client/errors.go"
    agent_structs_config_entry_exports_go["agent/structs/config_entry_exports.go"
    command_flags_flag_slice_value_go["command/flags/flag_slice_value.go"
    command_kv_put_kv_put_go["command/kv/put/kv_put.go"
    internal_go_sso_oidcauth_util_go["internal/go-sso/oidcauth/util.go"
    agent_configentry_compare_go["agent/configentry/compare.go"
    agent_grpc_external_services_peerstream_server_go["agent/grpc-external/services/peerstream/server.go"
    agent_grpc_middleware_testutil_testservice_fake_service_go["agent/grpc-middleware/testutil/testservice/fake_service.go"
    agent_proxycfg_glue_glue_go["agent/proxycfg-glue/glue.go"
    agent_structs_acl_templated_policy_go["agent/structs/acl_templated_policy.go"
    agent_consul_enterprise_server_ce_go["agent/consul/enterprise_server_ce.go"
    agent_envoyextensions_builtin_wasm_structs_go["agent/envoyextensions/builtin/wasm/structs.go"
    agent_prepared_query_endpoint_go["agent/prepared_query_endpoint.go"
    agent_proxycfg_state_go["agent/proxycfg/state.go"
    internal_controller_controllermock_mock_CacheIDModifier_go["internal/controller/controllermock/mock_CacheIDModifier.go"
    internal_multicluster_internal_types_types_go["internal/multicluster/internal/types/types.go"
    testing_deployer_topology_relationships_go["testing/deployer/topology/relationships.go"
    agent_cache_types_catalog_services_go["agent/cache-types/catalog_services.go"
    agent_consul_health_endpoint_ce_go["agent/consul/health_endpoint_ce.go"
    agent_consul_state_query_ce_go["agent/consul/state/query_ce.go"
    agent_hcp_scada_scada_go["agent/hcp/scada/scada.go"
    agent_proxycfg_sources_local_mock_ConfigManager_go["agent/proxycfg-sources/local/mock_ConfigManager.go"
    command_acl_acl_go["command/acl/acl.go"
    command_acl_policy_policy_go["command/acl/policy/policy.go"
    envoyextensions_extensioncommon_basic_envoy_extender_go["envoyextensions/extensioncommon/basic_envoy_extender.go"
    agent_consul_controller_controller_go["agent/consul/controller/controller.go"
    agent_peering_endpoint_go["agent/peering_endpoint.go"
    command_intention_check_check_go["command/intention/check/check.go"
    internal_controller_cache_cachemock_mock_Query_go["internal/controller/cache/cachemock/mock_Query.go"
    grpcmocks_proto_public_pbdns_mock_DNSServiceClient_go["grpcmocks/proto-public/pbdns/mock_DNSServiceClient.go"
    lib_eof_go["lib/eof.go"
    proto_public_annotations_ratelimit_ratelimit_json_gen_go["proto-public/annotations/ratelimit/ratelimit_json.gen.go"
    agent_consul_operator_backend_go["agent/consul/operator_backend.go"
    agent_grpc_internal_pipe_go["agent/grpc-internal/pipe.go"
    command_connect_expose_expose_go["command/connect/expose/expose.go"
    proto_public_pbmulticluster_v2_partition_exported_services_json_gen_go["proto-public/pbmulticluster/v2/partition_exported_services_json.gen.go"
    agent_check_go["agent/check.go"
    agent_connect_csr_go["agent/connect/csr.go"
    agent_hcp_config_mock_CloudConfig_go["agent/hcp/config/mock_CloudConfig.go"
    lib_ttlcache_eviction_go["lib/ttlcache/eviction.go"
    proto_private_pbpeerstream_convert_go["proto/private/pbpeerstream/convert.go"
    proto_public_pbmulticluster_v2_exported_services_consumer_json_gen_go["proto-public/pbmulticluster/v2/exported_services_consumer_json.gen.go"
    test_integration_consul_container_libs_service_service_go["test/integration/consul-container/libs/service/service.go"
    agent_connect_common_names_go["agent/connect/common_names.go"
    agent_consul_server_overview_go["agent/consul/server_overview.go"
    agent_event_endpoint_go["agent/event_endpoint.go"
    agent_proxycfg_glue_service_http_checks_go["agent/proxycfg-glue/service_http_checks.go"
    agent_submatview_store_go["agent/submatview/store.go"
    command_acl_token_list_token_list_go["command/acl/token/list/token_list.go"
    command_kv_exp_kv_export_go["command/kv/exp/kv_export.go"
    proto_private_pbconfigentry_config_entry_grpc_pb_go["proto/private/pbconfigentry/config_entry_grpc.pb.go"
    agent_consul_gateways_controller_gateways_ce_go["agent/consul/gateways/controller_gateways_ce.go"
    agent_uiserver_ui_template_data_go["agent/uiserver/ui_template_data.go"
    command_acl_role_delete_role_delete_go["command/acl/role/delete/role_delete.go"
    internal_resource_tenancy_go["internal/resource/tenancy.go"
    command_connect_ca_get_connect_ca_get_go["command/connect/ca/get/connect_ca_get.go"
    command_config_read_config_read_go["command/config/read/config_read.go"
    envoyextensions_extensioncommon_upstream_envoy_extender_go["envoyextensions/extensioncommon/upstream_envoy_extender.go"
    internal_controller_cache_cachemock_mock_ResourceIterator_go["internal/controller/cache/cachemock/mock_ResourceIterator.go"
    internal_controller_dependency_cache_go["internal/controller/dependency/cache.go"
    testing_deployer_sprawl_peering_go["testing/deployer/sprawl/peering.go"
    agent_consul_server_lookup_go["agent/consul/server_lookup.go"
    agent_consul_state_delay_ce_go["agent/consul/state/delay_ce.go"
    agent_grpc_external_services_acl_server_go["agent/grpc-external/services/acl/server.go"
    agent_connect_uri_mesh_gateway_go["agent/connect/uri_mesh_gateway.go"
    agent_consul_internal_endpoint_go["agent/consul/internal_endpoint.go"
    agent_grpc_middleware_recovery_go["agent/grpc-middleware/recovery.go"
    agent_hcp_mock_Manager_go["agent/hcp/mock_Manager.go"
    agent_structs_auto_encrypt_go["agent/structs/auto_encrypt.go"
    agent_xds_locality_policy_go["agent/xds/locality_policy.go"
    command_connect_redirecttraffic_redirect_traffic_go["command/connect/redirecttraffic/redirect_traffic.go"
    internal_controller_manager_go["internal/controller/manager.go"
    agent_health_endpoint_go["agent/health_endpoint.go"
    agent_proxycfg_sources_local_config_source_go["agent/proxycfg-sources/local/config_source.go"
    api_namespace_go["api/namespace.go"
    command_lock_util_unix_go["command/lock/util_unix.go"
    internal_resource_mappers_bimapper_bimapper_go["internal/resource/mappers/bimapper/bimapper.go"
    lib_map_walker_go["lib/map_walker.go"
    proto_private_pbservice_service_gen_go["proto/private/pbservice/service.gen.go"
    agent_configentry_discoverychain_go["agent/configentry/discoverychain.go"
    agent_dns_ce_go["agent/dns_ce.go"
    agent_proxycfg_data_sources_ce_go["agent/proxycfg/data_sources_ce.go"
    agent_uiserver_buffered_file_go["agent/uiserver/buffered_file.go"
    test_integ_peering_commontopo_commontopo_go["test-integ/peering_commontopo/commontopo.go"
    agent_config_runtime_go["agent/config/runtime.go"
    agent_consul_logging_go["agent/consul/logging.go"
    agent_proxycfg_terminating_gateway_go["agent/proxycfg/terminating_gateway.go"
    command_resource_client_helper_go["command/resource/client/helper.go"
    agent_config_limits_windows_go["agent/config/limits_windows.go"
    agent_consul_server_grpc_go["agent/consul/server_grpc.go"
    agent_consul_state_config_entry_schema_go["agent/consul/state/config_entry_schema.go"
    internal_testing_golden_golden_go["internal/testing/golden/golden.go"
    testing_deployer_sprawl_acl_rules_go["testing/deployer/sprawl/acl_rules.go"
    agent_consul_config_go["agent/consul/config.go"
    agent_consul_xdscapacity_capacity_go["agent/consul/xdscapacity/capacity.go"
    command_operator_utilization_utilization_go["command/operator/utilization/utilization.go"
    internal_gossip_librtt_rtt_go["internal/gossip/librtt/rtt.go"
    internal_resource_authz_ce_go["internal/resource/authz_ce.go"
    internal_multicluster_internal_controllers_v1compat_mock_AggregatedConfig_go["internal/multicluster/internal/controllers/v1compat/mock_AggregatedConfig.go"
    sdk_testutil_retry_timer_go["sdk/testutil/retry/timer.go"
    testing_deployer_topology_topology_go["testing/deployer/topology/topology.go"
    agent_cache_types_service_dump_go["agent/cache-types/service_dump.go"
    agent_xds_jwt_authn_go["agent/xds/jwt_authn.go"
    api_config_entry_file_system_certificate_go["api/config_entry_file_system_certificate.go"
    command_intention_intention_go["command/intention/intention.go"
    internal_controller_controllermock_mock_Reconciler_go["internal/controller/controllermock/mock_Reconciler.go"
    internal_tools_protoc_gen_grpc_clone_e2e_proto_service_grpc_pb_go["internal/tools/protoc-gen-grpc-clone/e2e/proto/service_grpc.pb.go"
    proto_public_pbdns_dns_deepcopy_gen_go["proto-public/pbdns/dns_deepcopy.gen.go"
    agent_envoyextensions_builtin_lua_lua_go["agent/envoyextensions/builtin/lua/lua.go"
    agent_hcp_discover_discover_go["agent/hcp/discover/discover.go"
    agent_proxycfg_data_sources_go["agent/proxycfg/data_sources.go"
    types_checks_go["types/checks.go"
    testing_deployer_sprawl_internal_tfgen_proxy_go["testing/deployer/sprawl/internal/tfgen/proxy.go"
    agent_consul_connect_ca_endpoint_go["agent/consul/connect_ca_endpoint.go"
    agent_consul_state_coordinate_ce_go["agent/consul/state/coordinate_ce.go"
    agent_consul_state_intention_ce_go["agent/consul/state/intention_ce.go"
    agent_consul_state_config_entry_intention_go["agent/consul/state/config_entry_intention.go"
    agent_retry_join_go["agent/retry_join.go"
    command_acl_authmethod_list_authmethod_list_go["command/acl/authmethod/list/authmethod_list.go"
    agent_checks_docker_windows_go["agent/checks/docker_windows.go"
    agent_metadata_server_go["agent/metadata/server.go"
    agent_pool_conn_go["agent/pool/conn.go"
    command_acl_role_role_go["command/acl/role/role.go"
    command_kv_imp_kv_import_go["command/kv/imp/kv_import.go"
    command_operator_usage_usage_go["command/operator/usage/usage.go"
    internal_protohcl_doc_go["internal/protohcl/doc.go"
    proto_private_pbdemo_v2_demo_pb_go["proto/private/pbdemo/v2/demo.pb.go"
    api_coordinate_go["api/coordinate.go"
    envoyextensions_xdscommon_xdscommon_go["envoyextensions/xdscommon/xdscommon.go"
    internal_controller_cache_index_indexmock_mock_ResourceIterator_go["internal/controller/cache/index/indexmock/mock_ResourceIterator.go"
    proto_public_pbdns_dns_cloning_grpc_pb_go["proto-public/pbdns/dns_cloning_grpc.pb.go"
    sdk_freeport_systemlimit_go["sdk/freeport/systemlimit.go"
    agent_structs_testing_go["agent/structs/testing.go"
    sdk_testutil_server_wrapper_go["sdk/testutil/server_wrapper.go"
    agent_rpcclient_common_go["agent/rpcclient/common.go"
    agent_xds_naming_naming_go["agent/xds/naming/naming.go"
    proto_public_pbresource_resource_pb_go["proto-public/pbresource/resource.pb.go"
    agent_grpc_external_options_go["agent/grpc-external/options.go"
    command_peering_exportedservices_exported_services_go["command/peering/exportedservices/exported_services.go"
    grpcmocks_proto_public_pbdataplane_mock_isGetEnvoyBootstrapParamsRequest_NodeSpec_go["grpcmocks/proto-public/pbdataplane/mock_isGetEnvoyBootstrapParamsRequest_NodeSpec.go"
    agent_consul_leader_peering_go["agent/consul/leader_peering.go"
    agent_structs_config_entry_mesh_ce_go["agent/structs/config_entry_mesh_ce.go"
    api_catalog_go["api/catalog.go"
    command_catalog_catalog_go["command/catalog/catalog.go"
    testing_deployer_sprawl_internal_secrets_store_go["testing/deployer/sprawl/internal/secrets/store.go"
    types_node_id_go["types/node_id.go"
    command_exec_exec_go["command/exec/exec.go"
    command_kv_get_kv_get_go["command/kv/get/kv_get.go"
    agent_consul_authmethod_testing_go["agent/consul/authmethod/testing.go"
    agent_consul_serf_filter_go["agent/consul/serf_filter.go"
    agent_grpc_external_services_resource_mock_TenancyBridge_go["agent/grpc-external/services/resource/mock_TenancyBridge.go"
    agent_token_persistence_go["agent/token/persistence.go"
    api_config_entry_go["api/config_entry.go"
    proto_private_pboperator_operator_pb_go["proto/private/pboperator/operator.pb.go"
    acl_policy_go["acl/policy.go"
    agent_auto_config_tls_go["agent/auto-config/tls.go"
    agent_consul_prepared_query_template_go["agent/consul/prepared_query/template.go"
    agent_consul_state_memdb_go["agent/consul/state/memdb.go"
    agent_grpc_external_services_resource_watch_go["agent/grpc-external/services/resource/watch.go"
    agent_hcp_scada_mock_Provider_go["agent/hcp/scada/mock_Provider.go"
    agent_structs_service_definition_go["agent/structs/service_definition.go"
    proto_public_pbacl_acl_json_gen_go["proto-public/pbacl/acl_json.gen.go"
    agent_consul_snapshot_endpoint_go["agent/consul/snapshot_endpoint.go"
    agent_leafcert_structs_go["agent/leafcert/structs.go"
    agent_structs_discovery_chain_ce_go["agent/structs/discovery_chain_ce.go"
    agent_token_store_ce_go["agent/token/store_ce.go"
    agent_uiserver_redirect_fs_go["agent/uiserver/redirect_fs.go"
    command_flags_merge_go["command/flags/merge.go"
    command_forceleave_forceleave_go["command/forceleave/forceleave.go"
    internal_controller_cache_indexers_indexersmock_mock_SingleIndexer_go["internal/controller/cache/indexers/indexersmock/mock_SingleIndexer.go"
    agent_checks_os_service_go["agent/checks/os_service.go"
    agent_consul_state_acl_events_go["agent/consul/state/acl_events.go"
    proto_private_pboperator_operator_gen_go["proto/private/pboperator/operator.gen.go"
    sdk_testutil_retry_counter_go["sdk/testutil/retry/counter.go"
    test_integration_consul_container_libs_topology_service_topology_go["test/integration/consul-container/libs/topology/service_topology.go"
    testing_deployer_sprawl_configentries_go["testing/deployer/sprawl/configentries.go"
    agent_connect_uri_agent_ce_go["agent/connect/uri_agent_ce.go"
    logging_syslog_go["logging/syslog.go"
    proto_private_pbautoconf_auto_config_go["proto/private/pbautoconf/auto_config.go"
    service_os_service_go["service_os/service.go"
    troubleshoot_ports_hostport_go["troubleshoot/ports/hostport.go"
    agent_consul_controller_reconciler_go["agent/consul/controller/reconciler.go"
    agent_consul_health_endpoint_go["agent/consul/health_endpoint.go"
    agent_consul_stream_subscription_go["agent/consul/stream/subscription.go"
    agent_rpc_peering_validate_go["agent/rpc/peering/validate.go"
    api_config_entry_sameness_group_go["api/config_entry_sameness_group.go"
    command_connect_envoy_exec_go["command/connect/envoy/exec.go"
    proto_private_pbservice_healthcheck_gen_go["proto/private/pbservice/healthcheck.gen.go"
    agent_consul_acl_server_go["agent/consul/acl_server.go"
    agent_grpc_internal_balancer_registry_go["agent/grpc-internal/balancer/registry.go"
    agent_hcp_telemetry_otlp_transform_go["agent/hcp/telemetry/otlp_transform.go"
    agent_hcp_config_config_go["agent/hcp/config/config.go"
    agent_log_drop_log_drop_go["agent/log-drop/log-drop.go"
    command_agent_agent_go["command/agent/agent.go"
    internal_multicluster_internal_types_computed_exported_services_go["internal/multicluster/internal/types/computed_exported_services.go"
    logging_logfile_go["logging/logfile.go"
    agent_consul_operator_autopilot_endpoint_go["agent/consul/operator_autopilot_endpoint.go"
    agent_structs_config_entry_inline_certificate_go["agent/structs/config_entry_inline_certificate.go"
    command_tls_ca_create_tls_ca_create_go["command/tls/ca/create/tls_ca_create.go"
    internal_controller_doc_go["internal/controller/doc.go"
    internal_multicluster_internal_types_namespace_exported_services_go["internal/multicluster/internal/types/namespace_exported_services.go"
    api_operator_utilization_go["api/operator_utilization.go"
    internal_go_sso_oidcauth_auth_go["internal/go-sso/oidcauth/auth.go"
    sdk_testutil_context_go["sdk/testutil/context.go"
    testing_deployer_sprawl_helpers_go["testing/deployer/sprawl/helpers.go"
    agent_consul_reporting_reporting_ce_go["agent/consul/reporting/reporting_ce.go"
    agent_consul_wanfed_pool_go["agent/consul/wanfed/pool.go"
    agent_proxycfg_glue_resolved_service_config_go["agent/proxycfg-glue/resolved_service_config.go"
    command_acl_authmethod_create_authmethod_create_ce_go["command/acl/authmethod/create/authmethod_create_ce.go"
    command_acl_policy_update_policy_update_go["command/acl/policy/update/policy_update.go"
    agent_consul_prepared_query_endpoint_go["agent/consul/prepared_query_endpoint.go"
    agent_dns_context_go["agent/dns/context.go"
    agent_grpc_external_services_resource_list_by_owner_go["agent/grpc-external/services/resource/list_by_owner.go"
    agent_grpc_external_services_peerstream_subscription_view_go["agent/grpc-external/services/peerstream/subscription_view.go"
    agent_structs_config_entry_jwt_provider_ce_go["agent/structs/config_entry_jwt_provider_ce.go"
    api_event_go["api/event.go"
    api_raw_go["api/raw.go"
    lib_channels_deliver_latest_go["lib/channels/deliver_latest.go"
    agent_cache_types_service_gateways_go["agent/cache-types/service_gateways.go"
    agent_config_default_ce_go["agent/config/default_ce.go"
    agent_consul_fsm_fsm_go["agent/consul/fsm/fsm.go"
    agent_hcp_bootstrap_config_loader_loader_go["agent/hcp/bootstrap/config-loader/loader.go"
    agent_structs_connect_ce_go["agent/structs/connect_ce.go"
    internal_controller_cache_indexers_indexersmock_mock_GetSingleRefOrID_go["internal/controller/cache/indexers/indexersmock/mock_GetSingleRefOrID.go"
    proto_private_pbconnect_connect_gen_go["proto/private/pbconnect/connect.gen.go"
    proto_private_pboperator_operator_grpc_pb_go["proto/private/pboperator/operator_grpc.pb.go"
    acl_acl_ce_go["acl/acl_ce.go"
    agent_consul_state_session_ce_go["agent/consul/state/session_ce.go"
    agent_consul_state_indexer_go["agent/consul/state/indexer.go"
    agent_proxycfg_glue_service_list_go["agent/proxycfg-glue/service_list.go"
    command_acl_policy_create_policy_create_go["command/acl/policy/create/policy_create.go"
    command_operator_autopilot_set_operator_autopilot_set_go["command/operator/autopilot/set/operator_autopilot_set.go"
    logging_logfile_solaris_go["logging/logfile_solaris.go"
    testing_deployer_sprawl_internal_tfgen_gen_go["testing/deployer/sprawl/internal/tfgen/gen.go"
    agent_cache_types_config_entry_go["agent/cache-types/config_entry.go"
    agent_consul_multilimiter_multilimiter_go["agent/consul/multilimiter/multilimiter.go"
    agent_proxycfg_glue_leafcerts_go["agent/proxycfg-glue/leafcerts.go"
    command_services_exportedservices_exported_services_go["command/services/exportedservices/exported_services.go"
    internal_resource_protoc_gen_json_shim_internal_generate_generate_go["internal/resource/protoc-gen-json-shim/internal/generate/generate.go"
    logging_logfile_windows_go["logging/logfile_windows.go"
    agent_consul_state_config_entry_exported_services_go["agent/consul/state/config_entry_exported_services.go"
    agent_structs_intention_ce_go["agent/structs/intention_ce.go"
    command_resource_client_grpc_resource_flags_go["command/resource/client/grpc-resource-flags.go"
    internal_resource_resourcetest_testing_go["internal/resource/resourcetest/testing.go"
    sdk_testutil_assertions_go["sdk/testutil/assertions.go"
    agent_consul_fsm_snapshot_go["agent/consul/fsm/snapshot.go"
    agent_consul_usagemetrics_usagemetrics_ce_go["agent/consul/usagemetrics/usagemetrics_ce.go"
    agent_grpc_external_services_dataplane_mock_ACLResolver_go["agent/grpc-external/services/dataplane/mock_ACLResolver.go"
    command_acl_authmethod_formatter_go["command/acl/authmethod/formatter.go"
    internal_controller_cache_index_iterator_go["internal/controller/cache/index/iterator.go"
    lib_hoststats_cpu_go["lib/hoststats/cpu.go"
    proto_public_pbmulticluster_v2_exported_services_consumer_pb_go["proto-public/pbmulticluster/v2/exported_services_consumer.pb.go"
    command_acl_templatedpolicy_list_templated_policy_list_go["command/acl/templatedpolicy/list/templated_policy_list.go"
    internal_storage_inmem_snapshot_go["internal/storage/inmem/snapshot.go"
    testing_deployer_sprawl_internal_tfgen_io_go["testing/deployer/sprawl/internal/tfgen/io.go"
    agent_consul_stream_event_go["agent/consul/stream/event.go"
    agent_structs_config_entry_intentions_ce_go["agent/structs/config_entry_intentions_ce.go"
    command_acl_templatedpolicy_formatter_go["command/acl/templatedpolicy/formatter.go"
    command_flags_config_go["command/flags/config.go"
    internal_resource_protoc_gen_resource_types_internal_generate_generate_go["internal/resource/protoc-gen-resource-types/internal/generate/generate.go"
    test_integration_consul_container_test_resource_http_api_client_client_go["test/integration/consul-container/test/resource/http_api/client/client.go"
    troubleshoot_validate_validate_go["troubleshoot/validate/validate.go"
    proto_public_pbacl_acl_cloning_grpc_pb_go["proto-public/pbacl/acl_cloning_grpc.pb.go"
    proto_public_pbserverdiscovery_serverdiscovery_json_gen_go["proto-public/pbserverdiscovery/serverdiscovery_json.gen.go"
    sdk_iptables_iptables_go["sdk/iptables/iptables.go"
    agent_grpc_external_server_go["agent/grpc-external/server.go"
    agent_remote_exec_go["agent/remote_exec.go"
    command_services_register_register_go["command/services/register/register.go"
    test_integration_consul_container_libs_cluster_container_go["test/integration/consul-container/libs/cluster/container.go"
    agent_acl_endpoint_go["agent/acl_endpoint.go"
    agent_hcp_scada_capabilities_go["agent/hcp/scada/capabilities.go"
    command_rtt_rtt_go["command/rtt/rtt.go"
    command_resource_resource_go["command/resource/resource.go"
    grpcmocks_proto_public_pbconnectca_mock_ConnectCAService_WatchRootsClient_go["grpcmocks/proto-public/pbconnectca/mock_ConnectCAService_WatchRootsClient.go"
    internal_multicluster_internal_types_helpers_ce_go["internal/multicluster/internal/types/helpers_ce.go"
    internal_resource_resourcetest_decode_go["internal/resource/resourcetest/decode.go"
    agent_consul_state_acl_go["agent/consul/state/acl.go"
    agent_grpc_external_testutils_mock_server_transport_stream_go["agent/grpc-external/testutils/mock_server_transport_stream.go"
    agent_structs_acl_templated_policy_ce_go["agent/structs/acl_templated_policy_ce.go"
    api_exported_services_go["api/exported_services.go"
    internal_go_sso_oidcauth_internal_strutil_util_go["internal/go-sso/oidcauth/internal/strutil/util.go"
    ipaddr_ipaddr_go["ipaddr/ipaddr.go"
    troubleshoot_proxy_z_xds_packages_go["troubleshoot/proxy/z_xds_packages.go"
    agent_grpc_external_services_peerstream_replication_go["agent/grpc-external/services/peerstream/replication.go"
    agent_hcp_bootstrap_constants_constants_go["agent/hcp/bootstrap/constants/constants.go"
    proto_public_pbresource_resource_json_gen_go["proto-public/pbresource/resource_json.gen.go"
    sdk_testutil_retry_doc_go["sdk/testutil/retry/doc.go"
    agent_consul_gateway_locator_go["agent/consul/gateway_locator.go"
    agent_consul_rate_handler_go["agent/consul/rate/handler.go"
    agent_consul_state_catalog_events_ce_go["agent/consul/state/catalog_events_ce.go"
    agent_grpc_external_querymeta_go["agent/grpc-external/querymeta.go"
    agent_xds_endpoints_go["agent/xds/endpoints.go"
    api_config_entry_exports_go["api/config_entry_exports.go"
    proto_private_pbpeerstream_types_go["proto/private/pbpeerstream/types.go"
    agent_consul_acl_endpoint_go["agent/consul/acl_endpoint.go"
    agent_hcp_mock_TelemetryProvider_go["agent/hcp/mock_TelemetryProvider.go"
    agent_proxycfg_naming_go["agent/proxycfg/naming.go"
    command_catalog_list_dc_catalog_list_datacenters_go["command/catalog/list/dc/catalog_list_datacenters.go"
    command_resource_client_grpc_client_go["command/resource/client/grpc-client.go"
    agent_checks_grpc_go["agent/checks/grpc.go"
    agent_consul_operator_endpoint_go["agent/consul/operator_endpoint.go"
    agent_grpc_external_services_acl_mock_Login_go["agent/grpc-external/services/acl/mock_Login.go"
    command_resource_read_grpc_read_go["command/resource/read-grpc/read.go"
    command_snapshot_snapshot_command_go["command/snapshot/snapshot_command.go"
    grpcmocks_proto_public_pbresource_mock_ResourceService_WatchListServer_go["grpcmocks/proto-public/pbresource/mock_ResourceService_WatchListServer.go"
    test_integration_consul_container_libs_assert_grpc_go["test/integration/consul-container/libs/assert/grpc.go"
    agent_auto_config_auto_config_ce_go["agent/auto-config/auto_config_ce.go"
    agent_config_default_go["agent/config/default.go"
    agent_structs_connect_go["agent/structs/connect.go"
    internal_controller_helper_go["internal/controller/helper.go"
    proto_private_pbcommon_common_gen_go["proto/private/pbcommon/common.gen.go"
    proto_private_pbservice_convert_ce_go["proto/private/pbservice/convert_ce.go"
    agent_cache_types_node_services_go["agent/cache-types/node_services.go"
    agent_denylist_go["agent/denylist.go"
    proto_private_pbconfigentry_config_entry_go["proto/private/pbconfigentry/config_entry.go"
    test_integration_consul_container_libs_utils_version_ce_go["test/integration/consul-container/libs/utils/version_ce.go"
    testing_deployer_sprawl_internal_tfgen_agent_go["testing/deployer/sprawl/internal/tfgen/agent.go"
    agent_consul_state_acl_schema_go["agent/consul/state/acl_schema.go"
    agent_grpc_external_services_peerstream_subscription_manager_go["agent/grpc-external/services/peerstream/subscription_manager.go"
    agent_grpc_external_services_resource_read_go["agent/grpc-external/services/resource/read.go"
    agent_grpc_middleware_auth_interceptor_go["agent/grpc-middleware/auth_interceptor.go"
    command_acl_policy_formatter_go["command/acl/policy/formatter.go"
    command_keyring_keyring_go["command/keyring/keyring.go"
    command_intention_helpers_go["command/intention/helpers.go"
    agent_connect_testing_ca_go["agent/connect/testing_ca.go"
    agent_consul_raft_rpc_go["agent/consul/raft_rpc.go"
    agent_federation_state_endpoint_go["agent/federation_state_endpoint.go"
    proto_private_pbsubscribe_subscribe_grpc_pb_go["proto/private/pbsubscribe/subscribe_grpc.pb.go"
    version_versiontest_versiontest_go["version/versiontest/versiontest.go"
    agent_consul_discoverychain_string_stack_go["agent/consul/discoverychain/string_stack.go"
    agent_consul_status_endpoint_go["agent/consul/status_endpoint.go"
    agent_grpc_external_services_resource_server_go["agent/grpc-external/services/resource/server.go"
    agent_proxycfg_sources_catalog_mock_ConfigManager_go["agent/proxycfg-sources/catalog/mock_ConfigManager.go"
    api_config_entry_status_go["api/config_entry_status.go"
    internal_controller_cache_index_index_go["internal/controller/cache/index/index.go"
    testing_deployer_topology_generate_go["testing/deployer/topology/generate.go"
    agent_cache_entry_go["agent/cache/entry.go"
    agent_catalog_endpoint_ce_go["agent/catalog_endpoint_ce.go"
    agent_metrics_go["agent/metrics.go"
    agent_structs_catalog_ce_go["agent/structs/catalog_ce.go"
    command_acl_authmethod_create_authmethod_create_go["command/acl/authmethod/create/authmethod_create.go"
    agent_auto_config_server_addr_go["agent/auto-config/server_addr.go"
    agent_pool_pool_go["agent/pool/pool.go"
    api_discovery_chain_go["api/discovery_chain.go"
    proto_private_prototest_golden_json_go["proto/private/prototest/golden_json.go"
    agent_consul_stream_event_snapshot_go["agent/consul/stream/event_snapshot.go"
    agent_dns_buffer_response_writer_go["agent/dns/buffer_response_writer.go"
    agent_token_store_go["agent/token/store.go"
    agent_xds_config_config_go["agent/xds/config/config.go"
    agent_structs_structs_deepcopy_ce_go["agent/structs/structs.deepcopy_ce.go"
    api_config_entry_mesh_go["api/config_entry_mesh.go"
    grpcmocks_proto_public_pbacl_mock_ACLServiceClient_go["grpcmocks/proto-public/pbacl/mock_ACLServiceClient.go"
    internal_protohcl_any_go["internal/protohcl/any.go"
    agent_consul_acl_authmethod_ce_go["agent/consul/acl_authmethod_ce.go"
    agent_consul_autopilotevents_ready_servers_events_go["agent/consul/autopilotevents/ready_servers_events.go"
    agent_setup_ce_go["agent/setup_ce.go"
    command_acl_role_update_role_update_go["command/acl/role/update/role_update.go"
    internal_resource_sort_go["internal/resource/sort.go"
    agent_agent_endpoint_ce_go["agent/agent_endpoint_ce.go"
    agent_connect_uri_mesh_gateway_ce_go["agent/connect/uri_mesh_gateway_ce.go"
    agent_consul_kvs_endpoint_go["agent/consul/kvs_endpoint.go"
    agent_debug_host_go["agent/debug/host.go"
    agent_xds_xds_go["agent/xds/xds.go"
    command_flags_usage_go["command/flags/usage.go"
    grpcmocks_proto_public_pbresource_mock_UnsafeResourceServiceServer_go["grpcmocks/proto-public/pbresource/mock_UnsafeResourceServiceServer.go"
    internal_controller_controllermock_mock_CustomDependencyMapper_go["internal/controller/controllermock/mock_CustomDependencyMapper.go"
    agent_consul_peering_backend_go["agent/consul/peering_backend.go"
    agent_consul_state_catalog_ce_go["agent/consul/state/catalog_ce.go"
    api_watch_plan_go["api/watch/plan.go"
    command_login_login_go["command/login/login.go"
    connect_proxy_listener_go["connect/proxy/listener.go"
    proto_public_pbserverdiscovery_serverdiscovery_pb_go["proto-public/pbserverdiscovery/serverdiscovery.pb.go"
    testing_deployer_sprawl_internal_tfgen_docker_go["testing/deployer/sprawl/internal/tfgen/docker.go"
    agent_http_register_go["agent/http_register.go"
    api_debug_go["api/debug.go"
    command_connect_envoy_flags_go["command/connect/envoy/flags.go"
    command_troubleshoot_troubleshoot_go["command/troubleshoot/troubleshoot.go"
    internal_controller_dependency_simple_go["internal/controller/dependency/simple.go"
    agent_connect_ca_provider_vault_go["agent/connect/ca/provider_vault.go"
    agent_discovery_chain_endpoint_go["agent/discovery_chain_endpoint.go"
    agent_proxycfg_api_gateway_ce_go["agent/proxycfg/api_gateway_ce.go"
    agent_rpc_middleware_interceptors_go["agent/rpc/middleware/interceptors.go"
    agent_setup_go["agent/setup.go"
    command_acl_authmethod_update_authmethod_update_go["command/acl/authmethod/update/authmethod_update.go"
    command_intention_create_create_go["command/intention/create/create.go"
    proto_private_pbpeerstream_peerstream_pb_go["proto/private/pbpeerstream/peerstream.pb.go"
    agent_consul_session_endpoint_go["agent/consul/session_endpoint.go"
    agent_xds_server_ce_go["agent/xds/server_ce.go"
    api_semaphore_go["api/semaphore.go"
    agent_envoyextensions_builtin_otel_access_logging_otel_access_logging_go["agent/envoyextensions/builtin/otel-access-logging/otel_access_logging.go"
    agent_structs_protobuf_compat_go["agent/structs/protobuf_compat.go"
    command_operator_autopilot_state_formatter_go["command/operator/autopilot/state/formatter.go"
    test_integration_consul_container_test_envoy_extensions_testdata_wasm_test_files_wasm_add_header_go["test/integration/consul-container/test/envoy_extensions/testdata/wasm_test_files/wasm_add_header.go"
    agent_grpc_external_services_peerstream_mock_ACLResolver_go["agent/grpc-external/services/peerstream/mock_ACLResolver.go"
    agent_hcp_testing_go["agent/hcp/testing.go"
    api_session_go["api/session.go"
    command_operator_autopilot_get_operator_autopilot_get_go["command/operator/autopilot/get/operator_autopilot_get.go"
    internal_tools_protoc_gen_consul_rate_limit_main_go["internal/tools/protoc-gen-consul-rate-limit/main.go"
    agent_consul_leader_metrics_go["agent/consul/leader_metrics.go"
    agent_consul_rate_metrics_go["agent/consul/rate/metrics.go"
    agent_consul_state_system_metadata_go["agent/consul/state/system_metadata.go"
    agent_grpc_external_services_serverdiscovery_server_go["agent/grpc-external/services/serverdiscovery/server.go"
    envoyextensions_extensioncommon_envoy_extender_go["envoyextensions/extensioncommon/envoy_extender.go"
    proto_public_pbdataplane_dataplane_json_gen_go["proto-public/pbdataplane/dataplane_json.gen.go"
    agent_consul_federation_state_endpoint_go["agent/consul/federation_state_endpoint.go"
    agent_consul_leader_connect_ca_go["agent/consul/leader_connect_ca.go"
    grpcmocks_proto_public_pbresource_mock_isWatchEvent_Event_go["grpcmocks/proto-public/pbresource/mock_isWatchEvent_Event.go"
    testing_deployer_util_internal_ipamutils_doc_go["testing/deployer/util/internal/ipamutils/doc.go"
    agent_consul_state_acl_ce_go["agent/consul/state/acl_ce.go"
    tlsutil_config_go["tlsutil/config.go"
    agent_consul_autopilot_go["agent/consul/autopilot.go"
    agent_structs_snapshot_go["agent/structs/snapshot.go"
    grpcmocks_proto_public_pbacl_mock_ACLServiceServer_go["grpcmocks/proto-public/pbacl/mock_ACLServiceServer.go"
    grpcmocks_proto_public_pbdataplane_mock_DataplaneServiceClient_go["grpcmocks/proto-public/pbdataplane/mock_DataplaneServiceClient.go"
    internal_resource_reference_go["internal/resource/reference.go"
    testing_deployer_sprawl_debug_go["testing/deployer/sprawl/debug.go"
    agent_cache_types_gateway_services_go["agent/cache-types/gateway_services.go"
    agent_configentry_config_entry_go["agent/configentry/config_entry.go"
    agent_proxycfg_glue_internal_service_dump_go["agent/proxycfg-glue/internal_service_dump.go"
    agent_proxycfg_glue_intentions_go["agent/proxycfg-glue/intentions.go"
    command_catalog_list_services_catalog_list_services_go["command/catalog/list/services/catalog_list_services.go"
    lib_stop_context_go["lib/stop_context.go"
    proto_public_pbresource_resource_deepcopy_gen_go["proto-public/pbresource/resource_deepcopy.gen.go"
    agent_config_flagset_go["agent/config/flagset.go"
    agent_grpc_middleware_rate_go["agent/grpc-middleware/rate.go"
    agent_xds_resources_go["agent/xds/resources.go"
    agent_xds_response_response_go["agent/xds/response/response.go"
    snapshot_snapshot_go["snapshot/snapshot.go"
    agent_consul_watch_server_local_go["agent/consul/watch/server_local.go"
    agent_grpc_external_services_peerstream_subscription_blocking_go["agent/grpc-external/services/peerstream/subscription_blocking.go"
    agent_xds_listeners_ingress_go["agent/xds/listeners_ingress.go"
    proto_private_pbcommon_common_go["proto/private/pbcommon/common.go"
    agent_cacheshim_cache_go["agent/cacheshim/cache.go"
    agent_hcp_bootstrap_testing_go["agent/hcp/bootstrap/testing.go"
    agent_mock_notify_go["agent/mock/notify.go"
    version_fips_go["version/fips.go"
    agent_exec_exec_windows_go["agent/exec/exec_windows.go"
    agent_hcp_manager_go["agent/hcp/manager.go"
    command_config_write_config_write_go["command/config/write/config_write.go"
    command_resource_helper_go["command/resource/helper.go"
    internal_resource_protoc_gen_deepcopy_internal_generate_generate_go["internal/resource/protoc-gen-deepcopy/internal/generate/generate.go"
    agent_acl_go["agent/acl.go"
    agent_consul_stream_event_publisher_go["agent/consul/stream/event_publisher.go"
    agent_dns_go["agent/dns.go"
    command_tls_tls_go["command/tls/tls.go"
    agent_consul_discoverychain_compile_ce_go["agent/consul/discoverychain/compile_ce.go"
    agent_grpc_external_services_acl_mock_TokenWriter_go["agent/grpc-external/services/acl/mock_TokenWriter.go"
    api_acl_go["api/acl.go"
    troubleshoot_proxy_validateupstream_go["troubleshoot/proxy/validateupstream.go"
    agent_cache_mock_Request_go["agent/cache/mock_Request.go"
    agent_consul_multilimiter_mock_RateLimiter_go["agent/consul/multilimiter/mock_RateLimiter.go"
    command_helpers_helpers_go["command/helpers/helpers.go"
    lib_decode_decode_go["lib/decode/decode.go"
    proto_public_pbdns_dns_json_gen_go["proto-public/pbdns/dns_json.gen.go"
    testing_deployer_sprawl_internal_runner_exec_go["testing/deployer/sprawl/internal/runner/exec.go"
    agent_cache_mock_Type_go["agent/cache/mock_Type.go"
    command_watch_watch_go["command/watch/watch.go"
    acl_resolver_danger_go["acl/resolver/danger.go"
    agent_auto_config_auto_encrypt_go["agent/auto-config/auto_encrypt.go"
    agent_http_go["agent/http.go"
    agent_xds_server_go["agent/xds/server.go"
    agent_blockingquery_mock_RequestOptions_go["agent/blockingquery/mock_RequestOptions.go"
    agent_consul_merge_ce_go["agent/consul/merge_ce.go"
    agent_coordinate_endpoint_go["agent/coordinate_endpoint.go"
    agent_structs_autopilot_ce_go["agent/structs/autopilot_ce.go"
    api_lock_go["api/lock.go"
    command_catalog_helpers_go["command/catalog/helpers.go"
    command_event_event_go["command/event/event.go"
    command_operator_raft_operator_raft_go["command/operator/raft/operator_raft.go"
    agent_consul_wanfed_wanfed_go["agent/consul/wanfed/wanfed.go"
    agent_proxycfg_upstreams_go["agent/proxycfg/upstreams.go"
    agent_structs_autopilot_go["agent/structs/autopilot.go"
    internal_controller_cache_indexers_indexersmock_mock_RefOrIDFetcher_go["internal/controller/cache/indexers/indexersmock/mock_RefOrIDFetcher.go"
    internal_resource_hooks_go["internal/resource/hooks.go"
    internal_tools_protoc_gen_grpc_clone_e2e_mock_Simple_FlowClient_go["internal/tools/protoc-gen-grpc-clone/e2e/mock_Simple_FlowClient.go"
    proto_private_pbservice_ids_go["proto/private/pbservice/ids.go"
    testing_deployer_util_internal_ipamutils_utils_go["testing/deployer/util/internal/ipamutils/utils.go"
    agent_consul_state_connect_ca_go["agent/consul/state/connect_ca.go"
    agent_log_drop_mock_Logger_go["agent/log-drop/mock_Logger.go"
    agent_proxycfg_naming_ce_go["agent/proxycfg/naming_ce.go"
    test_integration_consul_container_libs_service_helpers_go["test/integration/consul-container/libs/service/helpers.go"
    test_integration_consul_container_libs_utils_version_go["test/integration/consul-container/libs/utils/version.go"
    troubleshoot_ports_troubleshoot_tcp_go["troubleshoot/ports/troubleshoot_tcp.go"
    agent_grpc_external_services_peerstream_stream_tracker_go["agent/grpc-external/services/peerstream/stream_tracker.go"
    agent_proxycfg_testing_peering_go["agent/proxycfg/testing_peering.go"
    grpcmocks_proto_public_pbresource_mock_ResourceServiceClient_go["grpcmocks/proto-public/pbresource/mock_ResourceServiceClient.go"
    proto_private_pbpeerstream_peerstream_grpc_pb_go["proto/private/pbpeerstream/peerstream_grpc.pb.go"
    test_integration_consul_container_libs_assert_service_go["test/integration/consul-container/libs/assert/service.go"
    agent_connect_ca_provider_vault_auth_jwt_go["agent/connect/ca/provider_vault_auth_jwt.go"
    command_acl_token_read_token_read_go["command/acl/token/read/token_read.go"
    sdk_freeport_ephemeral_linux_go["sdk/freeport/ephemeral_linux.go"
    agent_cache_types_rpc_go["agent/cache-types/rpc.go"
    agent_consul_rtt_go["agent/consul/rtt.go"
    agent_grpc_external_services_resource_mock_Registry_go["agent/grpc-external/services/resource/mock_Registry.go"
    agent_structs_config_entry_ce_go["agent/structs/config_entry_ce.go"
    agent_structs_config_entry_gateways_go["agent/structs/config_entry_gateways.go"
    api_watch_funcs_go["api/watch/funcs.go"
    connect_service_go["connect/service.go"
    internal_go_sso_oidcauth_oidcauthtest_testing_go["internal/go-sso/oidcauth/oidcauthtest/testing.go"
    agent_envoyextensions_builtin_aws_lambda_aws_lambda_go["agent/envoyextensions/builtin/aws-lambda/aws_lambda.go"
    agent_nodeid_go["agent/nodeid.go"
    agent_uiserver_buf_index_fs_go["agent/uiserver/buf_index_fs.go"
    proto_private_pbconnect_connect_pb_go["proto/private/pbconnect/connect.pb.go"
    sdk_iptables_iptables_executor_unsupported_go["sdk/iptables/iptables_executor_unsupported.go"
    test_integration_connect_envoy_test_sds_server_sds_go["test/integration/connect/envoy/test-sds-server/sds.go"
    agent_consul_acl_client_go["agent/consul/acl_client.go"
    agent_consul_configentry_backend_go["agent/consul/configentry_backend.go"
    agent_consul_state_query_go["agent/consul/state/query.go"
    api_config_entry_gateways_go["api/config_entry_gateways.go"
    command_keygen_keygen_go["command/keygen/keygen.go"
    command_maint_maint_go["command/maint/maint.go"
    ipaddr_detect_go["ipaddr/detect.go"
    sdk_iptables_iptables_executor_linux_go["sdk/iptables/iptables_executor_linux.go"
    agent_xds_testing_go["agent/xds/testing.go"
    lib_cluster_go["lib/cluster.go"
    agent_envoyextensions_builtin_property_override_property_override_go["agent/envoyextensions/builtin/property-override/property_override.go"
    grpcmocks_proto_public_pbresource_mock_ResourceServiceServer_go["grpcmocks/proto-public/pbresource/mock_ResourceServiceServer.go"
    agent_connect_ca_provider_consul_go["agent/connect/ca/provider_consul.go"
    agent_consul_options_ce_go["agent/consul/options_ce.go"
    agent_proxycfg_connect_proxy_go["agent/proxycfg/connect_proxy.go"
    agent_utilization_endpoint_go["agent/utilization_endpoint.go"
    agent_xds_locality_policy_ce_go["agent/xds/locality_policy_ce.go"
    lib_retry_retry_go["lib/retry/retry.go"
    proto_public_annotations_ratelimit_ratelimit_deepcopy_gen_go["proto-public/annotations/ratelimit/ratelimit_deepcopy.gen.go"
    agent_structs_config_entry_status_go["agent/structs/config_entry_status.go"
    agent_xds_failover_policy_go["agent/xds/failover_policy.go"
    api_config_entry_inline_certificate_go["api/config_entry_inline_certificate.go"
    api_peering_go["api/peering.go"
    internal_controller_cache_index_txn_go["internal/controller/cache/index/txn.go"
    proto_public_pbacl_acl_deepcopy_gen_go["proto-public/pbacl/acl_deepcopy.gen.go"
    agent_consul_state_graveyard_go["agent/consul/state/graveyard.go"
    api_watch_watch_go["api/watch/watch.go"
    lib_template_hil_go["lib/template/hil.go"
    agent_consul_auth_token_writer_ce_go["agent/consul/auth/token_writer_ce.go"
    agent_consul_state_coordinate_go["agent/consul/state/coordinate.go"
    agent_consul_usagemetrics_usagemetrics_go["agent/consul/usagemetrics/usagemetrics.go"
    agent_grpc_external_services_connectca_mock_CAManager_go["agent/grpc-external/services/connectca/mock_CAManager.go"
    grpcmocks_proto_public_pbconnectca_mock_ConnectCAService_WatchRootsServer_go["grpcmocks/proto-public/pbconnectca/mock_ConnectCAService_WatchRootsServer.go"
    internal_controller_cache_indexers_indexersmock_mock_FromArgs_go["internal/controller/cache/indexers/indexersmock/mock_FromArgs.go"
    proto_public_pbmulticluster_v2_namespace_exported_services_pb_go["proto-public/pbmulticluster/v2/namespace_exported_services.pb.go"
    agent_auto_config_config_ce_go["agent/auto-config/config_ce.go"
    agent_consul_state_catalog_schema_go["agent/consul/state/catalog_schema.go"
    api_operator_area_go["api/operator_area.go"
    command_intention_list_intention_list_go["command/intention/list/intention_list.go"
    command_version_version_go["command/version/version.go"
    acl_policy_authorizer_go["acl/policy_authorizer.go"
    agent_cache_watch_go["agent/cache/watch.go"
    agent_consul_fsm_log_verification_chunking_shim_go["agent/consul/fsm/log_verification_chunking_shim.go"
    agent_consul_server_ce_go["agent/consul/server_ce.go"
    agent_local_state_go["agent/local/state.go"
    agent_xds_listeners_apigateway_go["agent/xds/listeners_apigateway.go"
    connect_proxy_testing_go["connect/proxy/testing.go"
    agent_xds_extensionruntime_runtime_config_go["agent/xds/extensionruntime/runtime_config.go"
    agent_grpc_external_services_connectca_mock_ACLResolver_go["agent/grpc-external/services/connectca/mock_ACLResolver.go"
    agent_grpc_external_services_peerstream_stream_resources_go["agent/grpc-external/services/peerstream/stream_resources.go"
    agent_rpc_middleware_recovery_go["agent/rpc/middleware/recovery.go"
    internal_controller_cache_indexers_id_indexer_go["internal/controller/cache/indexers/id_indexer.go"
    internal_multicluster_internal_types_decoded_go["internal/multicluster/internal/types/decoded.go"
    internal_resource_protoc_gen_json_shim_main_go["internal/resource/protoc-gen-json-shim/main.go"
    internal_resource_resource_go["internal/resource/resource.go"
    internal_storage_raft_forwarding_go["internal/storage/raft/forwarding.go"
    agent_connect_testing_spiffe_go["agent/connect/testing_spiffe.go"
    agent_grpc_external_services_resource_delete_go["agent/grpc-external/services/resource/delete.go"
    command_services_export_export_go["command/services/export/export.go"
    internal_controller_cache_client_go["internal/controller/cache/client.go"
    proto_private_pbconfigentry_config_entry_gen_go["proto/private/pbconfigentry/config_entry.gen.go"
    proto_private_prototest_testing_go["proto/private/prototest/testing.go"
    proto_private_pbservice_healthcheck_pb_go["proto/private/pbservice/healthcheck.pb.go"
    proto_public_pbconnectca_ca_grpc_pb_go["proto-public/pbconnectca/ca_grpc.pb.go"
    agent_consul_txn_endpoint_go["agent/consul/txn_endpoint.go"
    agent_grpc_external_services_peerstream_subscription_state_go["agent/grpc-external/services/peerstream/subscription_state.go"
    agent_signal_unix_go["agent/signal_unix.go"
    agent_submatview_handler_go["agent/submatview/handler.go"
    command_acl_token_create_token_create_go["command/acl/token/create/token_create.go"
    internal_protohcl_blocks_go["internal/protohcl/blocks.go"
    testing_deployer_sprawl_internal_tfgen_prelude_go["testing/deployer/sprawl/internal/tfgen/prelude.go"
    acl_errors_go["acl/errors.go"
    agent_config_config_go["agent/config/config.go"
    testing_deployer_topology_images_go["testing/deployer/topology/images.go"
    agent_cache_types_intention_upstreams_go["agent/cache-types/intention_upstreams.go"
    agent_signal_windows_go["agent/signal_windows.go"
    agent_status_endpoint_go["agent/status_endpoint.go"
    command_acl_authmethod_update_authmethod_update_ce_go["command/acl/authmethod/update/authmethod_update_ce.go"
    command_acl_role_read_role_read_go["command/acl/role/read/role_read.go"
    grpcmocks_proto_public_pbserverdiscovery_mock_ServerDiscoveryServiceServer_go["grpcmocks/proto-public/pbserverdiscovery/mock_ServerDiscoveryServiceServer.go"
    logging_monitor_monitor_go["logging/monitor/monitor.go"
    troubleshoot_proxy_utils_go["troubleshoot/proxy/utils.go"
    acl_static_authorizer_go["acl/static_authorizer.go"
    agent_acl_ce_go["agent/acl_ce.go"
    agent_consul_server_go["agent/consul/server.go"
    agent_proxycfg_snapshot_go["agent/proxycfg/snapshot.go"
    agent_proxycfg_sources_local_sync_go["agent/proxycfg-sources/local/sync.go"
    agent_structs_discovery_chain_go["agent/structs/discovery_chain.go"
    internal_protohcl_attributes_go["internal/protohcl/attributes.go"
    proto_private_pbservice_node_pb_go["proto/private/pbservice/node.pb.go"
    agent_cache_types_discovery_chain_go["agent/cache-types/discovery_chain.go"
    agent_hcp_telemetry_otel_exporter_go["agent/hcp/telemetry/otel_exporter.go"
    agent_proxycfg_glue_health_blocking_go["agent/proxycfg-glue/health_blocking.go"
    agent_snapshot_endpoint_go["agent/snapshot_endpoint.go"
    agent_structs_check_type_go["agent/structs/check_type.go"
    agent_submatview_mock_ACLResolver_go["agent/submatview/mock_ACLResolver.go"
    command_resource_client_grpc_flags_go["command/resource/client/grpc-flags.go"
    proto_public_annotations_ratelimit_ratelimit_pb_go["proto-public/annotations/ratelimit/ratelimit.pb.go"
    agent_consul_acl_replication_go["agent/consul/acl_replication.go"
    command_connect_envoy_exec_windows_go["command/connect/envoy/exec_windows.go"
    command_resource_delete_grpc_delete_go["command/resource/delete-grpc/delete.go"
    logging_log_levels_go["logging/log_levels.go"
    sdk_freeport_systemlimit_windows_go["sdk/freeport/systemlimit_windows.go"
    test_integration_consul_container_libs_cluster_config_go["test/integration/consul-container/libs/cluster/config.go"
    test_integration_consul_container_libs_utils_utils_go["test/integration/consul-container/libs/utils/utils.go"
    test_integration_consul_container_libs_utils_docker_go["test/integration/consul-container/libs/utils/docker.go"
    agent_consul_state_usage_go["agent/consul/state/usage.go"
    agent_grpc_middleware_stats_go["agent/grpc-middleware/stats.go"
    api_connect_go["api/connect.go"
    proto_public_pbconnectca_ca_deepcopy_gen_go["proto-public/pbconnectca/ca_deepcopy.gen.go"
    troubleshoot_proxy_stats_go["troubleshoot/proxy/stats.go"
    agent_consul_leader_connect_ca_ce_go["agent/consul/leader_connect_ca_ce.go"
    agent_structs_peering_go["agent/structs/peering.go"
    api_kv_go["api/kv.go"
    proto_private_pbcommon_convert_pbstruct_go["proto/private/pbcommon/convert_pbstruct.go"
    agent_consul_enterprise_config_ce_go["agent/consul/enterprise_config_ce.go"
    agent_proxycfg_glue_health_go["agent/proxycfg-glue/health.go"
    ai_annotate_py[""
    command_snapshot_restore_snapshot_restore_go["command/snapshot/restore/snapshot_restore.go"
    connect_proxy_conn_go["connect/proxy/conn.go"
    internal_controller_cache_indexers_ref_indexer_go["internal/controller/cache/indexers/ref_indexer.go"
    acl_validation_go["acl/validation.go"
    agent_consul_authmethod_testauth_testing_ce_go["agent/consul/authmethod/testauth/testing_ce.go"
    agent_connect_uri_signing_go["agent/connect/uri_signing.go"
    agent_leafcert_leafcert_test_helpers_go["agent/leafcert/leafcert_test_helpers.go"
    agent_submatview_rpc_materializer_go["agent/submatview/rpc_materializer.go"
    api_operator_raft_go["api/operator_raft.go"
    internal_controller_dependency_higher_order_go["internal/controller/dependency/higher_order.go"
    agent_cache_types_health_services_go["agent/cache-types/health_services.go"
    agent_consul_discoverychain_compile_go["agent/consul/discoverychain/compile.go"
    agent_consul_state_config_entry_sameness_group_ce_go["agent/consul/state/config_entry_sameness_group_ce.go"
    agent_structs_config_entry_sameness_group_go["agent/structs/config_entry_sameness_group.go"
    proto_public_pbmulticluster_v2_namespace_exported_services_json_gen_go["proto-public/pbmulticluster/v2/namespace_exported_services_json.gen.go"
    agent_cache_types_intention_upstreams_destination_go["agent/cache-types/intention_upstreams_destination.go"
    agent_consul_authmethod_testauth_testing_go["agent/consul/authmethod/testauth/testing.go"
    agent_consul_operator_usage_endpoint_go["agent/consul/operator_usage_endpoint.go"
    agent_consul_state_config_entry_sameness_group_go["agent/consul/state/config_entry_sameness_group.go"
    agent_structs_config_entry_discoverychain_go["agent/structs/config_entry_discoverychain.go"
    agent_structs_check_definition_go["agent/structs/check_definition.go"
    internal_radix_doc_go["internal/radix/doc.go"
    proto_private_pbdemo_v1_demo_pb_go["proto/private/pbdemo/v1/demo.pb.go"
    agent_connect_sni_go["agent/connect/sni.go"
    agent_proxycfg_sources_local_local_go["agent/proxycfg-sources/local/local.go"
    command_peering_peering_go["command/peering/peering.go"
    proto_private_pbservice_convert_go["proto/private/pbservice/convert.go"
    proto_public_pbdataplane_dataplane_pb_go["proto-public/pbdataplane/dataplane.pb.go"
    proto_public_pbdns_dns_grpc_pb_go["proto-public/pbdns/dns_grpc.pb.go"
    test_integration_consul_container_libs_assert_common_go["test/integration/consul-container/libs/assert/common.go"
    testing_deployer_sprawl_network_area_ce_go["testing/deployer/sprawl/network_area_ce.go"
    agent_consul_acl_token_exp_go["agent/consul/acl_token_exp.go"
    agent_consul_system_metadata_go["agent/consul/system_metadata.go"
    agent_structs_connect_ca_go["agent/structs/connect_ca.go"
    command_acl_role_list_role_list_go["command/acl/role/list/role_list.go"
    command_acl_token_update_token_update_go["command/acl/token/update/token_update.go"
    internal_tools_proto_gen_rpc_glue_e2e_source_pb_go["internal/tools/proto-gen-rpc-glue/e2e/source.pb.go"
    sentinel_scope_go["sentinel/scope.go"
    agent_blockingquery_mock_FSMServer_go["agent/blockingquery/mock_FSMServer.go"
    agent_consul_acl_go["agent/consul/acl.go"
    agent_grpc_external_services_resource_mock_Backend_go["agent/grpc-external/services/resource/mock_Backend.go"
    agent_grpc_internal_services_subscribe_subscribe_go["agent/grpc-internal/services/subscribe/subscribe.go"
    agent_submatview_local_materializer_go["agent/submatview/local_materializer.go"
    logging_names_go["logging/names.go"
    command_connect_proxy_register_go["command/connect/proxy/register.go"
    internal_tools_protoc_gen_grpc_clone_e2e_proto_cloning_stream_pb_go["internal/tools/protoc-gen-grpc-clone/e2e/proto/cloning_stream.pb.go"
    logging_logger_go["logging/logger.go"
    test_integration_consul_container_libs_utils_helpers_go["test/integration/consul-container/libs/utils/helpers.go"
    agent_blockingquery_mock_ResponseMeta_go["agent/blockingquery/mock_ResponseMeta.go"
    agent_grpc_external_services_resource_list_go["agent/grpc-external/services/resource/list.go"
    agent_proxycfg_testing_go["agent/proxycfg/testing.go"
    agent_structs_config_entry_jwt_provider_go["agent/structs/config_entry_jwt_provider.go"
    command_resource_client_grpc_config_go["command/resource/client/grpc-config.go"
    lib_mutex_mutex_go["lib/mutex/mutex.go"
    test_integration_consul_container_libs_utils_debug_go["test/integration/consul-container/libs/utils/debug.go"
    testing_deployer_sprawl_internal_tfgen_res_go["testing/deployer/sprawl/internal/tfgen/res.go"
    agent_auto_config_auto_config_go["agent/auto-config/auto_config.go"
    agent_consul_state_peering_go["agent/consul/state/peering.go"
    agent_proxycfg_ingress_gateway_go["agent/proxycfg/ingress_gateway.go"
    internal_protohcl_decoder_go["internal/protohcl/decoder.go"
    internal_storage_inmem_store_go["internal/storage/inmem/store.go"
    internal_tools_proto_gen_rpc_glue_e2e_consul_proto_pbcommon_common_go["internal/tools/proto-gen-rpc-glue/e2e/consul/proto/pbcommon/common.go"
    agent_consul_prepared_query_walk_go["agent/consul/prepared_query/walk.go"
    agent_proxycfg_proxycfg_go["agent/proxycfg/proxycfg.go"
    agent_structs_structs_deepcopy_go["agent/structs/structs.deepcopy.go"
    command_connect_envoy_pipe_bootstrap_connect_envoy_pipe_bootstrap_go["command/connect/envoy/pipe-bootstrap/connect_envoy_pipe-bootstrap.go"
    command_helpers_decode_shim_go["command/helpers/decode_shim.go"
    lib_useragent_go["lib/useragent.go"
    agent_consul_subscribe_backend_go["agent/consul/subscribe_backend.go"
    agent_grpc_external_testutils_acl_go["agent/grpc-external/testutils/acl.go"
    proto_public_pbmulticluster_v2_partition_exported_services_deepcopy_gen_go["proto-public/pbmulticluster/v2/partition_exported_services_deepcopy.gen.go"
    sdk_testutil_io_go["sdk/testutil/io.go"
    sdk_testutil_retry_interface_go["sdk/testutil/retry/interface.go"
    test_integration_consul_container_libs_utils_tenancy_go["test/integration/consul-container/libs/utils/tenancy.go"
    agent_config_file_watcher_go["agent/config/file_watcher.go"
    agent_consul_leader_go["agent/consul/leader.go"
    agent_structs_prepared_query_go["agent/structs/prepared_query.go"
    agent_structs_testing_intention_go["agent/structs/testing_intention.go"
    command_acl_token_token_go["command/acl/token/token.go"
    agent_consul_state_tombstone_gc_go["agent/consul/state/tombstone_gc.go"
    agent_proxycfg_testing_terminating_gateway_go["agent/proxycfg/testing_terminating_gateway.go"
    command_snapshot_save_snapshot_save_go["command/snapshot/save/snapshot_save.go"
    internal_controller_cache_index_indexmock_mock_resourceIterable_go["internal/controller/cache/index/indexmock/mock_resourceIterable.go"
    internal_multicluster_internal_controllers_register_go["internal/multicluster/internal/controllers/register.go"
    internal_storage_inmem_schema_go["internal/storage/inmem/schema.go"
    internal_tools_protoc_gen_grpc_clone_main_go["internal/tools/protoc-gen-grpc-clone/main.go"
    logging_logfile_darwinandfreebsd_go["logging/logfile_darwinandfreebsd.go"
    api_operator_usage_go["api/operator_usage.go"
    proto_private_pbautoconf_auto_config_pb_go["proto/private/pbautoconf/auto_config.pb.go"
    sdk_testutil_retry_retry_go["sdk/testutil/retry/retry.go"
    agent_structs_envoy_extension_go["agent/structs/envoy_extension.go"
    agent_xds_delta_go["agent/xds/delta.go"
    command_connect_envoy_bootstrap_config_go["command/connect/envoy/bootstrap_config.go"
    command_tls_cert_tls_cert_go["command/tls/cert/tls_cert.go"
    internal_resource_errors_go["internal/resource/errors.go"
    agent_config_merge_go["agent/config/merge.go"
    agent_consul_controller_queue_defer_go["agent/consul/controller/queue/defer.go"
    agent_consul_util_go["agent/consul/util.go"
    agent_leafcert_cached_roots_go["agent/leafcert/cached_roots.go"
    agent_proxycfg_mesh_gateway_go["agent/proxycfg/mesh_gateway.go"
    command_resource_delete_delete_go["command/resource/delete/delete.go"
    internal_tools_protoc_gen_grpc_clone_e2e_mock_SimpleClient_go["internal/tools/protoc-gen-grpc-clone/e2e/mock_SimpleClient.go"
    test_integration_consul_container_test_gateways_tenancy_ce_go["test/integration/consul-container/test/gateways/tenancy_ce.go"
    agent_checks_alias_go["agent/checks/alias.go"
    agent_consul_controller_doc_go["agent/consul/controller/doc.go"
    agent_operator_endpoint_go["agent/operator_endpoint.go"
    agent_rpcclient_health_view_go["agent/rpcclient/health/view.go"
    agent_uiserver_uiserver_go["agent/uiserver/uiserver.go"
    troubleshoot_proxy_certs_go["troubleshoot/proxy/certs.go"
    agent_connect_uri_service_ce_go["agent/connect/uri_service_ce.go"
    agent_grpc_external_services_resource_mutate_and_validate_go["agent/grpc-external/services/resource/mutate_and_validate.go"
    agent_hcp_client_http_client_go["agent/hcp/client/http_client.go"
    command_validate_validate_go["command/validate/validate.go"
    proto_public_pbconnectca_cloning_stream_pb_go["proto-public/pbconnectca/cloning_stream.pb.go"
    agent_agent_endpoint_go["agent/agent_endpoint.go"
    agent_consul_state_peering_ce_go["agent/consul/state/peering_ce.go"
    grpcmocks_proto_public_pbserverdiscovery_mock_ServerDiscoveryService_WatchServersClient_go["grpcmocks/proto-public/pbserverdiscovery/mock_ServerDiscoveryService_WatchServersClient.go"
    internal_controller_cache_index_convenience_go["internal/controller/cache/index/convenience.go"
    proto_private_pbpeering_peering_ce_go["proto/private/pbpeering/peering_ce.go"
    agent_grpc_external_services_acl_mock_Validator_go["agent/grpc-external/services/acl/mock_Validator.go"
    agent_local_testing_go["agent/local/testing.go"
    command_reload_reload_go["command/reload/reload.go"
    lib_math_go["lib/math.go"
    proto_private_pbconnect_connect_go["proto/private/pbconnect/connect.go"
    agent_cache_types_options_go["agent/cache-types/options.go"
    agent_consul_state_intention_go["agent/consul/state/intention.go"
    agent_grpc_external_services_peerstream_testing_go["agent/grpc-external/services/peerstream/testing.go"
    agent_proxycfg_testing_mesh_gateway_go["agent/proxycfg/testing_mesh_gateway.go"
    agent_structs_aclfilter_filter_go["agent/structs/aclfilter/filter.go"
    command_resource_list_grpc_list_go["command/resource/list-grpc/list.go"
    command_troubleshoot_ports_troubleshoot_ports_go["command/troubleshoot/ports/troubleshoot_ports.go"
    internal_go_sso_oidcauth_oidc_go["internal/go-sso/oidcauth/oidc.go"
    agent_grpc_external_services_resource_mock_ACLResolver_go["agent/grpc-external/services/resource/mock_ACLResolver.go"
    agent_proxycfg_sources_catalog_config_source_go["agent/proxycfg-sources/catalog/config_source.go"
    internal_tools_protoc_gen_grpc_clone_e2e_proto_service_cloning_grpc_pb_go["internal/tools/protoc-gen-grpc-clone/e2e/proto/service_cloning_grpc.pb.go"
    testing_deployer_sprawl_catalog_go["testing/deployer/sprawl/catalog.go"
    agent_envoyextensions_registered_extensions_go["agent/envoyextensions/registered_extensions.go"
    agent_hcp_client_metrics_client_go["agent/hcp/client/metrics_client.go"
    command_lock_util_windows_go["command/lock/util_windows.go"
    command_troubleshoot_proxy_troubleshoot_proxy_go["command/troubleshoot/proxy/troubleshoot_proxy.go"
    internal_storage_inmem_watch_go["internal/storage/inmem/watch.go"
    snapshot_archive_go["snapshot/archive.go"
    agent_grpc_external_services_resource_server_ce_go["agent/grpc-external/services/resource/server_ce.go"
    agent_structs_federation_state_go["agent/structs/federation_state.go"
    command_acl_role_create_role_create_go["command/acl/role/create/role_create.go"
    command_join_join_go["command/join/join.go"
    proto_private_pbpeerstream_peerstream_go["proto/private/pbpeerstream/peerstream.go"
    proto_public_pbmulticluster_v2_exported_services_deepcopy_gen_go["proto-public/pbmulticluster/v2/exported_services_deepcopy.gen.go"
    types_tls_go["types/tls.go"
    agent_consul_state_census_utilization_go["agent/consul/state/census_utilization.go"
    agent_proxycfg_glue_config_entry_go["agent/proxycfg-glue/config_entry.go"
    agent_rpcclient_configentry_configentry_go["agent/rpcclient/configentry/configentry.go"
    internal_controller_watch_go["internal/controller/watch.go"
    internal_multicluster_internal_controllers_v1compat_controller_go["internal/multicluster/internal/controllers/v1compat/controller.go"
    internal_protohcl_oneof_go["internal/protohcl/oneof.go"
    internal_tools_proto_gen_rpc_glue_main_go["internal/tools/proto-gen-rpc-glue/main.go"
    test_integration_consul_container_libs_utils_defer_go["test/integration/consul-container/libs/utils/defer.go"
    agent_consul_leader_registrator_v1_go["agent/consul/leader_registrator_v1.go"
    agent_grpc_external_services_connectca_watch_roots_go["agent/grpc-external/services/connectca/watch_roots.go"
    agent_proxycfg_glue_gateway_services_go["agent/proxycfg-glue/gateway_services.go"
    command_connect_connect_go["command/connect/connect.go"
    command_config_list_config_list_go["command/config/list/config_list.go"
    command_connect_envoy_exec_unsupported_go["command/connect/envoy/exec_unsupported.go"
    internal_multicluster_exports_go["internal/multicluster/exports.go"
    test_integration_consul_container_libs_cluster_cluster_go["test/integration/consul-container/libs/cluster/cluster.go"
    agent_xds_platform_net_fallback_go["agent/xds/platform/net_fallback.go"
    agent_cache_types_catalog_service_list_go["agent/cache-types/catalog_service_list.go"
    agent_consul_raft_handle_go["agent/consul/raft_handle.go"
    agent_grpc_external_services_connectca_sign_go["agent/grpc-external/services/connectca/sign.go"
    agent_rpc_peering_service_go["agent/rpc/peering/service.go"
    agent_structs_dns_go["agent/structs/dns.go"
    grpcmocks_proto_public_pbserverdiscovery_mock_UnsafeServerDiscoveryServiceServer_go["grpcmocks/proto-public/pbserverdiscovery/mock_UnsafeServerDiscoveryServiceServer.go"
    internal_radix_radix_go["internal/radix/radix.go"
    internal_tools_proto_gen_rpc_glue_e2e_consul_agent_structs_structs_go["internal/tools/proto-gen-rpc-glue/e2e/consul/agent/structs/structs.go"
    agent_consul_state_usage_ce_go["agent/consul/state/usage_ce.go"
    command_services_config_go["command/services/config.go"
    internal_controller_dependency_dependencymock_mock_DependencyTransform_go["internal/controller/dependency/dependencymock/mock_DependencyTransform.go"
    lib_hoststats_collector_go["lib/hoststats/collector.go"
    proto_public_pbresource_annotations_pb_go["proto-public/pbresource/annotations.pb.go"
    agent_consul_auth_login_go["agent/consul/auth/login.go"
    agent_consul_fsm_snapshot_ce_go["agent/consul/fsm/snapshot_ce.go"
    agent_consul_state_config_entry_go["agent/consul/state/config_entry.go"
    agent_consul_servercert_manager_go["agent/consul/servercert/manager.go"
    agent_grpc_external_services_peerstream_health_snapshot_go["agent/grpc-external/services/peerstream/health_snapshot.go"
    agent_xds_z_xds_packages_go["agent/xds/z_xds_packages.go"
    command_services_deregister_deregister_go["command/services/deregister/deregister.go"
    sentinel_sentinel_ce_go["sentinel/sentinel_ce.go"
    troubleshoot_proxy_troubleshoot_proxy_go["troubleshoot/proxy/troubleshoot_proxy.go"
    agent_consul_replication_go["agent/consul/replication.go"
    agent_grpc_internal_tracker_go["agent/grpc-internal/tracker.go"
    agent_xds_clusters_go["agent/xds/clusters.go"
    api_status_go["api/status.go"
    command_operator_operator_go["command/operator/operator.go"
    internal_controller_cache_errors_go["internal/controller/cache/errors.go"
    internal_resource_protoc_gen_deepcopy_main_go["internal/resource/protoc-gen-deepcopy/main.go"
    proto_public_pbconnectca_ca_cloning_grpc_pb_go["proto-public/pbconnectca/ca_cloning_grpc.pb.go"
    agent_cache_types_exported_peered_services_go["agent/cache-types/exported_peered_services.go"
    agent_config_config_ce_go["agent/config/config_ce.go"
    agent_consul_auto_config_endpoint_go["agent/consul/auto_config_endpoint.go"
    agent_consul_leader_intentions_ce_go["agent/consul/leader_intentions_ce.go"
    agent_consul_prepared_query_endpoint_ce_go["agent/consul/prepared_query_endpoint_ce.go"
    command_peering_establish_establish_go["command/peering/establish/establish.go"
    proto_private_pbstatus_status_pb_go["proto/private/pbstatus/status.pb.go"
    acl_resolver_result_go["acl/resolver/result.go"
    agent_config_agent_limits_go["agent/config/agent_limits.go"
    agent_consul_state_graveyard_ce_go["agent/consul/state/graveyard_ce.go"
    agent_consul_state_kvs_ce_go["agent/consul/state/kvs_ce.go"
    agent_grpc_internal_listener_go["agent/grpc-internal/listener.go"
    command_resource_apply_apply_go["command/resource/apply/apply.go"
    internal_protohcl_well_known_types_go["internal/protohcl/well_known_types.go"
    internal_resource_http_http_go["internal/resource/http/http.go"
    agent_cache_types_service_checks_go["agent/cache-types/service_checks.go"
    agent_intentions_endpoint_go["agent/intentions_endpoint.go"
    command_connect_proxy_flag_upstreams_go["command/connect/proxy/flag_upstreams.go"
    proto_public_pbresource_resource_grpc_pb_go["proto-public/pbresource/resource_grpc.pb.go"
    test_integration_consul_container_test_consul_envoy_version_consul_envoy_version_go["test/integration/consul-container/test/consul_envoy_version/consul_envoy_version.go"
    agent_consul_server_register_go["agent/consul/server_register.go"
    agent_consul_tenancy_bridge_go["agent/consul/tenancy_bridge.go"
    agent_grpc_external_services_acl_login_go["agent/grpc-external/services/acl/login.go"
    agent_hcp_telemetry_provider_go["agent/hcp/telemetry_provider.go"
    lib_stringslice_stringslice_go["lib/stringslice/stringslice.go"
    agent_consul_peering_backend_ce_go["agent/consul/peering_backend_ce.go"
    command_operator_autopilot_operator_autopilot_go["command/operator/autopilot/operator_autopilot.go"
    command_troubleshoot_upstreams_troubleshoot_upstreams_go["command/troubleshoot/upstreams/troubleshoot_upstreams.go"
    grpcmocks_proto_public_pbconnectca_mock_ConnectCAServiceServer_go["grpcmocks/proto-public/pbconnectca/mock_ConnectCAServiceServer.go"
    main_go["main.go"
    agent_config_deprecated_go["agent/config/deprecated.go"
    agent_consul_config_replication_go["agent/consul/config_replication.go"
    proto_public_pbdataplane_dataplane_deepcopy_gen_go["proto-public/pbdataplane/dataplane_deepcopy.gen.go"
    agent_cache_type_go["agent/cache/type.go"
    agent_consul_config_ce_go["agent/consul/config_ce.go"
    agent_consul_state_autopilot_go["agent/consul/state/autopilot.go"
    agent_grpc_external_services_serverdiscovery_watch_servers_go["agent/grpc-external/services/serverdiscovery/watch_servers.go"
    agent_systemd_notify_go["agent/systemd/notify.go"
    command_config_config_go["command/config/config.go"
    command_login_aws_go["command/login/aws.go"
    lib_path_go["lib/path.go"
    agent_consul_authmethod_authmethods_ce_go["agent/consul/authmethod/authmethods_ce.go"
    agent_consul_config_endpoint_go["agent/consul/config_endpoint.go"
    agent_xds_secrets_go["agent/xds/secrets.go"
    command_config_delete_config_delete_go["command/config/delete/config_delete.go"
    command_debug_debug_go["command/debug/debug.go"
    internal_testing_errors_errors_go["internal/testing/errors/errors.go"
    internal_storage_storage_go["internal/storage/storage.go"
    proto_private_pbpeering_peering_gen_go["proto/private/pbpeering/peering.gen.go"
    acl_authorizer_ce_go["acl/authorizer_ce.go"
    agent_consul_state_catalog_events_go["agent/consul/state/catalog_events.go"
    agent_consul_state_connect_ca_events_go["agent/consul/state/connect_ca_events.go"
    agent_hcp_bootstrap_bootstrap_go["agent/hcp/bootstrap/bootstrap.go"
    command_monitor_monitor_go["command/monitor/monitor.go"
    lib_hoststats_host_go["lib/hoststats/host.go"
    proto_private_pbsubscribe_subscribe_go["proto/private/pbsubscribe/subscribe.go"
    sdk_testutil_types_go["sdk/testutil/types.go"
    command_acl_token_formatter_go["command/acl/token/formatter.go"
    envoyextensions_extensioncommon_basic_extension_adapter_go["envoyextensions/extensioncommon/basic_extension_adapter.go"
    internal_controller_cache_cachemock_mock_DecodedResourceIterator_go["internal/controller/cache/cachemock/mock_DecodedResourceIterator.go"
    agent_consul_autopilot_ce_go["agent/consul/autopilot_ce.go"
    agent_consul_authmethod_ssoauth_sso_go["agent/consul/authmethod/ssoauth/sso.go"
    internal_controller_cache_index_interfaces_go["internal/controller/cache/index/interfaces.go"
    test_integration_consul_container_test_upgrade_common_go["test/integration/consul-container/test/upgrade/common.go"
    agent_consul_state_config_entry_intention_ce_go["agent/consul/state/config_entry_intention_ce.go"
    agent_proxycfg_testing_upstreams_ce_go["agent/proxycfg/testing_upstreams_ce.go"
    command_connect_envoy_exec_unix_go["command/connect/envoy/exec_unix.go"
    internal_controller_cache_indexers_indexersmock_mock_BoundReferences_go["internal/controller/cache/indexers/indexersmock/mock_BoundReferences.go"
    internal_resource_demo_demo_go["internal/resource/demo/demo.go"
    proto_public_pbserverdiscovery_serverdiscovery_grpc_pb_go["proto-public/pbserverdiscovery/serverdiscovery_grpc.pb.go"
    agent_grpc_external_forward_go["agent/grpc-external/forward.go"
    agent_hcp_telemetry_custom_metrics_go["agent/hcp/telemetry/custom_metrics.go"
    agent_rpc_middleware_rate_limit_mappings_go["agent/rpc/middleware/rate_limit_mappings.go"
    agent_structs_structs_go["agent/structs/structs.go"
    agent_watch_handler_go["agent/watch_handler.go"
    connect_proxy_proxy_go["connect/proxy/proxy.go"
    agent_cache_types_testing_go["agent/cache-types/testing.go"
    agent_connect_uri_agent_go["agent/connect/uri_agent.go"
    agent_consul_controller_queue_queue_go["agent/consul/controller/queue/queue.go"
    agent_envoyextensions_builtin_ext_authz_structs_go["agent/envoyextensions/builtin/ext-authz/structs.go"
    lib_json_go["lib/json.go"
    lib_translate_go["lib/translate.go"
    proto_public_pbmulticluster_v2_computed_exported_services_deepcopy_gen_go["proto-public/pbmulticluster/v2/computed_exported_services_deepcopy.gen.go"
    agent_cache_types_catalog_datacenters_go["agent/cache-types/catalog_datacenters.go"
    agent_notify_go["agent/notify.go"
    api_config_entry_discoverychain_go["api/config_entry_discoverychain.go"
    command_tls_cert_create_tls_cert_create_go["command/tls/cert/create/tls_cert_create.go"
    testing_deployer_topology_naming_shim_go["testing/deployer/topology/naming_shim.go"
    testing_deployer_util_net_go["testing/deployer/util/net.go"
    agent_grpc_middleware_testutil_fake_sink_go["agent/grpc-middleware/testutil/fake_sink.go"
    agent_user_event_go["agent/user_event.go"
    api_config_entry_intentions_go["api/config_entry_intentions.go"
    command_acl_authmethod_authmethod_go["command/acl/authmethod/authmethod.go"
    command_services_services_go["command/services/services.go"
    internal_go_sso_oidcauth_config_go["internal/go-sso/oidcauth/config.go"
    agent_configentry_doc_go["agent/configentry/doc.go"
    agent_exec_exec_go["agent/exec/exec.go"
    agent_structs_connect_proxy_config_ce_go["agent/structs/connect_proxy_config_ce.go"
    agent_xds_platform_net_linux_go["agent/xds/platform/net_linux.go"
    api_operator_license_go["api/operator_license.go"
    proto_private_pbconfig_config_pb_go["proto/private/pbconfig/config.pb.go"
    agent_consul_authmethod_awsauth_aws_go["agent/consul/authmethod/awsauth/aws.go"
    internal_controller_dependency_dependencymock_mock_CacheIDModifier_go["internal/controller/dependency/dependencymock/mock_CacheIDModifier.go"
    proto_private_pbdemo_v1_resources_rtypes_go["proto/private/pbdemo/v1/resources.rtypes.go"
    testing_deployer_sprawl_internal_tfgen_dns_go["testing/deployer/sprawl/internal/tfgen/dns.go"
    agent_config_endpoint_go["agent/config_endpoint.go"
    agent_consul_acl_server_ce_go["agent/consul/acl_server_ce.go"
    agent_enterprise_delegate_ce_go["agent/enterprise_delegate_ce.go"
    agent_sidecar_service_go["agent/sidecar_service.go"
    api_connect_intention_go["api/connect_intention.go"
    internal_resource_authz_go["internal/resource/authz.go"
    internal_resource_resourcetest_acls_go["internal/resource/resourcetest/acls.go"
    testing_deployer_sprawl_details_go["testing/deployer/sprawl/details.go"
    agent_hcp_telemetry_doc_go["agent/hcp/telemetry/doc.go"
    agent_leafcert_cert_go["agent/leafcert/cert.go"
    command_resource_read_read_go["command/resource/read/read.go"
    internal_controller_controllermock_mock_Lease_go["internal/controller/controllermock/mock_Lease.go"
    internal_controller_controllermock_mock_DependencyMapper_go["internal/controller/controllermock/mock_DependencyMapper.go"
    agent_ae_ae_go["agent/ae/ae.go"
    agent_leafcert_roots_go["agent/leafcert/roots.go"
    agent_translate_addr_go["agent/translate_addr.go"
    test_integration_consul_container_libs_cluster_app_go["test/integration/consul-container/libs/cluster/app.go"
    acl_acl_go["acl/acl.go"
    command_acl_policy_delete_policy_delete_go["command/acl/policy/delete/policy_delete.go"
    internal_resource_equality_go["internal/resource/equality.go"
    proto_private_pbconfigentry_config_entry_pb_go["proto/private/pbconfigentry/config_entry.pb.go"
    version_version_go["version/version.go"
    agent_consul_client_serf_go["agent/consul/client_serf.go"
    agent_consul_federation_state_replication_go["agent/consul/federation_state_replication.go"
    agent_proxycfg_sources_catalog_mock_SessionLimiter_go["agent/proxycfg-sources/catalog/mock_SessionLimiter.go"
    agent_txn_endpoint_go["agent/txn_endpoint.go"
    agent_config_config_deepcopy_go["agent/config/config.deepcopy.go"
    agent_consul_rpc_go["agent/consul/rpc.go"
    command_catalog_helpers_ce_go["command/catalog/helpers_ce.go"
    command_registry_go["command/registry.go"
    internal_controller_controllermock_mock_DependencyTransform_go["internal/controller/controllermock/mock_DependencyTransform.go"
    agent_connect_parsing_go["agent/connect/parsing.go"
    agent_consul_reporting_reportingmock_mock_ServerDelegate_go["agent/consul/reporting/reportingmock/mock_ServerDelegate.go"
    agent_grpc_external_limiter_limiter_go["agent/grpc-external/limiter/limiter.go"
    command_tls_ca_tls_ca_go["command/tls/ca/tls_ca.go"
    agent_connect_uri_service_go["agent/connect/uri_service.go"
    grpcmocks_proto_public_pbacl_mock_UnsafeACLServiceServer_go["grpcmocks/proto-public/pbacl/mock_UnsafeACLServiceServer.go"
    lib_file_atomic_go["lib/file/atomic.go"
    acl_policy_merger_go["acl/policy_merger.go"
    agent_ae_trigger_go["agent/ae/trigger.go"
    agent_apiserver_go["agent/apiserver.go"
    agent_grpc_external_services_dns_server_go["agent/grpc-external/services/dns/server.go"
    internal_resource_resourcetest_fs_go["internal/resource/resourcetest/fs.go"
    agent_consul_gateways_controller_gateways_go["agent/consul/gateways/controller_gateways.go"
    agent_rpc_peering_testing_go["agent/rpc/peering/testing.go"
    command_catalog_list_nodes_catalog_list_nodes_go["command/catalog/list/nodes/catalog_list_nodes.go"
    internal_controller_cache_indexers_decoded_indexer_go["internal/controller/cache/indexers/decoded_indexer.go"
    acl_MockAuthorizer_go["acl/MockAuthorizer.go"
    agent_grpc_external_services_dataplane_get_supported_features_go["agent/grpc-external/services/dataplane/get_supported_features.go"
    agent_leafcert_util_go["agent/leafcert/util.go"
    internal_protohcl_testproto_example_pb_go["internal/protohcl/testproto/example.pb.go"
    agent_checks_docker_go["agent/checks/docker.go"
    agent_consul_merge_go["agent/consul/merge.go"
    agent_proxycfg_testing_connect_proxy_go["agent/proxycfg/testing_connect_proxy.go"
    proto_private_pbpeering_peering_pb_go["proto/private/pbpeering/peering.pb.go"
    acl_enterprisemeta_ce_go["acl/enterprisemeta_ce.go"
    agent_consul_rate_handler_ce_go["agent/consul/rate/handler_ce.go"
    agent_envoyextensions_builtin_ext_authz_ext_authz_go["agent/envoyextensions/builtin/ext-authz/ext_authz.go"
    agent_grpc_external_services_dataplane_get_envoy_bootstrap_params_go["agent/grpc-external/services/dataplane/get_envoy_bootstrap_params.go"
    agent_structs_system_metadata_go["agent/structs/system_metadata.go"
    internal_resource_registry_ce_go["internal/resource/registry_ce.go"
    internal_resource_tombstone_go["internal/resource/tombstone.go"
    proto_public_pbmulticluster_v2_partition_exported_services_pb_go["proto-public/pbmulticluster/v2/partition_exported_services.pb.go"
    agent_hcp_client_telemetry_config_go["agent/hcp/client/telemetry_config.go"
    agent_leafcert_signer_netrpc_go["agent/leafcert/signer_netrpc.go"
    agent_proxycfg_glue_trust_bundle_go["agent/proxycfg-glue/trust_bundle.go"
    api_internal_go["api/internal.go"
    command_leave_leave_go["command/leave/leave.go"
    agent_blockingquery_blockingquery_go["agent/blockingquery/blockingquery.go"
    agent_consul_auto_encrypt_endpoint_go["agent/consul/auto_encrypt_endpoint.go"
    agent_consul_state_schema_ce_go["agent/consul/state/schema_ce.go"
    agent_grpc_internal_client_go["agent/grpc-internal/client.go"
    agent_leafcert_watch_go["agent/leafcert/watch.go"
    agent_xds_jwt_authn_ce_go["agent/xds/jwt_authn_ce.go"
    api_operator_segment_go["api/operator_segment.go"
    internal_resource_resourcetest_client_go["internal/resource/resourcetest/client.go"
    agent_leafcert_leafcert_go["agent/leafcert/leafcert.go"
    command_acl_bootstrap_bootstrap_go["command/acl/bootstrap/bootstrap.go"
    lib_testhelpers_testhelpers_go["lib/testhelpers/testhelpers.go"
    proto_public_pbdataplane_dataplane_cloning_grpc_pb_go["proto-public/pbdataplane/dataplane_cloning_grpc.pb.go"
    proto_public_pbserverdiscovery_serverdiscovery_deepcopy_gen_go["proto-public/pbserverdiscovery/serverdiscovery_deepcopy.gen.go"
    sdk_testutil_server_methods_go["sdk/testutil/server_methods.go"
    testing_deployer_topology_ids_go["testing/deployer/topology/ids.go"
    agent_grpc_external_services_resource_testing_testing_ce_go["agent/grpc-external/services/resource/testing/testing_ce.go"
    command_acl_templatedpolicy_preview_templated_policy_preview_go["command/acl/templatedpolicy/preview/templated_policy_preview.go"
    command_operator_raft_transferleader_transfer_leader_go["command/operator/raft/transferleader/transfer_leader.go"
    grpcmocks_proto_public_pbserverdiscovery_mock_ServerDiscoveryService_WatchServersServer_go["grpcmocks/proto-public/pbserverdiscovery/mock_ServerDiscoveryService_WatchServersServer.go"
    internal_controller_custom_watch_go["internal/controller/custom_watch.go"
    internal_protohcl_primitives_go["internal/protohcl/primitives.go"
    internal_tools_proto_gen_rpc_glue_e2e_consul_proto_pbcommon_common_pb_go["internal/tools/proto-gen-rpc-glue/e2e/consul/proto/pbcommon/common.pb.go"
    proto_public_pbmulticluster_v2_computed_exported_services_json_gen_go["proto-public/pbmulticluster/v2/computed_exported_services_json.gen.go"
    proto_public_pbresource_annotations_json_gen_go["proto-public/pbresource/annotations_json.gen.go"
    agent_consul_state_schema_go["agent/consul/state/schema.go"
    command_acl_acl_helpers_go["command/acl/acl_helpers.go"
    command_agent_startup_logger_go["command/agent/startup_logger.go"
    test_integration_consul_container_libs_cluster_network_go["test/integration/consul-container/libs/cluster/network.go"
    testing_deployer_topology_default_versions_go["testing/deployer/topology/default_versions.go"
    tlsutil_mock_go["tlsutil/mock.go"
    agent_consul_authmethod_authmethods_go["agent/consul/authmethod/authmethods.go"
    agent_consul_state_prepared_query_go["agent/consul/state/prepared_query.go"
    agent_pool_peek_go["agent/pool/peek.go"
    internal_resource_demo_controller_go["internal/resource/demo/controller.go"
    internal_resource_bound_refs_go["internal/resource/bound_refs.go"
    internal_storage_conformance_conformance_go["internal/storage/conformance/conformance.go"
    proto_public_pbconnectca_ca_pb_go["proto-public/pbconnectca/ca.pb.go"
    test_integ_upgrade_l7_traffic_management_common_go["test-integ/upgrade/l7_traffic_management/common.go"
    agent_agent_go["agent/agent.go"
    agent_connect_ca_common_go["agent/connect/ca/common.go"
    agent_grpc_internal_resolver_resolver_go["agent/grpc-internal/resolver/resolver.go"
    agent_structs_catalog_go["agent/structs/catalog.go"
    agent_structs_config_entry_go["agent/structs/config_entry.go"
    api_operator_keyring_go["api/operator_keyring.go"
    command_info_info_go["command/info/info.go"
    internal_resource_registry_go["internal/resource/registry.go"
    agent_agent_ce_go["agent/agent_ce.go"
    agent_consul_auth_token_writer_go["agent/consul/auth/token_writer.go"
    agent_consul_state_catalog_go["agent/consul/state/catalog.go"
    agent_hcp_client_client_go["agent/hcp/client/client.go"
    agent_hcp_client_mock_Client_go["agent/hcp/client/mock_Client.go"
    agent_structs_acl_go["agent/structs/acl.go"
    api_config_entry_jwt_provider_go["api/config_entry_jwt_provider.go"
    internal_resource_refkey_go["internal/resource/refkey.go"
    agent_proxycfg_sources_catalog_mock_Watcher_go["agent/proxycfg-sources/catalog/mock_Watcher.go"
    agent_service_manager_go["agent/service_manager.go"
    api_partition_go["api/partition.go"
    command_registry_ce_go["command/registry_ce.go"
    internal_controller_cache_cachemock_mock_ReadOnlyCache_go["internal/controller/cache/cachemock/mock_ReadOnlyCache.go"
    internal_controller_cache_indexers_indexersmock_mock_MultiIndexer_go["internal/controller/cache/indexers/indexersmock/mock_MultiIndexer.go"
    internal_resource_stringer_go["internal/resource/stringer.go"
    internal_tools_protoc_gen_grpc_clone_e2e_proto_service_pb_go["internal/tools/protoc-gen-grpc-clone/e2e/proto/service.pb.go"
    agent_consul_state_prepared_query_index_go["agent/consul/state/prepared_query_index.go"
    agent_grpc_external_services_serverdiscovery_mock_ACLResolver_go["agent/grpc-external/services/serverdiscovery/mock_ACLResolver.go"
    grpcmocks_proto_public_pbdns_mock_UnsafeDNSServiceServer_go["grpcmocks/proto-public/pbdns/mock_UnsafeDNSServiceServer.go"
    internal_go_sso_oidcauth_jwt_go["internal/go-sso/oidcauth/jwt.go"
    internal_tools_protoc_gen_grpc_clone_internal_generate_generate_go["internal/tools/protoc-gen-grpc-clone/internal/generate/generate.go"
    proto_private_pbstorage_raft_pb_go["proto/private/pbstorage/raft.pb.go"
    proto_public_pbresource_resource_cloning_grpc_pb_go["proto-public/pbresource/resource_cloning_grpc.pb.go"
    service_os_service_windows_go["service_os/service_windows.go"
    agent_auto_config_persist_go["agent/auto-config/persist.go"
    agent_checks_check_go["agent/checks/check.go"
    agent_grpc_external_services_dataplane_server_go["agent/grpc-external/services/dataplane/server.go"
    api_agent_go["api/agent.go"
    command_cli_formatting_go["command/cli/formatting.go"
    logging_grpc_go["logging/grpc.go"
    testing_deployer_util_files_go["testing/deployer/util/files.go"
    acl_policy_authorizer_ce_go["acl/policy_authorizer_ce.go"
    agent_connect_x509_patch_go["agent/connect/x509_patch.go"
    agent_consul_tenancy_bridge_ce_go["agent/consul/tenancy_bridge_ce.go"
    agent_envoyextensions_registered_extensions_ce_go["agent/envoyextensions/registered_extensions_ce.go"
    agent_grpc_internal_balancer_balancer_go["agent/grpc-internal/balancer/balancer.go"
    internal_protohcl_unmarshal_go["internal/protohcl/unmarshal.go"
    test_integration_consul_container_libs_topology_peering_topology_go["test/integration/consul-container/libs/topology/peering_topology.go"
    agent_cache_types_resolved_service_config_go["agent/cache-types/resolved_service_config.go"
    agent_configentry_service_config_go["agent/configentry/service_config.go"
    agent_consul_enterprise_client_ce_go["agent/consul/enterprise_client_ce.go"
    agent_consul_state_txn_go["agent/consul/state/txn.go"
    proto_public_pbserverdiscovery_serverdiscovery_cloning_grpc_pb_go["proto-public/pbserverdiscovery/serverdiscovery_cloning_grpc.pb.go"
    tools_internal_grpc_proxy_main_go["tools/internal-grpc-proxy/main.go"
    troubleshoot_proxy_upstreams_go["troubleshoot/proxy/upstreams.go"
    agent_consul_state_config_entry_ce_go["agent/consul/state/config_entry_ce.go"
    agent_consul_state_mock_publishFuncType_go["agent/consul/state/mock_publishFuncType.go"
    agent_grpc_external_services_resource_write_status_go["agent/grpc-external/services/resource/write_status.go"
    agent_proxycfg_internal_watch_watchmap_go["agent/proxycfg/internal/watch/watchmap.go"
    agent_structs_config_entry_intentions_go["agent/structs/config_entry_intentions.go"
    proto_public_pbresource_annotations_deepcopy_gen_go["proto-public/pbresource/annotations_deepcopy.gen.go"
    agent_auto_config_run_go["agent/auto-config/run.go"
    agent_proxycfg_testing_upstreams_go["agent/proxycfg/testing_upstreams.go"
    command_login_login_ce_go["command/login/login_ce.go"
    internal_controller_runner_go["internal/controller/runner.go"
    agent_consul_server_log_verification_go["agent/consul/server_log_verification.go"
    agent_envoyextensions_builtin_wasm_wasm_go["agent/envoyextensions/builtin/wasm/wasm.go"
    command_connect_envoy_bootstrap_tpl_go["command/connect/envoy/bootstrap_tpl.go"
    command_peering_list_list_go["command/peering/list/list.go"
    agent_consul_state_operations_ce_go["agent/consul/state/operations_ce.go"
    agent_consul_stream_noop_go["agent/consul/stream/noop.go"
    agent_reload_go["agent/reload.go"
    command_operator_raft_listpeers_operator_raft_list_go["command/operator/raft/listpeers/operator_raft_list.go"
    testing_deployer_sprawl_resources_go["testing/deployer/sprawl/resources.go"
    agent_consul_coordinate_endpoint_go["agent/consul/coordinate_endpoint.go"
    agent_proxycfg_api_gateway_go["agent/proxycfg/api_gateway.go"
    agent_proxycfg_testing_tproxy_go["agent/proxycfg/testing_tproxy.go"
    command_peering_read_read_go["command/peering/read/read.go"
    logging_gated_writer_go["logging/gated_writer.go"
    proto_public_pbdataplane_dataplane_grpc_pb_go["proto-public/pbdataplane/dataplane_grpc.pb.go"
    test_integration_consul_container_libs_cluster_encryption_go["test/integration/consul-container/libs/cluster/encryption.go"
    test_integration_consul_container_libs_service_log_go["test/integration/consul-container/libs/service/log.go"
    agent_config_ratelimited_file_watcher_go["agent/config/ratelimited_file_watcher.go"
    agent_xds_listeners_go["agent/xds/listeners.go"
    proto_public_pbmulticluster_v2_computed_exported_services_pb_go["proto-public/pbmulticluster/v2/computed_exported_services.pb.go"
    testing_deployer_sprawl_boot_go["testing/deployer/sprawl/boot.go"
    testing_deployer_util_consul_go["testing/deployer/util/consul.go"
    agent_connect_ca_provider_vault_auth_gcp_go["agent/connect/ca/provider_vault_auth_gcp.go"
    agent_consul_configentry_backend_ce_go["agent/consul/configentry_backend_ce.go"
    api_content_type_go["api/content_type.go"
    logging_logfile_linux_go["logging/logfile_linux.go"
    proto_private_pbservice_node_gen_go["proto/private/pbservice/node.gen.go"
    agent_checks_os_service_windows_go["agent/checks/os_service_windows.go"
    agent_connect_ca_testing_go["agent/connect/ca/testing.go"
    agent_grpc_external_services_resource_testing_testing_go["agent/grpc-external/services/resource/testing/testing.go"
    agent_proxycfg_proxycfg_deepcopy_go["agent/proxycfg/proxycfg.deepcopy.go"
    agent_structs_structs_ce_go["agent/structs/structs_ce.go"
    testing_deployer_util_v2_go["testing/deployer/util/v2.go"
    acl_policy_merger_ce_go["acl/policy_merger_ce.go"
    agent_cache_types_peered_upstreams_go["agent/cache-types/peered_upstreams.go"
    agent_consul_reporting_reportingmock_mock_StateDelegate_go["agent/consul/reporting/reportingmock/mock_StateDelegate.go"
    agent_hcp_telemetry_gauge_store_go["agent/hcp/telemetry/gauge_store.go"
    api_health_go["api/health.go"
    command_acl_token_delete_token_delete_go["command/acl/token/delete/token_delete.go"
    command_operator_raft_removepeer_operator_raft_remove_go["command/operator/raft/removepeer/operator_raft_remove.go"
    internal_controller_supervisor_go["internal/controller/supervisor.go"
    agent_consul_rate_mock_RequestLimitsHandler_go["agent/consul/rate/mock_RequestLimitsHandler.go"
    agent_structs_txn_go["agent/structs/txn.go"
    internal_controller_testing_go["internal/controller/testing.go"
    internal_multicluster_internal_types_exported_services_go["internal/multicluster/internal/types/exported_services.go"
    internal_resource_decode_go["internal/resource/decode.go"
    internal_tools_protoc_gen_consul_rate_limit_postprocess_main_go["internal/tools/protoc-gen-consul-rate-limit/postprocess/main.go"
    lib_semaphore_semaphore_go["lib/semaphore/semaphore.go"
    proto_private_pbautoconf_auto_config_ce_go["proto/private/pbautoconf/auto_config_ce.go"
    agent_cache_request_go["agent/cache/request.go"
    agent_config_doc_go["agent/config/doc.go"
    agent_connect_ca_provider_vault_auth_alicloud_go["agent/connect/ca/provider_vault_auth_alicloud.go"
    internal_controller_cache_cachemock_mock_WriteCache_go["internal/controller/cache/cachemock/mock_WriteCache.go"
    internal_storage_inmem_event_index_go["internal/storage/inmem/event_index.go"
    test_integ_upgrade_basic_common_go["test-integ/upgrade/basic/common.go"
    testing_deployer_util_v2_decode_go["testing/deployer/util/v2_decode.go"
    agent_connect_ca_provider_vault_auth_azure_go["agent/connect/ca/provider_vault_auth_azure.go"
    agent_grpc_internal_handler_go["agent/grpc-internal/handler.go"
    agent_xds_rbac_go["agent/xds/rbac.go"
    proto_public_pbmulticluster_v2_namespace_exported_services_deepcopy_gen_go["proto-public/pbmulticluster/v2/namespace_exported_services_deepcopy.gen.go"
    command_intention_match_match_go["command/intention/match/match.go"
    command_snapshot_decode_snapshot_decode_go["command/snapshot/decode/snapshot_decode.go"
    internal_resourcehcl_any_go["internal/resourcehcl/any.go"
    proto_public_pbmulticluster_v2_exported_services_pb_go["proto-public/pbmulticluster/v2/exported_services.pb.go"
    proto_public_pbmulticluster_v2_exported_services_consumer_deepcopy_gen_go["proto-public/pbmulticluster/v2/exported_services_consumer_deepcopy.gen.go"
    proto_public_pbdns_dns_pb_go["proto-public/pbdns/dns.pb.go"
    testing_deployer_topology_util_go["testing/deployer/topology/util.go"
    agent_consul_fsm_decode_downgrade_go["agent/consul/fsm/decode_downgrade.go"
    agent_grpc_external_testutils_fsm_go["agent/grpc-external/testutils/fsm.go"
    agent_keyring_go["agent/keyring.go"
    internal_resourcehcl_naming_go["internal/resourcehcl/naming.go"
    acl_testing_go["acl/testing.go"
    agent_consul_state_state_store_go{"agent/consul/state/state_store.go"
    agent_kvs_endpoint_go["agent/kvs_endpoint.go"
    agent_structs_testing_catalog_go["agent/structs/testing_catalog.go"
    grpcmocks_proto_public_pbconnectca_mock_ConnectCAServiceClient_go["grpcmocks/proto-public/pbconnectca/mock_ConnectCAServiceClient.go"
    lib_maps_maps_go["lib/maps/maps.go"
    testing_deployer_sprawl_acl_go["testing/deployer/sprawl/acl.go"
    testing_deployer_sprawl_tls_go["testing/deployer/sprawl/tls.go"
    agent_connect_generate_go["agent/connect/generate.go"
    agent_grpc_external_testutils_server_go["agent/grpc-external/testutils/server.go"
    agent_proxycfg_glue_discovery_chain_go["agent/proxycfg-glue/discovery_chain.go"
    agent_proxycfg_glue_federation_state_list_mesh_gateways_go["agent/proxycfg-glue/federation_state_list_mesh_gateways.go"
    internal_resource_resourcetest_require_go["internal/resource/resourcetest/require.go"
    agent_cache_types_connect_ca_root_go["agent/cache-types/connect_ca_root.go"
    agent_config_segment_ce_go["agent/config/segment_ce.go"
    agent_configentry_resolve_go["agent/configentry/resolve.go"
    agent_consul_state_catalog_schema_deepcopy_go["agent/consul/state/catalog_schema.deepcopy.go"
    command_lock_lock_go["command/lock/lock.go"
    sdk_testutil_retry_retryer_go["sdk/testutil/retry/retryer.go"
    agent_cache_types_federation_state_list_gateways_go["agent/cache-types/federation_state_list_gateways.go"
    agent_consul_session_ttl_go["agent/consul/session_ttl.go"
    command_intention_delete_delete_go["command/intention/delete/delete.go"
    command_peering_delete_delete_go["command/peering/delete/delete.go"
    connect_resolver_go["connect/resolver.go"
    envoyextensions_xdscommon_proxysupport_go["envoyextensions/xdscommon/proxysupport.go"
    sentinel_evaluator_go["sentinel/evaluator.go"
    sdk_testutil_server_go["sdk/testutil/server.go"
    testing_deployer_sprawl_grpc_go["testing/deployer/sprawl/grpc.go"
    agent_connect_ca_provider_vault_auth_k8s_go["agent/connect/ca/provider_vault_auth_k8s.go"
    agent_consul_state_config_entry_exported_services_ce_go["agent/consul/state/config_entry_exported_services_ce.go"
    agent_consul_state_events_go["agent/consul/state/events.go"
    command_resource_client_client_go["command/resource/client/client.go"
    testing_deployer_topology_compile_go["testing/deployer/topology/compile.go"
    acl_chained_authorizer_go["acl/chained_authorizer.go"
    agent_consul_stream_event_buffer_go["agent/consul/stream/event_buffer.go"
    agent_consul_stats_fetcher_go["agent/consul/stats_fetcher.go"
    agent_structs_intention_go["agent/structs/intention.go"
    internal_controller_controller_go["internal/controller/controller.go"
    internal_storage_inmem_backend_go["internal/storage/inmem/backend.go"
    agent_envoyextensions_builtin_otel_access_logging_structs_go["agent/envoyextensions/builtin/otel-access-logging/structs.go"
    agent_proxycfg_glue_exported_peered_services_go["agent/proxycfg-glue/exported_peered_services.go"
    internal_gossip_libserf_serf_go["internal/gossip/libserf/serf.go"
    agent_leafcert_generate_go["agent/leafcert/generate.go"
    command_acl_token_clone_token_clone_go["command/acl/token/clone/token_clone.go"
    internal_controller_cache_index_errors_go["internal/controller/cache/index/errors.go"
    internal_controller_controllermock_mock_task_go["internal/controller/controllermock/mock_task.go"
    acl_errors_ce_go["acl/errors_ce.go"
    agent_netutil_network_go["agent/netutil/network.go"
    agent_metrics_testing_go["agent/metrics/testing.go"
    agent_structs_config_entry_sameness_group_ce_go["agent/structs/config_entry_sameness_group_ce.go"
    agent_xds_configfetcher_config_fetcher_go["agent/xds/configfetcher/config_fetcher.go"
    api_operator_audit_go["api/operator_audit.go"
    internal_controller_cache_kind_go["internal/controller/cache/kind.go"
    agent_connect_ca_provider_vault_auth_approle_go["agent/connect/ca/provider_vault_auth_approle.go"
    agent_consul_acl_endpoint_ce_go["agent/consul/acl_endpoint_ce.go"
    agent_grpc_middleware_testutil_testservice_simple_pb_go["agent/grpc-middleware/testutil/testservice/simple.pb.go"
    agent_structs_connect_proxy_config_go["agent/structs/connect_proxy_config.go"
    api_connect_ca_go["api/connect_ca.go"
    api_txn_go["api/txn.go"
    internal_controller_cache_index_indexmock_mock_MultiIndexer_go["internal/controller/cache/index/indexmock/mock_MultiIndexer.go"
    agent_consul_catalog_endpoint_go["agent/consul/catalog_endpoint.go"
    agent_consul_intention_endpoint_go["agent/consul/intention_endpoint.go"
    logging_logfile_openbsd_go["logging/logfile_openbsd.go"
    agent_consul_server_serf_go["agent/consul/server_serf.go"
    agent_xds_accesslogs_accesslogs_go["agent/xds/accesslogs/accesslogs.go"
    command_resource_client_usage_go["command/resource/client/usage.go"
    internal_controller_cache_decoded_go["internal/controller/cache/decoded.go"
    agent_consul_state_federation_state_go["agent/consul/state/federation_state.go"
    agent_proxycfg_testing_api_gateway_go["agent/proxycfg/testing_api_gateway.go"
    agent_rpcclient_health_health_go["agent/rpcclient/health/health.go"
    agent_structs_config_entry_discoverychain_ce_go["agent/structs/config_entry_discoverychain_ce.go"
    testrpc_wait_go["testrpc/wait.go"
    agent_connect_ca_mock_Provider_go["agent/connect/ca/mock_Provider.go"
    agent_consul_type_registry_go["agent/consul/type_registry.go"
    agent_structs_errors_go["agent/structs/errors.go"
    command_connect_ca_set_connect_ca_set_go["command/connect/ca/set/connect_ca_set.go"
    connect_certgen_certgen_go["connect/certgen/certgen.go"
    proto_private_pbcommon_common_ce_go["proto/private/pbcommon/common_ce.go"
    acl_policy_ce_go["acl/policy_ce.go"
    command_cli_cli_go["command/cli/cli.go"
    test_integration_consul_container_libs_assert_peering_go["test/integration/consul-container/libs/assert/peering.go"
    agent_consul_authmethod_kubeauth_k8s_ce_go["agent/consul/authmethod/kubeauth/k8s_ce.go"
    agent_consul_options_go["agent/consul/options.go"
    agent_consul_v2_config_entry_exports_shim_go["agent/consul/v2_config_entry_exports_shim.go"
    agent_operator_endpoint_ce_go["agent/operator_endpoint_ce.go"
    command_operator_autopilot_state_operator_autopilot_state_go["command/operator/autopilot/state/operator_autopilot_state.go"
    lib_hoststats_metrics_go["lib/hoststats/metrics.go"
    test_integration_consul_container_libs_service_connect_go["test/integration/consul-container/libs/service/connect.go"
    agent_cache_types_trust_bundle_go["agent/cache-types/trust_bundle.go"
    agent_cache_types_trust_bundles_go["agent/cache-types/trust_bundles.go"
    agent_grpc_middleware_rate_limit_mappings_gen_go["agent/grpc-middleware/rate_limit_mappings.gen.go"
    agent_proxycfg_testing_ingress_gateway_go["agent/proxycfg/testing_ingress_gateway.go"
    agent_structs_config_entry_exports_ce_go["agent/structs/config_entry_exports_ce.go"
    api_prepared_query_go["api/prepared_query.go"
    command_peering_generate_generate_go["command/peering/generate/generate.go"
    command_version_formatter_go["command/version/formatter.go"
    internal_multicluster_internal_types_partition_exported_services_go["internal/multicluster/internal/types/partition_exported_services.go"
    internal_resource_resourcetest_validation_go["internal/resource/resourcetest/validation.go"
    agent_cache_types_mock_RPC_go["agent/cache-types/mock_RPC.go"
    agent_consul_leader_log_verification_go["agent/consul/leader_log_verification.go"
    agent_grpc_external_services_configentry_server_go["agent/grpc-external/services/configentry/server.go"
    agent_grpc_external_utils_go["agent/grpc-external/utils.go"
    agent_hcp_telemetry_otel_sink_go["agent/hcp/telemetry/otel_sink.go"
    agent_proxycfg_glue_intentions_ce_go["agent/proxycfg-glue/intentions_ce.go"
    agent_xds_failover_policy_ce_go["agent/xds/failover_policy_ce.go"
    agent_xds_protocol_trace_go["agent/xds/protocol_trace.go"
    agent_cache_types_peerings_go["agent/cache-types/peerings.go"
    agent_checks_docker_unix_go["agent/checks/docker_unix.go"
    agent_consul_acl_ce_go["agent/consul/acl_ce.go"
    agent_consul_reporting_reporting_go["agent/consul/reporting/reporting.go"
    agent_util_go["agent/util.go"
    command_acl_authmethod_read_authmethod_read_go["command/acl/authmethod/read/authmethod_read.go"
    command_resource_apply_grpc_apply_go["command/resource/apply-grpc/apply.go"
    command_resource_resource_grpc_go["command/resource/resource-grpc.go"
    agent_consul_acl_replication_types_go["agent/consul/acl_replication_types.go"
    agent_consul_authmethod_kubeauth_testing_go["agent/consul/authmethod/kubeauth/testing.go"
    agent_consul_session_timers_go["agent/consul/session_timers.go"
    agent_grpc_middleware_handshake_go["agent/grpc-middleware/handshake.go"
    internal_dnsutil_dns_go["internal/dnsutil/dns.go"
    proto_private_pbacl_acl_go["proto/private/pbacl/acl.go"
    proto_public_pbconnectca_ca_json_gen_go["proto-public/pbconnectca/ca_json.gen.go"
    proto_private_pbservice_service_pb_go["proto/private/pbservice/service.pb.go"
    agent_consul_auth_mock_ACLCache_go["agent/consul/auth/mock_ACLCache.go"
    agent_hcp_testserver_main_go["agent/hcp/testserver/main.go"
    agent_structs_acl_cache_go["agent/structs/acl_cache.go"
    command_intention_get_get_go["command/intention/get/get.go"
    command_resource_list_list_go["command/resource/list/list.go"
    lib_telemetry_go["lib/telemetry.go"
    proto_private_pbdemo_v2_resources_rtypes_go["proto/private/pbdemo/v2/resources.rtypes.go"
    test_integration_consul_container_libs_assert_envoy_go["test/integration/consul-container/libs/assert/envoy.go"
    agent_consul_authmethod_kubeauth_k8s_go["agent/consul/authmethod/kubeauth/k8s.go"
    agent_consul_state_kvs_go["agent/consul/state/kvs.go"
    agent_structs_testing_service_definition_go["agent/structs/testing_service_definition.go"
    api_config_entry_rate_limit_ip_go["api/config_entry_rate_limit_ip.go"
    api_operator_go["api/operator.go"
    connect_testing_go["connect/testing.go"
    grpcmocks_proto_public_pbserverdiscovery_mock_ServerDiscoveryServiceClient_go["grpcmocks/proto-public/pbserverdiscovery/mock_ServerDiscoveryServiceClient.go"
    internal_resource_acls_go["internal/resource/acls.go"
    agent_connect_ca_provider_aws_go["agent/connect/ca/provider_aws.go"
    agent_rpcclient_configentry_view_go["agent/rpcclient/configentry/view.go"
    agent_session_endpoint_go["agent/session_endpoint.go"
    command_kv_kv_go["command/kv/kv.go"
    internal_controller_controllermock_mock_Initializer_go["internal/controller/controllermock/mock_Initializer.go"
    sdk_testutil_testlog_go["sdk/testutil/testlog.go"
    testing_deployer_sprawl_consul_go["testing/deployer/sprawl/consul.go"
    troubleshoot_ports_troubleshoot_ports_go["troubleshoot/ports/troubleshoot_ports.go"
    acl_authorizer_go["acl/authorizer.go"
    agent_structs_testing_connect_proxy_config_go["agent/structs/testing_connect_proxy_config.go"
    command_members_members_go["command/members/members.go"
    proto_private_pbstorage_raft_grpc_pb_go["proto/private/pbstorage/raft_grpc.pb.go"
    proto_public_pbmulticluster_v2_resources_rtypes_go["proto-public/pbmulticluster/v2/resources.rtypes.go"
    test_integration_consul_container_libs_service_gateway_go["test/integration/consul-container/libs/service/gateway.go"
    test_integration_consul_container_libs_utils_retry_go["test/integration/consul-container/libs/utils/retry.go"
    agent_cache_cache_go["agent/cache/cache.go"
    agent_connect_uri_go["agent/connect/uri.go"
    agent_http_ce_go["agent/http_ce.go"
    agent_structs_config_entry_mesh_go["agent/structs/config_entry_mesh.go"
    command_acl_authmethod_delete_authmethod_delete_go["command/acl/authmethod/delete/authmethod_delete.go"
    command_connect_envoy_envoy_go["command/connect/envoy/envoy.go"
    command_flags_flag_map_value_go["command/flags/flag_map_value.go"
    proto_public_pbacl_acl_pb_go["proto-public/pbacl/acl.pb.go"
    agent_connect_ca_provider_vault_auth_aws_go["agent/connect/ca/provider_vault_auth_aws.go"
    agent_proxycfg_glue_peered_upstreams_go["agent/proxycfg-glue/peered_upstreams.go"
    agent_testagent_go["agent/testagent.go"
    command_kv_impexp_kvimpexp_go["command/kv/impexp/kvimpexp.go"
    envoyextensions_extensioncommon_runtime_config_go["envoyextensions/extensioncommon/runtime_config.go"
    internal_controller_cache_clone_go["internal/controller/cache/clone.go"
    test_integration_consul_container_libs_cluster_log_go["test/integration/consul-container/libs/cluster/log.go"
    agent_consul_discoverychain_testing_go["agent/consul/discoverychain/testing.go"
    agent_checks_os_service_unix_go["agent/checks/os_service_unix.go"
    agent_config_flags_go["agent/config/flags.go"
    agent_configentry_merge_service_config_go["agent/configentry/merge_service_config.go"
    internal_controller_cache_index_indexmock_mock_SingleIndexer_go["internal/controller/cache/index/indexmock/mock_SingleIndexer.go"
    internal_resourcehcl_unmarshal_go["internal/resourcehcl/unmarshal.go"
    test_integration_consul_container_libs_cluster_agent_go["test/integration/consul-container/libs/cluster/agent.go"
    testing_deployer_sprawl_sprawltest_sprawltest_go["testing/deployer/sprawl/sprawltest/sprawltest.go"
    agent_proxycfg_mesh_gateway_ce_go["agent/proxycfg/mesh_gateway_ce.go"
    command_acl_policy_list_policy_list_go["command/acl/policy/list/policy_list.go"
    connect_tls_go["connect/tls.go"
    internal_multicluster_internal_types_helpers_go["internal/multicluster/internal/types/helpers.go"
    proto_private_pbpeering_peering_grpc_pb_go["proto/private/pbpeering/peering_grpc.pb.go"
    types_area_go["types/area.go"
    agent_cache_types_prepared_query_go["agent/cache-types/prepared_query.go"
    internal_storage_raft_backend_go["internal/storage/raft/backend.go"
    lib_uuid_go["lib/uuid.go"
    proto_public_pbmulticluster_v2_exported_services_json_gen_go["proto-public/pbmulticluster/v2/exported_services_json.gen.go"
    test_integration_consul_container_libs_service_examples_go["test/integration/consul-container/libs/service/examples.go"
    agent_grpc_external_services_resource_write_go["agent/grpc-external/services/resource/write.go"
    internal_controller_cache_cache_go["internal/controller/cache/cache.go"
    grpcmocks_proto_public_pbdataplane_mock_UnsafeDataplaneServiceServer_go["grpcmocks/proto-public/pbdataplane/mock_UnsafeDataplaneServiceServer.go"
    internal_protohcl_naming_go["internal/protohcl/naming.go"
    sdk_testutil_retry_run_go["sdk/testutil/retry/run.go"
    agent_cache_types_intention_match_go["agent/cache-types/intention_match.go"
    agent_cache_types_catalog_list_services_go["agent/cache-types/catalog_list_services.go"
    agent_consul_auto_config_backend_go["agent/consul/auto_config_backend.go"
    agent_consul_controller_queue_rate_go["agent/consul/controller/queue/rate.go"
    agent_consul_operator_raft_endpoint_go["agent/consul/operator_raft_endpoint.go"
    agent_proxycfg_glue_peering_list_go["agent/proxycfg-glue/peering_list.go"
    envoyextensions_extensioncommon_resources_go["envoyextensions/extensioncommon/resources.go"
    internal_resource_reaper_controller_go["internal/resource/reaper/controller.go"
    agent_consul_client_go["agent/consul/client.go"
    agent_structs_config_entry_apigw_jwt_ce_go["agent/structs/config_entry_apigw_jwt_ce.go"
    agent_structs_identity_go["agent/structs/identity.go"
    internal_go_sso_oidcauth_oidcjwt_go["internal/go-sso/oidcauth/oidcjwt.go"
    sdk_freeport_ephemeral_darwin_go["sdk/freeport/ephemeral_darwin.go"
    testing_deployer_sprawl_internal_tfgen_digest_go["testing/deployer/sprawl/internal/tfgen/digest.go"
    test_integration_consul_container_test_resource_http_api_helper_go["test/integration/consul-container/test/resource/http_api/helper.go"
    troubleshoot_ports_troubleshoot_protocol_go["troubleshoot/ports/troubleshoot_protocol.go"
    agent_consul_fsm_decode_ce_go["agent/consul/fsm/decode_ce.go"
    agent_consul_discoverychain_gateway_go["agent/consul/discoverychain/gateway.go"
    agent_exec_exec_unix_go["agent/exec/exec_unix.go"
    agent_structs_acl_ce_go["agent/structs/acl_ce.go"
    agent_structs_config_entry_file_system_certificate_go["agent/structs/config_entry_file_system_certificate.go"
    command_acl_policy_read_policy_read_go["command/acl/policy/read/policy_read.go"
    internal_protohcl_cty_go["internal/protohcl/cty.go"
    internal_resource_resourcetest_tenancy_go["internal/resource/resourcetest/tenancy.go"
    agent_consul_discovery_chain_endpoint_go["agent/consul/discovery_chain_endpoint.go"
    agent_grpc_middleware_testutil_testservice_simple_grpc_pb_go["agent/grpc-middleware/testutil/testservice/simple_grpc.pb.go"
    command_acl_role_formatter_go["command/acl/role/formatter.go"
    agent_consul_acl_authmethod_go["agent/consul/acl_authmethod.go"
    agent_grpc_internal_resolver_registry_go["agent/grpc-internal/resolver/registry.go"
    envoyextensions_xdscommon_envoy_versioning_go["envoyextensions/xdscommon/envoy_versioning.go"
    testing_deployer_sprawl_internal_tfgen_tfgen_go["testing/deployer/sprawl/internal/tfgen/tfgen.go"
    testing_deployer_sprawl_sprawl_go["testing/deployer/sprawl/sprawl.go"
    agent_auto_config_config_translate_go["agent/auto-config/config_translate.go"
    agent_cache_testing_go["agent/cache/testing.go"
    api_api_go["api/api.go"
    command_intention_format_go["command/intention/format.go"
    internal_controller_dependency_transform_go["internal/controller/dependency/transform.go"
    internal_resource_protoc_gen_resource_types_main_go["internal/resource/protoc-gen-resource-types/main.go"
    lib_strings_go["lib/strings.go"
    proto_private_pbcommon_common_pb_go["proto/private/pbcommon/common.pb.go"
    agent_consul_flood_go["agent/consul/flood.go"
    agent_consul_fsm_commands_ce_go["agent/consul/fsm/commands_ce.go"
    agent_consul_server_metadata_go["agent/consul/server_metadata.go"
    agent_grpc_internal_services_subscribe_logger_go["agent/grpc-internal/services/subscribe/logger.go"
    agent_submatview_materializer_go["agent/submatview/materializer.go"
    proto_public_pbresource_cloning_stream_pb_go["proto-public/pbresource/cloning_stream.pb.go"
    sdk_freeport_ephemeral_fallback_go["sdk/freeport/ephemeral_fallback.go"
    agent_consul_context_go["agent/consul/context.go"
    agent_proxycfg_glue_intention_upstreams_go["agent/proxycfg-glue/intention_upstreams.go"
    command_acl_agenttokens_agent_tokens_go["command/acl/agenttokens/agent_tokens.go"
    command_operator_usage_instances_usage_instances_go["command/operator/usage/instances/usage_instances.go"
    internal_controller_cache_index_indexmock_mock_Indexer_go["internal/controller/cache/index/indexmock/mock_Indexer.go"
    proto_public_pbserverdiscovery_cloning_stream_pb_go["proto-public/pbserverdiscovery/cloning_stream.pb.go"
    agent_hcp_deps_go["agent/hcp/deps.go"
    command_connect_envoy_exec_supported_go["command/connect/envoy/exec_supported.go"
    command_connect_proxy_proxy_go["command/connect/proxy/proxy.go"
    grpcmocks_proto_public_pbconnectca_mock_UnsafeConnectCAServiceServer_go["grpcmocks/proto-public/pbconnectca/mock_UnsafeConnectCAServiceServer.go"
    proto_private_pbpeering_peering_go["proto/private/pbpeering/peering.go"
    agent_config_runtime_ce_go["agent/config/runtime_ce.go"
    agent_consul_filter_go["agent/consul/filter.go"
    agent_consul_leader_federation_state_ae_go["agent/consul/leader_federation_state_ae.go"
    grpcmocks_proto_public_pbdataplane_mock_DataplaneServiceServer_go["grpcmocks/proto-public/pbdataplane/mock_DataplaneServiceServer.go"
    internal_controller_lease_go["internal/controller/lease.go"
    proto_public_pbacl_acl_grpc_pb_go["proto-public/pbacl/acl_grpc.pb.go"
    sdk_freeport_freeport_go["sdk/freeport/freeport.go"
    agent_consul_server_connect_go["agent/consul/server_connect.go"
    agent_pool_peeked_conn_go["agent/pool/peeked_conn.go"
    agent_structs_operator_go["agent/structs/operator.go"
    command_acl_templatedpolicy_read_templated_policy_read_go["command/acl/templatedpolicy/read/templated_policy_read.go"
    command_flags_http_go["command/flags/http.go"
    command_operator_usage_instances_usage_instances_ce_go["command/operator/usage/instances/usage_instances_ce.go"
    proto_private_pbconfigentry_config_entry_ce_go["proto/private/pbconfigentry/config_entry_ce.go"
    agent_connect_ca_provider_vault_auth_go["agent/connect/ca/provider_vault_auth.go"
    agent_connect_uri_server_go["agent/connect/uri_server.go"
    agent_connect_ca_endpoint_go["agent/connect_ca_endpoint.go"
    agent_consul_authmethod_ssoauth_sso_ce_go["agent/consul/authmethod/ssoauth/sso_ce.go"
    agent_proxycfg_config_snapshot_glue_go["agent/proxycfg/config_snapshot_glue.go"
    tlsutil_generate_go["tlsutil/generate.go"
    agent_connect_authz_go["agent/connect/authz.go"
    agent_consul_state_config_entry_events_go["agent/consul/state/config_entry_events.go"
    agent_proxycfg_testing_ce_go["agent/proxycfg/testing_ce.go"
    agent_xds_testcommon_testcommon_go["agent/xds/testcommon/testcommon.go"
    proto_private_pbacl_acl_pb_go["proto/private/pbacl/acl.pb.go"
    agent_rpc_operator_service_go["agent/rpc/operator/service.go"
    command_snapshot_inspect_formatter_go["command/snapshot/inspect/formatter.go"
    proto_private_pbsubscribe_subscribe_pb_go["proto/private/pbsubscribe/subscribe.pb.go"
    agent_config_limits_go["agent/config/limits.go"
    agent_connect_ca_provider_go["agent/connect/ca/provider.go"
    agent_consul_leader_intentions_go["agent/consul/leader_intentions.go"
    grpcmocks_proto_public_pbresource_mock_ResourceService_WatchListClient_go["grpcmocks/proto-public/pbresource/mock_ResourceService_WatchListClient.go"
    agent_ui_endpoint_go["agent/ui_endpoint.go"
    internal_controller_cache_cachemock_mock_Cache_go["internal/controller/cache/cachemock/mock_Cache.go"
    internal_controller_dependencies_go["internal/controller/dependencies.go"
    testing_deployer_sprawl_internal_tfgen_nodes_go["testing/deployer/sprawl/internal/tfgen/nodes.go"
    agent_consul_leader_connect_go["agent/consul/leader_connect.go"
    agent_auto_config_config_go["agent/auto-config/config.go"
    agent_connect_ca_provider_consul_config_go["agent/connect/ca/provider_consul_config.go"
    agent_consul_stream_string_types_go["agent/consul/stream/string_types.go"
    agent_catalog_endpoint_go["agent/catalog_endpoint.go"
    agent_envoyextensions_builtin_property_override_structpatcher_go["agent/envoyextensions/builtin/property-override/structpatcher.go"
    agent_grpc_external_services_connectca_server_go["agent/grpc-external/services/connectca/server.go"
    api_operator_autopilot_go["api/operator_autopilot.go"
    command_snapshot_inspect_snapshot_inspect_go["command/snapshot/inspect/snapshot_inspect.go"
    test_integration_consul_container_libs_service_common_go["test/integration/consul-container/libs/service/common.go"
    test_integration_consul_container_libs_cluster_dataplane_go["test/integration/consul-container/libs/cluster/dataplane.go"
    testing_deployer_sprawl_ent_go["testing/deployer/sprawl/ent.go"
    agent_consul_segment_ce_go["agent/consul/segment_ce.go"
    agent_proxycfg_manager_go["agent/proxycfg/manager.go"
    api_snapshot_go["api/snapshot.go"
    command_acl_templatedpolicy_templated_policy_go["command/acl/templatedpolicy/templated_policy.go"
    command_kv_del_kv_delete_go["command/kv/del/kv_delete.go"


    classDef style0 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style1 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style2 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style3 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style4 fill:#4CAF50,stroke:#333,stroke-width:2px
    classDef style5 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style6 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style7 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style8 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style9 fill:#90CAF9,stroke:#333,stroke-width:2px
    classDef style10 fill:#90CAF9,stroke:#333,stroke-width:2px

    class command_connect_ca_ca_go style1
    class connect_proxy_config_go style2
    class grpcmocks_proto_public_pbdns_mock_DNSServiceServer_go style7
    class agent_consul_state_session_go style3
    class agent_hcp_client_errors_go style3
    class agent_structs_config_entry_exports_go style3
    class command_flags_flag_slice_value_go style1
    class command_kv_put_kv_put_go style1
    class internal_go_sso_oidcauth_util_go style7
    class agent_configentry_compare_go style3
    class agent_grpc_external_services_peerstream_server_go style8
    class agent_grpc_middleware_testutil_testservice_fake_service_go style3
    class agent_proxycfg_glue_glue_go style3
    class agent_structs_acl_templated_policy_go style3
    class agent_consul_enterprise_server_ce_go style3
    class agent_envoyextensions_builtin_wasm_structs_go style3
    class agent_prepared_query_endpoint_go style3
    class agent_proxycfg_state_go style3
    class internal_controller_controllermock_mock_CacheIDModifier_go style7
    class internal_multicluster_internal_types_types_go style7
    class testing_deployer_topology_relationships_go style7
    class agent_cache_types_catalog_services_go style3
    class agent_consul_health_endpoint_ce_go style3
    class agent_consul_state_query_ce_go style3
    class agent_hcp_scada_scada_go style3
    class agent_proxycfg_sources_local_mock_ConfigManager_go style3
    class command_acl_acl_go style1
    class command_acl_policy_policy_go style1
    class envoyextensions_extensioncommon_basic_envoy_extender_go style7
    class agent_consul_controller_controller_go style3
    class agent_peering_endpoint_go style3
    class command_intention_check_check_go style1
    class internal_controller_cache_cachemock_mock_Query_go style7
    class grpcmocks_proto_public_pbdns_mock_DNSServiceClient_go style7
    class lib_eof_go style7
    class proto_public_annotations_ratelimit_ratelimit_json_gen_go style7
    class agent_consul_operator_backend_go style3
    class agent_grpc_internal_pipe_go style3
    class command_connect_expose_expose_go style1
    class proto_public_pbmulticluster_v2_partition_exported_services_json_gen_go style7
    class agent_check_go style3
    class agent_connect_csr_go style3
    class agent_hcp_config_mock_CloudConfig_go style3
    class lib_ttlcache_eviction_go style7
    class proto_private_pbpeerstream_convert_go style7
    class proto_public_pbmulticluster_v2_exported_services_consumer_json_gen_go style7
    class test_integration_consul_container_libs_service_service_go style7
    class agent_connect_common_names_go style3
    class agent_consul_server_overview_go style3
    class agent_event_endpoint_go style3
    class agent_proxycfg_glue_service_http_checks_go style3
    class agent_submatview_store_go style3
    class command_acl_token_list_token_list_go style1
    class command_kv_exp_kv_export_go style1
    class proto_private_pbconfigentry_config_entry_grpc_pb_go style7
    class agent_consul_gateways_controller_gateways_ce_go style3
    class agent_uiserver_ui_template_data_go style3
    class command_acl_role_delete_role_delete_go style1
    class internal_resource_tenancy_go style7
    class command_connect_ca_get_connect_ca_get_go style1
    class command_config_read_config_read_go style1
    class envoyextensions_extensioncommon_upstream_envoy_extender_go style7
    class internal_controller_cache_cachemock_mock_ResourceIterator_go style7
    class internal_controller_dependency_cache_go style7
    class testing_deployer_sprawl_peering_go style7
    class agent_consul_server_lookup_go style3
    class agent_consul_state_delay_ce_go style3
    class agent_grpc_external_services_acl_server_go style3
    class agent_connect_uri_mesh_gateway_go style3
    class agent_consul_internal_endpoint_go style3
    class agent_grpc_middleware_recovery_go style3
    class agent_hcp_mock_Manager_go style3
    class agent_structs_auto_encrypt_go style3
    class agent_xds_locality_policy_go style3
    class command_connect_redirecttraffic_redirect_traffic_go style1
    class internal_controller_manager_go style7
    class agent_health_endpoint_go style3
    class agent_proxycfg_sources_local_config_source_go style3
    class api_namespace_go style4
    class command_lock_util_unix_go style1
    class internal_resource_mappers_bimapper_bimapper_go style7
    class lib_map_walker_go style7
    class proto_private_pbservice_service_gen_go style7
    class agent_configentry_discoverychain_go style3
    class agent_dns_ce_go style3
    class agent_proxycfg_data_sources_ce_go style3
    class agent_uiserver_buffered_file_go style3
    class test_integ_peering_commontopo_commontopo_go style7
    class agent_config_runtime_go style3
    class agent_consul_logging_go style3
    class agent_proxycfg_terminating_gateway_go style3
    class command_resource_client_helper_go style1
    class agent_config_limits_windows_go style3
    class agent_consul_server_grpc_go style3
    class agent_consul_state_config_entry_schema_go style3
    class internal_testing_golden_golden_go style7
    class testing_deployer_sprawl_acl_rules_go style7
    class agent_consul_config_go style3
    class agent_consul_xdscapacity_capacity_go style3
    class command_operator_utilization_utilization_go style1
    class internal_gossip_librtt_rtt_go style7
    class internal_resource_authz_ce_go style7
    class internal_multicluster_internal_controllers_v1compat_mock_AggregatedConfig_go style7
    class sdk_testutil_retry_timer_go style7
    class testing_deployer_topology_topology_go style7
    class agent_cache_types_service_dump_go style3
    class agent_xds_jwt_authn_go style3
    class api_config_entry_file_system_certificate_go style4
    class command_intention_intention_go style1
    class internal_controller_controllermock_mock_Reconciler_go style7
    class internal_tools_protoc_gen_grpc_clone_e2e_proto_service_grpc_pb_go style7
    class proto_public_pbdns_dns_deepcopy_gen_go style7
    class agent_envoyextensions_builtin_lua_lua_go style3
    class agent_hcp_discover_discover_go style3
    class agent_proxycfg_data_sources_go style3
    class types_checks_go style4
    class testing_deployer_sprawl_internal_tfgen_proxy_go style7
    class agent_consul_connect_ca_endpoint_go style3
    class agent_consul_state_coordinate_ce_go style3
    class agent_consul_state_intention_ce_go style3
    class agent_consul_state_config_entry_intention_go style3
    class agent_retry_join_go style3
    class command_acl_authmethod_list_authmethod_list_go style1
    class agent_checks_docker_windows_go style3
    class agent_metadata_server_go style3
    class agent_pool_conn_go style3
    class command_acl_role_role_go style1
    class command_kv_imp_kv_import_go style1
    class command_operator_usage_usage_go style1
    class internal_protohcl_doc_go style7
    class proto_private_pbdemo_v2_demo_pb_go style7
    class api_coordinate_go style4
    class envoyextensions_xdscommon_xdscommon_go style7
    class internal_controller_cache_index_indexmock_mock_ResourceIterator_go style7
    class proto_public_pbdns_dns_cloning_grpc_pb_go style7
    class sdk_freeport_systemlimit_go style7
    class agent_structs_testing_go style3
    class sdk_testutil_server_wrapper_go style7
    class agent_rpcclient_common_go style3
    class agent_xds_naming_naming_go style3
    class proto_public_pbresource_resource_pb_go style7
    class agent_grpc_external_options_go style3
    class command_peering_exportedservices_exported_services_go style1
    class grpcmocks_proto_public_pbdataplane_mock_isGetEnvoyBootstrapParamsRequest_NodeSpec_go style7
    class agent_consul_leader_peering_go style3
    class agent_structs_config_entry_mesh_ce_go style3
    class api_catalog_go style4
    class command_catalog_catalog_go style1
    class testing_deployer_sprawl_internal_secrets_store_go style7
    class types_node_id_go style4
    class command_exec_exec_go style1
    class command_kv_get_kv_get_go style1
    class agent_consul_authmethod_testing_go style3
    class agent_consul_serf_filter_go style3
    class agent_grpc_external_services_resource_mock_TenancyBridge_go style3
    class agent_token_persistence_go style3
    class api_config_entry_go style4
    class proto_private_pboperator_operator_pb_go style7
    class acl_policy_go style9
    class agent_auto_config_tls_go style3
    class agent_consul_prepared_query_template_go style3
    class agent_consul_state_memdb_go style3
    class agent_grpc_external_services_resource_watch_go style3
    class agent_hcp_scada_mock_Provider_go style3
    class agent_structs_service_definition_go style4
    class proto_public_pbacl_acl_json_gen_go style7
    class agent_consul_snapshot_endpoint_go style3
    class agent_leafcert_structs_go style3
    class agent_structs_discovery_chain_ce_go style3
    class agent_token_store_ce_go style3
    class agent_uiserver_redirect_fs_go style3
    class command_flags_merge_go style1
    class command_forceleave_forceleave_go style1
    class internal_controller_cache_indexers_indexersmock_mock_SingleIndexer_go style7
    class agent_checks_os_service_go style3
    class agent_consul_state_acl_events_go style3
    class proto_private_pboperator_operator_gen_go style7
    class sdk_testutil_retry_counter_go style7
    class test_integration_consul_container_libs_topology_service_topology_go style7
    class testing_deployer_sprawl_configentries_go style7
    class agent_connect_uri_agent_ce_go style3
    class logging_syslog_go style7
    class proto_private_pbautoconf_auto_config_go style7
    class service_os_service_go style7
    class troubleshoot_ports_hostport_go style7
    class agent_consul_controller_reconciler_go style3
    class agent_consul_health_endpoint_go style10
    class agent_consul_stream_subscription_go style3
    class agent_rpc_peering_validate_go style3
    class api_config_entry_sameness_group_go style4
    class command_connect_envoy_exec_go style1
    class proto_private_pbservice_healthcheck_gen_go style7
    class agent_consul_acl_server_go style3
    class agent_grpc_internal_balancer_registry_go style3
    class agent_hcp_telemetry_otlp_transform_go style3
    class agent_hcp_config_config_go style3
    class agent_log_drop_log_drop_go style3
    class command_agent_agent_go style1
    class internal_multicluster_internal_types_computed_exported_services_go style7
    class logging_logfile_go style7
    class agent_consul_operator_autopilot_endpoint_go style3
    class agent_structs_config_entry_inline_certificate_go style3
    class command_tls_ca_create_tls_ca_create_go style1
    class internal_controller_doc_go style7
    class internal_multicluster_internal_types_namespace_exported_services_go style7
    class api_operator_utilization_go style4
    class internal_go_sso_oidcauth_auth_go style7
    class sdk_testutil_context_go style7
    class testing_deployer_sprawl_helpers_go style7
    class agent_consul_reporting_reporting_ce_go style3
    class agent_consul_wanfed_pool_go style3
    class agent_proxycfg_glue_resolved_service_config_go style3
    class command_acl_authmethod_create_authmethod_create_ce_go style1
    class command_acl_policy_update_policy_update_go style1
    class agent_consul_prepared_query_endpoint_go style3
    class agent_dns_context_go style3
    class agent_grpc_external_services_resource_list_by_owner_go style3
    class agent_grpc_external_services_peerstream_subscription_view_go style3
    class agent_structs_config_entry_jwt_provider_ce_go style3
    class api_event_go style4
    class api_raw_go style4
    class lib_channels_deliver_latest_go style7
    class agent_cache_types_service_gateways_go style3
    class agent_config_default_ce_go style3
    class agent_consul_fsm_fsm_go style5
    class agent_hcp_bootstrap_config_loader_loader_go style3
    class agent_structs_connect_ce_go style3
    class internal_controller_cache_indexers_indexersmock_mock_GetSingleRefOrID_go style7
    class proto_private_pbconnect_connect_gen_go style7
    class proto_private_pboperator_operator_grpc_pb_go style7
    class acl_acl_ce_go style9
    class agent_consul_state_session_ce_go style3
    class agent_consul_state_indexer_go style3
    class agent_proxycfg_glue_service_list_go style3
    class command_acl_policy_create_policy_create_go style1
    class command_operator_autopilot_set_operator_autopilot_set_go style1
    class logging_logfile_solaris_go style7
    class testing_deployer_sprawl_internal_tfgen_gen_go style7
    class agent_cache_types_config_entry_go style3
    class agent_consul_multilimiter_multilimiter_go style3
    class agent_proxycfg_glue_leafcerts_go style3
    class command_services_exportedservices_exported_services_go style1
    class internal_resource_protoc_gen_json_shim_internal_generate_generate_go style7
    class logging_logfile_windows_go style7
    class agent_consul_state_config_entry_exported_services_go style3
    class agent_structs_intention_ce_go style3
    class command_resource_client_grpc_resource_flags_go style1
    class internal_resource_resourcetest_testing_go style7
    class sdk_testutil_assertions_go style7
    class agent_consul_fsm_snapshot_go style3
    class agent_consul_usagemetrics_usagemetrics_ce_go style3
    class agent_grpc_external_services_dataplane_mock_ACLResolver_go style3
    class command_acl_authmethod_formatter_go style1
    class internal_controller_cache_index_iterator_go style7
    class lib_hoststats_cpu_go style7
    class proto_public_pbmulticluster_v2_exported_services_consumer_pb_go style7
    class command_acl_templatedpolicy_list_templated_policy_list_go style1
    class internal_storage_inmem_snapshot_go style7
    class testing_deployer_sprawl_internal_tfgen_io_go style7
    class agent_consul_stream_event_go style3
    class agent_structs_config_entry_intentions_ce_go style3
    class command_acl_templatedpolicy_formatter_go style1
    class command_flags_config_go style1
    class internal_resource_protoc_gen_resource_types_internal_generate_generate_go style7
    class test_integration_consul_container_test_resource_http_api_client_client_go style7
    class troubleshoot_validate_validate_go style7
    class proto_public_pbacl_acl_cloning_grpc_pb_go style7
    class proto_public_pbserverdiscovery_serverdiscovery_json_gen_go style7
    class sdk_iptables_iptables_go style7
    class agent_grpc_external_server_go style8
    class agent_remote_exec_go style3
    class command_services_register_register_go style1
    class test_integration_consul_container_libs_cluster_container_go style7
    class agent_acl_endpoint_go style3
    class agent_hcp_scada_capabilities_go style3
    class command_rtt_rtt_go style1
    class command_resource_resource_go style1
    class grpcmocks_proto_public_pbconnectca_mock_ConnectCAService_WatchRootsClient_go style7
    class internal_multicluster_internal_types_helpers_ce_go style7
    class internal_resource_resourcetest_decode_go style7
    class agent_consul_state_acl_go style3
    class agent_grpc_external_testutils_mock_server_transport_stream_go style3
    class agent_structs_acl_templated_policy_ce_go style3
    class api_exported_services_go style4
    class internal_go_sso_oidcauth_internal_strutil_util_go style7
    class ipaddr_ipaddr_go style7
    class troubleshoot_proxy_z_xds_packages_go style7
    class agent_grpc_external_services_peerstream_replication_go style3
    class agent_hcp_bootstrap_constants_constants_go style3
    class proto_public_pbresource_resource_json_gen_go style7
    class sdk_testutil_retry_doc_go style7
    class agent_consul_gateway_locator_go style3
    class agent_consul_rate_handler_go style3
    class agent_consul_state_catalog_events_ce_go style3
    class agent_grpc_external_querymeta_go style3
    class agent_xds_endpoints_go style3
    class api_config_entry_exports_go style4
    class proto_private_pbpeerstream_types_go style7
    class agent_consul_acl_endpoint_go style10
    class agent_hcp_mock_TelemetryProvider_go style3
    class agent_proxycfg_naming_go style3
    class command_catalog_list_dc_catalog_list_datacenters_go style1
    class command_resource_client_grpc_client_go style1
    class agent_checks_grpc_go style3
    class agent_consul_operator_endpoint_go style3
    class agent_grpc_external_services_acl_mock_Login_go style3
    class command_resource_read_grpc_read_go style1
    class command_snapshot_snapshot_command_go style1
    class grpcmocks_proto_public_pbresource_mock_ResourceService_WatchListServer_go style7
    class test_integration_consul_container_libs_assert_grpc_go style7
    class agent_auto_config_auto_config_ce_go style3
    class agent_config_default_go style3
    class agent_structs_connect_go style3
    class internal_controller_helper_go style7
    class proto_private_pbcommon_common_gen_go style7
    class proto_private_pbservice_convert_ce_go style7
    class agent_cache_types_node_services_go style3
    class agent_denylist_go style3
    class proto_private_pbconfigentry_config_entry_go style7
    class test_integration_consul_container_libs_utils_version_ce_go style7
    class testing_deployer_sprawl_internal_tfgen_agent_go style7
    class agent_consul_state_acl_schema_go style3
    class agent_grpc_external_services_peerstream_subscription_manager_go style3
    class agent_grpc_external_services_resource_read_go style3
    class agent_grpc_middleware_auth_interceptor_go style3
    class command_acl_policy_formatter_go style1
    class command_keyring_keyring_go style1
    class command_intention_helpers_go style1
    class agent_connect_testing_ca_go style3
    class agent_consul_raft_rpc_go style3
    class agent_federation_state_endpoint_go style3
    class proto_private_pbsubscribe_subscribe_grpc_pb_go style7
    class version_versiontest_versiontest_go style7
    class agent_consul_discoverychain_string_stack_go style3
    class agent_consul_status_endpoint_go style3
    class agent_grpc_external_services_resource_server_go style8
    class agent_proxycfg_sources_catalog_mock_ConfigManager_go style3
    class api_config_entry_status_go style4
    class internal_controller_cache_index_index_go style7
    class testing_deployer_topology_generate_go style7
    class agent_cache_entry_go style3
    class agent_catalog_endpoint_ce_go style3
    class agent_metrics_go style3
    class agent_structs_catalog_ce_go style3
    class command_acl_authmethod_create_authmethod_create_go style1
    class agent_auto_config_server_addr_go style3
    class agent_pool_pool_go style3
    class api_discovery_chain_go style4
    class proto_private_prototest_golden_json_go style7
    class agent_consul_stream_event_snapshot_go style3
    class agent_dns_buffer_response_writer_go style3
    class agent_token_store_go style3
    class agent_xds_config_config_go style3
    class agent_structs_structs_deepcopy_ce_go style3
    class api_config_entry_mesh_go style4
    class grpcmocks_proto_public_pbacl_mock_ACLServiceClient_go style7
    class internal_protohcl_any_go style7
    class agent_consul_acl_authmethod_ce_go style3
    class agent_consul_autopilotevents_ready_servers_events_go style3
    class agent_setup_ce_go style3
    class command_acl_role_update_role_update_go style1
    class internal_resource_sort_go style7
    class agent_agent_endpoint_ce_go style3
    class agent_connect_uri_mesh_gateway_ce_go style3
    class agent_consul_kvs_endpoint_go style3
    class agent_debug_host_go style3
    class agent_xds_xds_go style3
    class command_flags_usage_go style1
    class grpcmocks_proto_public_pbresource_mock_UnsafeResourceServiceServer_go style7
    class internal_controller_controllermock_mock_CustomDependencyMapper_go style7
    class agent_consul_peering_backend_go style3
    class agent_consul_state_catalog_ce_go style3
    class api_watch_plan_go style4
    class command_login_login_go style1
    class connect_proxy_listener_go style2
    class proto_public_pbserverdiscovery_serverdiscovery_pb_go style7
    class testing_deployer_sprawl_internal_tfgen_docker_go style7
    class agent_http_register_go style3
    class api_debug_go style4
    class command_connect_envoy_flags_go style1
    class command_troubleshoot_troubleshoot_go style1
    class internal_controller_dependency_simple_go style7
    class agent_connect_ca_provider_vault_go style3
    class agent_discovery_chain_endpoint_go style3
    class agent_proxycfg_api_gateway_ce_go style3
    class agent_rpc_middleware_interceptors_go style3
    class agent_setup_go style3
    class command_acl_authmethod_update_authmethod_update_go style1
    class command_intention_create_create_go style1
    class proto_private_pbpeerstream_peerstream_pb_go style7
    class agent_consul_session_endpoint_go style3
    class agent_xds_server_ce_go style3
    class api_semaphore_go style4
    class agent_envoyextensions_builtin_otel_access_logging_otel_access_logging_go style3
    class agent_structs_protobuf_compat_go style3
    class command_operator_autopilot_state_formatter_go style1
    class test_integration_consul_container_test_envoy_extensions_testdata_wasm_test_files_wasm_add_header_go style7
    class agent_grpc_external_services_peerstream_mock_ACLResolver_go style3
    class agent_hcp_testing_go style3
    class api_session_go style4
    class command_operator_autopilot_get_operator_autopilot_get_go style1
    class internal_tools_protoc_gen_consul_rate_limit_main_go style7
    class agent_consul_leader_metrics_go style3
    class agent_consul_rate_metrics_go style3
    class agent_consul_state_system_metadata_go style3
    class agent_grpc_external_services_serverdiscovery_server_go style3
    class envoyextensions_extensioncommon_envoy_extender_go style7
    class proto_public_pbdataplane_dataplane_json_gen_go style7
    class agent_consul_federation_state_endpoint_go style3
    class agent_consul_leader_connect_ca_go style3
    class grpcmocks_proto_public_pbresource_mock_isWatchEvent_Event_go style7
    class testing_deployer_util_internal_ipamutils_doc_go style7
    class agent_consul_state_acl_ce_go style3
    class tlsutil_config_go style7
    class agent_consul_autopilot_go style3
    class agent_structs_snapshot_go style3
    class grpcmocks_proto_public_pbacl_mock_ACLServiceServer_go style7
    class grpcmocks_proto_public_pbdataplane_mock_DataplaneServiceClient_go style7
    class internal_resource_reference_go style7
    class testing_deployer_sprawl_debug_go style7
    class agent_cache_types_gateway_services_go style3
    class agent_configentry_config_entry_go style3
    class agent_proxycfg_glue_internal_service_dump_go style3
    class agent_proxycfg_glue_intentions_go style3
    class command_catalog_list_services_catalog_list_services_go style1
    class lib_stop_context_go style7
    class proto_public_pbresource_resource_deepcopy_gen_go style7
    class agent_config_flagset_go style3
    class agent_grpc_middleware_rate_go style3
    class agent_xds_resources_go style3
    class agent_xds_response_response_go style3
    class snapshot_snapshot_go style7
    class agent_consul_watch_server_local_go style3
    class agent_grpc_external_services_peerstream_subscription_blocking_go style3
    class agent_xds_listeners_ingress_go style3
    class proto_private_pbcommon_common_go style7
    class agent_cacheshim_cache_go style3
    class agent_hcp_bootstrap_testing_go style3
    class agent_mock_notify_go style3
    class version_fips_go style7
    class agent_exec_exec_windows_go style3
    class agent_hcp_manager_go style3
    class command_config_write_config_write_go style1
    class command_resource_helper_go style1
    class internal_resource_protoc_gen_deepcopy_internal_generate_generate_go style7
    class agent_acl_go style3
    class agent_consul_stream_event_publisher_go style3
    class agent_dns_go style3
    class command_tls_tls_go style1
    class agent_consul_discoverychain_compile_ce_go style3
    class agent_grpc_external_services_acl_mock_TokenWriter_go style3
    class api_acl_go style4
    class troubleshoot_proxy_validateupstream_go style7
    class agent_cache_mock_Request_go style3
    class agent_consul_multilimiter_mock_RateLimiter_go style3
    class command_helpers_helpers_go style1
    class lib_decode_decode_go style7
    class proto_public_pbdns_dns_json_gen_go style7
    class testing_deployer_sprawl_internal_runner_exec_go style7
    class agent_cache_mock_Type_go style3
    class command_watch_watch_go style1
    class acl_resolver_danger_go style9
    class agent_auto_config_auto_encrypt_go style3
    class agent_http_go style3
    class agent_xds_server_go style2
    class agent_blockingquery_mock_RequestOptions_go style3
    class agent_consul_merge_ce_go style3
    class agent_coordinate_endpoint_go style3
    class agent_structs_autopilot_ce_go style3
    class api_lock_go style4
    class command_catalog_helpers_go style1
    class command_event_event_go style1
    class command_operator_raft_operator_raft_go style1
    class agent_consul_wanfed_wanfed_go style3
    class agent_proxycfg_upstreams_go style3
    class agent_structs_autopilot_go style3
    class internal_controller_cache_indexers_indexersmock_mock_RefOrIDFetcher_go style7
    class internal_resource_hooks_go style7
    class internal_tools_protoc_gen_grpc_clone_e2e_mock_Simple_FlowClient_go style7
    class proto_private_pbservice_ids_go style7
    class testing_deployer_util_internal_ipamutils_utils_go style7
    class agent_consul_state_connect_ca_go style3
    class agent_log_drop_mock_Logger_go style3
    class agent_proxycfg_naming_ce_go style3
    class test_integration_consul_container_libs_service_helpers_go style7
    class test_integration_consul_container_libs_utils_version_go style7
    class troubleshoot_ports_troubleshoot_tcp_go style7
    class agent_grpc_external_services_peerstream_stream_tracker_go style3
    class agent_proxycfg_testing_peering_go style3
    class grpcmocks_proto_public_pbresource_mock_ResourceServiceClient_go style7
    class proto_private_pbpeerstream_peerstream_grpc_pb_go style7
    class test_integration_consul_container_libs_assert_service_go style7
    class agent_connect_ca_provider_vault_auth_jwt_go style3
    class command_acl_token_read_token_read_go style1
    class sdk_freeport_ephemeral_linux_go style7
    class agent_cache_types_rpc_go style3
    class agent_consul_rtt_go style3
    class agent_grpc_external_services_resource_mock_Registry_go style3
    class agent_structs_config_entry_ce_go style3
    class agent_structs_config_entry_gateways_go style3
    class api_watch_funcs_go style4
    class connect_service_go style2
    class internal_go_sso_oidcauth_oidcauthtest_testing_go style7
    class agent_envoyextensions_builtin_aws_lambda_aws_lambda_go style3
    class agent_nodeid_go style3
    class agent_uiserver_buf_index_fs_go style3
    class proto_private_pbconnect_connect_pb_go style7
    class sdk_iptables_iptables_executor_unsupported_go style7
    class test_integration_connect_envoy_test_sds_server_sds_go style2
    class agent_consul_acl_client_go style3
    class agent_consul_configentry_backend_go style3
    class agent_consul_state_query_go style3
    class api_config_entry_gateways_go style4
    class command_keygen_keygen_go style1
    class command_maint_maint_go style1
    class ipaddr_detect_go style7
    class sdk_iptables_iptables_executor_linux_go style7
    class agent_xds_testing_go style3
    class lib_cluster_go style7
    class agent_envoyextensions_builtin_property_override_property_override_go style3
    class grpcmocks_proto_public_pbresource_mock_ResourceServiceServer_go style7
    class agent_connect_ca_provider_consul_go style3
    class agent_consul_options_ce_go style3
    class agent_proxycfg_connect_proxy_go style3
    class agent_utilization_endpoint_go style3
    class agent_xds_locality_policy_ce_go style3
    class lib_retry_retry_go style7
    class proto_public_annotations_ratelimit_ratelimit_deepcopy_gen_go style7
    class agent_structs_config_entry_status_go style3
    class agent_xds_failover_policy_go style3
    class api_config_entry_inline_certificate_go style4
    class api_peering_go style4
    class internal_controller_cache_index_txn_go style7
    class proto_public_pbacl_acl_deepcopy_gen_go style7
    class agent_consul_state_graveyard_go style3
    class api_watch_watch_go style4
    class lib_template_hil_go style7
    class agent_consul_auth_token_writer_ce_go style3
    class agent_consul_state_coordinate_go style3
    class agent_consul_usagemetrics_usagemetrics_go style3
    class agent_grpc_external_services_connectca_mock_CAManager_go style3
    class grpcmocks_proto_public_pbconnectca_mock_ConnectCAService_WatchRootsServer_go style7
    class internal_controller_cache_indexers_indexersmock_mock_FromArgs_go style7
    class proto_public_pbmulticluster_v2_namespace_exported_services_pb_go style7
    class agent_auto_config_config_ce_go style3
    class agent_consul_state_catalog_schema_go style3
    class api_operator_area_go style4
    class command_intention_list_intention_list_go style1
    class command_version_version_go style1
    class acl_policy_authorizer_go style9
    class agent_cache_watch_go style3
    class agent_consul_fsm_log_verification_chunking_shim_go style3
    class agent_consul_server_ce_go style3
    class agent_local_state_go style3
    class agent_xds_listeners_apigateway_go style3
    class connect_proxy_testing_go style2
    class agent_xds_extensionruntime_runtime_config_go style3
    class agent_grpc_external_services_connectca_mock_ACLResolver_go style3
    class agent_grpc_external_services_peerstream_stream_resources_go style3
    class agent_rpc_middleware_recovery_go style3
    class internal_controller_cache_indexers_id_indexer_go style7
    class internal_multicluster_internal_types_decoded_go style7
    class internal_resource_protoc_gen_json_shim_main_go style7
    class internal_resource_resource_go style7
    class internal_storage_raft_forwarding_go style7
    class agent_connect_testing_spiffe_go style3
    class agent_grpc_external_services_resource_delete_go style3
    class command_services_export_export_go style1
    class internal_controller_cache_client_go style7
    class proto_private_pbconfigentry_config_entry_gen_go style7
    class proto_private_prototest_testing_go style7
    class proto_private_pbservice_healthcheck_pb_go style7
    class proto_public_pbconnectca_ca_grpc_pb_go style7
    class agent_consul_txn_endpoint_go style3
    class agent_grpc_external_services_peerstream_subscription_state_go style3
    class agent_signal_unix_go style3
    class agent_submatview_handler_go style3
    class command_acl_token_create_token_create_go style1
    class internal_protohcl_blocks_go style7
    class testing_deployer_sprawl_internal_tfgen_prelude_go style7
    class acl_errors_go style9
    class agent_config_config_go style3
    class testing_deployer_topology_images_go style7
    class agent_cache_types_intention_upstreams_go style3
    class agent_signal_windows_go style3
    class agent_status_endpoint_go style3
    class command_acl_authmethod_update_authmethod_update_ce_go style1
    class command_acl_role_read_role_read_go style1
    class grpcmocks_proto_public_pbserverdiscovery_mock_ServerDiscoveryServiceServer_go style7
    class logging_monitor_monitor_go style7
    class troubleshoot_proxy_utils_go style7
    class acl_static_authorizer_go style9
    class agent_acl_ce_go style3
    class agent_consul_server_go style5
    class agent_proxycfg_snapshot_go style3
    class agent_proxycfg_sources_local_sync_go style3
    class agent_structs_discovery_chain_go style3
    class internal_protohcl_attributes_go style7
    class proto_private_pbservice_node_pb_go style7
    class agent_cache_types_discovery_chain_go style3
    class agent_hcp_telemetry_otel_exporter_go style3
    class agent_proxycfg_glue_health_blocking_go style3
    class agent_snapshot_endpoint_go style3
    class agent_structs_check_type_go style3
    class agent_submatview_mock_ACLResolver_go style3
    class command_resource_client_grpc_flags_go style1
    class proto_public_annotations_ratelimit_ratelimit_pb_go style7
    class agent_consul_acl_replication_go style3
    class command_connect_envoy_exec_windows_go style1
    class command_resource_delete_grpc_delete_go style1
    class logging_log_levels_go style7
    class sdk_freeport_systemlimit_windows_go style7
    class test_integration_consul_container_libs_cluster_config_go style7
    class test_integration_consul_container_libs_utils_utils_go style7
    class test_integration_consul_container_libs_utils_docker_go style7
    class agent_consul_state_usage_go style3
    class agent_grpc_middleware_stats_go style3
    class api_connect_go style4
    class proto_public_pbconnectca_ca_deepcopy_gen_go style7
    class troubleshoot_proxy_stats_go style7
    class agent_consul_leader_connect_ca_ce_go style3
    class agent_structs_peering_go style4
    class api_kv_go style4
    class proto_private_pbcommon_convert_pbstruct_go style7
    class agent_consul_enterprise_config_ce_go style3
    class agent_proxycfg_glue_health_go style3
    class command_snapshot_restore_snapshot_restore_go style1
    class connect_proxy_conn_go style2
    class internal_controller_cache_indexers_ref_indexer_go style7
    class acl_validation_go style9
    class agent_consul_authmethod_testauth_testing_ce_go style3
    class agent_connect_uri_signing_go style3
    class agent_leafcert_leafcert_test_helpers_go style3
    class agent_submatview_rpc_materializer_go style3
    class api_operator_raft_go style4
    class internal_controller_dependency_higher_order_go style7
    class agent_cache_types_health_services_go style3
    class agent_consul_discoverychain_compile_go style3
    class agent_consul_state_config_entry_sameness_group_ce_go style3
    class agent_structs_config_entry_sameness_group_go style3
    class proto_public_pbmulticluster_v2_namespace_exported_services_json_gen_go style7
    class agent_cache_types_intention_upstreams_destination_go style3
    class agent_consul_authmethod_testauth_testing_go style3
    class agent_consul_operator_usage_endpoint_go style3
    class agent_consul_state_config_entry_sameness_group_go style3
    class agent_structs_config_entry_discoverychain_go style3
    class agent_structs_check_definition_go style4
    class internal_radix_doc_go style7
    class proto_private_pbdemo_v1_demo_pb_go style7
    class agent_connect_sni_go style3
    class agent_proxycfg_sources_local_local_go style3
    class command_peering_peering_go style1
    class proto_private_pbservice_convert_go style7
    class proto_public_pbdataplane_dataplane_pb_go style7
    class proto_public_pbdns_dns_grpc_pb_go style7
    class test_integration_consul_container_libs_assert_common_go style7
    class testing_deployer_sprawl_network_area_ce_go style7
    class agent_consul_acl_token_exp_go style3
    class agent_consul_system_metadata_go style3
    class agent_structs_connect_ca_go style3
    class command_acl_role_list_role_list_go style1
    class command_acl_token_update_token_update_go style1
    class internal_tools_proto_gen_rpc_glue_e2e_source_pb_go style7
    class sentinel_scope_go style7
    class agent_blockingquery_mock_FSMServer_go style3
    class agent_consul_acl_go style3
    class agent_grpc_external_services_resource_mock_Backend_go style3
    class agent_grpc_internal_services_subscribe_subscribe_go style3
    class agent_submatview_local_materializer_go style3
    class logging_names_go style7
    class command_connect_proxy_register_go style1
    class internal_tools_protoc_gen_grpc_clone_e2e_proto_cloning_stream_pb_go style7
    class logging_logger_go style7
    class test_integration_consul_container_libs_utils_helpers_go style7
    class agent_blockingquery_mock_ResponseMeta_go style3
    class agent_grpc_external_services_resource_list_go style3
    class agent_proxycfg_testing_go style3
    class agent_structs_config_entry_jwt_provider_go style3
    class command_resource_client_grpc_config_go style1
    class lib_mutex_mutex_go style7
    class test_integration_consul_container_libs_utils_debug_go style7
    class testing_deployer_sprawl_internal_tfgen_res_go style7
    class agent_auto_config_auto_config_go style3
    class agent_consul_state_peering_go style3
    class agent_proxycfg_ingress_gateway_go style3
    class internal_protohcl_decoder_go style7
    class internal_storage_inmem_store_go style7
    class internal_tools_proto_gen_rpc_glue_e2e_consul_proto_pbcommon_common_go style7
    class agent_consul_prepared_query_walk_go style3
    class agent_proxycfg_proxycfg_go style3
    class agent_structs_structs_deepcopy_go style3
    class command_connect_envoy_pipe_bootstrap_connect_envoy_pipe_bootstrap_go style1
    class command_helpers_decode_shim_go style1
    class lib_useragent_go style7
    class agent_consul_subscribe_backend_go style3
    class agent_grpc_external_testutils_acl_go style3
    class proto_public_pbmulticluster_v2_partition_exported_services_deepcopy_gen_go style7
    class sdk_testutil_io_go style7
    class sdk_testutil_retry_interface_go style7
    class test_integration_consul_container_libs_utils_tenancy_go style7
    class agent_config_file_watcher_go style3
    class agent_structs_prepared_query_go style3
    class agent_structs_testing_intention_go style3
    class command_acl_token_token_go style1
    class agent_consul_state_tombstone_gc_go style3
    class agent_proxycfg_testing_terminating_gateway_go style3
    class command_snapshot_save_snapshot_save_go style1
    class internal_controller_cache_index_indexmock_mock_resourceIterable_go style7
    class internal_multicluster_internal_controllers_register_go style7
    class internal_storage_inmem_schema_go style7
    class internal_tools_protoc_gen_grpc_clone_main_go style7
    class logging_logfile_darwinandfreebsd_go style7
    class api_operator_usage_go style4
    class proto_private_pbautoconf_auto_config_pb_go style7
    class sdk_testutil_retry_retry_go style7
    class agent_structs_envoy_extension_go style3
    class agent_xds_delta_go style3
    class command_connect_envoy_bootstrap_config_go style1
    class command_tls_cert_tls_cert_go style1
    class internal_resource_errors_go style7
    class agent_config_merge_go style3
    class agent_consul_controller_queue_defer_go style3
    class agent_consul_util_go style3
    class agent_leafcert_cached_roots_go style3
    class agent_proxycfg_mesh_gateway_go style3
    class command_resource_delete_delete_go style1
    class internal_tools_protoc_gen_grpc_clone_e2e_mock_SimpleClient_go style7
    class test_integration_consul_container_test_gateways_tenancy_ce_go style7
    class agent_checks_alias_go style3
    class agent_consul_controller_doc_go style3
    class agent_operator_endpoint_go style3
    class agent_rpcclient_health_view_go style3
    class agent_uiserver_uiserver_go style3
    class troubleshoot_proxy_certs_go style7
    class agent_connect_uri_service_ce_go style3
    class agent_grpc_external_services_resource_mutate_and_validate_go style3
    class agent_hcp_client_http_client_go style3
    class command_validate_validate_go style1
    class proto_public_pbconnectca_cloning_stream_pb_go style7
    class agent_agent_endpoint_go style4
    class agent_consul_state_peering_ce_go style3
    class grpcmocks_proto_public_pbserverdiscovery_mock_ServerDiscoveryService_WatchServersClient_go style7
    class internal_controller_cache_index_convenience_go style7
    class proto_private_pbpeering_peering_ce_go style7
    class agent_grpc_external_services_acl_mock_Validator_go style3
    class agent_local_testing_go style3
    class command_reload_reload_go style1
    class lib_math_go style7
    class proto_private_pbconnect_connect_go style7
    class agent_cache_types_options_go style3
    class agent_consul_state_intention_go style3
    class agent_grpc_external_services_peerstream_testing_go style3
    class agent_proxycfg_testing_mesh_gateway_go style3
    class agent_structs_aclfilter_filter_go style3
    class command_resource_list_grpc_list_go style1
    class command_troubleshoot_ports_troubleshoot_ports_go style1
    class internal_go_sso_oidcauth_oidc_go style7
    class agent_grpc_external_services_resource_mock_ACLResolver_go style3
    class agent_proxycfg_sources_catalog_config_source_go style3
    class internal_tools_protoc_gen_grpc_clone_e2e_proto_service_cloning_grpc_pb_go style7
    class testing_deployer_sprawl_catalog_go style7
    class agent_envoyextensions_registered_extensions_go style3
    class agent_hcp_client_metrics_client_go style3
    class command_lock_util_windows_go style1
    class command_troubleshoot_proxy_troubleshoot_proxy_go style1
    class internal_storage_inmem_watch_go style7
    class snapshot_archive_go style7
    class agent_grpc_external_services_resource_server_ce_go style3
    class agent_structs_federation_state_go style3
    class command_acl_role_create_role_create_go style1
    class command_join_join_go style1
    class proto_private_pbpeerstream_peerstream_go style7
    class proto_public_pbmulticluster_v2_exported_services_deepcopy_gen_go style7
    class types_tls_go style7
    class agent_consul_state_census_utilization_go style3
    class agent_proxycfg_glue_config_entry_go style3
    class agent_rpcclient_configentry_configentry_go style3
    class internal_controller_watch_go style7
    class internal_multicluster_internal_controllers_v1compat_controller_go style7
    class internal_protohcl_oneof_go style7
    class internal_tools_proto_gen_rpc_glue_main_go style7
    class test_integration_consul_container_libs_utils_defer_go style7
    class agent_consul_leader_registrator_v1_go style3
    class agent_grpc_external_services_connectca_watch_roots_go style3
    class agent_proxycfg_glue_gateway_services_go style3
    class command_connect_connect_go style1
    class command_config_list_config_list_go style1
    class command_connect_envoy_exec_unsupported_go style1
    class internal_multicluster_exports_go style7
    class test_integration_consul_container_libs_cluster_cluster_go style7
    class agent_xds_platform_net_fallback_go style3
    class agent_cache_types_catalog_service_list_go style3
    class agent_consul_raft_handle_go style3
    class agent_grpc_external_services_connectca_sign_go style3
    class agent_rpc_peering_service_go style3
    class agent_structs_dns_go style3
    class grpcmocks_proto_public_pbserverdiscovery_mock_UnsafeServerDiscoveryServiceServer_go style7
    class internal_radix_radix_go style7
    class internal_tools_proto_gen_rpc_glue_e2e_consul_agent_structs_structs_go style7
    class agent_consul_state_usage_ce_go style3
    class command_services_config_go style1
    class internal_controller_dependency_dependencymock_mock_DependencyTransform_go style7
    class lib_hoststats_collector_go style7
    class proto_public_pbresource_annotations_pb_go style7
    class agent_consul_auth_login_go style3
    class agent_consul_fsm_snapshot_ce_go style3
    class agent_consul_state_config_entry_go style3
    class agent_consul_servercert_manager_go style3
    class agent_grpc_external_services_peerstream_health_snapshot_go style3
    class agent_xds_z_xds_packages_go style3
    class command_services_deregister_deregister_go style1
    class sentinel_sentinel_ce_go style7
    class troubleshoot_proxy_troubleshoot_proxy_go style7
    class agent_consul_replication_go style3
    class agent_grpc_internal_tracker_go style3
    class agent_xds_clusters_go style3
    class api_status_go style4
    class command_operator_operator_go style1
    class internal_controller_cache_errors_go style7
    class internal_resource_protoc_gen_deepcopy_main_go style7
    class proto_public_pbconnectca_ca_cloning_grpc_pb_go style7
    class agent_cache_types_exported_peered_services_go style3
    class agent_config_config_ce_go style3
    class agent_consul_auto_config_endpoint_go style3
    class agent_consul_leader_intentions_ce_go style3
    class agent_consul_prepared_query_endpoint_ce_go style3
    class command_peering_establish_establish_go style1
    class proto_private_pbstatus_status_pb_go style7
    class acl_resolver_result_go style9
    class agent_config_agent_limits_go style3
    class agent_consul_state_graveyard_ce_go style3
    class agent_consul_state_kvs_ce_go style3
    class agent_grpc_internal_listener_go style3
    class command_resource_apply_apply_go style1
    class internal_protohcl_well_known_types_go style7
    class internal_resource_http_http_go style7
    class agent_cache_types_service_checks_go style3
    class agent_intentions_endpoint_go style3
    class command_connect_proxy_flag_upstreams_go style1
    class proto_public_pbresource_resource_grpc_pb_go style7
    class test_integration_consul_container_test_consul_envoy_version_consul_envoy_version_go style7
    class agent_consul_server_register_go style3
    class agent_consul_tenancy_bridge_go style3
    class agent_grpc_external_services_acl_login_go style3
    class agent_hcp_telemetry_provider_go style3
    class lib_stringslice_stringslice_go style7
    class agent_consul_peering_backend_ce_go style3
    class command_operator_autopilot_operator_autopilot_go style1
    class command_troubleshoot_upstreams_troubleshoot_upstreams_go style1
    class grpcmocks_proto_public_pbconnectca_mock_ConnectCAServiceServer_go style7
    class main_go style7
    class agent_config_deprecated_go style3
    class agent_consul_config_replication_go style3
    class proto_public_pbdataplane_dataplane_deepcopy_gen_go style7
    class agent_cache_type_go style3
    class agent_consul_config_ce_go style3
    class agent_consul_state_autopilot_go style3
    class agent_grpc_external_services_serverdiscovery_watch_servers_go style3
    class agent_systemd_notify_go style3
    class command_config_config_go style1
    class command_login_aws_go style1
    class lib_path_go style7
    class agent_consul_authmethod_authmethods_ce_go style3
    class agent_consul_config_endpoint_go style3
    class agent_xds_secrets_go style3
    class command_config_delete_config_delete_go style1
    class command_debug_debug_go style1
    class internal_testing_errors_errors_go style7
    class internal_storage_storage_go style7
    class proto_private_pbpeering_peering_gen_go style7
    class acl_authorizer_ce_go style9
    class agent_consul_state_catalog_events_go style3
    class agent_consul_state_connect_ca_events_go style3
    class agent_hcp_bootstrap_bootstrap_go style3
    class command_monitor_monitor_go style1
    class lib_hoststats_host_go style7
    class proto_private_pbsubscribe_subscribe_go style7
    class sdk_testutil_types_go style7
    class command_acl_token_formatter_go style1
    class envoyextensions_extensioncommon_basic_extension_adapter_go style7
    class internal_controller_cache_cachemock_mock_DecodedResourceIterator_go style7
    class agent_consul_autopilot_ce_go style3
    class agent_consul_authmethod_ssoauth_sso_go style3
    class internal_controller_cache_index_interfaces_go style7
    class test_integration_consul_container_test_upgrade_common_go style7
    class agent_consul_state_config_entry_intention_ce_go style3
    class agent_proxycfg_testing_upstreams_ce_go style3
    class command_connect_envoy_exec_unix_go style1
    class internal_controller_cache_indexers_indexersmock_mock_BoundReferences_go style7
    class internal_resource_demo_demo_go style7
    class proto_public_pbserverdiscovery_serverdiscovery_grpc_pb_go style7
    class agent_grpc_external_forward_go style3
    class agent_hcp_telemetry_custom_metrics_go style3
    class agent_rpc_middleware_rate_limit_mappings_go style3
    class agent_structs_structs_go style4
    class agent_watch_handler_go style3
    class connect_proxy_proxy_go style2
    class agent_cache_types_testing_go style3
    class agent_connect_uri_agent_go style3
    class agent_consul_controller_queue_queue_go style3
    class agent_envoyextensions_builtin_ext_authz_structs_go style3
    class lib_json_go style7
    class lib_translate_go style7
    class proto_public_pbmulticluster_v2_computed_exported_services_deepcopy_gen_go style7
    class agent_cache_types_catalog_datacenters_go style3
    class agent_notify_go style3
    class api_config_entry_discoverychain_go style4
    class command_tls_cert_create_tls_cert_create_go style1
    class testing_deployer_topology_naming_shim_go style7
    class testing_deployer_util_net_go style7
    class agent_grpc_middleware_testutil_fake_sink_go style3
    class agent_user_event_go style3
    class api_config_entry_intentions_go style4
    class command_acl_authmethod_authmethod_go style1
    class command_services_services_go style1
    class internal_go_sso_oidcauth_config_go style7
    class agent_configentry_doc_go style3
    class agent_exec_exec_go style3
    class agent_structs_connect_proxy_config_ce_go style3
    class agent_xds_platform_net_linux_go style3
    class api_operator_license_go style4
    class proto_private_pbconfig_config_pb_go style7
    class agent_consul_authmethod_awsauth_aws_go style3
    class internal_controller_dependency_dependencymock_mock_CacheIDModifier_go style7
    class proto_private_pbdemo_v1_resources_rtypes_go style7
    class testing_deployer_sprawl_internal_tfgen_dns_go style7
    class agent_config_endpoint_go style3
    class agent_consul_acl_server_ce_go style3
    class agent_enterprise_delegate_ce_go style3
    class agent_sidecar_service_go style3
    class api_connect_intention_go style4
    class internal_resource_authz_go style7
    class internal_resource_resourcetest_acls_go style7
    class testing_deployer_sprawl_details_go style7
    class agent_hcp_telemetry_doc_go style3
    class agent_leafcert_cert_go style3
    class command_resource_read_read_go style1
    class internal_controller_controllermock_mock_Lease_go style7
    class internal_controller_controllermock_mock_DependencyMapper_go style7
    class agent_ae_ae_go style3
    class agent_leafcert_roots_go style3
    class agent_translate_addr_go style3
    class test_integration_consul_container_libs_cluster_app_go style7
    class acl_acl_go style9
    class command_acl_policy_delete_policy_delete_go style1
    class internal_resource_equality_go style7
    class proto_private_pbconfigentry_config_entry_pb_go style7
    class version_version_go style7
    class agent_consul_client_serf_go style3
    class agent_consul_federation_state_replication_go style3
    class agent_proxycfg_sources_catalog_mock_SessionLimiter_go style3
    class agent_txn_endpoint_go style3
    class agent_config_config_deepcopy_go style3
    class agent_consul_rpc_go style3
    class command_catalog_helpers_ce_go style1
    class command_registry_go style1
    class internal_controller_controllermock_mock_DependencyTransform_go style7
    class agent_connect_parsing_go style3
    class agent_consul_reporting_reportingmock_mock_ServerDelegate_go style3
    class agent_grpc_external_limiter_limiter_go style3
    class command_tls_ca_tls_ca_go style1
    class agent_connect_uri_service_go style3
    class grpcmocks_proto_public_pbacl_mock_UnsafeACLServiceServer_go style7
    class lib_file_atomic_go style7
    class acl_policy_merger_go style9
    class agent_ae_trigger_go style3
    class agent_apiserver_go style4
    class agent_grpc_external_services_dns_server_go style8
    class internal_resource_resourcetest_fs_go style7
    class agent_consul_gateways_controller_gateways_go style3
    class agent_rpc_peering_testing_go style3
    class command_catalog_list_nodes_catalog_list_nodes_go style1
    class internal_controller_cache_indexers_decoded_indexer_go style7
    class acl_MockAuthorizer_go style9
    class agent_grpc_external_services_dataplane_get_supported_features_go style3
    class agent_leafcert_util_go style3
    class internal_protohcl_testproto_example_pb_go style7
    class agent_checks_docker_go style3
    class agent_consul_merge_go style3
    class agent_proxycfg_testing_connect_proxy_go style3
    class proto_private_pbpeering_peering_pb_go style7
    class acl_enterprisemeta_ce_go style9
    class agent_consul_rate_handler_ce_go style3
    class agent_envoyextensions_builtin_ext_authz_ext_authz_go style3
    class agent_grpc_external_services_dataplane_get_envoy_bootstrap_params_go style3
    class agent_structs_system_metadata_go style3
    class internal_resource_registry_ce_go style7
    class internal_resource_tombstone_go style7
    class proto_public_pbmulticluster_v2_partition_exported_services_pb_go style7
    class agent_hcp_client_telemetry_config_go style3
    class agent_leafcert_signer_netrpc_go style3
    class agent_proxycfg_glue_trust_bundle_go style3
    class api_internal_go style4
    class command_leave_leave_go style1
    class agent_blockingquery_blockingquery_go style3
    class agent_consul_auto_encrypt_endpoint_go style3
    class agent_consul_state_schema_ce_go style3
    class agent_grpc_internal_client_go style3
    class agent_leafcert_watch_go style3
    class agent_xds_jwt_authn_ce_go style3
    class api_operator_segment_go style4
    class internal_resource_resourcetest_client_go style7
    class agent_leafcert_leafcert_go style3
    class command_acl_bootstrap_bootstrap_go style1
    class lib_testhelpers_testhelpers_go style7
    class proto_public_pbdataplane_dataplane_cloning_grpc_pb_go style7
    class proto_public_pbserverdiscovery_serverdiscovery_deepcopy_gen_go style7
    class sdk_testutil_server_methods_go style7
    class testing_deployer_topology_ids_go style7
    class agent_grpc_external_services_resource_testing_testing_ce_go style3
    class command_acl_templatedpolicy_preview_templated_policy_preview_go style1
    class command_operator_raft_transferleader_transfer_leader_go style1
    class grpcmocks_proto_public_pbserverdiscovery_mock_ServerDiscoveryService_WatchServersServer_go style7
    class internal_controller_custom_watch_go style7
    class internal_protohcl_primitives_go style7
    class internal_tools_proto_gen_rpc_glue_e2e_consul_proto_pbcommon_common_pb_go style7
    class proto_public_pbmulticluster_v2_computed_exported_services_json_gen_go style7
    class proto_public_pbresource_annotations_json_gen_go style7
    class agent_consul_state_schema_go style3
    class command_acl_acl_helpers_go style1
    class command_agent_startup_logger_go style1
    class test_integration_consul_container_libs_cluster_network_go style7
    class testing_deployer_topology_default_versions_go style7
    class tlsutil_mock_go style0
    class agent_consul_authmethod_authmethods_go style3
    class agent_consul_state_prepared_query_go style3
    class agent_pool_peek_go style3
    class internal_resource_demo_controller_go style7
    class internal_resource_bound_refs_go style7
    class internal_storage_conformance_conformance_go style7
    class proto_public_pbconnectca_ca_pb_go style7
    class test_integ_upgrade_l7_traffic_management_common_go style7
    class agent_agent_go style3
    class agent_connect_ca_common_go style3
    class agent_grpc_internal_resolver_resolver_go style3
    class agent_structs_catalog_go style3
    class agent_structs_config_entry_go style4
    class api_operator_keyring_go style4
    class command_info_info_go style1
    class internal_resource_registry_go style7
    class agent_agent_ce_go style3
    class agent_consul_auth_token_writer_go style3
    class agent_consul_state_catalog_go style3
    class agent_hcp_client_client_go style3
    class agent_hcp_client_mock_Client_go style3
    class agent_structs_acl_go style3
    class api_config_entry_jwt_provider_go style4
    class internal_resource_refkey_go style7
    class agent_proxycfg_sources_catalog_mock_Watcher_go style3
    class agent_service_manager_go style3
    class api_partition_go style4
    class command_registry_ce_go style1
    class internal_controller_cache_cachemock_mock_ReadOnlyCache_go style7
    class internal_controller_cache_indexers_indexersmock_mock_MultiIndexer_go style7
    class internal_resource_stringer_go style7
    class internal_tools_protoc_gen_grpc_clone_e2e_proto_service_pb_go style7
    class agent_consul_state_prepared_query_index_go style3
    class agent_grpc_external_services_serverdiscovery_mock_ACLResolver_go style3
    class grpcmocks_proto_public_pbdns_mock_UnsafeDNSServiceServer_go style7
    class internal_go_sso_oidcauth_jwt_go style7
    class internal_tools_protoc_gen_grpc_clone_internal_generate_generate_go style7
    class proto_private_pbstorage_raft_pb_go style7
    class proto_public_pbresource_resource_cloning_grpc_pb_go style7
    class service_os_service_windows_go style7
    class agent_auto_config_persist_go style3
    class agent_checks_check_go style3
    class agent_grpc_external_services_dataplane_server_go style8
    class api_agent_go style4
    class command_cli_formatting_go style1
    class logging_grpc_go style7
    class testing_deployer_util_files_go style7
    class acl_policy_authorizer_ce_go style9
    class agent_connect_x509_patch_go style3
    class agent_consul_tenancy_bridge_ce_go style3
    class agent_envoyextensions_registered_extensions_ce_go style3
    class agent_grpc_internal_balancer_balancer_go style3
    class internal_protohcl_unmarshal_go style7
    class test_integration_consul_container_libs_topology_peering_topology_go style7
    class agent_cache_types_resolved_service_config_go style3
    class agent_configentry_service_config_go style3
    class agent_consul_enterprise_client_ce_go style3
    class agent_consul_state_txn_go style3
    class proto_public_pbserverdiscovery_serverdiscovery_cloning_grpc_pb_go style7
    class tools_internal_grpc_proxy_main_go style7
    class troubleshoot_proxy_upstreams_go style7
    class agent_consul_state_config_entry_ce_go style3
    class agent_consul_state_mock_publishFuncType_go style3
    class agent_grpc_external_services_resource_write_status_go style3
    class agent_proxycfg_internal_watch_watchmap_go style3
    class agent_structs_config_entry_intentions_go style3
    class proto_public_pbresource_annotations_deepcopy_gen_go style7
    class agent_auto_config_run_go style3
    class agent_proxycfg_testing_upstreams_go style3
    class command_login_login_ce_go style1
    class internal_controller_runner_go style7
    class agent_consul_server_log_verification_go style3
    class agent_envoyextensions_builtin_wasm_wasm_go style3
    class command_connect_envoy_bootstrap_tpl_go style1
    class command_peering_list_list_go style1
    class agent_consul_state_operations_ce_go style3
    class agent_consul_stream_noop_go style3
    class agent_reload_go style3
    class command_operator_raft_listpeers_operator_raft_list_go style1
    class testing_deployer_sprawl_resources_go style7
    class agent_consul_coordinate_endpoint_go style3
    class agent_proxycfg_api_gateway_go style3
    class agent_proxycfg_testing_tproxy_go style3
    class command_peering_read_read_go style1
    class logging_gated_writer_go style7
    class proto_public_pbdataplane_dataplane_grpc_pb_go style7
    class test_integration_consul_container_libs_cluster_encryption_go style7
    class test_integration_consul_container_libs_service_log_go style7
    class agent_config_ratelimited_file_watcher_go style3
    class agent_xds_listeners_go style3
    class proto_public_pbmulticluster_v2_computed_exported_services_pb_go style7
    class testing_deployer_sprawl_boot_go style7
    class testing_deployer_util_consul_go style7
    class agent_connect_ca_provider_vault_auth_gcp_go style3
    class agent_consul_configentry_backend_ce_go style3
    class api_content_type_go style4
    class logging_logfile_linux_go style7
    class proto_private_pbservice_node_gen_go style7
    class agent_checks_os_service_windows_go style3
    class agent_connect_ca_testing_go style3
    class agent_grpc_external_services_resource_testing_testing_go style3
    class agent_proxycfg_proxycfg_deepcopy_go style3
    class agent_structs_structs_ce_go style3
    class testing_deployer_util_v2_go style7
    class acl_policy_merger_ce_go style9
    class agent_cache_types_peered_upstreams_go style3
    class agent_consul_reporting_reportingmock_mock_StateDelegate_go style3
    class agent_hcp_telemetry_gauge_store_go style3
    class api_health_go style4
    class command_acl_token_delete_token_delete_go style1
    class command_operator_raft_removepeer_operator_raft_remove_go style1
    class internal_controller_supervisor_go style7
    class agent_consul_rate_mock_RequestLimitsHandler_go style3
    class agent_structs_txn_go style3
    class internal_controller_testing_go style7
    class internal_multicluster_internal_types_exported_services_go style7
    class internal_resource_decode_go style7
    class internal_tools_protoc_gen_consul_rate_limit_postprocess_main_go style7
    class lib_semaphore_semaphore_go style7
    class proto_private_pbautoconf_auto_config_ce_go style7
    class agent_cache_request_go style3
    class agent_config_doc_go style3
    class agent_connect_ca_provider_vault_auth_alicloud_go style3
    class internal_controller_cache_cachemock_mock_WriteCache_go style7
    class internal_storage_inmem_event_index_go style7
    class test_integ_upgrade_basic_common_go style7
    class testing_deployer_util_v2_decode_go style7
    class agent_connect_ca_provider_vault_auth_azure_go style3
    class agent_grpc_internal_handler_go style8
    class agent_xds_rbac_go style3
    class proto_public_pbmulticluster_v2_namespace_exported_services_deepcopy_gen_go style7
    class command_intention_match_match_go style1
    class command_snapshot_decode_snapshot_decode_go style1
    class internal_resourcehcl_any_go style7
    class proto_public_pbmulticluster_v2_exported_services_pb_go style7
    class proto_public_pbmulticluster_v2_exported_services_consumer_deepcopy_gen_go style7
    class proto_public_pbdns_dns_pb_go style7
    class testing_deployer_topology_util_go style7
    class agent_consul_fsm_decode_downgrade_go style3
    class agent_grpc_external_testutils_fsm_go style3
    class agent_keyring_go style3
    class internal_resourcehcl_naming_go style7
    class acl_testing_go style9
    class agent_consul_state_state_store_go style6
    class agent_kvs_endpoint_go style3
    class agent_structs_testing_catalog_go style3
    class grpcmocks_proto_public_pbconnectca_mock_ConnectCAServiceClient_go style7
    class lib_maps_maps_go style7
    class testing_deployer_sprawl_acl_go style7
    class testing_deployer_sprawl_tls_go style7
    class agent_connect_generate_go style3
    class agent_grpc_external_testutils_server_go style3
    class agent_proxycfg_glue_discovery_chain_go style3
    class agent_proxycfg_glue_federation_state_list_mesh_gateways_go style3
    class internal_resource_resourcetest_require_go style7
    class agent_cache_types_connect_ca_root_go style3
    class agent_config_segment_ce_go style3
    class agent_configentry_resolve_go style3
    class agent_consul_state_catalog_schema_deepcopy_go style3
    class command_lock_lock_go style1
    class sdk_testutil_retry_retryer_go style7
    class agent_cache_types_federation_state_list_gateways_go style3
    class agent_consul_session_ttl_go style3
    class command_intention_delete_delete_go style1
    class command_peering_delete_delete_go style1
    class connect_resolver_go style2
    class envoyextensions_xdscommon_proxysupport_go style7
    class sentinel_evaluator_go style7
    class sdk_testutil_server_go style7
    class testing_deployer_sprawl_grpc_go style7
    class agent_connect_ca_provider_vault_auth_k8s_go style3
    class agent_consul_state_config_entry_exported_services_ce_go style3
    class agent_consul_state_events_go style3
    class command_resource_client_client_go style1
    class testing_deployer_topology_compile_go style7
    class acl_chained_authorizer_go style9
    class agent_consul_stream_event_buffer_go style3
    class agent_consul_stats_fetcher_go style3
    class agent_structs_intention_go style4
    class internal_controller_controller_go style7
    class internal_storage_inmem_backend_go style7
    class agent_envoyextensions_builtin_otel_access_logging_structs_go style3
    class agent_proxycfg_glue_exported_peered_services_go style3
    class internal_gossip_libserf_serf_go style7
    class agent_leafcert_generate_go style3
    class command_acl_token_clone_token_clone_go style1
    class internal_controller_cache_index_errors_go style7
    class internal_controller_controllermock_mock_task_go style7
    class acl_errors_ce_go style9
    class agent_netutil_network_go style3
    class agent_metrics_testing_go style3
    class agent_structs_config_entry_sameness_group_ce_go style3
    class agent_xds_configfetcher_config_fetcher_go style3
    class api_operator_audit_go style4
    class internal_controller_cache_kind_go style7
    class agent_connect_ca_provider_vault_auth_approle_go style3
    class agent_consul_acl_endpoint_ce_go style3
    class agent_grpc_middleware_testutil_testservice_simple_pb_go style3
    class agent_structs_connect_proxy_config_go style3
    class api_connect_ca_go style4
    class api_txn_go style4
    class internal_controller_cache_index_indexmock_mock_MultiIndexer_go style7
    class agent_consul_catalog_endpoint_go style10
    class agent_consul_intention_endpoint_go style3
    class logging_logfile_openbsd_go style7
    class agent_consul_server_serf_go style3
    class agent_xds_accesslogs_accesslogs_go style3
    class command_resource_client_usage_go style1
    class internal_controller_cache_decoded_go style7
    class agent_consul_state_federation_state_go style3
    class agent_proxycfg_testing_api_gateway_go style3
    class agent_rpcclient_health_health_go style3
    class agent_structs_config_entry_discoverychain_ce_go style3
    class testrpc_wait_go style7
    class agent_connect_ca_mock_Provider_go style3
    class agent_consul_type_registry_go style3
    class agent_structs_errors_go style3
    class command_connect_ca_set_connect_ca_set_go style1
    class connect_certgen_certgen_go style2
    class proto_private_pbcommon_common_ce_go style7
    class acl_policy_ce_go style9
    class command_cli_cli_go style1
    class test_integration_consul_container_libs_assert_peering_go style7
    class agent_consul_authmethod_kubeauth_k8s_ce_go style3
    class agent_consul_options_go style3
    class agent_consul_v2_config_entry_exports_shim_go style3
    class agent_operator_endpoint_ce_go style3
    class command_operator_autopilot_state_operator_autopilot_state_go style1
    class lib_hoststats_metrics_go style7
    class test_integration_consul_container_libs_service_connect_go style7
    class agent_cache_types_trust_bundle_go style3
    class agent_cache_types_trust_bundles_go style3
    class agent_grpc_middleware_rate_limit_mappings_gen_go style3
    class agent_proxycfg_testing_ingress_gateway_go style3
    class agent_structs_config_entry_exports_ce_go style3
    class api_prepared_query_go style4
    class command_peering_generate_generate_go style1
    class command_version_formatter_go style1
    class internal_multicluster_internal_types_partition_exported_services_go style7
    class internal_resource_resourcetest_validation_go style7
    class agent_cache_types_mock_RPC_go style3
    class agent_consul_leader_log_verification_go style3
    class agent_grpc_external_services_configentry_server_go style3
    class agent_grpc_external_utils_go style3
    class agent_hcp_telemetry_otel_sink_go style3
    class agent_proxycfg_glue_intentions_ce_go style3
    class agent_xds_failover_policy_ce_go style3
    class agent_xds_protocol_trace_go style3
    class agent_cache_types_peerings_go style3
    class agent_checks_docker_unix_go style3
    class agent_consul_acl_ce_go style3
    class agent_consul_reporting_reporting_go style3
    class agent_util_go style3
    class command_acl_authmethod_read_authmethod_read_go style1
    class command_resource_apply_grpc_apply_go style1
    class command_resource_resource_grpc_go style1
    class agent_consul_acl_replication_types_go style3
    class agent_consul_authmethod_kubeauth_testing_go style3
    class agent_consul_session_timers_go style3
    class agent_grpc_middleware_handshake_go style3
    class internal_dnsutil_dns_go style7
    class proto_private_pbacl_acl_go style7
    class proto_public_pbconnectca_ca_json_gen_go style7
    class proto_private_pbservice_service_pb_go style7
    class agent_consul_auth_mock_ACLCache_go style3
    class agent_hcp_testserver_main_go style3
    class agent_structs_acl_cache_go style3
    class command_intention_get_get_go style1
    class command_resource_list_list_go style1
    class lib_telemetry_go style7
    class proto_private_pbdemo_v2_resources_rtypes_go style7
    class test_integration_consul_container_libs_assert_envoy_go style7
    class agent_consul_authmethod_kubeauth_k8s_go style3
    class agent_consul_state_kvs_go style3
    class agent_structs_testing_service_definition_go style3
    class api_config_entry_rate_limit_ip_go style4
    class api_operator_go style4
    class connect_testing_go style2
    class grpcmocks_proto_public_pbserverdiscovery_mock_ServerDiscoveryServiceClient_go style7
    class internal_resource_acls_go style7
    class agent_connect_ca_provider_aws_go style3
    class agent_rpcclient_configentry_view_go style3
    class agent_session_endpoint_go style3
    class command_kv_kv_go style1
    class internal_controller_controllermock_mock_Initializer_go style7
    class sdk_testutil_testlog_go style7
    class testing_deployer_sprawl_consul_go style7
    class troubleshoot_ports_troubleshoot_ports_go style7
    class acl_authorizer_go style9
    class agent_structs_testing_connect_proxy_config_go style3
    class command_members_members_go style1
    class proto_private_pbstorage_raft_grpc_pb_go style7
    class proto_public_pbmulticluster_v2_resources_rtypes_go style7
    class test_integration_consul_container_libs_service_gateway_go style7
    class test_integration_consul_container_libs_utils_retry_go style7
    class agent_cache_cache_go style3
    class agent_connect_uri_go style3
    class agent_http_ce_go style3
    class agent_structs_config_entry_mesh_go style3
    class command_acl_authmethod_delete_authmethod_delete_go style1
    class command_connect_envoy_envoy_go style1
    class command_flags_flag_map_value_go style1
    class proto_public_pbacl_acl_pb_go style7
    class agent_connect_ca_provider_vault_auth_aws_go style3
    class agent_proxycfg_glue_peered_upstreams_go style3
    class agent_testagent_go style3
    class command_kv_impexp_kvimpexp_go style1
    class envoyextensions_extensioncommon_runtime_config_go style7
    class internal_controller_cache_clone_go style7
    class test_integration_consul_container_libs_cluster_log_go style7
    class agent_consul_discoverychain_testing_go style3
    class agent_checks_os_service_unix_go style3
    class agent_config_flags_go style3
    class agent_configentry_merge_service_config_go style3
    class internal_controller_cache_index_indexmock_mock_SingleIndexer_go style7
    class internal_resourcehcl_unmarshal_go style7
    class test_integration_consul_container_libs_cluster_agent_go style7
    class testing_deployer_sprawl_sprawltest_sprawltest_go style7
    class agent_proxycfg_mesh_gateway_ce_go style3
    class command_acl_policy_list_policy_list_go style1
    class connect_tls_go style2
    class internal_multicluster_internal_types_helpers_go style7
    class proto_private_pbpeering_peering_grpc_pb_go style7
    class types_area_go style7
    class agent_cache_types_prepared_query_go style3
    class internal_storage_raft_backend_go style7
    class lib_uuid_go style7
    class proto_public_pbmulticluster_v2_exported_services_json_gen_go style7
    class test_integration_consul_container_libs_service_examples_go style7
    class agent_grpc_external_services_resource_write_go style3
    class internal_controller_cache_cache_go style7
    class grpcmocks_proto_public_pbdataplane_mock_UnsafeDataplaneServiceServer_go style7
    class internal_protohcl_naming_go style7
    class sdk_testutil_retry_run_go style7
    class agent_cache_types_intention_match_go style3
    class agent_cache_types_catalog_list_services_go style3
    class agent_consul_auto_config_backend_go style3
    class agent_consul_controller_queue_rate_go style3
    class agent_consul_operator_raft_endpoint_go style3
    class agent_proxycfg_glue_peering_list_go style3
    class envoyextensions_extensioncommon_resources_go style7
    class internal_resource_reaper_controller_go style7
    class agent_consul_client_go style3
    class agent_structs_config_entry_apigw_jwt_ce_go style3
    class agent_structs_identity_go style3
    class internal_go_sso_oidcauth_oidcjwt_go style7
    class sdk_freeport_ephemeral_darwin_go style7
    class testing_deployer_sprawl_internal_tfgen_digest_go style7
    class test_integration_consul_container_test_resource_http_api_helper_go style7
    class troubleshoot_ports_troubleshoot_protocol_go style7
    class agent_consul_fsm_decode_ce_go style3
    class agent_consul_discoverychain_gateway_go style3
    class agent_exec_exec_unix_go style3
    class agent_structs_acl_ce_go style3
    class agent_structs_config_entry_file_system_certificate_go style3
    class command_acl_policy_read_policy_read_go style1
    class internal_protohcl_cty_go style7
    class internal_resource_resourcetest_tenancy_go style7
    class agent_consul_discovery_chain_endpoint_go style3
    class agent_grpc_middleware_testutil_testservice_simple_grpc_pb_go style3
    class command_acl_role_formatter_go style1
    class agent_consul_acl_authmethod_go style3
    class agent_grpc_internal_resolver_registry_go style3
    class envoyextensions_xdscommon_envoy_versioning_go style7
    class testing_deployer_sprawl_internal_tfgen_tfgen_go style7
    class testing_deployer_sprawl_sprawl_go style7
    class agent_auto_config_config_translate_go style3
    class agent_cache_testing_go style3
    class api_api_go style4
    class command_intention_format_go style1
    class internal_controller_dependency_transform_go style7
    class internal_resource_protoc_gen_resource_types_main_go style7
    class lib_strings_go style7
    class proto_private_pbcommon_common_pb_go style7
    class agent_consul_flood_go style3
    class agent_consul_fsm_commands_ce_go style3
    class agent_consul_server_metadata_go style3
    class agent_grpc_internal_services_subscribe_logger_go style3
    class agent_submatview_materializer_go style3
    class proto_public_pbresource_cloning_stream_pb_go style7
    class sdk_freeport_ephemeral_fallback_go style7
    class agent_consul_context_go style3
    class agent_proxycfg_glue_intention_upstreams_go style3
    class command_acl_agenttokens_agent_tokens_go style1
    class command_operator_usage_instances_usage_instances_go style1
    class internal_controller_cache_index_indexmock_mock_Indexer_go style7
    class proto_public_pbserverdiscovery_cloning_stream_pb_go style7
    class agent_hcp_deps_go style3
    class command_connect_envoy_exec_supported_go style1
    class command_connect_proxy_proxy_go style1
    class grpcmocks_proto_public_pbconnectca_mock_UnsafeConnectCAServiceServer_go style7
    class proto_private_pbpeering_peering_go style7
    class agent_config_runtime_ce_go style3
    class agent_consul_filter_go style3
    class agent_consul_leader_federation_state_ae_go style3
    class grpcmocks_proto_public_pbdataplane_mock_DataplaneServiceServer_go style7
    class internal_controller_lease_go style7
    class proto_public_pbacl_acl_grpc_pb_go style7
    class sdk_freeport_freeport_go style7
    class agent_consul_server_connect_go style3
    class agent_pool_peeked_conn_go style3
    class agent_structs_operator_go style3
    class command_acl_templatedpolicy_read_templated_policy_read_go style1
    class command_flags_http_go style1
    class command_operator_usage_instances_usage_instances_ce_go style1
    class proto_private_pbconfigentry_config_entry_ce_go style7
    class agent_connect_ca_provider_vault_auth_go style3
    class agent_connect_uri_server_go style3
    class agent_connect_ca_endpoint_go style3
    class agent_consul_authmethod_ssoauth_sso_ce_go style3
    class agent_proxycfg_config_snapshot_glue_go style3
    class tlsutil_generate_go style0
    class agent_connect_authz_go style3
    class agent_consul_state_config_entry_events_go style3
    class agent_proxycfg_testing_ce_go style3
    class agent_xds_testcommon_testcommon_go style3
    class proto_private_pbacl_acl_pb_go style7
    class agent_rpc_operator_service_go style3
    class command_snapshot_inspect_formatter_go style1
    class proto_private_pbsubscribe_subscribe_pb_go style7
    class agent_config_limits_go style3
    class agent_connect_ca_provider_go style3
    class agent_consul_leader_intentions_go style3
    class grpcmocks_proto_public_pbresource_mock_ResourceService_WatchListClient_go style7
    class agent_ui_endpoint_go style3
    class internal_controller_cache_cachemock_mock_Cache_go style7
    class internal_controller_dependencies_go style7
    class testing_deployer_sprawl_internal_tfgen_nodes_go style7
    class agent_consul_leader_connect_go style3
    class agent_auto_config_config_go style3
    class agent_connect_ca_provider_consul_config_go style3
    class agent_consul_stream_string_types_go style3
    class agent_catalog_endpoint_go style3
    class agent_envoyextensions_builtin_property_override_structpatcher_go style3
    class agent_grpc_external_services_connectca_server_go style3
    class api_operator_autopilot_go style4
    class command_snapshot_inspect_snapshot_inspect_go style1
    class test_integration_consul_container_libs_service_common_go style7
    class test_integration_consul_container_libs_cluster_dataplane_go style7
    class testing_deployer_sprawl_ent_go style7
    class agent_consul_segment_ce_go style3
    class agent_proxycfg_manager_go style2
    class api_snapshot_go style4
    class command_acl_templatedpolicy_templated_policy_go style1
    class command_kv_del_kv_delete_go style1

```
