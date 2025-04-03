// Code generated from ClickHouseParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package clickhouse // ClickHouseParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ClickHouseParser.
type ClickHouseParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ClickHouseParser#queryStmt.
	VisitQueryStmt(ctx *QueryStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ctes.
	VisitCtes(ctx *CtesContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#namedQuery.
	VisitNamedQuery(ctx *NamedQueryContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#columnAliases.
	VisitColumnAliases(ctx *ColumnAliasesContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableStmt.
	VisitAlterTableStmt(ctx *AlterTableStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseAddColumn.
	VisitAlterTableClauseAddColumn(ctx *AlterTableClauseAddColumnContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseAddIndex.
	VisitAlterTableClauseAddIndex(ctx *AlterTableClauseAddIndexContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseAddProjection.
	VisitAlterTableClauseAddProjection(ctx *AlterTableClauseAddProjectionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseAttach.
	VisitAlterTableClauseAttach(ctx *AlterTableClauseAttachContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseClearColumn.
	VisitAlterTableClauseClearColumn(ctx *AlterTableClauseClearColumnContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseClearIndex.
	VisitAlterTableClauseClearIndex(ctx *AlterTableClauseClearIndexContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseClearProjection.
	VisitAlterTableClauseClearProjection(ctx *AlterTableClauseClearProjectionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseComment.
	VisitAlterTableClauseComment(ctx *AlterTableClauseCommentContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseDelete.
	VisitAlterTableClauseDelete(ctx *AlterTableClauseDeleteContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseDetach.
	VisitAlterTableClauseDetach(ctx *AlterTableClauseDetachContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseDropColumn.
	VisitAlterTableClauseDropColumn(ctx *AlterTableClauseDropColumnContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseDropIndex.
	VisitAlterTableClauseDropIndex(ctx *AlterTableClauseDropIndexContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseDropProjection.
	VisitAlterTableClauseDropProjection(ctx *AlterTableClauseDropProjectionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseDropPartition.
	VisitAlterTableClauseDropPartition(ctx *AlterTableClauseDropPartitionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseFreezePartition.
	VisitAlterTableClauseFreezePartition(ctx *AlterTableClauseFreezePartitionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseMaterializeIndex.
	VisitAlterTableClauseMaterializeIndex(ctx *AlterTableClauseMaterializeIndexContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseMaterializeProjection.
	VisitAlterTableClauseMaterializeProjection(ctx *AlterTableClauseMaterializeProjectionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseModifyCodec.
	VisitAlterTableClauseModifyCodec(ctx *AlterTableClauseModifyCodecContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseModifyComment.
	VisitAlterTableClauseModifyComment(ctx *AlterTableClauseModifyCommentContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseModifyRemove.
	VisitAlterTableClauseModifyRemove(ctx *AlterTableClauseModifyRemoveContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseModify.
	VisitAlterTableClauseModify(ctx *AlterTableClauseModifyContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseModifyOrderBy.
	VisitAlterTableClauseModifyOrderBy(ctx *AlterTableClauseModifyOrderByContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseModifyTTL.
	VisitAlterTableClauseModifyTTL(ctx *AlterTableClauseModifyTTLContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseMovePartition.
	VisitAlterTableClauseMovePartition(ctx *AlterTableClauseMovePartitionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseRemoveTTL.
	VisitAlterTableClauseRemoveTTL(ctx *AlterTableClauseRemoveTTLContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseRename.
	VisitAlterTableClauseRename(ctx *AlterTableClauseRenameContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseReplace.
	VisitAlterTableClauseReplace(ctx *AlterTableClauseReplaceContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AlterTableClauseUpdate.
	VisitAlterTableClauseUpdate(ctx *AlterTableClauseUpdateContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#assignmentExprList.
	VisitAssignmentExprList(ctx *AssignmentExprListContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#assignmentExpr.
	VisitAssignmentExpr(ctx *AssignmentExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#tableColumnPropertyType.
	VisitTableColumnPropertyType(ctx *TableColumnPropertyTypeContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#partitionClause.
	VisitPartitionClause(ctx *PartitionClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#AttachDictionaryStmt.
	VisitAttachDictionaryStmt(ctx *AttachDictionaryStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#checkStmt.
	VisitCheckStmt(ctx *CheckStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#CreateDatabaseStmt.
	VisitCreateDatabaseStmt(ctx *CreateDatabaseStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#CreateDictionaryStmt.
	VisitCreateDictionaryStmt(ctx *CreateDictionaryStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#CreateLiveViewStmt.
	VisitCreateLiveViewStmt(ctx *CreateLiveViewStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#CreateMaterializedViewStmt.
	VisitCreateMaterializedViewStmt(ctx *CreateMaterializedViewStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#CreateTableStmt.
	VisitCreateTableStmt(ctx *CreateTableStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#CreateViewStmt.
	VisitCreateViewStmt(ctx *CreateViewStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#dictionarySchemaClause.
	VisitDictionarySchemaClause(ctx *DictionarySchemaClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#dictionaryAttrDfnt.
	VisitDictionaryAttrDfnt(ctx *DictionaryAttrDfntContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#dictionaryEngineClause.
	VisitDictionaryEngineClause(ctx *DictionaryEngineClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#dictionaryPrimaryKeyClause.
	VisitDictionaryPrimaryKeyClause(ctx *DictionaryPrimaryKeyClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#dictionaryArgExpr.
	VisitDictionaryArgExpr(ctx *DictionaryArgExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#sourceClause.
	VisitSourceClause(ctx *SourceClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#lifetimeClause.
	VisitLifetimeClause(ctx *LifetimeClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#layoutClause.
	VisitLayoutClause(ctx *LayoutClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#rangeClause.
	VisitRangeClause(ctx *RangeClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#dictionarySettingsClause.
	VisitDictionarySettingsClause(ctx *DictionarySettingsClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#clusterClause.
	VisitClusterClause(ctx *ClusterClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#uuidClause.
	VisitUuidClause(ctx *UuidClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#destinationClause.
	VisitDestinationClause(ctx *DestinationClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#subqueryClause.
	VisitSubqueryClause(ctx *SubqueryClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#SchemaDescriptionClause.
	VisitSchemaDescriptionClause(ctx *SchemaDescriptionClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#SchemaAsTableClause.
	VisitSchemaAsTableClause(ctx *SchemaAsTableClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#SchemaAsFunctionClause.
	VisitSchemaAsFunctionClause(ctx *SchemaAsFunctionClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#engineClause.
	VisitEngineClause(ctx *EngineClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#partitionByClause.
	VisitPartitionByClause(ctx *PartitionByClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#primaryKeyClause.
	VisitPrimaryKeyClause(ctx *PrimaryKeyClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#sampleByClause.
	VisitSampleByClause(ctx *SampleByClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ttlClause.
	VisitTtlClause(ctx *TtlClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#engineExpr.
	VisitEngineExpr(ctx *EngineExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#TableElementExprColumn.
	VisitTableElementExprColumn(ctx *TableElementExprColumnContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#TableElementExprConstraint.
	VisitTableElementExprConstraint(ctx *TableElementExprConstraintContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#TableElementExprIndex.
	VisitTableElementExprIndex(ctx *TableElementExprIndexContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#TableElementExprProjection.
	VisitTableElementExprProjection(ctx *TableElementExprProjectionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#tableColumnDfnt.
	VisitTableColumnDfnt(ctx *TableColumnDfntContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#tableColumnPropertyExpr.
	VisitTableColumnPropertyExpr(ctx *TableColumnPropertyExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#tableIndexDfnt.
	VisitTableIndexDfnt(ctx *TableIndexDfntContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#tableProjectionDfnt.
	VisitTableProjectionDfnt(ctx *TableProjectionDfntContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#codecExpr.
	VisitCodecExpr(ctx *CodecExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#codecArgExpr.
	VisitCodecArgExpr(ctx *CodecArgExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ttlExpr.
	VisitTtlExpr(ctx *TtlExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#describeStmt.
	VisitDescribeStmt(ctx *DescribeStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#DropDatabaseStmt.
	VisitDropDatabaseStmt(ctx *DropDatabaseStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#DropTableStmt.
	VisitDropTableStmt(ctx *DropTableStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#deleteStmt.
	VisitDeleteStmt(ctx *DeleteStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ExistsDatabaseStmt.
	VisitExistsDatabaseStmt(ctx *ExistsDatabaseStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ExistsTableStmt.
	VisitExistsTableStmt(ctx *ExistsTableStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ExplainASTStmt.
	VisitExplainASTStmt(ctx *ExplainASTStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ExplainSyntaxStmt.
	VisitExplainSyntaxStmt(ctx *ExplainSyntaxStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#insertStmt.
	VisitInsertStmt(ctx *InsertStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#columnsClause.
	VisitColumnsClause(ctx *ColumnsClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#DataClauseFormat.
	VisitDataClauseFormat(ctx *DataClauseFormatContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#DataClauseValues.
	VisitDataClauseValues(ctx *DataClauseValuesContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#DataClauseSelect.
	VisitDataClauseSelect(ctx *DataClauseSelectContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#KillMutationStmt.
	VisitKillMutationStmt(ctx *KillMutationStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#optimizeStmt.
	VisitOptimizeStmt(ctx *OptimizeStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#renameStmt.
	VisitRenameStmt(ctx *RenameStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#projectionSelectStmt.
	VisitProjectionSelectStmt(ctx *ProjectionSelectStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#selectUnionStmt.
	VisitSelectUnionStmt(ctx *SelectUnionStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#selectStmtWithParens.
	VisitSelectStmtWithParens(ctx *SelectStmtWithParensContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#selectStmt.
	VisitSelectStmt(ctx *SelectStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#topClause.
	VisitTopClause(ctx *TopClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#arrayJoinClause.
	VisitArrayJoinClause(ctx *ArrayJoinClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#windowClause.
	VisitWindowClause(ctx *WindowClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#prewhereClause.
	VisitPrewhereClause(ctx *PrewhereClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#projectionOrderByClause.
	VisitProjectionOrderByClause(ctx *ProjectionOrderByClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#limitByClause.
	VisitLimitByClause(ctx *LimitByClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#settingsClause.
	VisitSettingsClause(ctx *SettingsClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#JoinExprOp.
	VisitJoinExprOp(ctx *JoinExprOpContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#JoinExprTable.
	VisitJoinExprTable(ctx *JoinExprTableContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#JoinExprParens.
	VisitJoinExprParens(ctx *JoinExprParensContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#JoinExprCrossOp.
	VisitJoinExprCrossOp(ctx *JoinExprCrossOpContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#JoinOpInner.
	VisitJoinOpInner(ctx *JoinOpInnerContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#JoinOpLeftRight.
	VisitJoinOpLeftRight(ctx *JoinOpLeftRightContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#JoinOpFull.
	VisitJoinOpFull(ctx *JoinOpFullContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#joinOpCross.
	VisitJoinOpCross(ctx *JoinOpCrossContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#joinConstraintClause.
	VisitJoinConstraintClause(ctx *JoinConstraintClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#sampleClause.
	VisitSampleClause(ctx *SampleClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#limitExpr.
	VisitLimitExpr(ctx *LimitExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#orderExprList.
	VisitOrderExprList(ctx *OrderExprListContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#orderExpr.
	VisitOrderExpr(ctx *OrderExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ratioExpr.
	VisitRatioExpr(ctx *RatioExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#settingExprList.
	VisitSettingExprList(ctx *SettingExprListContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#settingExpr.
	VisitSettingExpr(ctx *SettingExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#windowExpr.
	VisitWindowExpr(ctx *WindowExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#winPartitionByClause.
	VisitWinPartitionByClause(ctx *WinPartitionByClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#winOrderByClause.
	VisitWinOrderByClause(ctx *WinOrderByClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#winFrameClause.
	VisitWinFrameClause(ctx *WinFrameClauseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#frameStart.
	VisitFrameStart(ctx *FrameStartContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#frameBetween.
	VisitFrameBetween(ctx *FrameBetweenContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#winFrameBound.
	VisitWinFrameBound(ctx *WinFrameBoundContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#setStmt.
	VisitSetStmt(ctx *SetStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#showCreateDatabaseStmt.
	VisitShowCreateDatabaseStmt(ctx *ShowCreateDatabaseStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#showCreateDictionaryStmt.
	VisitShowCreateDictionaryStmt(ctx *ShowCreateDictionaryStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#showCreateTableStmt.
	VisitShowCreateTableStmt(ctx *ShowCreateTableStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#showDatabasesStmt.
	VisitShowDatabasesStmt(ctx *ShowDatabasesStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#showDictionariesStmt.
	VisitShowDictionariesStmt(ctx *ShowDictionariesStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#showTablesStmt.
	VisitShowTablesStmt(ctx *ShowTablesStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#systemStmt.
	VisitSystemStmt(ctx *SystemStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#truncateStmt.
	VisitTruncateStmt(ctx *TruncateStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#useStmt.
	VisitUseStmt(ctx *UseStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#watchStmt.
	VisitWatchStmt(ctx *WatchStmtContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnTypeExprSimple.
	VisitColumnTypeExprSimple(ctx *ColumnTypeExprSimpleContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnTypeExprNested.
	VisitColumnTypeExprNested(ctx *ColumnTypeExprNestedContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnTypeExprEnum.
	VisitColumnTypeExprEnum(ctx *ColumnTypeExprEnumContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnTypeExprComplex.
	VisitColumnTypeExprComplex(ctx *ColumnTypeExprComplexContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnTypeExprParam.
	VisitColumnTypeExprParam(ctx *ColumnTypeExprParamContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#columnExprList.
	VisitColumnExprList(ctx *ColumnExprListContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnsExprAsterisk.
	VisitColumnsExprAsterisk(ctx *ColumnsExprAsteriskContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnsExprSubquery.
	VisitColumnsExprSubquery(ctx *ColumnsExprSubqueryContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnsExprColumn.
	VisitColumnsExprColumn(ctx *ColumnsExprColumnContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprTernaryOp.
	VisitColumnExprTernaryOp(ctx *ColumnExprTernaryOpContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprAlias.
	VisitColumnExprAlias(ctx *ColumnExprAliasContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprExtract.
	VisitColumnExprExtract(ctx *ColumnExprExtractContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprNegate.
	VisitColumnExprNegate(ctx *ColumnExprNegateContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprSubquery.
	VisitColumnExprSubquery(ctx *ColumnExprSubqueryContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprLiteral.
	VisitColumnExprLiteral(ctx *ColumnExprLiteralContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprArray.
	VisitColumnExprArray(ctx *ColumnExprArrayContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprSubstring.
	VisitColumnExprSubstring(ctx *ColumnExprSubstringContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprCast.
	VisitColumnExprCast(ctx *ColumnExprCastContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprOr.
	VisitColumnExprOr(ctx *ColumnExprOrContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprPrecedence1.
	VisitColumnExprPrecedence1(ctx *ColumnExprPrecedence1Context) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprPrecedence2.
	VisitColumnExprPrecedence2(ctx *ColumnExprPrecedence2Context) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprPrecedence3.
	VisitColumnExprPrecedence3(ctx *ColumnExprPrecedence3Context) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprInterval.
	VisitColumnExprInterval(ctx *ColumnExprIntervalContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprIsNull.
	VisitColumnExprIsNull(ctx *ColumnExprIsNullContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprWinFunctionTarget.
	VisitColumnExprWinFunctionTarget(ctx *ColumnExprWinFunctionTargetContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprTrim.
	VisitColumnExprTrim(ctx *ColumnExprTrimContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprTuple.
	VisitColumnExprTuple(ctx *ColumnExprTupleContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprArrayAccess.
	VisitColumnExprArrayAccess(ctx *ColumnExprArrayAccessContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprBetween.
	VisitColumnExprBetween(ctx *ColumnExprBetweenContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprParens.
	VisitColumnExprParens(ctx *ColumnExprParensContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprTimestamp.
	VisitColumnExprTimestamp(ctx *ColumnExprTimestampContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprAnd.
	VisitColumnExprAnd(ctx *ColumnExprAndContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprTupleAccess.
	VisitColumnExprTupleAccess(ctx *ColumnExprTupleAccessContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprCase.
	VisitColumnExprCase(ctx *ColumnExprCaseContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprDate.
	VisitColumnExprDate(ctx *ColumnExprDateContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprNot.
	VisitColumnExprNot(ctx *ColumnExprNotContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprWinFunction.
	VisitColumnExprWinFunction(ctx *ColumnExprWinFunctionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprIdentifier.
	VisitColumnExprIdentifier(ctx *ColumnExprIdentifierContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprFunction.
	VisitColumnExprFunction(ctx *ColumnExprFunctionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#ColumnExprAsterisk.
	VisitColumnExprAsterisk(ctx *ColumnExprAsteriskContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#columnArgList.
	VisitColumnArgList(ctx *ColumnArgListContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#columnArgExpr.
	VisitColumnArgExpr(ctx *ColumnArgExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#columnLambdaExpr.
	VisitColumnLambdaExpr(ctx *ColumnLambdaExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#columnIdentifier.
	VisitColumnIdentifier(ctx *ColumnIdentifierContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#nestedIdentifier.
	VisitNestedIdentifier(ctx *NestedIdentifierContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#emptyListExpr.
	VisitEmptyListExpr(ctx *EmptyListExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#TableExprIdentifier.
	VisitTableExprIdentifier(ctx *TableExprIdentifierContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#TableExprSubquery.
	VisitTableExprSubquery(ctx *TableExprSubqueryContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#TableExprAlias.
	VisitTableExprAlias(ctx *TableExprAliasContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#TableExprFunction.
	VisitTableExprFunction(ctx *TableExprFunctionContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#tableFunctionExpr.
	VisitTableFunctionExpr(ctx *TableFunctionExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#tableIdentifier.
	VisitTableIdentifier(ctx *TableIdentifierContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#tableArgList.
	VisitTableArgList(ctx *TableArgListContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#tableArgExpr.
	VisitTableArgExpr(ctx *TableArgExprContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#databaseIdentifier.
	VisitDatabaseIdentifier(ctx *DatabaseIdentifierContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#floatingLiteral.
	VisitFloatingLiteral(ctx *FloatingLiteralContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#interval.
	VisitInterval(ctx *IntervalContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#keywordForAlias.
	VisitKeywordForAlias(ctx *KeywordForAliasContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#identifierOrNull.
	VisitIdentifierOrNull(ctx *IdentifierOrNullContext) interface{}

	// Visit a parse tree produced by ClickHouseParser#enumValue.
	VisitEnumValue(ctx *EnumValueContext) interface{}
}
