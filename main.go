package main

import "goretech/comparison/sqlInjections"

func main() {
	// 10er examples
	sqlInjections.UnprepStmtArg()
	sqlInjections.UnprepStmtConst()
	sqlInjections.UnprepStmtConstDecl()
	sqlInjections.UnprepStmtConstDeclUntypedString()
	sqlInjections.UnprepStmtConstDeclTypedString()
	sqlInjections.UnprepStmtConstIsConstRetVal()
	sqlInjections.UnprepStmtConstIsConstRetValFlowSens()
	sqlInjections.UnprepStmtConstDeclChangeAfterQueryFunCall()

	// 20er examples
	sqlInjections.UnprepStmtConstDeclUntypedStringQueryRow()
	sqlInjections.UnprepStmtDeclStringQueryRow()
	sqlInjections.UnprepStmtFormatedString()
	sqlInjections.UnprepStmtConstDeclSmallKeywords()
}