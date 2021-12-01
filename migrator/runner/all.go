package runner

import (
	// Import these packages to trigger the registration.
	_ "github.com/stackrox/rox/migrator/migrations/m_44_to_m_45_rocksdb_clusters"
	_ "github.com/stackrox/rox/migrator/migrations/m_45_to_m_46_imagecveedge"
	_ "github.com/stackrox/rox/migrator/migrations/m_46_to_m_47_compliance_in_rocksdb"
	_ "github.com/stackrox/rox/migrator/migrations/m_47_to_m_48_drop_bolt_buckets"
	_ "github.com/stackrox/rox/migrator/migrations/m_48_to_m_49_externalize_compliance_domains"
	_ "github.com/stackrox/rox/migrator/migrations/m_49_to_m_50_deprecate_email_label_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_50_to_m_51_default_integration_health"
	_ "github.com/stackrox/rox/migrator/migrations/m_51_to_m_52_remove_invalid_transactions"
	_ "github.com/stackrox/rox/migrator/migrations/m_52_to_m_53_aws_endpoints"
	_ "github.com/stackrox/rox/migrator/migrations/m_53_to_m_54_enable_exec_webhook"
	_ "github.com/stackrox/rox/migrator/migrations/m_54_to_m_55_deprecate_policy_whitelist"
	_ "github.com/stackrox/rox/migrator/migrations/m_55_to_m_56_node_scanning_empty"
	_ "github.com/stackrox/rox/migrator/migrations/m_56_to_m_57_compliance_policy_categories"
	_ "github.com/stackrox/rox/migrator/migrations/m_57_to_m_58_update_run_secrets_volume_policy_regex"
	_ "github.com/stackrox/rox/migrator/migrations/m_58_to_m_59_node_scanning_flag_on"
	_ "github.com/stackrox/rox/migrator/migrations/m_59_to_m_60_add_docker_cis_category_to_existing"
	_ "github.com/stackrox/rox/migrator/migrations/m_60_to_m_61_update_network_management_policy_regex"
	_ "github.com/stackrox/rox/migrator/migrations/m_61_to_m_62_multiple_cve_types"
	_ "github.com/stackrox/rox/migrator/migrations/m_62_to_m_63_splunk_source_type"
	_ "github.com/stackrox/rox/migrator/migrations/m_63_to_m_64_exclude_some_openshift_operators_from_policies"
	_ "github.com/stackrox/rox/migrator/migrations/m_64_to_m_65_detect_openshift4_cluster_on_exec_webhooks"
	_ "github.com/stackrox/rox/migrator/migrations/m_65_to_m_66_policy_bug_fixes"
	_ "github.com/stackrox/rox/migrator/migrations/m_66_to_m_67_missing_policy_migrations"
	_ "github.com/stackrox/rox/migrator/migrations/m_67_to_m_68_exclude_pdcsi_from_mount_propagation"
	_ "github.com/stackrox/rox/migrator/migrations/m_68_to_m_69_update_global_access_roles"
	_ "github.com/stackrox/rox/migrator/migrations/m_69_to_m_70_add_xmrig_to_crypto_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_70_to_m_71_disable_audit_log_collection"
	_ "github.com/stackrox/rox/migrator/migrations/m_71_to_m_72_delete_namespacesac_bucket"
	_ "github.com/stackrox/rox/migrator/migrations/m_72_to_m_73_change_roles_to_sac_v2"
	_ "github.com/stackrox/rox/migrator/migrations/m_73_to_m_74_runtime_policy_event_source"
	_ "github.com/stackrox/rox/migrator/migrations/m_74_to_m_75_severity_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_75_to_m_76_exclude_compliance_operator_dnf_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_76_to_m_77_move_roles_to_rocksdb"
	_ "github.com/stackrox/rox/migrator/migrations/m_77_to_m_78_mitre"
	_ "github.com/stackrox/rox/migrator/migrations/m_78_to_m_79_exclude_openshift_sdn_host_pids_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_79_to_m_80_more_openshift_exclusions"
	_ "github.com/stackrox/rox/migrator/migrations/m_80_to_m_81_rm_demo_policies"
	_ "github.com/stackrox/rox/migrator/migrations/m_81_to_m_82_modify_docker_policies"
	_ "github.com/stackrox/rox/migrator/migrations/m_82_to_m_83_default_pol_flag"
	_ "github.com/stackrox/rox/migrator/migrations/m_83_to_m_84_mitre_fixes"
	_ "github.com/stackrox/rox/migrator/migrations/m_84_to_m_85_exclude_compliance_op_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_85_to_m_86_apktools_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_86_to_m_87_microdnf_policy"
	_ "github.com/stackrox/rox/migrator/migrations/m_87_to_m_88_central_secret_policy"
)
