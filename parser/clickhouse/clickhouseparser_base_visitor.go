// Code generated from ClickHouseParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package clickhouse // ClickHouseParser
import "github.com/antlr4-go/antlr/v4"

type BaseClickHouseParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseClickHouseParserVisitor) VisitQueryStmt(ctx *QueryStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCtes(ctx *CtesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitNamedQuery(ctx *NamedQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnAliases(ctx *ColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableStmt(ctx *AlterTableStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseAddColumn(ctx *AlterTableClauseAddColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseAddIndex(ctx *AlterTableClauseAddIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseAddProjection(ctx *AlterTableClauseAddProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseAttach(ctx *AlterTableClauseAttachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseClearColumn(ctx *AlterTableClauseClearColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseClearIndex(ctx *AlterTableClauseClearIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseClearProjection(ctx *AlterTableClauseClearProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseComment(ctx *AlterTableClauseCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseDelete(ctx *AlterTableClauseDeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseDetach(ctx *AlterTableClauseDetachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseDropColumn(ctx *AlterTableClauseDropColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseDropIndex(ctx *AlterTableClauseDropIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseDropProjection(ctx *AlterTableClauseDropProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseDropPartition(ctx *AlterTableClauseDropPartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseFreezePartition(ctx *AlterTableClauseFreezePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseMaterializeIndex(ctx *AlterTableClauseMaterializeIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseMaterializeProjection(ctx *AlterTableClauseMaterializeProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseModifyCodec(ctx *AlterTableClauseModifyCodecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseModifyComment(ctx *AlterTableClauseModifyCommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseModifyRemove(ctx *AlterTableClauseModifyRemoveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseModify(ctx *AlterTableClauseModifyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseModifyOrderBy(ctx *AlterTableClauseModifyOrderByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseModifyTTL(ctx *AlterTableClauseModifyTTLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseMovePartition(ctx *AlterTableClauseMovePartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseRemoveTTL(ctx *AlterTableClauseRemoveTTLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseRename(ctx *AlterTableClauseRenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseReplace(ctx *AlterTableClauseReplaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlterTableClauseUpdate(ctx *AlterTableClauseUpdateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAssignmentExprList(ctx *AssignmentExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAssignmentExpr(ctx *AssignmentExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableColumnPropertyType(ctx *TableColumnPropertyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitPartitionClause(ctx *PartitionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAttachDictionaryStmt(ctx *AttachDictionaryStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCheckStmt(ctx *CheckStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCreateDatabaseStmt(ctx *CreateDatabaseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCreateDictionaryStmt(ctx *CreateDictionaryStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCreateLiveViewStmt(ctx *CreateLiveViewStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCreateMaterializedViewStmt(ctx *CreateMaterializedViewStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCreateTableStmt(ctx *CreateTableStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCreateViewStmt(ctx *CreateViewStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDictionarySchemaClause(ctx *DictionarySchemaClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDictionaryAttrDfnt(ctx *DictionaryAttrDfntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDictionaryEngineClause(ctx *DictionaryEngineClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDictionaryPrimaryKeyClause(ctx *DictionaryPrimaryKeyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDictionaryArgExpr(ctx *DictionaryArgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSourceClause(ctx *SourceClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitLifetimeClause(ctx *LifetimeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitLayoutClause(ctx *LayoutClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitRangeClause(ctx *RangeClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDictionarySettingsClause(ctx *DictionarySettingsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitClusterClause(ctx *ClusterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitUuidClause(ctx *UuidClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDestinationClause(ctx *DestinationClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSubqueryClause(ctx *SubqueryClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSchemaDescriptionClause(ctx *SchemaDescriptionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSchemaAsTableClause(ctx *SchemaAsTableClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSchemaAsFunctionClause(ctx *SchemaAsFunctionClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitEngineClause(ctx *EngineClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitPartitionByClause(ctx *PartitionByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitPrimaryKeyClause(ctx *PrimaryKeyClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSampleByClause(ctx *SampleByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTtlClause(ctx *TtlClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitEngineExpr(ctx *EngineExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableElementExprColumn(ctx *TableElementExprColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableElementExprConstraint(ctx *TableElementExprConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableElementExprIndex(ctx *TableElementExprIndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableElementExprProjection(ctx *TableElementExprProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableColumnDfnt(ctx *TableColumnDfntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableColumnPropertyExpr(ctx *TableColumnPropertyExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableIndexDfnt(ctx *TableIndexDfntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableProjectionDfnt(ctx *TableProjectionDfntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCodecExpr(ctx *CodecExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitCodecArgExpr(ctx *CodecArgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTtlExpr(ctx *TtlExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDescribeStmt(ctx *DescribeStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDropDatabaseStmt(ctx *DropDatabaseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDropTableStmt(ctx *DropTableStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDeleteStmt(ctx *DeleteStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitExistsDatabaseStmt(ctx *ExistsDatabaseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitExistsTableStmt(ctx *ExistsTableStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitExplainASTStmt(ctx *ExplainASTStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitExplainSyntaxStmt(ctx *ExplainSyntaxStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitInsertStmt(ctx *InsertStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnsClause(ctx *ColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDataClauseFormat(ctx *DataClauseFormatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDataClauseValues(ctx *DataClauseValuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDataClauseSelect(ctx *DataClauseSelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitKillMutationStmt(ctx *KillMutationStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitOptimizeStmt(ctx *OptimizeStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitRenameStmt(ctx *RenameStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitProjectionSelectStmt(ctx *ProjectionSelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSelectUnionStmt(ctx *SelectUnionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSelectStmtWithParens(ctx *SelectStmtWithParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSelectStmt(ctx *SelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTopClause(ctx *TopClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitArrayJoinClause(ctx *ArrayJoinClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitPrewhereClause(ctx *PrewhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitProjectionOrderByClause(ctx *ProjectionOrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitLimitByClause(ctx *LimitByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSettingsClause(ctx *SettingsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitJoinExprOp(ctx *JoinExprOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitJoinExprTable(ctx *JoinExprTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitJoinExprParens(ctx *JoinExprParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitJoinExprCrossOp(ctx *JoinExprCrossOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitJoinOpInner(ctx *JoinOpInnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitJoinOpLeftRight(ctx *JoinOpLeftRightContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitJoinOpFull(ctx *JoinOpFullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitJoinOpCross(ctx *JoinOpCrossContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitJoinConstraintClause(ctx *JoinConstraintClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSampleClause(ctx *SampleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitLimitExpr(ctx *LimitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitOrderExprList(ctx *OrderExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitOrderExpr(ctx *OrderExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitRatioExpr(ctx *RatioExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSettingExprList(ctx *SettingExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSettingExpr(ctx *SettingExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitWindowExpr(ctx *WindowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitWinPartitionByClause(ctx *WinPartitionByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitWinOrderByClause(ctx *WinOrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitWinFrameClause(ctx *WinFrameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitFrameStart(ctx *FrameStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitFrameBetween(ctx *FrameBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitWinFrameBound(ctx *WinFrameBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSetStmt(ctx *SetStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitShowCreateDatabaseStmt(ctx *ShowCreateDatabaseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitShowCreateDictionaryStmt(ctx *ShowCreateDictionaryStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitShowCreateTableStmt(ctx *ShowCreateTableStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitShowDatabasesStmt(ctx *ShowDatabasesStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitShowDictionariesStmt(ctx *ShowDictionariesStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitShowTablesStmt(ctx *ShowTablesStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitSystemStmt(ctx *SystemStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTruncateStmt(ctx *TruncateStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitUseStmt(ctx *UseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitWatchStmt(ctx *WatchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnTypeExprSimple(ctx *ColumnTypeExprSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnTypeExprNested(ctx *ColumnTypeExprNestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnTypeExprEnum(ctx *ColumnTypeExprEnumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnTypeExprComplex(ctx *ColumnTypeExprComplexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnTypeExprParam(ctx *ColumnTypeExprParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprList(ctx *ColumnExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnsExprAsterisk(ctx *ColumnsExprAsteriskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnsExprSubquery(ctx *ColumnsExprSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnsExprColumn(ctx *ColumnsExprColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprTernaryOp(ctx *ColumnExprTernaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprAlias(ctx *ColumnExprAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprExtract(ctx *ColumnExprExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprNegate(ctx *ColumnExprNegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprSubquery(ctx *ColumnExprSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprLiteral(ctx *ColumnExprLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprArray(ctx *ColumnExprArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprSubstring(ctx *ColumnExprSubstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprCast(ctx *ColumnExprCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprOr(ctx *ColumnExprOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprPrecedence1(ctx *ColumnExprPrecedence1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprPrecedence2(ctx *ColumnExprPrecedence2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprPrecedence3(ctx *ColumnExprPrecedence3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprInterval(ctx *ColumnExprIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprIsNull(ctx *ColumnExprIsNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprWinFunctionTarget(ctx *ColumnExprWinFunctionTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprTrim(ctx *ColumnExprTrimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprTuple(ctx *ColumnExprTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprArrayAccess(ctx *ColumnExprArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprBetween(ctx *ColumnExprBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprParens(ctx *ColumnExprParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprTimestamp(ctx *ColumnExprTimestampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprAnd(ctx *ColumnExprAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprTupleAccess(ctx *ColumnExprTupleAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprCase(ctx *ColumnExprCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprDate(ctx *ColumnExprDateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprNot(ctx *ColumnExprNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprWinFunction(ctx *ColumnExprWinFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprIdentifier(ctx *ColumnExprIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprFunction(ctx *ColumnExprFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnExprAsterisk(ctx *ColumnExprAsteriskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnArgList(ctx *ColumnArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnArgExpr(ctx *ColumnArgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnLambdaExpr(ctx *ColumnLambdaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitColumnIdentifier(ctx *ColumnIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitNestedIdentifier(ctx *NestedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitEmptyListExpr(ctx *EmptyListExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableExprIdentifier(ctx *TableExprIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableExprSubquery(ctx *TableExprSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableExprAlias(ctx *TableExprAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableExprFunction(ctx *TableExprFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableFunctionExpr(ctx *TableFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableIdentifier(ctx *TableIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableArgList(ctx *TableArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitTableArgExpr(ctx *TableArgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitDatabaseIdentifier(ctx *DatabaseIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitFloatingLiteral(ctx *FloatingLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitNumberLiteral(ctx *NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitKeywordForAlias(ctx *KeywordForAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitIdentifierOrNull(ctx *IdentifierOrNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseClickHouseParserVisitor) VisitEnumValue(ctx *EnumValueContext) interface{} {
	return v.VisitChildren(ctx)
}
