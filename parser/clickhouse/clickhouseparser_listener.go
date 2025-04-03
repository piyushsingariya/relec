// Code generated from ClickHouseParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package clickhouse // ClickHouseParser
import "github.com/antlr4-go/antlr/v4"

// ClickHouseParserListener is a complete listener for a parse tree produced by ClickHouseParser.
type ClickHouseParserListener interface {
	antlr.ParseTreeListener

	// EnterQueryStmt is called when entering the queryStmt production.
	EnterQueryStmt(c *QueryStmtContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterCtes is called when entering the ctes production.
	EnterCtes(c *CtesContext)

	// EnterNamedQuery is called when entering the namedQuery production.
	EnterNamedQuery(c *NamedQueryContext)

	// EnterColumnAliases is called when entering the columnAliases production.
	EnterColumnAliases(c *ColumnAliasesContext)

	// EnterAlterTableStmt is called when entering the AlterTableStmt production.
	EnterAlterTableStmt(c *AlterTableStmtContext)

	// EnterAlterTableClauseAddColumn is called when entering the AlterTableClauseAddColumn production.
	EnterAlterTableClauseAddColumn(c *AlterTableClauseAddColumnContext)

	// EnterAlterTableClauseAddIndex is called when entering the AlterTableClauseAddIndex production.
	EnterAlterTableClauseAddIndex(c *AlterTableClauseAddIndexContext)

	// EnterAlterTableClauseAddProjection is called when entering the AlterTableClauseAddProjection production.
	EnterAlterTableClauseAddProjection(c *AlterTableClauseAddProjectionContext)

	// EnterAlterTableClauseAttach is called when entering the AlterTableClauseAttach production.
	EnterAlterTableClauseAttach(c *AlterTableClauseAttachContext)

	// EnterAlterTableClauseClearColumn is called when entering the AlterTableClauseClearColumn production.
	EnterAlterTableClauseClearColumn(c *AlterTableClauseClearColumnContext)

	// EnterAlterTableClauseClearIndex is called when entering the AlterTableClauseClearIndex production.
	EnterAlterTableClauseClearIndex(c *AlterTableClauseClearIndexContext)

	// EnterAlterTableClauseClearProjection is called when entering the AlterTableClauseClearProjection production.
	EnterAlterTableClauseClearProjection(c *AlterTableClauseClearProjectionContext)

	// EnterAlterTableClauseComment is called when entering the AlterTableClauseComment production.
	EnterAlterTableClauseComment(c *AlterTableClauseCommentContext)

	// EnterAlterTableClauseDelete is called when entering the AlterTableClauseDelete production.
	EnterAlterTableClauseDelete(c *AlterTableClauseDeleteContext)

	// EnterAlterTableClauseDetach is called when entering the AlterTableClauseDetach production.
	EnterAlterTableClauseDetach(c *AlterTableClauseDetachContext)

	// EnterAlterTableClauseDropColumn is called when entering the AlterTableClauseDropColumn production.
	EnterAlterTableClauseDropColumn(c *AlterTableClauseDropColumnContext)

	// EnterAlterTableClauseDropIndex is called when entering the AlterTableClauseDropIndex production.
	EnterAlterTableClauseDropIndex(c *AlterTableClauseDropIndexContext)

	// EnterAlterTableClauseDropProjection is called when entering the AlterTableClauseDropProjection production.
	EnterAlterTableClauseDropProjection(c *AlterTableClauseDropProjectionContext)

	// EnterAlterTableClauseDropPartition is called when entering the AlterTableClauseDropPartition production.
	EnterAlterTableClauseDropPartition(c *AlterTableClauseDropPartitionContext)

	// EnterAlterTableClauseFreezePartition is called when entering the AlterTableClauseFreezePartition production.
	EnterAlterTableClauseFreezePartition(c *AlterTableClauseFreezePartitionContext)

	// EnterAlterTableClauseMaterializeIndex is called when entering the AlterTableClauseMaterializeIndex production.
	EnterAlterTableClauseMaterializeIndex(c *AlterTableClauseMaterializeIndexContext)

	// EnterAlterTableClauseMaterializeProjection is called when entering the AlterTableClauseMaterializeProjection production.
	EnterAlterTableClauseMaterializeProjection(c *AlterTableClauseMaterializeProjectionContext)

	// EnterAlterTableClauseModifyCodec is called when entering the AlterTableClauseModifyCodec production.
	EnterAlterTableClauseModifyCodec(c *AlterTableClauseModifyCodecContext)

	// EnterAlterTableClauseModifyComment is called when entering the AlterTableClauseModifyComment production.
	EnterAlterTableClauseModifyComment(c *AlterTableClauseModifyCommentContext)

	// EnterAlterTableClauseModifyRemove is called when entering the AlterTableClauseModifyRemove production.
	EnterAlterTableClauseModifyRemove(c *AlterTableClauseModifyRemoveContext)

	// EnterAlterTableClauseModify is called when entering the AlterTableClauseModify production.
	EnterAlterTableClauseModify(c *AlterTableClauseModifyContext)

	// EnterAlterTableClauseModifyOrderBy is called when entering the AlterTableClauseModifyOrderBy production.
	EnterAlterTableClauseModifyOrderBy(c *AlterTableClauseModifyOrderByContext)

	// EnterAlterTableClauseModifyTTL is called when entering the AlterTableClauseModifyTTL production.
	EnterAlterTableClauseModifyTTL(c *AlterTableClauseModifyTTLContext)

	// EnterAlterTableClauseMovePartition is called when entering the AlterTableClauseMovePartition production.
	EnterAlterTableClauseMovePartition(c *AlterTableClauseMovePartitionContext)

	// EnterAlterTableClauseRemoveTTL is called when entering the AlterTableClauseRemoveTTL production.
	EnterAlterTableClauseRemoveTTL(c *AlterTableClauseRemoveTTLContext)

	// EnterAlterTableClauseRename is called when entering the AlterTableClauseRename production.
	EnterAlterTableClauseRename(c *AlterTableClauseRenameContext)

	// EnterAlterTableClauseReplace is called when entering the AlterTableClauseReplace production.
	EnterAlterTableClauseReplace(c *AlterTableClauseReplaceContext)

	// EnterAlterTableClauseUpdate is called when entering the AlterTableClauseUpdate production.
	EnterAlterTableClauseUpdate(c *AlterTableClauseUpdateContext)

	// EnterAssignmentExprList is called when entering the assignmentExprList production.
	EnterAssignmentExprList(c *AssignmentExprListContext)

	// EnterAssignmentExpr is called when entering the assignmentExpr production.
	EnterAssignmentExpr(c *AssignmentExprContext)

	// EnterTableColumnPropertyType is called when entering the tableColumnPropertyType production.
	EnterTableColumnPropertyType(c *TableColumnPropertyTypeContext)

	// EnterPartitionClause is called when entering the partitionClause production.
	EnterPartitionClause(c *PartitionClauseContext)

	// EnterAttachDictionaryStmt is called when entering the AttachDictionaryStmt production.
	EnterAttachDictionaryStmt(c *AttachDictionaryStmtContext)

	// EnterCheckStmt is called when entering the checkStmt production.
	EnterCheckStmt(c *CheckStmtContext)

	// EnterCreateDatabaseStmt is called when entering the CreateDatabaseStmt production.
	EnterCreateDatabaseStmt(c *CreateDatabaseStmtContext)

	// EnterCreateDictionaryStmt is called when entering the CreateDictionaryStmt production.
	EnterCreateDictionaryStmt(c *CreateDictionaryStmtContext)

	// EnterCreateLiveViewStmt is called when entering the CreateLiveViewStmt production.
	EnterCreateLiveViewStmt(c *CreateLiveViewStmtContext)

	// EnterCreateMaterializedViewStmt is called when entering the CreateMaterializedViewStmt production.
	EnterCreateMaterializedViewStmt(c *CreateMaterializedViewStmtContext)

	// EnterCreateTableStmt is called when entering the CreateTableStmt production.
	EnterCreateTableStmt(c *CreateTableStmtContext)

	// EnterCreateViewStmt is called when entering the CreateViewStmt production.
	EnterCreateViewStmt(c *CreateViewStmtContext)

	// EnterDictionarySchemaClause is called when entering the dictionarySchemaClause production.
	EnterDictionarySchemaClause(c *DictionarySchemaClauseContext)

	// EnterDictionaryAttrDfnt is called when entering the dictionaryAttrDfnt production.
	EnterDictionaryAttrDfnt(c *DictionaryAttrDfntContext)

	// EnterDictionaryEngineClause is called when entering the dictionaryEngineClause production.
	EnterDictionaryEngineClause(c *DictionaryEngineClauseContext)

	// EnterDictionaryPrimaryKeyClause is called when entering the dictionaryPrimaryKeyClause production.
	EnterDictionaryPrimaryKeyClause(c *DictionaryPrimaryKeyClauseContext)

	// EnterDictionaryArgExpr is called when entering the dictionaryArgExpr production.
	EnterDictionaryArgExpr(c *DictionaryArgExprContext)

	// EnterSourceClause is called when entering the sourceClause production.
	EnterSourceClause(c *SourceClauseContext)

	// EnterLifetimeClause is called when entering the lifetimeClause production.
	EnterLifetimeClause(c *LifetimeClauseContext)

	// EnterLayoutClause is called when entering the layoutClause production.
	EnterLayoutClause(c *LayoutClauseContext)

	// EnterRangeClause is called when entering the rangeClause production.
	EnterRangeClause(c *RangeClauseContext)

	// EnterDictionarySettingsClause is called when entering the dictionarySettingsClause production.
	EnterDictionarySettingsClause(c *DictionarySettingsClauseContext)

	// EnterClusterClause is called when entering the clusterClause production.
	EnterClusterClause(c *ClusterClauseContext)

	// EnterUuidClause is called when entering the uuidClause production.
	EnterUuidClause(c *UuidClauseContext)

	// EnterDestinationClause is called when entering the destinationClause production.
	EnterDestinationClause(c *DestinationClauseContext)

	// EnterSubqueryClause is called when entering the subqueryClause production.
	EnterSubqueryClause(c *SubqueryClauseContext)

	// EnterSchemaDescriptionClause is called when entering the SchemaDescriptionClause production.
	EnterSchemaDescriptionClause(c *SchemaDescriptionClauseContext)

	// EnterSchemaAsTableClause is called when entering the SchemaAsTableClause production.
	EnterSchemaAsTableClause(c *SchemaAsTableClauseContext)

	// EnterSchemaAsFunctionClause is called when entering the SchemaAsFunctionClause production.
	EnterSchemaAsFunctionClause(c *SchemaAsFunctionClauseContext)

	// EnterEngineClause is called when entering the engineClause production.
	EnterEngineClause(c *EngineClauseContext)

	// EnterPartitionByClause is called when entering the partitionByClause production.
	EnterPartitionByClause(c *PartitionByClauseContext)

	// EnterPrimaryKeyClause is called when entering the primaryKeyClause production.
	EnterPrimaryKeyClause(c *PrimaryKeyClauseContext)

	// EnterSampleByClause is called when entering the sampleByClause production.
	EnterSampleByClause(c *SampleByClauseContext)

	// EnterTtlClause is called when entering the ttlClause production.
	EnterTtlClause(c *TtlClauseContext)

	// EnterEngineExpr is called when entering the engineExpr production.
	EnterEngineExpr(c *EngineExprContext)

	// EnterTableElementExprColumn is called when entering the TableElementExprColumn production.
	EnterTableElementExprColumn(c *TableElementExprColumnContext)

	// EnterTableElementExprConstraint is called when entering the TableElementExprConstraint production.
	EnterTableElementExprConstraint(c *TableElementExprConstraintContext)

	// EnterTableElementExprIndex is called when entering the TableElementExprIndex production.
	EnterTableElementExprIndex(c *TableElementExprIndexContext)

	// EnterTableElementExprProjection is called when entering the TableElementExprProjection production.
	EnterTableElementExprProjection(c *TableElementExprProjectionContext)

	// EnterTableColumnDfnt is called when entering the tableColumnDfnt production.
	EnterTableColumnDfnt(c *TableColumnDfntContext)

	// EnterTableColumnPropertyExpr is called when entering the tableColumnPropertyExpr production.
	EnterTableColumnPropertyExpr(c *TableColumnPropertyExprContext)

	// EnterTableIndexDfnt is called when entering the tableIndexDfnt production.
	EnterTableIndexDfnt(c *TableIndexDfntContext)

	// EnterTableProjectionDfnt is called when entering the tableProjectionDfnt production.
	EnterTableProjectionDfnt(c *TableProjectionDfntContext)

	// EnterCodecExpr is called when entering the codecExpr production.
	EnterCodecExpr(c *CodecExprContext)

	// EnterCodecArgExpr is called when entering the codecArgExpr production.
	EnterCodecArgExpr(c *CodecArgExprContext)

	// EnterTtlExpr is called when entering the ttlExpr production.
	EnterTtlExpr(c *TtlExprContext)

	// EnterDescribeStmt is called when entering the describeStmt production.
	EnterDescribeStmt(c *DescribeStmtContext)

	// EnterDropDatabaseStmt is called when entering the DropDatabaseStmt production.
	EnterDropDatabaseStmt(c *DropDatabaseStmtContext)

	// EnterDropTableStmt is called when entering the DropTableStmt production.
	EnterDropTableStmt(c *DropTableStmtContext)

	// EnterDeleteStmt is called when entering the deleteStmt production.
	EnterDeleteStmt(c *DeleteStmtContext)

	// EnterExistsDatabaseStmt is called when entering the ExistsDatabaseStmt production.
	EnterExistsDatabaseStmt(c *ExistsDatabaseStmtContext)

	// EnterExistsTableStmt is called when entering the ExistsTableStmt production.
	EnterExistsTableStmt(c *ExistsTableStmtContext)

	// EnterExplainASTStmt is called when entering the ExplainASTStmt production.
	EnterExplainASTStmt(c *ExplainASTStmtContext)

	// EnterExplainSyntaxStmt is called when entering the ExplainSyntaxStmt production.
	EnterExplainSyntaxStmt(c *ExplainSyntaxStmtContext)

	// EnterInsertStmt is called when entering the insertStmt production.
	EnterInsertStmt(c *InsertStmtContext)

	// EnterColumnsClause is called when entering the columnsClause production.
	EnterColumnsClause(c *ColumnsClauseContext)

	// EnterDataClauseFormat is called when entering the DataClauseFormat production.
	EnterDataClauseFormat(c *DataClauseFormatContext)

	// EnterDataClauseValues is called when entering the DataClauseValues production.
	EnterDataClauseValues(c *DataClauseValuesContext)

	// EnterDataClauseSelect is called when entering the DataClauseSelect production.
	EnterDataClauseSelect(c *DataClauseSelectContext)

	// EnterKillMutationStmt is called when entering the KillMutationStmt production.
	EnterKillMutationStmt(c *KillMutationStmtContext)

	// EnterOptimizeStmt is called when entering the optimizeStmt production.
	EnterOptimizeStmt(c *OptimizeStmtContext)

	// EnterRenameStmt is called when entering the renameStmt production.
	EnterRenameStmt(c *RenameStmtContext)

	// EnterProjectionSelectStmt is called when entering the projectionSelectStmt production.
	EnterProjectionSelectStmt(c *ProjectionSelectStmtContext)

	// EnterSelectUnionStmt is called when entering the selectUnionStmt production.
	EnterSelectUnionStmt(c *SelectUnionStmtContext)

	// EnterSelectStmtWithParens is called when entering the selectStmtWithParens production.
	EnterSelectStmtWithParens(c *SelectStmtWithParensContext)

	// EnterSelectStmt is called when entering the selectStmt production.
	EnterSelectStmt(c *SelectStmtContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterTopClause is called when entering the topClause production.
	EnterTopClause(c *TopClauseContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterArrayJoinClause is called when entering the arrayJoinClause production.
	EnterArrayJoinClause(c *ArrayJoinClauseContext)

	// EnterWindowClause is called when entering the windowClause production.
	EnterWindowClause(c *WindowClauseContext)

	// EnterPrewhereClause is called when entering the prewhereClause production.
	EnterPrewhereClause(c *PrewhereClauseContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterProjectionOrderByClause is called when entering the projectionOrderByClause production.
	EnterProjectionOrderByClause(c *ProjectionOrderByClauseContext)

	// EnterLimitByClause is called when entering the limitByClause production.
	EnterLimitByClause(c *LimitByClauseContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterSettingsClause is called when entering the settingsClause production.
	EnterSettingsClause(c *SettingsClauseContext)

	// EnterJoinExprOp is called when entering the JoinExprOp production.
	EnterJoinExprOp(c *JoinExprOpContext)

	// EnterJoinExprTable is called when entering the JoinExprTable production.
	EnterJoinExprTable(c *JoinExprTableContext)

	// EnterJoinExprParens is called when entering the JoinExprParens production.
	EnterJoinExprParens(c *JoinExprParensContext)

	// EnterJoinExprCrossOp is called when entering the JoinExprCrossOp production.
	EnterJoinExprCrossOp(c *JoinExprCrossOpContext)

	// EnterJoinOpInner is called when entering the JoinOpInner production.
	EnterJoinOpInner(c *JoinOpInnerContext)

	// EnterJoinOpLeftRight is called when entering the JoinOpLeftRight production.
	EnterJoinOpLeftRight(c *JoinOpLeftRightContext)

	// EnterJoinOpFull is called when entering the JoinOpFull production.
	EnterJoinOpFull(c *JoinOpFullContext)

	// EnterJoinOpCross is called when entering the joinOpCross production.
	EnterJoinOpCross(c *JoinOpCrossContext)

	// EnterJoinConstraintClause is called when entering the joinConstraintClause production.
	EnterJoinConstraintClause(c *JoinConstraintClauseContext)

	// EnterSampleClause is called when entering the sampleClause production.
	EnterSampleClause(c *SampleClauseContext)

	// EnterLimitExpr is called when entering the limitExpr production.
	EnterLimitExpr(c *LimitExprContext)

	// EnterOrderExprList is called when entering the orderExprList production.
	EnterOrderExprList(c *OrderExprListContext)

	// EnterOrderExpr is called when entering the orderExpr production.
	EnterOrderExpr(c *OrderExprContext)

	// EnterRatioExpr is called when entering the ratioExpr production.
	EnterRatioExpr(c *RatioExprContext)

	// EnterSettingExprList is called when entering the settingExprList production.
	EnterSettingExprList(c *SettingExprListContext)

	// EnterSettingExpr is called when entering the settingExpr production.
	EnterSettingExpr(c *SettingExprContext)

	// EnterWindowExpr is called when entering the windowExpr production.
	EnterWindowExpr(c *WindowExprContext)

	// EnterWinPartitionByClause is called when entering the winPartitionByClause production.
	EnterWinPartitionByClause(c *WinPartitionByClauseContext)

	// EnterWinOrderByClause is called when entering the winOrderByClause production.
	EnterWinOrderByClause(c *WinOrderByClauseContext)

	// EnterWinFrameClause is called when entering the winFrameClause production.
	EnterWinFrameClause(c *WinFrameClauseContext)

	// EnterFrameStart is called when entering the frameStart production.
	EnterFrameStart(c *FrameStartContext)

	// EnterFrameBetween is called when entering the frameBetween production.
	EnterFrameBetween(c *FrameBetweenContext)

	// EnterWinFrameBound is called when entering the winFrameBound production.
	EnterWinFrameBound(c *WinFrameBoundContext)

	// EnterSetStmt is called when entering the setStmt production.
	EnterSetStmt(c *SetStmtContext)

	// EnterShowCreateDatabaseStmt is called when entering the showCreateDatabaseStmt production.
	EnterShowCreateDatabaseStmt(c *ShowCreateDatabaseStmtContext)

	// EnterShowCreateDictionaryStmt is called when entering the showCreateDictionaryStmt production.
	EnterShowCreateDictionaryStmt(c *ShowCreateDictionaryStmtContext)

	// EnterShowCreateTableStmt is called when entering the showCreateTableStmt production.
	EnterShowCreateTableStmt(c *ShowCreateTableStmtContext)

	// EnterShowDatabasesStmt is called when entering the showDatabasesStmt production.
	EnterShowDatabasesStmt(c *ShowDatabasesStmtContext)

	// EnterShowDictionariesStmt is called when entering the showDictionariesStmt production.
	EnterShowDictionariesStmt(c *ShowDictionariesStmtContext)

	// EnterShowTablesStmt is called when entering the showTablesStmt production.
	EnterShowTablesStmt(c *ShowTablesStmtContext)

	// EnterSystemStmt is called when entering the systemStmt production.
	EnterSystemStmt(c *SystemStmtContext)

	// EnterTruncateStmt is called when entering the truncateStmt production.
	EnterTruncateStmt(c *TruncateStmtContext)

	// EnterUseStmt is called when entering the useStmt production.
	EnterUseStmt(c *UseStmtContext)

	// EnterWatchStmt is called when entering the watchStmt production.
	EnterWatchStmt(c *WatchStmtContext)

	// EnterColumnTypeExprSimple is called when entering the ColumnTypeExprSimple production.
	EnterColumnTypeExprSimple(c *ColumnTypeExprSimpleContext)

	// EnterColumnTypeExprNested is called when entering the ColumnTypeExprNested production.
	EnterColumnTypeExprNested(c *ColumnTypeExprNestedContext)

	// EnterColumnTypeExprEnum is called when entering the ColumnTypeExprEnum production.
	EnterColumnTypeExprEnum(c *ColumnTypeExprEnumContext)

	// EnterColumnTypeExprComplex is called when entering the ColumnTypeExprComplex production.
	EnterColumnTypeExprComplex(c *ColumnTypeExprComplexContext)

	// EnterColumnTypeExprParam is called when entering the ColumnTypeExprParam production.
	EnterColumnTypeExprParam(c *ColumnTypeExprParamContext)

	// EnterColumnExprList is called when entering the columnExprList production.
	EnterColumnExprList(c *ColumnExprListContext)

	// EnterColumnsExprAsterisk is called when entering the ColumnsExprAsterisk production.
	EnterColumnsExprAsterisk(c *ColumnsExprAsteriskContext)

	// EnterColumnsExprSubquery is called when entering the ColumnsExprSubquery production.
	EnterColumnsExprSubquery(c *ColumnsExprSubqueryContext)

	// EnterColumnsExprColumn is called when entering the ColumnsExprColumn production.
	EnterColumnsExprColumn(c *ColumnsExprColumnContext)

	// EnterColumnExprTernaryOp is called when entering the ColumnExprTernaryOp production.
	EnterColumnExprTernaryOp(c *ColumnExprTernaryOpContext)

	// EnterColumnExprAlias is called when entering the ColumnExprAlias production.
	EnterColumnExprAlias(c *ColumnExprAliasContext)

	// EnterColumnExprExtract is called when entering the ColumnExprExtract production.
	EnterColumnExprExtract(c *ColumnExprExtractContext)

	// EnterColumnExprNegate is called when entering the ColumnExprNegate production.
	EnterColumnExprNegate(c *ColumnExprNegateContext)

	// EnterColumnExprSubquery is called when entering the ColumnExprSubquery production.
	EnterColumnExprSubquery(c *ColumnExprSubqueryContext)

	// EnterColumnExprLiteral is called when entering the ColumnExprLiteral production.
	EnterColumnExprLiteral(c *ColumnExprLiteralContext)

	// EnterColumnExprArray is called when entering the ColumnExprArray production.
	EnterColumnExprArray(c *ColumnExprArrayContext)

	// EnterColumnExprSubstring is called when entering the ColumnExprSubstring production.
	EnterColumnExprSubstring(c *ColumnExprSubstringContext)

	// EnterColumnExprCast is called when entering the ColumnExprCast production.
	EnterColumnExprCast(c *ColumnExprCastContext)

	// EnterColumnExprOr is called when entering the ColumnExprOr production.
	EnterColumnExprOr(c *ColumnExprOrContext)

	// EnterColumnExprPrecedence1 is called when entering the ColumnExprPrecedence1 production.
	EnterColumnExprPrecedence1(c *ColumnExprPrecedence1Context)

	// EnterColumnExprPrecedence2 is called when entering the ColumnExprPrecedence2 production.
	EnterColumnExprPrecedence2(c *ColumnExprPrecedence2Context)

	// EnterColumnExprPrecedence3 is called when entering the ColumnExprPrecedence3 production.
	EnterColumnExprPrecedence3(c *ColumnExprPrecedence3Context)

	// EnterColumnExprInterval is called when entering the ColumnExprInterval production.
	EnterColumnExprInterval(c *ColumnExprIntervalContext)

	// EnterColumnExprIsNull is called when entering the ColumnExprIsNull production.
	EnterColumnExprIsNull(c *ColumnExprIsNullContext)

	// EnterColumnExprWinFunctionTarget is called when entering the ColumnExprWinFunctionTarget production.
	EnterColumnExprWinFunctionTarget(c *ColumnExprWinFunctionTargetContext)

	// EnterColumnExprTrim is called when entering the ColumnExprTrim production.
	EnterColumnExprTrim(c *ColumnExprTrimContext)

	// EnterColumnExprTuple is called when entering the ColumnExprTuple production.
	EnterColumnExprTuple(c *ColumnExprTupleContext)

	// EnterColumnExprArrayAccess is called when entering the ColumnExprArrayAccess production.
	EnterColumnExprArrayAccess(c *ColumnExprArrayAccessContext)

	// EnterColumnExprBetween is called when entering the ColumnExprBetween production.
	EnterColumnExprBetween(c *ColumnExprBetweenContext)

	// EnterColumnExprParens is called when entering the ColumnExprParens production.
	EnterColumnExprParens(c *ColumnExprParensContext)

	// EnterColumnExprTimestamp is called when entering the ColumnExprTimestamp production.
	EnterColumnExprTimestamp(c *ColumnExprTimestampContext)

	// EnterColumnExprAnd is called when entering the ColumnExprAnd production.
	EnterColumnExprAnd(c *ColumnExprAndContext)

	// EnterColumnExprTupleAccess is called when entering the ColumnExprTupleAccess production.
	EnterColumnExprTupleAccess(c *ColumnExprTupleAccessContext)

	// EnterColumnExprCase is called when entering the ColumnExprCase production.
	EnterColumnExprCase(c *ColumnExprCaseContext)

	// EnterColumnExprDate is called when entering the ColumnExprDate production.
	EnterColumnExprDate(c *ColumnExprDateContext)

	// EnterColumnExprNot is called when entering the ColumnExprNot production.
	EnterColumnExprNot(c *ColumnExprNotContext)

	// EnterColumnExprWinFunction is called when entering the ColumnExprWinFunction production.
	EnterColumnExprWinFunction(c *ColumnExprWinFunctionContext)

	// EnterColumnExprIdentifier is called when entering the ColumnExprIdentifier production.
	EnterColumnExprIdentifier(c *ColumnExprIdentifierContext)

	// EnterColumnExprFunction is called when entering the ColumnExprFunction production.
	EnterColumnExprFunction(c *ColumnExprFunctionContext)

	// EnterColumnExprAsterisk is called when entering the ColumnExprAsterisk production.
	EnterColumnExprAsterisk(c *ColumnExprAsteriskContext)

	// EnterColumnArgList is called when entering the columnArgList production.
	EnterColumnArgList(c *ColumnArgListContext)

	// EnterColumnArgExpr is called when entering the columnArgExpr production.
	EnterColumnArgExpr(c *ColumnArgExprContext)

	// EnterColumnLambdaExpr is called when entering the columnLambdaExpr production.
	EnterColumnLambdaExpr(c *ColumnLambdaExprContext)

	// EnterColumnIdentifier is called when entering the columnIdentifier production.
	EnterColumnIdentifier(c *ColumnIdentifierContext)

	// EnterNestedIdentifier is called when entering the nestedIdentifier production.
	EnterNestedIdentifier(c *NestedIdentifierContext)

	// EnterEmptyListExpr is called when entering the emptyListExpr production.
	EnterEmptyListExpr(c *EmptyListExprContext)

	// EnterTableExprIdentifier is called when entering the TableExprIdentifier production.
	EnterTableExprIdentifier(c *TableExprIdentifierContext)

	// EnterTableExprSubquery is called when entering the TableExprSubquery production.
	EnterTableExprSubquery(c *TableExprSubqueryContext)

	// EnterTableExprAlias is called when entering the TableExprAlias production.
	EnterTableExprAlias(c *TableExprAliasContext)

	// EnterTableExprFunction is called when entering the TableExprFunction production.
	EnterTableExprFunction(c *TableExprFunctionContext)

	// EnterTableFunctionExpr is called when entering the tableFunctionExpr production.
	EnterTableFunctionExpr(c *TableFunctionExprContext)

	// EnterTableIdentifier is called when entering the tableIdentifier production.
	EnterTableIdentifier(c *TableIdentifierContext)

	// EnterTableArgList is called when entering the tableArgList production.
	EnterTableArgList(c *TableArgListContext)

	// EnterTableArgExpr is called when entering the tableArgExpr production.
	EnterTableArgExpr(c *TableArgExprContext)

	// EnterDatabaseIdentifier is called when entering the databaseIdentifier production.
	EnterDatabaseIdentifier(c *DatabaseIdentifierContext)

	// EnterFloatingLiteral is called when entering the floatingLiteral production.
	EnterFloatingLiteral(c *FloatingLiteralContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterKeywordForAlias is called when entering the keywordForAlias production.
	EnterKeywordForAlias(c *KeywordForAliasContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifierOrNull is called when entering the identifierOrNull production.
	EnterIdentifierOrNull(c *IdentifierOrNullContext)

	// EnterEnumValue is called when entering the enumValue production.
	EnterEnumValue(c *EnumValueContext)

	// ExitQueryStmt is called when exiting the queryStmt production.
	ExitQueryStmt(c *QueryStmtContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitCtes is called when exiting the ctes production.
	ExitCtes(c *CtesContext)

	// ExitNamedQuery is called when exiting the namedQuery production.
	ExitNamedQuery(c *NamedQueryContext)

	// ExitColumnAliases is called when exiting the columnAliases production.
	ExitColumnAliases(c *ColumnAliasesContext)

	// ExitAlterTableStmt is called when exiting the AlterTableStmt production.
	ExitAlterTableStmt(c *AlterTableStmtContext)

	// ExitAlterTableClauseAddColumn is called when exiting the AlterTableClauseAddColumn production.
	ExitAlterTableClauseAddColumn(c *AlterTableClauseAddColumnContext)

	// ExitAlterTableClauseAddIndex is called when exiting the AlterTableClauseAddIndex production.
	ExitAlterTableClauseAddIndex(c *AlterTableClauseAddIndexContext)

	// ExitAlterTableClauseAddProjection is called when exiting the AlterTableClauseAddProjection production.
	ExitAlterTableClauseAddProjection(c *AlterTableClauseAddProjectionContext)

	// ExitAlterTableClauseAttach is called when exiting the AlterTableClauseAttach production.
	ExitAlterTableClauseAttach(c *AlterTableClauseAttachContext)

	// ExitAlterTableClauseClearColumn is called when exiting the AlterTableClauseClearColumn production.
	ExitAlterTableClauseClearColumn(c *AlterTableClauseClearColumnContext)

	// ExitAlterTableClauseClearIndex is called when exiting the AlterTableClauseClearIndex production.
	ExitAlterTableClauseClearIndex(c *AlterTableClauseClearIndexContext)

	// ExitAlterTableClauseClearProjection is called when exiting the AlterTableClauseClearProjection production.
	ExitAlterTableClauseClearProjection(c *AlterTableClauseClearProjectionContext)

	// ExitAlterTableClauseComment is called when exiting the AlterTableClauseComment production.
	ExitAlterTableClauseComment(c *AlterTableClauseCommentContext)

	// ExitAlterTableClauseDelete is called when exiting the AlterTableClauseDelete production.
	ExitAlterTableClauseDelete(c *AlterTableClauseDeleteContext)

	// ExitAlterTableClauseDetach is called when exiting the AlterTableClauseDetach production.
	ExitAlterTableClauseDetach(c *AlterTableClauseDetachContext)

	// ExitAlterTableClauseDropColumn is called when exiting the AlterTableClauseDropColumn production.
	ExitAlterTableClauseDropColumn(c *AlterTableClauseDropColumnContext)

	// ExitAlterTableClauseDropIndex is called when exiting the AlterTableClauseDropIndex production.
	ExitAlterTableClauseDropIndex(c *AlterTableClauseDropIndexContext)

	// ExitAlterTableClauseDropProjection is called when exiting the AlterTableClauseDropProjection production.
	ExitAlterTableClauseDropProjection(c *AlterTableClauseDropProjectionContext)

	// ExitAlterTableClauseDropPartition is called when exiting the AlterTableClauseDropPartition production.
	ExitAlterTableClauseDropPartition(c *AlterTableClauseDropPartitionContext)

	// ExitAlterTableClauseFreezePartition is called when exiting the AlterTableClauseFreezePartition production.
	ExitAlterTableClauseFreezePartition(c *AlterTableClauseFreezePartitionContext)

	// ExitAlterTableClauseMaterializeIndex is called when exiting the AlterTableClauseMaterializeIndex production.
	ExitAlterTableClauseMaterializeIndex(c *AlterTableClauseMaterializeIndexContext)

	// ExitAlterTableClauseMaterializeProjection is called when exiting the AlterTableClauseMaterializeProjection production.
	ExitAlterTableClauseMaterializeProjection(c *AlterTableClauseMaterializeProjectionContext)

	// ExitAlterTableClauseModifyCodec is called when exiting the AlterTableClauseModifyCodec production.
	ExitAlterTableClauseModifyCodec(c *AlterTableClauseModifyCodecContext)

	// ExitAlterTableClauseModifyComment is called when exiting the AlterTableClauseModifyComment production.
	ExitAlterTableClauseModifyComment(c *AlterTableClauseModifyCommentContext)

	// ExitAlterTableClauseModifyRemove is called when exiting the AlterTableClauseModifyRemove production.
	ExitAlterTableClauseModifyRemove(c *AlterTableClauseModifyRemoveContext)

	// ExitAlterTableClauseModify is called when exiting the AlterTableClauseModify production.
	ExitAlterTableClauseModify(c *AlterTableClauseModifyContext)

	// ExitAlterTableClauseModifyOrderBy is called when exiting the AlterTableClauseModifyOrderBy production.
	ExitAlterTableClauseModifyOrderBy(c *AlterTableClauseModifyOrderByContext)

	// ExitAlterTableClauseModifyTTL is called when exiting the AlterTableClauseModifyTTL production.
	ExitAlterTableClauseModifyTTL(c *AlterTableClauseModifyTTLContext)

	// ExitAlterTableClauseMovePartition is called when exiting the AlterTableClauseMovePartition production.
	ExitAlterTableClauseMovePartition(c *AlterTableClauseMovePartitionContext)

	// ExitAlterTableClauseRemoveTTL is called when exiting the AlterTableClauseRemoveTTL production.
	ExitAlterTableClauseRemoveTTL(c *AlterTableClauseRemoveTTLContext)

	// ExitAlterTableClauseRename is called when exiting the AlterTableClauseRename production.
	ExitAlterTableClauseRename(c *AlterTableClauseRenameContext)

	// ExitAlterTableClauseReplace is called when exiting the AlterTableClauseReplace production.
	ExitAlterTableClauseReplace(c *AlterTableClauseReplaceContext)

	// ExitAlterTableClauseUpdate is called when exiting the AlterTableClauseUpdate production.
	ExitAlterTableClauseUpdate(c *AlterTableClauseUpdateContext)

	// ExitAssignmentExprList is called when exiting the assignmentExprList production.
	ExitAssignmentExprList(c *AssignmentExprListContext)

	// ExitAssignmentExpr is called when exiting the assignmentExpr production.
	ExitAssignmentExpr(c *AssignmentExprContext)

	// ExitTableColumnPropertyType is called when exiting the tableColumnPropertyType production.
	ExitTableColumnPropertyType(c *TableColumnPropertyTypeContext)

	// ExitPartitionClause is called when exiting the partitionClause production.
	ExitPartitionClause(c *PartitionClauseContext)

	// ExitAttachDictionaryStmt is called when exiting the AttachDictionaryStmt production.
	ExitAttachDictionaryStmt(c *AttachDictionaryStmtContext)

	// ExitCheckStmt is called when exiting the checkStmt production.
	ExitCheckStmt(c *CheckStmtContext)

	// ExitCreateDatabaseStmt is called when exiting the CreateDatabaseStmt production.
	ExitCreateDatabaseStmt(c *CreateDatabaseStmtContext)

	// ExitCreateDictionaryStmt is called when exiting the CreateDictionaryStmt production.
	ExitCreateDictionaryStmt(c *CreateDictionaryStmtContext)

	// ExitCreateLiveViewStmt is called when exiting the CreateLiveViewStmt production.
	ExitCreateLiveViewStmt(c *CreateLiveViewStmtContext)

	// ExitCreateMaterializedViewStmt is called when exiting the CreateMaterializedViewStmt production.
	ExitCreateMaterializedViewStmt(c *CreateMaterializedViewStmtContext)

	// ExitCreateTableStmt is called when exiting the CreateTableStmt production.
	ExitCreateTableStmt(c *CreateTableStmtContext)

	// ExitCreateViewStmt is called when exiting the CreateViewStmt production.
	ExitCreateViewStmt(c *CreateViewStmtContext)

	// ExitDictionarySchemaClause is called when exiting the dictionarySchemaClause production.
	ExitDictionarySchemaClause(c *DictionarySchemaClauseContext)

	// ExitDictionaryAttrDfnt is called when exiting the dictionaryAttrDfnt production.
	ExitDictionaryAttrDfnt(c *DictionaryAttrDfntContext)

	// ExitDictionaryEngineClause is called when exiting the dictionaryEngineClause production.
	ExitDictionaryEngineClause(c *DictionaryEngineClauseContext)

	// ExitDictionaryPrimaryKeyClause is called when exiting the dictionaryPrimaryKeyClause production.
	ExitDictionaryPrimaryKeyClause(c *DictionaryPrimaryKeyClauseContext)

	// ExitDictionaryArgExpr is called when exiting the dictionaryArgExpr production.
	ExitDictionaryArgExpr(c *DictionaryArgExprContext)

	// ExitSourceClause is called when exiting the sourceClause production.
	ExitSourceClause(c *SourceClauseContext)

	// ExitLifetimeClause is called when exiting the lifetimeClause production.
	ExitLifetimeClause(c *LifetimeClauseContext)

	// ExitLayoutClause is called when exiting the layoutClause production.
	ExitLayoutClause(c *LayoutClauseContext)

	// ExitRangeClause is called when exiting the rangeClause production.
	ExitRangeClause(c *RangeClauseContext)

	// ExitDictionarySettingsClause is called when exiting the dictionarySettingsClause production.
	ExitDictionarySettingsClause(c *DictionarySettingsClauseContext)

	// ExitClusterClause is called when exiting the clusterClause production.
	ExitClusterClause(c *ClusterClauseContext)

	// ExitUuidClause is called when exiting the uuidClause production.
	ExitUuidClause(c *UuidClauseContext)

	// ExitDestinationClause is called when exiting the destinationClause production.
	ExitDestinationClause(c *DestinationClauseContext)

	// ExitSubqueryClause is called when exiting the subqueryClause production.
	ExitSubqueryClause(c *SubqueryClauseContext)

	// ExitSchemaDescriptionClause is called when exiting the SchemaDescriptionClause production.
	ExitSchemaDescriptionClause(c *SchemaDescriptionClauseContext)

	// ExitSchemaAsTableClause is called when exiting the SchemaAsTableClause production.
	ExitSchemaAsTableClause(c *SchemaAsTableClauseContext)

	// ExitSchemaAsFunctionClause is called when exiting the SchemaAsFunctionClause production.
	ExitSchemaAsFunctionClause(c *SchemaAsFunctionClauseContext)

	// ExitEngineClause is called when exiting the engineClause production.
	ExitEngineClause(c *EngineClauseContext)

	// ExitPartitionByClause is called when exiting the partitionByClause production.
	ExitPartitionByClause(c *PartitionByClauseContext)

	// ExitPrimaryKeyClause is called when exiting the primaryKeyClause production.
	ExitPrimaryKeyClause(c *PrimaryKeyClauseContext)

	// ExitSampleByClause is called when exiting the sampleByClause production.
	ExitSampleByClause(c *SampleByClauseContext)

	// ExitTtlClause is called when exiting the ttlClause production.
	ExitTtlClause(c *TtlClauseContext)

	// ExitEngineExpr is called when exiting the engineExpr production.
	ExitEngineExpr(c *EngineExprContext)

	// ExitTableElementExprColumn is called when exiting the TableElementExprColumn production.
	ExitTableElementExprColumn(c *TableElementExprColumnContext)

	// ExitTableElementExprConstraint is called when exiting the TableElementExprConstraint production.
	ExitTableElementExprConstraint(c *TableElementExprConstraintContext)

	// ExitTableElementExprIndex is called when exiting the TableElementExprIndex production.
	ExitTableElementExprIndex(c *TableElementExprIndexContext)

	// ExitTableElementExprProjection is called when exiting the TableElementExprProjection production.
	ExitTableElementExprProjection(c *TableElementExprProjectionContext)

	// ExitTableColumnDfnt is called when exiting the tableColumnDfnt production.
	ExitTableColumnDfnt(c *TableColumnDfntContext)

	// ExitTableColumnPropertyExpr is called when exiting the tableColumnPropertyExpr production.
	ExitTableColumnPropertyExpr(c *TableColumnPropertyExprContext)

	// ExitTableIndexDfnt is called when exiting the tableIndexDfnt production.
	ExitTableIndexDfnt(c *TableIndexDfntContext)

	// ExitTableProjectionDfnt is called when exiting the tableProjectionDfnt production.
	ExitTableProjectionDfnt(c *TableProjectionDfntContext)

	// ExitCodecExpr is called when exiting the codecExpr production.
	ExitCodecExpr(c *CodecExprContext)

	// ExitCodecArgExpr is called when exiting the codecArgExpr production.
	ExitCodecArgExpr(c *CodecArgExprContext)

	// ExitTtlExpr is called when exiting the ttlExpr production.
	ExitTtlExpr(c *TtlExprContext)

	// ExitDescribeStmt is called when exiting the describeStmt production.
	ExitDescribeStmt(c *DescribeStmtContext)

	// ExitDropDatabaseStmt is called when exiting the DropDatabaseStmt production.
	ExitDropDatabaseStmt(c *DropDatabaseStmtContext)

	// ExitDropTableStmt is called when exiting the DropTableStmt production.
	ExitDropTableStmt(c *DropTableStmtContext)

	// ExitDeleteStmt is called when exiting the deleteStmt production.
	ExitDeleteStmt(c *DeleteStmtContext)

	// ExitExistsDatabaseStmt is called when exiting the ExistsDatabaseStmt production.
	ExitExistsDatabaseStmt(c *ExistsDatabaseStmtContext)

	// ExitExistsTableStmt is called when exiting the ExistsTableStmt production.
	ExitExistsTableStmt(c *ExistsTableStmtContext)

	// ExitExplainASTStmt is called when exiting the ExplainASTStmt production.
	ExitExplainASTStmt(c *ExplainASTStmtContext)

	// ExitExplainSyntaxStmt is called when exiting the ExplainSyntaxStmt production.
	ExitExplainSyntaxStmt(c *ExplainSyntaxStmtContext)

	// ExitInsertStmt is called when exiting the insertStmt production.
	ExitInsertStmt(c *InsertStmtContext)

	// ExitColumnsClause is called when exiting the columnsClause production.
	ExitColumnsClause(c *ColumnsClauseContext)

	// ExitDataClauseFormat is called when exiting the DataClauseFormat production.
	ExitDataClauseFormat(c *DataClauseFormatContext)

	// ExitDataClauseValues is called when exiting the DataClauseValues production.
	ExitDataClauseValues(c *DataClauseValuesContext)

	// ExitDataClauseSelect is called when exiting the DataClauseSelect production.
	ExitDataClauseSelect(c *DataClauseSelectContext)

	// ExitKillMutationStmt is called when exiting the KillMutationStmt production.
	ExitKillMutationStmt(c *KillMutationStmtContext)

	// ExitOptimizeStmt is called when exiting the optimizeStmt production.
	ExitOptimizeStmt(c *OptimizeStmtContext)

	// ExitRenameStmt is called when exiting the renameStmt production.
	ExitRenameStmt(c *RenameStmtContext)

	// ExitProjectionSelectStmt is called when exiting the projectionSelectStmt production.
	ExitProjectionSelectStmt(c *ProjectionSelectStmtContext)

	// ExitSelectUnionStmt is called when exiting the selectUnionStmt production.
	ExitSelectUnionStmt(c *SelectUnionStmtContext)

	// ExitSelectStmtWithParens is called when exiting the selectStmtWithParens production.
	ExitSelectStmtWithParens(c *SelectStmtWithParensContext)

	// ExitSelectStmt is called when exiting the selectStmt production.
	ExitSelectStmt(c *SelectStmtContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitTopClause is called when exiting the topClause production.
	ExitTopClause(c *TopClauseContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitArrayJoinClause is called when exiting the arrayJoinClause production.
	ExitArrayJoinClause(c *ArrayJoinClauseContext)

	// ExitWindowClause is called when exiting the windowClause production.
	ExitWindowClause(c *WindowClauseContext)

	// ExitPrewhereClause is called when exiting the prewhereClause production.
	ExitPrewhereClause(c *PrewhereClauseContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitProjectionOrderByClause is called when exiting the projectionOrderByClause production.
	ExitProjectionOrderByClause(c *ProjectionOrderByClauseContext)

	// ExitLimitByClause is called when exiting the limitByClause production.
	ExitLimitByClause(c *LimitByClauseContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitSettingsClause is called when exiting the settingsClause production.
	ExitSettingsClause(c *SettingsClauseContext)

	// ExitJoinExprOp is called when exiting the JoinExprOp production.
	ExitJoinExprOp(c *JoinExprOpContext)

	// ExitJoinExprTable is called when exiting the JoinExprTable production.
	ExitJoinExprTable(c *JoinExprTableContext)

	// ExitJoinExprParens is called when exiting the JoinExprParens production.
	ExitJoinExprParens(c *JoinExprParensContext)

	// ExitJoinExprCrossOp is called when exiting the JoinExprCrossOp production.
	ExitJoinExprCrossOp(c *JoinExprCrossOpContext)

	// ExitJoinOpInner is called when exiting the JoinOpInner production.
	ExitJoinOpInner(c *JoinOpInnerContext)

	// ExitJoinOpLeftRight is called when exiting the JoinOpLeftRight production.
	ExitJoinOpLeftRight(c *JoinOpLeftRightContext)

	// ExitJoinOpFull is called when exiting the JoinOpFull production.
	ExitJoinOpFull(c *JoinOpFullContext)

	// ExitJoinOpCross is called when exiting the joinOpCross production.
	ExitJoinOpCross(c *JoinOpCrossContext)

	// ExitJoinConstraintClause is called when exiting the joinConstraintClause production.
	ExitJoinConstraintClause(c *JoinConstraintClauseContext)

	// ExitSampleClause is called when exiting the sampleClause production.
	ExitSampleClause(c *SampleClauseContext)

	// ExitLimitExpr is called when exiting the limitExpr production.
	ExitLimitExpr(c *LimitExprContext)

	// ExitOrderExprList is called when exiting the orderExprList production.
	ExitOrderExprList(c *OrderExprListContext)

	// ExitOrderExpr is called when exiting the orderExpr production.
	ExitOrderExpr(c *OrderExprContext)

	// ExitRatioExpr is called when exiting the ratioExpr production.
	ExitRatioExpr(c *RatioExprContext)

	// ExitSettingExprList is called when exiting the settingExprList production.
	ExitSettingExprList(c *SettingExprListContext)

	// ExitSettingExpr is called when exiting the settingExpr production.
	ExitSettingExpr(c *SettingExprContext)

	// ExitWindowExpr is called when exiting the windowExpr production.
	ExitWindowExpr(c *WindowExprContext)

	// ExitWinPartitionByClause is called when exiting the winPartitionByClause production.
	ExitWinPartitionByClause(c *WinPartitionByClauseContext)

	// ExitWinOrderByClause is called when exiting the winOrderByClause production.
	ExitWinOrderByClause(c *WinOrderByClauseContext)

	// ExitWinFrameClause is called when exiting the winFrameClause production.
	ExitWinFrameClause(c *WinFrameClauseContext)

	// ExitFrameStart is called when exiting the frameStart production.
	ExitFrameStart(c *FrameStartContext)

	// ExitFrameBetween is called when exiting the frameBetween production.
	ExitFrameBetween(c *FrameBetweenContext)

	// ExitWinFrameBound is called when exiting the winFrameBound production.
	ExitWinFrameBound(c *WinFrameBoundContext)

	// ExitSetStmt is called when exiting the setStmt production.
	ExitSetStmt(c *SetStmtContext)

	// ExitShowCreateDatabaseStmt is called when exiting the showCreateDatabaseStmt production.
	ExitShowCreateDatabaseStmt(c *ShowCreateDatabaseStmtContext)

	// ExitShowCreateDictionaryStmt is called when exiting the showCreateDictionaryStmt production.
	ExitShowCreateDictionaryStmt(c *ShowCreateDictionaryStmtContext)

	// ExitShowCreateTableStmt is called when exiting the showCreateTableStmt production.
	ExitShowCreateTableStmt(c *ShowCreateTableStmtContext)

	// ExitShowDatabasesStmt is called when exiting the showDatabasesStmt production.
	ExitShowDatabasesStmt(c *ShowDatabasesStmtContext)

	// ExitShowDictionariesStmt is called when exiting the showDictionariesStmt production.
	ExitShowDictionariesStmt(c *ShowDictionariesStmtContext)

	// ExitShowTablesStmt is called when exiting the showTablesStmt production.
	ExitShowTablesStmt(c *ShowTablesStmtContext)

	// ExitSystemStmt is called when exiting the systemStmt production.
	ExitSystemStmt(c *SystemStmtContext)

	// ExitTruncateStmt is called when exiting the truncateStmt production.
	ExitTruncateStmt(c *TruncateStmtContext)

	// ExitUseStmt is called when exiting the useStmt production.
	ExitUseStmt(c *UseStmtContext)

	// ExitWatchStmt is called when exiting the watchStmt production.
	ExitWatchStmt(c *WatchStmtContext)

	// ExitColumnTypeExprSimple is called when exiting the ColumnTypeExprSimple production.
	ExitColumnTypeExprSimple(c *ColumnTypeExprSimpleContext)

	// ExitColumnTypeExprNested is called when exiting the ColumnTypeExprNested production.
	ExitColumnTypeExprNested(c *ColumnTypeExprNestedContext)

	// ExitColumnTypeExprEnum is called when exiting the ColumnTypeExprEnum production.
	ExitColumnTypeExprEnum(c *ColumnTypeExprEnumContext)

	// ExitColumnTypeExprComplex is called when exiting the ColumnTypeExprComplex production.
	ExitColumnTypeExprComplex(c *ColumnTypeExprComplexContext)

	// ExitColumnTypeExprParam is called when exiting the ColumnTypeExprParam production.
	ExitColumnTypeExprParam(c *ColumnTypeExprParamContext)

	// ExitColumnExprList is called when exiting the columnExprList production.
	ExitColumnExprList(c *ColumnExprListContext)

	// ExitColumnsExprAsterisk is called when exiting the ColumnsExprAsterisk production.
	ExitColumnsExprAsterisk(c *ColumnsExprAsteriskContext)

	// ExitColumnsExprSubquery is called when exiting the ColumnsExprSubquery production.
	ExitColumnsExprSubquery(c *ColumnsExprSubqueryContext)

	// ExitColumnsExprColumn is called when exiting the ColumnsExprColumn production.
	ExitColumnsExprColumn(c *ColumnsExprColumnContext)

	// ExitColumnExprTernaryOp is called when exiting the ColumnExprTernaryOp production.
	ExitColumnExprTernaryOp(c *ColumnExprTernaryOpContext)

	// ExitColumnExprAlias is called when exiting the ColumnExprAlias production.
	ExitColumnExprAlias(c *ColumnExprAliasContext)

	// ExitColumnExprExtract is called when exiting the ColumnExprExtract production.
	ExitColumnExprExtract(c *ColumnExprExtractContext)

	// ExitColumnExprNegate is called when exiting the ColumnExprNegate production.
	ExitColumnExprNegate(c *ColumnExprNegateContext)

	// ExitColumnExprSubquery is called when exiting the ColumnExprSubquery production.
	ExitColumnExprSubquery(c *ColumnExprSubqueryContext)

	// ExitColumnExprLiteral is called when exiting the ColumnExprLiteral production.
	ExitColumnExprLiteral(c *ColumnExprLiteralContext)

	// ExitColumnExprArray is called when exiting the ColumnExprArray production.
	ExitColumnExprArray(c *ColumnExprArrayContext)

	// ExitColumnExprSubstring is called when exiting the ColumnExprSubstring production.
	ExitColumnExprSubstring(c *ColumnExprSubstringContext)

	// ExitColumnExprCast is called when exiting the ColumnExprCast production.
	ExitColumnExprCast(c *ColumnExprCastContext)

	// ExitColumnExprOr is called when exiting the ColumnExprOr production.
	ExitColumnExprOr(c *ColumnExprOrContext)

	// ExitColumnExprPrecedence1 is called when exiting the ColumnExprPrecedence1 production.
	ExitColumnExprPrecedence1(c *ColumnExprPrecedence1Context)

	// ExitColumnExprPrecedence2 is called when exiting the ColumnExprPrecedence2 production.
	ExitColumnExprPrecedence2(c *ColumnExprPrecedence2Context)

	// ExitColumnExprPrecedence3 is called when exiting the ColumnExprPrecedence3 production.
	ExitColumnExprPrecedence3(c *ColumnExprPrecedence3Context)

	// ExitColumnExprInterval is called when exiting the ColumnExprInterval production.
	ExitColumnExprInterval(c *ColumnExprIntervalContext)

	// ExitColumnExprIsNull is called when exiting the ColumnExprIsNull production.
	ExitColumnExprIsNull(c *ColumnExprIsNullContext)

	// ExitColumnExprWinFunctionTarget is called when exiting the ColumnExprWinFunctionTarget production.
	ExitColumnExprWinFunctionTarget(c *ColumnExprWinFunctionTargetContext)

	// ExitColumnExprTrim is called when exiting the ColumnExprTrim production.
	ExitColumnExprTrim(c *ColumnExprTrimContext)

	// ExitColumnExprTuple is called when exiting the ColumnExprTuple production.
	ExitColumnExprTuple(c *ColumnExprTupleContext)

	// ExitColumnExprArrayAccess is called when exiting the ColumnExprArrayAccess production.
	ExitColumnExprArrayAccess(c *ColumnExprArrayAccessContext)

	// ExitColumnExprBetween is called when exiting the ColumnExprBetween production.
	ExitColumnExprBetween(c *ColumnExprBetweenContext)

	// ExitColumnExprParens is called when exiting the ColumnExprParens production.
	ExitColumnExprParens(c *ColumnExprParensContext)

	// ExitColumnExprTimestamp is called when exiting the ColumnExprTimestamp production.
	ExitColumnExprTimestamp(c *ColumnExprTimestampContext)

	// ExitColumnExprAnd is called when exiting the ColumnExprAnd production.
	ExitColumnExprAnd(c *ColumnExprAndContext)

	// ExitColumnExprTupleAccess is called when exiting the ColumnExprTupleAccess production.
	ExitColumnExprTupleAccess(c *ColumnExprTupleAccessContext)

	// ExitColumnExprCase is called when exiting the ColumnExprCase production.
	ExitColumnExprCase(c *ColumnExprCaseContext)

	// ExitColumnExprDate is called when exiting the ColumnExprDate production.
	ExitColumnExprDate(c *ColumnExprDateContext)

	// ExitColumnExprNot is called when exiting the ColumnExprNot production.
	ExitColumnExprNot(c *ColumnExprNotContext)

	// ExitColumnExprWinFunction is called when exiting the ColumnExprWinFunction production.
	ExitColumnExprWinFunction(c *ColumnExprWinFunctionContext)

	// ExitColumnExprIdentifier is called when exiting the ColumnExprIdentifier production.
	ExitColumnExprIdentifier(c *ColumnExprIdentifierContext)

	// ExitColumnExprFunction is called when exiting the ColumnExprFunction production.
	ExitColumnExprFunction(c *ColumnExprFunctionContext)

	// ExitColumnExprAsterisk is called when exiting the ColumnExprAsterisk production.
	ExitColumnExprAsterisk(c *ColumnExprAsteriskContext)

	// ExitColumnArgList is called when exiting the columnArgList production.
	ExitColumnArgList(c *ColumnArgListContext)

	// ExitColumnArgExpr is called when exiting the columnArgExpr production.
	ExitColumnArgExpr(c *ColumnArgExprContext)

	// ExitColumnLambdaExpr is called when exiting the columnLambdaExpr production.
	ExitColumnLambdaExpr(c *ColumnLambdaExprContext)

	// ExitColumnIdentifier is called when exiting the columnIdentifier production.
	ExitColumnIdentifier(c *ColumnIdentifierContext)

	// ExitNestedIdentifier is called when exiting the nestedIdentifier production.
	ExitNestedIdentifier(c *NestedIdentifierContext)

	// ExitEmptyListExpr is called when exiting the emptyListExpr production.
	ExitEmptyListExpr(c *EmptyListExprContext)

	// ExitTableExprIdentifier is called when exiting the TableExprIdentifier production.
	ExitTableExprIdentifier(c *TableExprIdentifierContext)

	// ExitTableExprSubquery is called when exiting the TableExprSubquery production.
	ExitTableExprSubquery(c *TableExprSubqueryContext)

	// ExitTableExprAlias is called when exiting the TableExprAlias production.
	ExitTableExprAlias(c *TableExprAliasContext)

	// ExitTableExprFunction is called when exiting the TableExprFunction production.
	ExitTableExprFunction(c *TableExprFunctionContext)

	// ExitTableFunctionExpr is called when exiting the tableFunctionExpr production.
	ExitTableFunctionExpr(c *TableFunctionExprContext)

	// ExitTableIdentifier is called when exiting the tableIdentifier production.
	ExitTableIdentifier(c *TableIdentifierContext)

	// ExitTableArgList is called when exiting the tableArgList production.
	ExitTableArgList(c *TableArgListContext)

	// ExitTableArgExpr is called when exiting the tableArgExpr production.
	ExitTableArgExpr(c *TableArgExprContext)

	// ExitDatabaseIdentifier is called when exiting the databaseIdentifier production.
	ExitDatabaseIdentifier(c *DatabaseIdentifierContext)

	// ExitFloatingLiteral is called when exiting the floatingLiteral production.
	ExitFloatingLiteral(c *FloatingLiteralContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitKeywordForAlias is called when exiting the keywordForAlias production.
	ExitKeywordForAlias(c *KeywordForAliasContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifierOrNull is called when exiting the identifierOrNull production.
	ExitIdentifierOrNull(c *IdentifierOrNullContext)

	// ExitEnumValue is called when exiting the enumValue production.
	ExitEnumValue(c *EnumValueContext)
}
