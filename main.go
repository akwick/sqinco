package main

import "github.com/akwick/sqinco/sqlInjections"

func main() {
	// 10er examples
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
	sqlInjections.WrappedFunc()
	sqlInjections.NotQuerry()
}
