// Code generated by "stringer"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[V21_1-0]
	_ = x[Start21_1PLUS-1]
	_ = x[Start21_2-2]
	_ = x[AcquisitionTypeInLeaseHistory-3]
	_ = x[SerializeViewUDTs-4]
	_ = x[ExpressionIndexes-5]
	_ = x[DeleteDeprecatedNamespaceTableDescriptorMigration-6]
	_ = x[FixDescriptors-7]
	_ = x[DatabaseRoleSettings-8]
	_ = x[TenantUsageTable-9]
	_ = x[SQLInstancesTable-10]
	_ = x[NewRetryableRangefeedErrors-11]
	_ = x[AlterSystemWebSessionsCreateIndexes-12]
	_ = x[SeparatedIntentsMigration-13]
	_ = x[PostSeparatedIntentsMigration-14]
	_ = x[RetryJobsWithExponentialBackoff-15]
	_ = x[AutoSpanConfigReconciliationJob-16]
	_ = x[DefaultPrivileges-17]
	_ = x[ZonesTableForSecondaryTenants-18]
	_ = x[UseKeyEncodeForHashShardedIndexes-19]
	_ = x[DatabasePlacementPolicy-20]
	_ = x[GeneratedAsIdentity-21]
	_ = x[OnUpdateExpressions-22]
	_ = x[SpanConfigurationsTable-23]
	_ = x[BoundedStaleness-24]
	_ = x[DateAndIntervalStyle-25]
	_ = x[TenantUsageSingleConsumptionColumn-26]
	_ = x[SQLStatsTables-27]
	_ = x[SQLStatsCompactionScheduledJob-28]
	_ = x[V21_2-29]
	_ = x[Start22_1-30]
	_ = x[TargetBytesAvoidExcess-31]
	_ = x[AvoidDrainingNames-32]
	_ = x[DrainingNamesMigration-33]
	_ = x[TraceIDDoesntImplyStructuredRecording-34]
	_ = x[AlterSystemTableStatisticsAddAvgSizeCol-35]
	_ = x[AlterSystemStmtDiagReqs-36]
	_ = x[MVCCAddSSTable-37]
	_ = x[InsertPublicSchemaNamespaceEntryOnRestore-38]
	_ = x[UnsplitRangesInAsyncGCJobs-39]
	_ = x[ValidateGrantOption-40]
	_ = x[PebbleFormatBlockPropertyCollector-41]
	_ = x[ProbeRequest-42]
	_ = x[SelectRPCsTakeTracingInfoInband-43]
	_ = x[PreSeedTenantSpanConfigs-44]
	_ = x[SeedTenantSpanConfigs-45]
	_ = x[PublicSchemasWithDescriptors-46]
}

const _Key_name = "V21_1Start21_1PLUSStart21_2AcquisitionTypeInLeaseHistorySerializeViewUDTsExpressionIndexesDeleteDeprecatedNamespaceTableDescriptorMigrationFixDescriptorsDatabaseRoleSettingsTenantUsageTableSQLInstancesTableNewRetryableRangefeedErrorsAlterSystemWebSessionsCreateIndexesSeparatedIntentsMigrationPostSeparatedIntentsMigrationRetryJobsWithExponentialBackoffAutoSpanConfigReconciliationJobDefaultPrivilegesZonesTableForSecondaryTenantsUseKeyEncodeForHashShardedIndexesDatabasePlacementPolicyGeneratedAsIdentityOnUpdateExpressionsSpanConfigurationsTableBoundedStalenessDateAndIntervalStyleTenantUsageSingleConsumptionColumnSQLStatsTablesSQLStatsCompactionScheduledJobV21_2Start22_1TargetBytesAvoidExcessAvoidDrainingNamesDrainingNamesMigrationTraceIDDoesntImplyStructuredRecordingAlterSystemTableStatisticsAddAvgSizeColAlterSystemStmtDiagReqsMVCCAddSSTableInsertPublicSchemaNamespaceEntryOnRestoreUnsplitRangesInAsyncGCJobsValidateGrantOptionPebbleFormatBlockPropertyCollectorProbeRequestSelectRPCsTakeTracingInfoInbandPreSeedTenantSpanConfigsSeedTenantSpanConfigsPublicSchemasWithDescriptors"

var _Key_index = [...]uint16{0, 5, 18, 27, 56, 73, 90, 139, 153, 173, 189, 206, 233, 268, 293, 322, 353, 384, 401, 430, 463, 486, 505, 524, 547, 563, 583, 617, 631, 661, 666, 675, 697, 715, 737, 774, 813, 836, 850, 891, 917, 936, 970, 982, 1013, 1037, 1058, 1086}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
