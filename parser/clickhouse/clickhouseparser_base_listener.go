// Code generated from ClickHouseParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package clickhouse // ClickHouseParser
import "github.com/antlr4-go/antlr/v4"

// BaseClickHouseParserListener is a complete listener for a parse tree produced by ClickHouseParser.
type BaseClickHouseParserListener struct{}

var _ ClickHouseParserListener = &BaseClickHouseParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseClickHouseParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseClickHouseParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseClickHouseParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseClickHouseParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQueryStmt is called when production queryStmt is entered.
func (s *BaseClickHouseParserListener) EnterQueryStmt(ctx *QueryStmtContext) {}

// ExitQueryStmt is called when production queryStmt is exited.
func (s *BaseClickHouseParserListener) ExitQueryStmt(ctx *QueryStmtContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseClickHouseParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseClickHouseParserListener) ExitQuery(ctx *QueryContext) {}

// EnterCtes is called when production ctes is entered.
func (s *BaseClickHouseParserListener) EnterCtes(ctx *CtesContext) {}

// ExitCtes is called when production ctes is exited.
func (s *BaseClickHouseParserListener) ExitCtes(ctx *CtesContext) {}

// EnterNamedQuery is called when production namedQuery is entered.
func (s *BaseClickHouseParserListener) EnterNamedQuery(ctx *NamedQueryContext) {}

// ExitNamedQuery is called when production namedQuery is exited.
func (s *BaseClickHouseParserListener) ExitNamedQuery(ctx *NamedQueryContext) {}

// EnterColumnAliases is called when production columnAliases is entered.
func (s *BaseClickHouseParserListener) EnterColumnAliases(ctx *ColumnAliasesContext) {}

// ExitColumnAliases is called when production columnAliases is exited.
func (s *BaseClickHouseParserListener) ExitColumnAliases(ctx *ColumnAliasesContext) {}

// EnterAlterTableStmt is called when production AlterTableStmt is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableStmt(ctx *AlterTableStmtContext) {}

// ExitAlterTableStmt is called when production AlterTableStmt is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableStmt(ctx *AlterTableStmtContext) {}

// EnterAlterTableClauseAddColumn is called when production AlterTableClauseAddColumn is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseAddColumn(ctx *AlterTableClauseAddColumnContext) {
}

// ExitAlterTableClauseAddColumn is called when production AlterTableClauseAddColumn is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseAddColumn(ctx *AlterTableClauseAddColumnContext) {
}

// EnterAlterTableClauseAddIndex is called when production AlterTableClauseAddIndex is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseAddIndex(ctx *AlterTableClauseAddIndexContext) {
}

// ExitAlterTableClauseAddIndex is called when production AlterTableClauseAddIndex is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseAddIndex(ctx *AlterTableClauseAddIndexContext) {
}

// EnterAlterTableClauseAddProjection is called when production AlterTableClauseAddProjection is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseAddProjection(ctx *AlterTableClauseAddProjectionContext) {
}

// ExitAlterTableClauseAddProjection is called when production AlterTableClauseAddProjection is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseAddProjection(ctx *AlterTableClauseAddProjectionContext) {
}

// EnterAlterTableClauseAttach is called when production AlterTableClauseAttach is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseAttach(ctx *AlterTableClauseAttachContext) {
}

// ExitAlterTableClauseAttach is called when production AlterTableClauseAttach is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseAttach(ctx *AlterTableClauseAttachContext) {
}

// EnterAlterTableClauseClearColumn is called when production AlterTableClauseClearColumn is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseClearColumn(ctx *AlterTableClauseClearColumnContext) {
}

// ExitAlterTableClauseClearColumn is called when production AlterTableClauseClearColumn is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseClearColumn(ctx *AlterTableClauseClearColumnContext) {
}

// EnterAlterTableClauseClearIndex is called when production AlterTableClauseClearIndex is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseClearIndex(ctx *AlterTableClauseClearIndexContext) {
}

// ExitAlterTableClauseClearIndex is called when production AlterTableClauseClearIndex is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseClearIndex(ctx *AlterTableClauseClearIndexContext) {
}

// EnterAlterTableClauseClearProjection is called when production AlterTableClauseClearProjection is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseClearProjection(ctx *AlterTableClauseClearProjectionContext) {
}

// ExitAlterTableClauseClearProjection is called when production AlterTableClauseClearProjection is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseClearProjection(ctx *AlterTableClauseClearProjectionContext) {
}

// EnterAlterTableClauseComment is called when production AlterTableClauseComment is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseComment(ctx *AlterTableClauseCommentContext) {
}

// ExitAlterTableClauseComment is called when production AlterTableClauseComment is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseComment(ctx *AlterTableClauseCommentContext) {
}

// EnterAlterTableClauseDelete is called when production AlterTableClauseDelete is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseDelete(ctx *AlterTableClauseDeleteContext) {
}

// ExitAlterTableClauseDelete is called when production AlterTableClauseDelete is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseDelete(ctx *AlterTableClauseDeleteContext) {
}

// EnterAlterTableClauseDetach is called when production AlterTableClauseDetach is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseDetach(ctx *AlterTableClauseDetachContext) {
}

// ExitAlterTableClauseDetach is called when production AlterTableClauseDetach is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseDetach(ctx *AlterTableClauseDetachContext) {
}

// EnterAlterTableClauseDropColumn is called when production AlterTableClauseDropColumn is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseDropColumn(ctx *AlterTableClauseDropColumnContext) {
}

// ExitAlterTableClauseDropColumn is called when production AlterTableClauseDropColumn is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseDropColumn(ctx *AlterTableClauseDropColumnContext) {
}

// EnterAlterTableClauseDropIndex is called when production AlterTableClauseDropIndex is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseDropIndex(ctx *AlterTableClauseDropIndexContext) {
}

// ExitAlterTableClauseDropIndex is called when production AlterTableClauseDropIndex is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseDropIndex(ctx *AlterTableClauseDropIndexContext) {
}

// EnterAlterTableClauseDropProjection is called when production AlterTableClauseDropProjection is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseDropProjection(ctx *AlterTableClauseDropProjectionContext) {
}

// ExitAlterTableClauseDropProjection is called when production AlterTableClauseDropProjection is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseDropProjection(ctx *AlterTableClauseDropProjectionContext) {
}

// EnterAlterTableClauseDropPartition is called when production AlterTableClauseDropPartition is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseDropPartition(ctx *AlterTableClauseDropPartitionContext) {
}

// ExitAlterTableClauseDropPartition is called when production AlterTableClauseDropPartition is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseDropPartition(ctx *AlterTableClauseDropPartitionContext) {
}

// EnterAlterTableClauseFreezePartition is called when production AlterTableClauseFreezePartition is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseFreezePartition(ctx *AlterTableClauseFreezePartitionContext) {
}

// ExitAlterTableClauseFreezePartition is called when production AlterTableClauseFreezePartition is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseFreezePartition(ctx *AlterTableClauseFreezePartitionContext) {
}

// EnterAlterTableClauseMaterializeIndex is called when production AlterTableClauseMaterializeIndex is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseMaterializeIndex(ctx *AlterTableClauseMaterializeIndexContext) {
}

// ExitAlterTableClauseMaterializeIndex is called when production AlterTableClauseMaterializeIndex is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseMaterializeIndex(ctx *AlterTableClauseMaterializeIndexContext) {
}

// EnterAlterTableClauseMaterializeProjection is called when production AlterTableClauseMaterializeProjection is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseMaterializeProjection(ctx *AlterTableClauseMaterializeProjectionContext) {
}

// ExitAlterTableClauseMaterializeProjection is called when production AlterTableClauseMaterializeProjection is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseMaterializeProjection(ctx *AlterTableClauseMaterializeProjectionContext) {
}

// EnterAlterTableClauseModifyCodec is called when production AlterTableClauseModifyCodec is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseModifyCodec(ctx *AlterTableClauseModifyCodecContext) {
}

// ExitAlterTableClauseModifyCodec is called when production AlterTableClauseModifyCodec is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseModifyCodec(ctx *AlterTableClauseModifyCodecContext) {
}

// EnterAlterTableClauseModifyComment is called when production AlterTableClauseModifyComment is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseModifyComment(ctx *AlterTableClauseModifyCommentContext) {
}

// ExitAlterTableClauseModifyComment is called when production AlterTableClauseModifyComment is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseModifyComment(ctx *AlterTableClauseModifyCommentContext) {
}

// EnterAlterTableClauseModifyRemove is called when production AlterTableClauseModifyRemove is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseModifyRemove(ctx *AlterTableClauseModifyRemoveContext) {
}

// ExitAlterTableClauseModifyRemove is called when production AlterTableClauseModifyRemove is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseModifyRemove(ctx *AlterTableClauseModifyRemoveContext) {
}

// EnterAlterTableClauseModify is called when production AlterTableClauseModify is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseModify(ctx *AlterTableClauseModifyContext) {
}

// ExitAlterTableClauseModify is called when production AlterTableClauseModify is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseModify(ctx *AlterTableClauseModifyContext) {
}

// EnterAlterTableClauseModifyOrderBy is called when production AlterTableClauseModifyOrderBy is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseModifyOrderBy(ctx *AlterTableClauseModifyOrderByContext) {
}

// ExitAlterTableClauseModifyOrderBy is called when production AlterTableClauseModifyOrderBy is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseModifyOrderBy(ctx *AlterTableClauseModifyOrderByContext) {
}

// EnterAlterTableClauseModifyTTL is called when production AlterTableClauseModifyTTL is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseModifyTTL(ctx *AlterTableClauseModifyTTLContext) {
}

// ExitAlterTableClauseModifyTTL is called when production AlterTableClauseModifyTTL is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseModifyTTL(ctx *AlterTableClauseModifyTTLContext) {
}

// EnterAlterTableClauseMovePartition is called when production AlterTableClauseMovePartition is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseMovePartition(ctx *AlterTableClauseMovePartitionContext) {
}

// ExitAlterTableClauseMovePartition is called when production AlterTableClauseMovePartition is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseMovePartition(ctx *AlterTableClauseMovePartitionContext) {
}

// EnterAlterTableClauseRemoveTTL is called when production AlterTableClauseRemoveTTL is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseRemoveTTL(ctx *AlterTableClauseRemoveTTLContext) {
}

// ExitAlterTableClauseRemoveTTL is called when production AlterTableClauseRemoveTTL is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseRemoveTTL(ctx *AlterTableClauseRemoveTTLContext) {
}

// EnterAlterTableClauseRename is called when production AlterTableClauseRename is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseRename(ctx *AlterTableClauseRenameContext) {
}

// ExitAlterTableClauseRename is called when production AlterTableClauseRename is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseRename(ctx *AlterTableClauseRenameContext) {
}

// EnterAlterTableClauseReplace is called when production AlterTableClauseReplace is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseReplace(ctx *AlterTableClauseReplaceContext) {
}

// ExitAlterTableClauseReplace is called when production AlterTableClauseReplace is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseReplace(ctx *AlterTableClauseReplaceContext) {
}

// EnterAlterTableClauseUpdate is called when production AlterTableClauseUpdate is entered.
func (s *BaseClickHouseParserListener) EnterAlterTableClauseUpdate(ctx *AlterTableClauseUpdateContext) {
}

// ExitAlterTableClauseUpdate is called when production AlterTableClauseUpdate is exited.
func (s *BaseClickHouseParserListener) ExitAlterTableClauseUpdate(ctx *AlterTableClauseUpdateContext) {
}

// EnterAssignmentExprList is called when production assignmentExprList is entered.
func (s *BaseClickHouseParserListener) EnterAssignmentExprList(ctx *AssignmentExprListContext) {}

// ExitAssignmentExprList is called when production assignmentExprList is exited.
func (s *BaseClickHouseParserListener) ExitAssignmentExprList(ctx *AssignmentExprListContext) {}

// EnterAssignmentExpr is called when production assignmentExpr is entered.
func (s *BaseClickHouseParserListener) EnterAssignmentExpr(ctx *AssignmentExprContext) {}

// ExitAssignmentExpr is called when production assignmentExpr is exited.
func (s *BaseClickHouseParserListener) ExitAssignmentExpr(ctx *AssignmentExprContext) {}

// EnterTableColumnPropertyType is called when production tableColumnPropertyType is entered.
func (s *BaseClickHouseParserListener) EnterTableColumnPropertyType(ctx *TableColumnPropertyTypeContext) {
}

// ExitTableColumnPropertyType is called when production tableColumnPropertyType is exited.
func (s *BaseClickHouseParserListener) ExitTableColumnPropertyType(ctx *TableColumnPropertyTypeContext) {
}

// EnterPartitionClause is called when production partitionClause is entered.
func (s *BaseClickHouseParserListener) EnterPartitionClause(ctx *PartitionClauseContext) {}

// ExitPartitionClause is called when production partitionClause is exited.
func (s *BaseClickHouseParserListener) ExitPartitionClause(ctx *PartitionClauseContext) {}

// EnterAttachDictionaryStmt is called when production AttachDictionaryStmt is entered.
func (s *BaseClickHouseParserListener) EnterAttachDictionaryStmt(ctx *AttachDictionaryStmtContext) {}

// ExitAttachDictionaryStmt is called when production AttachDictionaryStmt is exited.
func (s *BaseClickHouseParserListener) ExitAttachDictionaryStmt(ctx *AttachDictionaryStmtContext) {}

// EnterCheckStmt is called when production checkStmt is entered.
func (s *BaseClickHouseParserListener) EnterCheckStmt(ctx *CheckStmtContext) {}

// ExitCheckStmt is called when production checkStmt is exited.
func (s *BaseClickHouseParserListener) ExitCheckStmt(ctx *CheckStmtContext) {}

// EnterCreateDatabaseStmt is called when production CreateDatabaseStmt is entered.
func (s *BaseClickHouseParserListener) EnterCreateDatabaseStmt(ctx *CreateDatabaseStmtContext) {}

// ExitCreateDatabaseStmt is called when production CreateDatabaseStmt is exited.
func (s *BaseClickHouseParserListener) ExitCreateDatabaseStmt(ctx *CreateDatabaseStmtContext) {}

// EnterCreateDictionaryStmt is called when production CreateDictionaryStmt is entered.
func (s *BaseClickHouseParserListener) EnterCreateDictionaryStmt(ctx *CreateDictionaryStmtContext) {}

// ExitCreateDictionaryStmt is called when production CreateDictionaryStmt is exited.
func (s *BaseClickHouseParserListener) ExitCreateDictionaryStmt(ctx *CreateDictionaryStmtContext) {}

// EnterCreateLiveViewStmt is called when production CreateLiveViewStmt is entered.
func (s *BaseClickHouseParserListener) EnterCreateLiveViewStmt(ctx *CreateLiveViewStmtContext) {}

// ExitCreateLiveViewStmt is called when production CreateLiveViewStmt is exited.
func (s *BaseClickHouseParserListener) ExitCreateLiveViewStmt(ctx *CreateLiveViewStmtContext) {}

// EnterCreateMaterializedViewStmt is called when production CreateMaterializedViewStmt is entered.
func (s *BaseClickHouseParserListener) EnterCreateMaterializedViewStmt(ctx *CreateMaterializedViewStmtContext) {
}

// ExitCreateMaterializedViewStmt is called when production CreateMaterializedViewStmt is exited.
func (s *BaseClickHouseParserListener) ExitCreateMaterializedViewStmt(ctx *CreateMaterializedViewStmtContext) {
}

// EnterCreateTableStmt is called when production CreateTableStmt is entered.
func (s *BaseClickHouseParserListener) EnterCreateTableStmt(ctx *CreateTableStmtContext) {}

// ExitCreateTableStmt is called when production CreateTableStmt is exited.
func (s *BaseClickHouseParserListener) ExitCreateTableStmt(ctx *CreateTableStmtContext) {}

// EnterCreateViewStmt is called when production CreateViewStmt is entered.
func (s *BaseClickHouseParserListener) EnterCreateViewStmt(ctx *CreateViewStmtContext) {}

// ExitCreateViewStmt is called when production CreateViewStmt is exited.
func (s *BaseClickHouseParserListener) ExitCreateViewStmt(ctx *CreateViewStmtContext) {}

// EnterDictionarySchemaClause is called when production dictionarySchemaClause is entered.
func (s *BaseClickHouseParserListener) EnterDictionarySchemaClause(ctx *DictionarySchemaClauseContext) {
}

// ExitDictionarySchemaClause is called when production dictionarySchemaClause is exited.
func (s *BaseClickHouseParserListener) ExitDictionarySchemaClause(ctx *DictionarySchemaClauseContext) {
}

// EnterDictionaryAttrDfnt is called when production dictionaryAttrDfnt is entered.
func (s *BaseClickHouseParserListener) EnterDictionaryAttrDfnt(ctx *DictionaryAttrDfntContext) {}

// ExitDictionaryAttrDfnt is called when production dictionaryAttrDfnt is exited.
func (s *BaseClickHouseParserListener) ExitDictionaryAttrDfnt(ctx *DictionaryAttrDfntContext) {}

// EnterDictionaryEngineClause is called when production dictionaryEngineClause is entered.
func (s *BaseClickHouseParserListener) EnterDictionaryEngineClause(ctx *DictionaryEngineClauseContext) {
}

// ExitDictionaryEngineClause is called when production dictionaryEngineClause is exited.
func (s *BaseClickHouseParserListener) ExitDictionaryEngineClause(ctx *DictionaryEngineClauseContext) {
}

// EnterDictionaryPrimaryKeyClause is called when production dictionaryPrimaryKeyClause is entered.
func (s *BaseClickHouseParserListener) EnterDictionaryPrimaryKeyClause(ctx *DictionaryPrimaryKeyClauseContext) {
}

// ExitDictionaryPrimaryKeyClause is called when production dictionaryPrimaryKeyClause is exited.
func (s *BaseClickHouseParserListener) ExitDictionaryPrimaryKeyClause(ctx *DictionaryPrimaryKeyClauseContext) {
}

// EnterDictionaryArgExpr is called when production dictionaryArgExpr is entered.
func (s *BaseClickHouseParserListener) EnterDictionaryArgExpr(ctx *DictionaryArgExprContext) {}

// ExitDictionaryArgExpr is called when production dictionaryArgExpr is exited.
func (s *BaseClickHouseParserListener) ExitDictionaryArgExpr(ctx *DictionaryArgExprContext) {}

// EnterSourceClause is called when production sourceClause is entered.
func (s *BaseClickHouseParserListener) EnterSourceClause(ctx *SourceClauseContext) {}

// ExitSourceClause is called when production sourceClause is exited.
func (s *BaseClickHouseParserListener) ExitSourceClause(ctx *SourceClauseContext) {}

// EnterLifetimeClause is called when production lifetimeClause is entered.
func (s *BaseClickHouseParserListener) EnterLifetimeClause(ctx *LifetimeClauseContext) {}

// ExitLifetimeClause is called when production lifetimeClause is exited.
func (s *BaseClickHouseParserListener) ExitLifetimeClause(ctx *LifetimeClauseContext) {}

// EnterLayoutClause is called when production layoutClause is entered.
func (s *BaseClickHouseParserListener) EnterLayoutClause(ctx *LayoutClauseContext) {}

// ExitLayoutClause is called when production layoutClause is exited.
func (s *BaseClickHouseParserListener) ExitLayoutClause(ctx *LayoutClauseContext) {}

// EnterRangeClause is called when production rangeClause is entered.
func (s *BaseClickHouseParserListener) EnterRangeClause(ctx *RangeClauseContext) {}

// ExitRangeClause is called when production rangeClause is exited.
func (s *BaseClickHouseParserListener) ExitRangeClause(ctx *RangeClauseContext) {}

// EnterDictionarySettingsClause is called when production dictionarySettingsClause is entered.
func (s *BaseClickHouseParserListener) EnterDictionarySettingsClause(ctx *DictionarySettingsClauseContext) {
}

// ExitDictionarySettingsClause is called when production dictionarySettingsClause is exited.
func (s *BaseClickHouseParserListener) ExitDictionarySettingsClause(ctx *DictionarySettingsClauseContext) {
}

// EnterClusterClause is called when production clusterClause is entered.
func (s *BaseClickHouseParserListener) EnterClusterClause(ctx *ClusterClauseContext) {}

// ExitClusterClause is called when production clusterClause is exited.
func (s *BaseClickHouseParserListener) ExitClusterClause(ctx *ClusterClauseContext) {}

// EnterUuidClause is called when production uuidClause is entered.
func (s *BaseClickHouseParserListener) EnterUuidClause(ctx *UuidClauseContext) {}

// ExitUuidClause is called when production uuidClause is exited.
func (s *BaseClickHouseParserListener) ExitUuidClause(ctx *UuidClauseContext) {}

// EnterDestinationClause is called when production destinationClause is entered.
func (s *BaseClickHouseParserListener) EnterDestinationClause(ctx *DestinationClauseContext) {}

// ExitDestinationClause is called when production destinationClause is exited.
func (s *BaseClickHouseParserListener) ExitDestinationClause(ctx *DestinationClauseContext) {}

// EnterSubqueryClause is called when production subqueryClause is entered.
func (s *BaseClickHouseParserListener) EnterSubqueryClause(ctx *SubqueryClauseContext) {}

// ExitSubqueryClause is called when production subqueryClause is exited.
func (s *BaseClickHouseParserListener) ExitSubqueryClause(ctx *SubqueryClauseContext) {}

// EnterSchemaDescriptionClause is called when production SchemaDescriptionClause is entered.
func (s *BaseClickHouseParserListener) EnterSchemaDescriptionClause(ctx *SchemaDescriptionClauseContext) {
}

// ExitSchemaDescriptionClause is called when production SchemaDescriptionClause is exited.
func (s *BaseClickHouseParserListener) ExitSchemaDescriptionClause(ctx *SchemaDescriptionClauseContext) {
}

// EnterSchemaAsTableClause is called when production SchemaAsTableClause is entered.
func (s *BaseClickHouseParserListener) EnterSchemaAsTableClause(ctx *SchemaAsTableClauseContext) {}

// ExitSchemaAsTableClause is called when production SchemaAsTableClause is exited.
func (s *BaseClickHouseParserListener) ExitSchemaAsTableClause(ctx *SchemaAsTableClauseContext) {}

// EnterSchemaAsFunctionClause is called when production SchemaAsFunctionClause is entered.
func (s *BaseClickHouseParserListener) EnterSchemaAsFunctionClause(ctx *SchemaAsFunctionClauseContext) {
}

// ExitSchemaAsFunctionClause is called when production SchemaAsFunctionClause is exited.
func (s *BaseClickHouseParserListener) ExitSchemaAsFunctionClause(ctx *SchemaAsFunctionClauseContext) {
}

// EnterEngineClause is called when production engineClause is entered.
func (s *BaseClickHouseParserListener) EnterEngineClause(ctx *EngineClauseContext) {}

// ExitEngineClause is called when production engineClause is exited.
func (s *BaseClickHouseParserListener) ExitEngineClause(ctx *EngineClauseContext) {}

// EnterPartitionByClause is called when production partitionByClause is entered.
func (s *BaseClickHouseParserListener) EnterPartitionByClause(ctx *PartitionByClauseContext) {}

// ExitPartitionByClause is called when production partitionByClause is exited.
func (s *BaseClickHouseParserListener) ExitPartitionByClause(ctx *PartitionByClauseContext) {}

// EnterPrimaryKeyClause is called when production primaryKeyClause is entered.
func (s *BaseClickHouseParserListener) EnterPrimaryKeyClause(ctx *PrimaryKeyClauseContext) {}

// ExitPrimaryKeyClause is called when production primaryKeyClause is exited.
func (s *BaseClickHouseParserListener) ExitPrimaryKeyClause(ctx *PrimaryKeyClauseContext) {}

// EnterSampleByClause is called when production sampleByClause is entered.
func (s *BaseClickHouseParserListener) EnterSampleByClause(ctx *SampleByClauseContext) {}

// ExitSampleByClause is called when production sampleByClause is exited.
func (s *BaseClickHouseParserListener) ExitSampleByClause(ctx *SampleByClauseContext) {}

// EnterTtlClause is called when production ttlClause is entered.
func (s *BaseClickHouseParserListener) EnterTtlClause(ctx *TtlClauseContext) {}

// ExitTtlClause is called when production ttlClause is exited.
func (s *BaseClickHouseParserListener) ExitTtlClause(ctx *TtlClauseContext) {}

// EnterEngineExpr is called when production engineExpr is entered.
func (s *BaseClickHouseParserListener) EnterEngineExpr(ctx *EngineExprContext) {}

// ExitEngineExpr is called when production engineExpr is exited.
func (s *BaseClickHouseParserListener) ExitEngineExpr(ctx *EngineExprContext) {}

// EnterTableElementExprColumn is called when production TableElementExprColumn is entered.
func (s *BaseClickHouseParserListener) EnterTableElementExprColumn(ctx *TableElementExprColumnContext) {
}

// ExitTableElementExprColumn is called when production TableElementExprColumn is exited.
func (s *BaseClickHouseParserListener) ExitTableElementExprColumn(ctx *TableElementExprColumnContext) {
}

// EnterTableElementExprConstraint is called when production TableElementExprConstraint is entered.
func (s *BaseClickHouseParserListener) EnterTableElementExprConstraint(ctx *TableElementExprConstraintContext) {
}

// ExitTableElementExprConstraint is called when production TableElementExprConstraint is exited.
func (s *BaseClickHouseParserListener) ExitTableElementExprConstraint(ctx *TableElementExprConstraintContext) {
}

// EnterTableElementExprIndex is called when production TableElementExprIndex is entered.
func (s *BaseClickHouseParserListener) EnterTableElementExprIndex(ctx *TableElementExprIndexContext) {
}

// ExitTableElementExprIndex is called when production TableElementExprIndex is exited.
func (s *BaseClickHouseParserListener) ExitTableElementExprIndex(ctx *TableElementExprIndexContext) {}

// EnterTableElementExprProjection is called when production TableElementExprProjection is entered.
func (s *BaseClickHouseParserListener) EnterTableElementExprProjection(ctx *TableElementExprProjectionContext) {
}

// ExitTableElementExprProjection is called when production TableElementExprProjection is exited.
func (s *BaseClickHouseParserListener) ExitTableElementExprProjection(ctx *TableElementExprProjectionContext) {
}

// EnterTableColumnDfnt is called when production tableColumnDfnt is entered.
func (s *BaseClickHouseParserListener) EnterTableColumnDfnt(ctx *TableColumnDfntContext) {}

// ExitTableColumnDfnt is called when production tableColumnDfnt is exited.
func (s *BaseClickHouseParserListener) ExitTableColumnDfnt(ctx *TableColumnDfntContext) {}

// EnterTableColumnPropertyExpr is called when production tableColumnPropertyExpr is entered.
func (s *BaseClickHouseParserListener) EnterTableColumnPropertyExpr(ctx *TableColumnPropertyExprContext) {
}

// ExitTableColumnPropertyExpr is called when production tableColumnPropertyExpr is exited.
func (s *BaseClickHouseParserListener) ExitTableColumnPropertyExpr(ctx *TableColumnPropertyExprContext) {
}

// EnterTableIndexDfnt is called when production tableIndexDfnt is entered.
func (s *BaseClickHouseParserListener) EnterTableIndexDfnt(ctx *TableIndexDfntContext) {}

// ExitTableIndexDfnt is called when production tableIndexDfnt is exited.
func (s *BaseClickHouseParserListener) ExitTableIndexDfnt(ctx *TableIndexDfntContext) {}

// EnterTableProjectionDfnt is called when production tableProjectionDfnt is entered.
func (s *BaseClickHouseParserListener) EnterTableProjectionDfnt(ctx *TableProjectionDfntContext) {}

// ExitTableProjectionDfnt is called when production tableProjectionDfnt is exited.
func (s *BaseClickHouseParserListener) ExitTableProjectionDfnt(ctx *TableProjectionDfntContext) {}

// EnterCodecExpr is called when production codecExpr is entered.
func (s *BaseClickHouseParserListener) EnterCodecExpr(ctx *CodecExprContext) {}

// ExitCodecExpr is called when production codecExpr is exited.
func (s *BaseClickHouseParserListener) ExitCodecExpr(ctx *CodecExprContext) {}

// EnterCodecArgExpr is called when production codecArgExpr is entered.
func (s *BaseClickHouseParserListener) EnterCodecArgExpr(ctx *CodecArgExprContext) {}

// ExitCodecArgExpr is called when production codecArgExpr is exited.
func (s *BaseClickHouseParserListener) ExitCodecArgExpr(ctx *CodecArgExprContext) {}

// EnterTtlExpr is called when production ttlExpr is entered.
func (s *BaseClickHouseParserListener) EnterTtlExpr(ctx *TtlExprContext) {}

// ExitTtlExpr is called when production ttlExpr is exited.
func (s *BaseClickHouseParserListener) ExitTtlExpr(ctx *TtlExprContext) {}

// EnterDescribeStmt is called when production describeStmt is entered.
func (s *BaseClickHouseParserListener) EnterDescribeStmt(ctx *DescribeStmtContext) {}

// ExitDescribeStmt is called when production describeStmt is exited.
func (s *BaseClickHouseParserListener) ExitDescribeStmt(ctx *DescribeStmtContext) {}

// EnterDropDatabaseStmt is called when production DropDatabaseStmt is entered.
func (s *BaseClickHouseParserListener) EnterDropDatabaseStmt(ctx *DropDatabaseStmtContext) {}

// ExitDropDatabaseStmt is called when production DropDatabaseStmt is exited.
func (s *BaseClickHouseParserListener) ExitDropDatabaseStmt(ctx *DropDatabaseStmtContext) {}

// EnterDropTableStmt is called when production DropTableStmt is entered.
func (s *BaseClickHouseParserListener) EnterDropTableStmt(ctx *DropTableStmtContext) {}

// ExitDropTableStmt is called when production DropTableStmt is exited.
func (s *BaseClickHouseParserListener) ExitDropTableStmt(ctx *DropTableStmtContext) {}

// EnterDeleteStmt is called when production deleteStmt is entered.
func (s *BaseClickHouseParserListener) EnterDeleteStmt(ctx *DeleteStmtContext) {}

// ExitDeleteStmt is called when production deleteStmt is exited.
func (s *BaseClickHouseParserListener) ExitDeleteStmt(ctx *DeleteStmtContext) {}

// EnterExistsDatabaseStmt is called when production ExistsDatabaseStmt is entered.
func (s *BaseClickHouseParserListener) EnterExistsDatabaseStmt(ctx *ExistsDatabaseStmtContext) {}

// ExitExistsDatabaseStmt is called when production ExistsDatabaseStmt is exited.
func (s *BaseClickHouseParserListener) ExitExistsDatabaseStmt(ctx *ExistsDatabaseStmtContext) {}

// EnterExistsTableStmt is called when production ExistsTableStmt is entered.
func (s *BaseClickHouseParserListener) EnterExistsTableStmt(ctx *ExistsTableStmtContext) {}

// ExitExistsTableStmt is called when production ExistsTableStmt is exited.
func (s *BaseClickHouseParserListener) ExitExistsTableStmt(ctx *ExistsTableStmtContext) {}

// EnterExplainASTStmt is called when production ExplainASTStmt is entered.
func (s *BaseClickHouseParserListener) EnterExplainASTStmt(ctx *ExplainASTStmtContext) {}

// ExitExplainASTStmt is called when production ExplainASTStmt is exited.
func (s *BaseClickHouseParserListener) ExitExplainASTStmt(ctx *ExplainASTStmtContext) {}

// EnterExplainSyntaxStmt is called when production ExplainSyntaxStmt is entered.
func (s *BaseClickHouseParserListener) EnterExplainSyntaxStmt(ctx *ExplainSyntaxStmtContext) {}

// ExitExplainSyntaxStmt is called when production ExplainSyntaxStmt is exited.
func (s *BaseClickHouseParserListener) ExitExplainSyntaxStmt(ctx *ExplainSyntaxStmtContext) {}

// EnterInsertStmt is called when production insertStmt is entered.
func (s *BaseClickHouseParserListener) EnterInsertStmt(ctx *InsertStmtContext) {}

// ExitInsertStmt is called when production insertStmt is exited.
func (s *BaseClickHouseParserListener) ExitInsertStmt(ctx *InsertStmtContext) {}

// EnterColumnsClause is called when production columnsClause is entered.
func (s *BaseClickHouseParserListener) EnterColumnsClause(ctx *ColumnsClauseContext) {}

// ExitColumnsClause is called when production columnsClause is exited.
func (s *BaseClickHouseParserListener) ExitColumnsClause(ctx *ColumnsClauseContext) {}

// EnterDataClauseFormat is called when production DataClauseFormat is entered.
func (s *BaseClickHouseParserListener) EnterDataClauseFormat(ctx *DataClauseFormatContext) {}

// ExitDataClauseFormat is called when production DataClauseFormat is exited.
func (s *BaseClickHouseParserListener) ExitDataClauseFormat(ctx *DataClauseFormatContext) {}

// EnterDataClauseValues is called when production DataClauseValues is entered.
func (s *BaseClickHouseParserListener) EnterDataClauseValues(ctx *DataClauseValuesContext) {}

// ExitDataClauseValues is called when production DataClauseValues is exited.
func (s *BaseClickHouseParserListener) ExitDataClauseValues(ctx *DataClauseValuesContext) {}

// EnterDataClauseSelect is called when production DataClauseSelect is entered.
func (s *BaseClickHouseParserListener) EnterDataClauseSelect(ctx *DataClauseSelectContext) {}

// ExitDataClauseSelect is called when production DataClauseSelect is exited.
func (s *BaseClickHouseParserListener) ExitDataClauseSelect(ctx *DataClauseSelectContext) {}

// EnterKillMutationStmt is called when production KillMutationStmt is entered.
func (s *BaseClickHouseParserListener) EnterKillMutationStmt(ctx *KillMutationStmtContext) {}

// ExitKillMutationStmt is called when production KillMutationStmt is exited.
func (s *BaseClickHouseParserListener) ExitKillMutationStmt(ctx *KillMutationStmtContext) {}

// EnterOptimizeStmt is called when production optimizeStmt is entered.
func (s *BaseClickHouseParserListener) EnterOptimizeStmt(ctx *OptimizeStmtContext) {}

// ExitOptimizeStmt is called when production optimizeStmt is exited.
func (s *BaseClickHouseParserListener) ExitOptimizeStmt(ctx *OptimizeStmtContext) {}

// EnterRenameStmt is called when production renameStmt is entered.
func (s *BaseClickHouseParserListener) EnterRenameStmt(ctx *RenameStmtContext) {}

// ExitRenameStmt is called when production renameStmt is exited.
func (s *BaseClickHouseParserListener) ExitRenameStmt(ctx *RenameStmtContext) {}

// EnterProjectionSelectStmt is called when production projectionSelectStmt is entered.
func (s *BaseClickHouseParserListener) EnterProjectionSelectStmt(ctx *ProjectionSelectStmtContext) {}

// ExitProjectionSelectStmt is called when production projectionSelectStmt is exited.
func (s *BaseClickHouseParserListener) ExitProjectionSelectStmt(ctx *ProjectionSelectStmtContext) {}

// EnterSelectUnionStmt is called when production selectUnionStmt is entered.
func (s *BaseClickHouseParserListener) EnterSelectUnionStmt(ctx *SelectUnionStmtContext) {}

// ExitSelectUnionStmt is called when production selectUnionStmt is exited.
func (s *BaseClickHouseParserListener) ExitSelectUnionStmt(ctx *SelectUnionStmtContext) {}

// EnterSelectStmtWithParens is called when production selectStmtWithParens is entered.
func (s *BaseClickHouseParserListener) EnterSelectStmtWithParens(ctx *SelectStmtWithParensContext) {}

// ExitSelectStmtWithParens is called when production selectStmtWithParens is exited.
func (s *BaseClickHouseParserListener) ExitSelectStmtWithParens(ctx *SelectStmtWithParensContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseClickHouseParserListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseClickHouseParserListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseClickHouseParserListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseClickHouseParserListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterTopClause is called when production topClause is entered.
func (s *BaseClickHouseParserListener) EnterTopClause(ctx *TopClauseContext) {}

// ExitTopClause is called when production topClause is exited.
func (s *BaseClickHouseParserListener) ExitTopClause(ctx *TopClauseContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseClickHouseParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseClickHouseParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterArrayJoinClause is called when production arrayJoinClause is entered.
func (s *BaseClickHouseParserListener) EnterArrayJoinClause(ctx *ArrayJoinClauseContext) {}

// ExitArrayJoinClause is called when production arrayJoinClause is exited.
func (s *BaseClickHouseParserListener) ExitArrayJoinClause(ctx *ArrayJoinClauseContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseClickHouseParserListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseClickHouseParserListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterPrewhereClause is called when production prewhereClause is entered.
func (s *BaseClickHouseParserListener) EnterPrewhereClause(ctx *PrewhereClauseContext) {}

// ExitPrewhereClause is called when production prewhereClause is exited.
func (s *BaseClickHouseParserListener) ExitPrewhereClause(ctx *PrewhereClauseContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseClickHouseParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseClickHouseParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseClickHouseParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseClickHouseParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseClickHouseParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseClickHouseParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseClickHouseParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseClickHouseParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterProjectionOrderByClause is called when production projectionOrderByClause is entered.
func (s *BaseClickHouseParserListener) EnterProjectionOrderByClause(ctx *ProjectionOrderByClauseContext) {
}

// ExitProjectionOrderByClause is called when production projectionOrderByClause is exited.
func (s *BaseClickHouseParserListener) ExitProjectionOrderByClause(ctx *ProjectionOrderByClauseContext) {
}

// EnterLimitByClause is called when production limitByClause is entered.
func (s *BaseClickHouseParserListener) EnterLimitByClause(ctx *LimitByClauseContext) {}

// ExitLimitByClause is called when production limitByClause is exited.
func (s *BaseClickHouseParserListener) ExitLimitByClause(ctx *LimitByClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseClickHouseParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseClickHouseParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterSettingsClause is called when production settingsClause is entered.
func (s *BaseClickHouseParserListener) EnterSettingsClause(ctx *SettingsClauseContext) {}

// ExitSettingsClause is called when production settingsClause is exited.
func (s *BaseClickHouseParserListener) ExitSettingsClause(ctx *SettingsClauseContext) {}

// EnterJoinExprOp is called when production JoinExprOp is entered.
func (s *BaseClickHouseParserListener) EnterJoinExprOp(ctx *JoinExprOpContext) {}

// ExitJoinExprOp is called when production JoinExprOp is exited.
func (s *BaseClickHouseParserListener) ExitJoinExprOp(ctx *JoinExprOpContext) {}

// EnterJoinExprTable is called when production JoinExprTable is entered.
func (s *BaseClickHouseParserListener) EnterJoinExprTable(ctx *JoinExprTableContext) {}

// ExitJoinExprTable is called when production JoinExprTable is exited.
func (s *BaseClickHouseParserListener) ExitJoinExprTable(ctx *JoinExprTableContext) {}

// EnterJoinExprParens is called when production JoinExprParens is entered.
func (s *BaseClickHouseParserListener) EnterJoinExprParens(ctx *JoinExprParensContext) {}

// ExitJoinExprParens is called when production JoinExprParens is exited.
func (s *BaseClickHouseParserListener) ExitJoinExprParens(ctx *JoinExprParensContext) {}

// EnterJoinExprCrossOp is called when production JoinExprCrossOp is entered.
func (s *BaseClickHouseParserListener) EnterJoinExprCrossOp(ctx *JoinExprCrossOpContext) {}

// ExitJoinExprCrossOp is called when production JoinExprCrossOp is exited.
func (s *BaseClickHouseParserListener) ExitJoinExprCrossOp(ctx *JoinExprCrossOpContext) {}

// EnterJoinOpInner is called when production JoinOpInner is entered.
func (s *BaseClickHouseParserListener) EnterJoinOpInner(ctx *JoinOpInnerContext) {}

// ExitJoinOpInner is called when production JoinOpInner is exited.
func (s *BaseClickHouseParserListener) ExitJoinOpInner(ctx *JoinOpInnerContext) {}

// EnterJoinOpLeftRight is called when production JoinOpLeftRight is entered.
func (s *BaseClickHouseParserListener) EnterJoinOpLeftRight(ctx *JoinOpLeftRightContext) {}

// ExitJoinOpLeftRight is called when production JoinOpLeftRight is exited.
func (s *BaseClickHouseParserListener) ExitJoinOpLeftRight(ctx *JoinOpLeftRightContext) {}

// EnterJoinOpFull is called when production JoinOpFull is entered.
func (s *BaseClickHouseParserListener) EnterJoinOpFull(ctx *JoinOpFullContext) {}

// ExitJoinOpFull is called when production JoinOpFull is exited.
func (s *BaseClickHouseParserListener) ExitJoinOpFull(ctx *JoinOpFullContext) {}

// EnterJoinOpCross is called when production joinOpCross is entered.
func (s *BaseClickHouseParserListener) EnterJoinOpCross(ctx *JoinOpCrossContext) {}

// ExitJoinOpCross is called when production joinOpCross is exited.
func (s *BaseClickHouseParserListener) ExitJoinOpCross(ctx *JoinOpCrossContext) {}

// EnterJoinConstraintClause is called when production joinConstraintClause is entered.
func (s *BaseClickHouseParserListener) EnterJoinConstraintClause(ctx *JoinConstraintClauseContext) {}

// ExitJoinConstraintClause is called when production joinConstraintClause is exited.
func (s *BaseClickHouseParserListener) ExitJoinConstraintClause(ctx *JoinConstraintClauseContext) {}

// EnterSampleClause is called when production sampleClause is entered.
func (s *BaseClickHouseParserListener) EnterSampleClause(ctx *SampleClauseContext) {}

// ExitSampleClause is called when production sampleClause is exited.
func (s *BaseClickHouseParserListener) ExitSampleClause(ctx *SampleClauseContext) {}

// EnterLimitExpr is called when production limitExpr is entered.
func (s *BaseClickHouseParserListener) EnterLimitExpr(ctx *LimitExprContext) {}

// ExitLimitExpr is called when production limitExpr is exited.
func (s *BaseClickHouseParserListener) ExitLimitExpr(ctx *LimitExprContext) {}

// EnterOrderExprList is called when production orderExprList is entered.
func (s *BaseClickHouseParserListener) EnterOrderExprList(ctx *OrderExprListContext) {}

// ExitOrderExprList is called when production orderExprList is exited.
func (s *BaseClickHouseParserListener) ExitOrderExprList(ctx *OrderExprListContext) {}

// EnterOrderExpr is called when production orderExpr is entered.
func (s *BaseClickHouseParserListener) EnterOrderExpr(ctx *OrderExprContext) {}

// ExitOrderExpr is called when production orderExpr is exited.
func (s *BaseClickHouseParserListener) ExitOrderExpr(ctx *OrderExprContext) {}

// EnterRatioExpr is called when production ratioExpr is entered.
func (s *BaseClickHouseParserListener) EnterRatioExpr(ctx *RatioExprContext) {}

// ExitRatioExpr is called when production ratioExpr is exited.
func (s *BaseClickHouseParserListener) ExitRatioExpr(ctx *RatioExprContext) {}

// EnterSettingExprList is called when production settingExprList is entered.
func (s *BaseClickHouseParserListener) EnterSettingExprList(ctx *SettingExprListContext) {}

// ExitSettingExprList is called when production settingExprList is exited.
func (s *BaseClickHouseParserListener) ExitSettingExprList(ctx *SettingExprListContext) {}

// EnterSettingExpr is called when production settingExpr is entered.
func (s *BaseClickHouseParserListener) EnterSettingExpr(ctx *SettingExprContext) {}

// ExitSettingExpr is called when production settingExpr is exited.
func (s *BaseClickHouseParserListener) ExitSettingExpr(ctx *SettingExprContext) {}

// EnterWindowExpr is called when production windowExpr is entered.
func (s *BaseClickHouseParserListener) EnterWindowExpr(ctx *WindowExprContext) {}

// ExitWindowExpr is called when production windowExpr is exited.
func (s *BaseClickHouseParserListener) ExitWindowExpr(ctx *WindowExprContext) {}

// EnterWinPartitionByClause is called when production winPartitionByClause is entered.
func (s *BaseClickHouseParserListener) EnterWinPartitionByClause(ctx *WinPartitionByClauseContext) {}

// ExitWinPartitionByClause is called when production winPartitionByClause is exited.
func (s *BaseClickHouseParserListener) ExitWinPartitionByClause(ctx *WinPartitionByClauseContext) {}

// EnterWinOrderByClause is called when production winOrderByClause is entered.
func (s *BaseClickHouseParserListener) EnterWinOrderByClause(ctx *WinOrderByClauseContext) {}

// ExitWinOrderByClause is called when production winOrderByClause is exited.
func (s *BaseClickHouseParserListener) ExitWinOrderByClause(ctx *WinOrderByClauseContext) {}

// EnterWinFrameClause is called when production winFrameClause is entered.
func (s *BaseClickHouseParserListener) EnterWinFrameClause(ctx *WinFrameClauseContext) {}

// ExitWinFrameClause is called when production winFrameClause is exited.
func (s *BaseClickHouseParserListener) ExitWinFrameClause(ctx *WinFrameClauseContext) {}

// EnterFrameStart is called when production frameStart is entered.
func (s *BaseClickHouseParserListener) EnterFrameStart(ctx *FrameStartContext) {}

// ExitFrameStart is called when production frameStart is exited.
func (s *BaseClickHouseParserListener) ExitFrameStart(ctx *FrameStartContext) {}

// EnterFrameBetween is called when production frameBetween is entered.
func (s *BaseClickHouseParserListener) EnterFrameBetween(ctx *FrameBetweenContext) {}

// ExitFrameBetween is called when production frameBetween is exited.
func (s *BaseClickHouseParserListener) ExitFrameBetween(ctx *FrameBetweenContext) {}

// EnterWinFrameBound is called when production winFrameBound is entered.
func (s *BaseClickHouseParserListener) EnterWinFrameBound(ctx *WinFrameBoundContext) {}

// ExitWinFrameBound is called when production winFrameBound is exited.
func (s *BaseClickHouseParserListener) ExitWinFrameBound(ctx *WinFrameBoundContext) {}

// EnterSetStmt is called when production setStmt is entered.
func (s *BaseClickHouseParserListener) EnterSetStmt(ctx *SetStmtContext) {}

// ExitSetStmt is called when production setStmt is exited.
func (s *BaseClickHouseParserListener) ExitSetStmt(ctx *SetStmtContext) {}

// EnterShowCreateDatabaseStmt is called when production showCreateDatabaseStmt is entered.
func (s *BaseClickHouseParserListener) EnterShowCreateDatabaseStmt(ctx *ShowCreateDatabaseStmtContext) {
}

// ExitShowCreateDatabaseStmt is called when production showCreateDatabaseStmt is exited.
func (s *BaseClickHouseParserListener) ExitShowCreateDatabaseStmt(ctx *ShowCreateDatabaseStmtContext) {
}

// EnterShowCreateDictionaryStmt is called when production showCreateDictionaryStmt is entered.
func (s *BaseClickHouseParserListener) EnterShowCreateDictionaryStmt(ctx *ShowCreateDictionaryStmtContext) {
}

// ExitShowCreateDictionaryStmt is called when production showCreateDictionaryStmt is exited.
func (s *BaseClickHouseParserListener) ExitShowCreateDictionaryStmt(ctx *ShowCreateDictionaryStmtContext) {
}

// EnterShowCreateTableStmt is called when production showCreateTableStmt is entered.
func (s *BaseClickHouseParserListener) EnterShowCreateTableStmt(ctx *ShowCreateTableStmtContext) {}

// ExitShowCreateTableStmt is called when production showCreateTableStmt is exited.
func (s *BaseClickHouseParserListener) ExitShowCreateTableStmt(ctx *ShowCreateTableStmtContext) {}

// EnterShowDatabasesStmt is called when production showDatabasesStmt is entered.
func (s *BaseClickHouseParserListener) EnterShowDatabasesStmt(ctx *ShowDatabasesStmtContext) {}

// ExitShowDatabasesStmt is called when production showDatabasesStmt is exited.
func (s *BaseClickHouseParserListener) ExitShowDatabasesStmt(ctx *ShowDatabasesStmtContext) {}

// EnterShowDictionariesStmt is called when production showDictionariesStmt is entered.
func (s *BaseClickHouseParserListener) EnterShowDictionariesStmt(ctx *ShowDictionariesStmtContext) {}

// ExitShowDictionariesStmt is called when production showDictionariesStmt is exited.
func (s *BaseClickHouseParserListener) ExitShowDictionariesStmt(ctx *ShowDictionariesStmtContext) {}

// EnterShowTablesStmt is called when production showTablesStmt is entered.
func (s *BaseClickHouseParserListener) EnterShowTablesStmt(ctx *ShowTablesStmtContext) {}

// ExitShowTablesStmt is called when production showTablesStmt is exited.
func (s *BaseClickHouseParserListener) ExitShowTablesStmt(ctx *ShowTablesStmtContext) {}

// EnterSystemStmt is called when production systemStmt is entered.
func (s *BaseClickHouseParserListener) EnterSystemStmt(ctx *SystemStmtContext) {}

// ExitSystemStmt is called when production systemStmt is exited.
func (s *BaseClickHouseParserListener) ExitSystemStmt(ctx *SystemStmtContext) {}

// EnterTruncateStmt is called when production truncateStmt is entered.
func (s *BaseClickHouseParserListener) EnterTruncateStmt(ctx *TruncateStmtContext) {}

// ExitTruncateStmt is called when production truncateStmt is exited.
func (s *BaseClickHouseParserListener) ExitTruncateStmt(ctx *TruncateStmtContext) {}

// EnterUseStmt is called when production useStmt is entered.
func (s *BaseClickHouseParserListener) EnterUseStmt(ctx *UseStmtContext) {}

// ExitUseStmt is called when production useStmt is exited.
func (s *BaseClickHouseParserListener) ExitUseStmt(ctx *UseStmtContext) {}

// EnterWatchStmt is called when production watchStmt is entered.
func (s *BaseClickHouseParserListener) EnterWatchStmt(ctx *WatchStmtContext) {}

// ExitWatchStmt is called when production watchStmt is exited.
func (s *BaseClickHouseParserListener) ExitWatchStmt(ctx *WatchStmtContext) {}

// EnterColumnTypeExprSimple is called when production ColumnTypeExprSimple is entered.
func (s *BaseClickHouseParserListener) EnterColumnTypeExprSimple(ctx *ColumnTypeExprSimpleContext) {}

// ExitColumnTypeExprSimple is called when production ColumnTypeExprSimple is exited.
func (s *BaseClickHouseParserListener) ExitColumnTypeExprSimple(ctx *ColumnTypeExprSimpleContext) {}

// EnterColumnTypeExprNested is called when production ColumnTypeExprNested is entered.
func (s *BaseClickHouseParserListener) EnterColumnTypeExprNested(ctx *ColumnTypeExprNestedContext) {}

// ExitColumnTypeExprNested is called when production ColumnTypeExprNested is exited.
func (s *BaseClickHouseParserListener) ExitColumnTypeExprNested(ctx *ColumnTypeExprNestedContext) {}

// EnterColumnTypeExprEnum is called when production ColumnTypeExprEnum is entered.
func (s *BaseClickHouseParserListener) EnterColumnTypeExprEnum(ctx *ColumnTypeExprEnumContext) {}

// ExitColumnTypeExprEnum is called when production ColumnTypeExprEnum is exited.
func (s *BaseClickHouseParserListener) ExitColumnTypeExprEnum(ctx *ColumnTypeExprEnumContext) {}

// EnterColumnTypeExprComplex is called when production ColumnTypeExprComplex is entered.
func (s *BaseClickHouseParserListener) EnterColumnTypeExprComplex(ctx *ColumnTypeExprComplexContext) {
}

// ExitColumnTypeExprComplex is called when production ColumnTypeExprComplex is exited.
func (s *BaseClickHouseParserListener) ExitColumnTypeExprComplex(ctx *ColumnTypeExprComplexContext) {}

// EnterColumnTypeExprParam is called when production ColumnTypeExprParam is entered.
func (s *BaseClickHouseParserListener) EnterColumnTypeExprParam(ctx *ColumnTypeExprParamContext) {}

// ExitColumnTypeExprParam is called when production ColumnTypeExprParam is exited.
func (s *BaseClickHouseParserListener) ExitColumnTypeExprParam(ctx *ColumnTypeExprParamContext) {}

// EnterColumnExprList is called when production columnExprList is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprList(ctx *ColumnExprListContext) {}

// ExitColumnExprList is called when production columnExprList is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprList(ctx *ColumnExprListContext) {}

// EnterColumnsExprAsterisk is called when production ColumnsExprAsterisk is entered.
func (s *BaseClickHouseParserListener) EnterColumnsExprAsterisk(ctx *ColumnsExprAsteriskContext) {}

// ExitColumnsExprAsterisk is called when production ColumnsExprAsterisk is exited.
func (s *BaseClickHouseParserListener) ExitColumnsExprAsterisk(ctx *ColumnsExprAsteriskContext) {}

// EnterColumnsExprSubquery is called when production ColumnsExprSubquery is entered.
func (s *BaseClickHouseParserListener) EnterColumnsExprSubquery(ctx *ColumnsExprSubqueryContext) {}

// ExitColumnsExprSubquery is called when production ColumnsExprSubquery is exited.
func (s *BaseClickHouseParserListener) ExitColumnsExprSubquery(ctx *ColumnsExprSubqueryContext) {}

// EnterColumnsExprColumn is called when production ColumnsExprColumn is entered.
func (s *BaseClickHouseParserListener) EnterColumnsExprColumn(ctx *ColumnsExprColumnContext) {}

// ExitColumnsExprColumn is called when production ColumnsExprColumn is exited.
func (s *BaseClickHouseParserListener) ExitColumnsExprColumn(ctx *ColumnsExprColumnContext) {}

// EnterColumnExprTernaryOp is called when production ColumnExprTernaryOp is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprTernaryOp(ctx *ColumnExprTernaryOpContext) {}

// ExitColumnExprTernaryOp is called when production ColumnExprTernaryOp is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprTernaryOp(ctx *ColumnExprTernaryOpContext) {}

// EnterColumnExprAlias is called when production ColumnExprAlias is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprAlias(ctx *ColumnExprAliasContext) {}

// ExitColumnExprAlias is called when production ColumnExprAlias is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprAlias(ctx *ColumnExprAliasContext) {}

// EnterColumnExprExtract is called when production ColumnExprExtract is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprExtract(ctx *ColumnExprExtractContext) {}

// ExitColumnExprExtract is called when production ColumnExprExtract is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprExtract(ctx *ColumnExprExtractContext) {}

// EnterColumnExprNegate is called when production ColumnExprNegate is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprNegate(ctx *ColumnExprNegateContext) {}

// ExitColumnExprNegate is called when production ColumnExprNegate is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprNegate(ctx *ColumnExprNegateContext) {}

// EnterColumnExprSubquery is called when production ColumnExprSubquery is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprSubquery(ctx *ColumnExprSubqueryContext) {}

// ExitColumnExprSubquery is called when production ColumnExprSubquery is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprSubquery(ctx *ColumnExprSubqueryContext) {}

// EnterColumnExprLiteral is called when production ColumnExprLiteral is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprLiteral(ctx *ColumnExprLiteralContext) {}

// ExitColumnExprLiteral is called when production ColumnExprLiteral is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprLiteral(ctx *ColumnExprLiteralContext) {}

// EnterColumnExprArray is called when production ColumnExprArray is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprArray(ctx *ColumnExprArrayContext) {}

// ExitColumnExprArray is called when production ColumnExprArray is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprArray(ctx *ColumnExprArrayContext) {}

// EnterColumnExprSubstring is called when production ColumnExprSubstring is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprSubstring(ctx *ColumnExprSubstringContext) {}

// ExitColumnExprSubstring is called when production ColumnExprSubstring is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprSubstring(ctx *ColumnExprSubstringContext) {}

// EnterColumnExprCast is called when production ColumnExprCast is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprCast(ctx *ColumnExprCastContext) {}

// ExitColumnExprCast is called when production ColumnExprCast is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprCast(ctx *ColumnExprCastContext) {}

// EnterColumnExprOr is called when production ColumnExprOr is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprOr(ctx *ColumnExprOrContext) {}

// ExitColumnExprOr is called when production ColumnExprOr is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprOr(ctx *ColumnExprOrContext) {}

// EnterColumnExprPrecedence1 is called when production ColumnExprPrecedence1 is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprPrecedence1(ctx *ColumnExprPrecedence1Context) {
}

// ExitColumnExprPrecedence1 is called when production ColumnExprPrecedence1 is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprPrecedence1(ctx *ColumnExprPrecedence1Context) {}

// EnterColumnExprPrecedence2 is called when production ColumnExprPrecedence2 is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprPrecedence2(ctx *ColumnExprPrecedence2Context) {
}

// ExitColumnExprPrecedence2 is called when production ColumnExprPrecedence2 is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprPrecedence2(ctx *ColumnExprPrecedence2Context) {}

// EnterColumnExprPrecedence3 is called when production ColumnExprPrecedence3 is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprPrecedence3(ctx *ColumnExprPrecedence3Context) {
}

// ExitColumnExprPrecedence3 is called when production ColumnExprPrecedence3 is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprPrecedence3(ctx *ColumnExprPrecedence3Context) {}

// EnterColumnExprInterval is called when production ColumnExprInterval is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprInterval(ctx *ColumnExprIntervalContext) {}

// ExitColumnExprInterval is called when production ColumnExprInterval is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprInterval(ctx *ColumnExprIntervalContext) {}

// EnterColumnExprIsNull is called when production ColumnExprIsNull is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprIsNull(ctx *ColumnExprIsNullContext) {}

// ExitColumnExprIsNull is called when production ColumnExprIsNull is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprIsNull(ctx *ColumnExprIsNullContext) {}

// EnterColumnExprWinFunctionTarget is called when production ColumnExprWinFunctionTarget is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprWinFunctionTarget(ctx *ColumnExprWinFunctionTargetContext) {
}

// ExitColumnExprWinFunctionTarget is called when production ColumnExprWinFunctionTarget is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprWinFunctionTarget(ctx *ColumnExprWinFunctionTargetContext) {
}

// EnterColumnExprTrim is called when production ColumnExprTrim is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprTrim(ctx *ColumnExprTrimContext) {}

// ExitColumnExprTrim is called when production ColumnExprTrim is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprTrim(ctx *ColumnExprTrimContext) {}

// EnterColumnExprTuple is called when production ColumnExprTuple is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprTuple(ctx *ColumnExprTupleContext) {}

// ExitColumnExprTuple is called when production ColumnExprTuple is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprTuple(ctx *ColumnExprTupleContext) {}

// EnterColumnExprArrayAccess is called when production ColumnExprArrayAccess is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprArrayAccess(ctx *ColumnExprArrayAccessContext) {
}

// ExitColumnExprArrayAccess is called when production ColumnExprArrayAccess is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprArrayAccess(ctx *ColumnExprArrayAccessContext) {}

// EnterColumnExprBetween is called when production ColumnExprBetween is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprBetween(ctx *ColumnExprBetweenContext) {}

// ExitColumnExprBetween is called when production ColumnExprBetween is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprBetween(ctx *ColumnExprBetweenContext) {}

// EnterColumnExprParens is called when production ColumnExprParens is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprParens(ctx *ColumnExprParensContext) {}

// ExitColumnExprParens is called when production ColumnExprParens is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprParens(ctx *ColumnExprParensContext) {}

// EnterColumnExprTimestamp is called when production ColumnExprTimestamp is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprTimestamp(ctx *ColumnExprTimestampContext) {}

// ExitColumnExprTimestamp is called when production ColumnExprTimestamp is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprTimestamp(ctx *ColumnExprTimestampContext) {}

// EnterColumnExprAnd is called when production ColumnExprAnd is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprAnd(ctx *ColumnExprAndContext) {}

// ExitColumnExprAnd is called when production ColumnExprAnd is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprAnd(ctx *ColumnExprAndContext) {}

// EnterColumnExprTupleAccess is called when production ColumnExprTupleAccess is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprTupleAccess(ctx *ColumnExprTupleAccessContext) {
}

// ExitColumnExprTupleAccess is called when production ColumnExprTupleAccess is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprTupleAccess(ctx *ColumnExprTupleAccessContext) {}

// EnterColumnExprCase is called when production ColumnExprCase is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprCase(ctx *ColumnExprCaseContext) {}

// ExitColumnExprCase is called when production ColumnExprCase is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprCase(ctx *ColumnExprCaseContext) {}

// EnterColumnExprDate is called when production ColumnExprDate is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprDate(ctx *ColumnExprDateContext) {}

// ExitColumnExprDate is called when production ColumnExprDate is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprDate(ctx *ColumnExprDateContext) {}

// EnterColumnExprNot is called when production ColumnExprNot is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprNot(ctx *ColumnExprNotContext) {}

// ExitColumnExprNot is called when production ColumnExprNot is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprNot(ctx *ColumnExprNotContext) {}

// EnterColumnExprWinFunction is called when production ColumnExprWinFunction is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprWinFunction(ctx *ColumnExprWinFunctionContext) {
}

// ExitColumnExprWinFunction is called when production ColumnExprWinFunction is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprWinFunction(ctx *ColumnExprWinFunctionContext) {}

// EnterColumnExprIdentifier is called when production ColumnExprIdentifier is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprIdentifier(ctx *ColumnExprIdentifierContext) {}

// ExitColumnExprIdentifier is called when production ColumnExprIdentifier is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprIdentifier(ctx *ColumnExprIdentifierContext) {}

// EnterColumnExprFunction is called when production ColumnExprFunction is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprFunction(ctx *ColumnExprFunctionContext) {}

// ExitColumnExprFunction is called when production ColumnExprFunction is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprFunction(ctx *ColumnExprFunctionContext) {}

// EnterColumnExprAsterisk is called when production ColumnExprAsterisk is entered.
func (s *BaseClickHouseParserListener) EnterColumnExprAsterisk(ctx *ColumnExprAsteriskContext) {}

// ExitColumnExprAsterisk is called when production ColumnExprAsterisk is exited.
func (s *BaseClickHouseParserListener) ExitColumnExprAsterisk(ctx *ColumnExprAsteriskContext) {}

// EnterColumnArgList is called when production columnArgList is entered.
func (s *BaseClickHouseParserListener) EnterColumnArgList(ctx *ColumnArgListContext) {}

// ExitColumnArgList is called when production columnArgList is exited.
func (s *BaseClickHouseParserListener) ExitColumnArgList(ctx *ColumnArgListContext) {}

// EnterColumnArgExpr is called when production columnArgExpr is entered.
func (s *BaseClickHouseParserListener) EnterColumnArgExpr(ctx *ColumnArgExprContext) {}

// ExitColumnArgExpr is called when production columnArgExpr is exited.
func (s *BaseClickHouseParserListener) ExitColumnArgExpr(ctx *ColumnArgExprContext) {}

// EnterColumnLambdaExpr is called when production columnLambdaExpr is entered.
func (s *BaseClickHouseParserListener) EnterColumnLambdaExpr(ctx *ColumnLambdaExprContext) {}

// ExitColumnLambdaExpr is called when production columnLambdaExpr is exited.
func (s *BaseClickHouseParserListener) ExitColumnLambdaExpr(ctx *ColumnLambdaExprContext) {}

// EnterColumnIdentifier is called when production columnIdentifier is entered.
func (s *BaseClickHouseParserListener) EnterColumnIdentifier(ctx *ColumnIdentifierContext) {}

// ExitColumnIdentifier is called when production columnIdentifier is exited.
func (s *BaseClickHouseParserListener) ExitColumnIdentifier(ctx *ColumnIdentifierContext) {}

// EnterNestedIdentifier is called when production nestedIdentifier is entered.
func (s *BaseClickHouseParserListener) EnterNestedIdentifier(ctx *NestedIdentifierContext) {}

// ExitNestedIdentifier is called when production nestedIdentifier is exited.
func (s *BaseClickHouseParserListener) ExitNestedIdentifier(ctx *NestedIdentifierContext) {}

// EnterEmptyListExpr is called when production emptyListExpr is entered.
func (s *BaseClickHouseParserListener) EnterEmptyListExpr(ctx *EmptyListExprContext) {}

// ExitEmptyListExpr is called when production emptyListExpr is exited.
func (s *BaseClickHouseParserListener) ExitEmptyListExpr(ctx *EmptyListExprContext) {}

// EnterTableExprIdentifier is called when production TableExprIdentifier is entered.
func (s *BaseClickHouseParserListener) EnterTableExprIdentifier(ctx *TableExprIdentifierContext) {}

// ExitTableExprIdentifier is called when production TableExprIdentifier is exited.
func (s *BaseClickHouseParserListener) ExitTableExprIdentifier(ctx *TableExprIdentifierContext) {}

// EnterTableExprSubquery is called when production TableExprSubquery is entered.
func (s *BaseClickHouseParserListener) EnterTableExprSubquery(ctx *TableExprSubqueryContext) {}

// ExitTableExprSubquery is called when production TableExprSubquery is exited.
func (s *BaseClickHouseParserListener) ExitTableExprSubquery(ctx *TableExprSubqueryContext) {}

// EnterTableExprAlias is called when production TableExprAlias is entered.
func (s *BaseClickHouseParserListener) EnterTableExprAlias(ctx *TableExprAliasContext) {}

// ExitTableExprAlias is called when production TableExprAlias is exited.
func (s *BaseClickHouseParserListener) ExitTableExprAlias(ctx *TableExprAliasContext) {}

// EnterTableExprFunction is called when production TableExprFunction is entered.
func (s *BaseClickHouseParserListener) EnterTableExprFunction(ctx *TableExprFunctionContext) {}

// ExitTableExprFunction is called when production TableExprFunction is exited.
func (s *BaseClickHouseParserListener) ExitTableExprFunction(ctx *TableExprFunctionContext) {}

// EnterTableFunctionExpr is called when production tableFunctionExpr is entered.
func (s *BaseClickHouseParserListener) EnterTableFunctionExpr(ctx *TableFunctionExprContext) {}

// ExitTableFunctionExpr is called when production tableFunctionExpr is exited.
func (s *BaseClickHouseParserListener) ExitTableFunctionExpr(ctx *TableFunctionExprContext) {}

// EnterTableIdentifier is called when production tableIdentifier is entered.
func (s *BaseClickHouseParserListener) EnterTableIdentifier(ctx *TableIdentifierContext) {}

// ExitTableIdentifier is called when production tableIdentifier is exited.
func (s *BaseClickHouseParserListener) ExitTableIdentifier(ctx *TableIdentifierContext) {}

// EnterTableArgList is called when production tableArgList is entered.
func (s *BaseClickHouseParserListener) EnterTableArgList(ctx *TableArgListContext) {}

// ExitTableArgList is called when production tableArgList is exited.
func (s *BaseClickHouseParserListener) ExitTableArgList(ctx *TableArgListContext) {}

// EnterTableArgExpr is called when production tableArgExpr is entered.
func (s *BaseClickHouseParserListener) EnterTableArgExpr(ctx *TableArgExprContext) {}

// ExitTableArgExpr is called when production tableArgExpr is exited.
func (s *BaseClickHouseParserListener) ExitTableArgExpr(ctx *TableArgExprContext) {}

// EnterDatabaseIdentifier is called when production databaseIdentifier is entered.
func (s *BaseClickHouseParserListener) EnterDatabaseIdentifier(ctx *DatabaseIdentifierContext) {}

// ExitDatabaseIdentifier is called when production databaseIdentifier is exited.
func (s *BaseClickHouseParserListener) ExitDatabaseIdentifier(ctx *DatabaseIdentifierContext) {}

// EnterFloatingLiteral is called when production floatingLiteral is entered.
func (s *BaseClickHouseParserListener) EnterFloatingLiteral(ctx *FloatingLiteralContext) {}

// ExitFloatingLiteral is called when production floatingLiteral is exited.
func (s *BaseClickHouseParserListener) ExitFloatingLiteral(ctx *FloatingLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseClickHouseParserListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseClickHouseParserListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseClickHouseParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseClickHouseParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseClickHouseParserListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseClickHouseParserListener) ExitInterval(ctx *IntervalContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseClickHouseParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseClickHouseParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterKeywordForAlias is called when production keywordForAlias is entered.
func (s *BaseClickHouseParserListener) EnterKeywordForAlias(ctx *KeywordForAliasContext) {}

// ExitKeywordForAlias is called when production keywordForAlias is exited.
func (s *BaseClickHouseParserListener) ExitKeywordForAlias(ctx *KeywordForAliasContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseClickHouseParserListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseClickHouseParserListener) ExitAlias(ctx *AliasContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseClickHouseParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseClickHouseParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifierOrNull is called when production identifierOrNull is entered.
func (s *BaseClickHouseParserListener) EnterIdentifierOrNull(ctx *IdentifierOrNullContext) {}

// ExitIdentifierOrNull is called when production identifierOrNull is exited.
func (s *BaseClickHouseParserListener) ExitIdentifierOrNull(ctx *IdentifierOrNullContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseClickHouseParserListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseClickHouseParserListener) ExitEnumValue(ctx *EnumValueContext) {}
