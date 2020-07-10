// Code generated by "stringer -type=Position -output ./position_string.go"; DO NOT EDIT.

package freefloating

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Start-0]
	_ = x[End-1]
	_ = x[Slash-2]
	_ = x[Colon-3]
	_ = x[SemiColon-4]
	_ = x[AltEnd-5]
	_ = x[Dollar-6]
	_ = x[Ampersand-7]
	_ = x[Name-8]
	_ = x[Prefix-9]
	_ = x[Key-10]
	_ = x[Var-11]
	_ = x[UseType-12]
	_ = x[ReturnType-13]
	_ = x[OptionalType-14]
	_ = x[CaseSeparator-15]
	_ = x[LexicalVars-16]
	_ = x[Params-17]
	_ = x[Ref-18]
	_ = x[Cast-19]
	_ = x[Expr-20]
	_ = x[InitExpr-21]
	_ = x[CondExpr-22]
	_ = x[IncExpr-23]
	_ = x[True-24]
	_ = x[Cond-25]
	_ = x[HaltCompiller-26]
	_ = x[Namespace-27]
	_ = x[Static-28]
	_ = x[Class-29]
	_ = x[Use-30]
	_ = x[While-31]
	_ = x[For-32]
	_ = x[Switch-33]
	_ = x[Break-34]
	_ = x[Foreach-35]
	_ = x[Declare-36]
	_ = x[Label-37]
	_ = x[Finally-38]
	_ = x[List-39]
	_ = x[Default-40]
	_ = x[If-41]
	_ = x[ElseIf-42]
	_ = x[Else-43]
	_ = x[Variadic-44]
	_ = x[Function-45]
	_ = x[DoubleArrow-46]
	_ = x[Alias-47]
	_ = x[As-48]
	_ = x[Equal-49]
	_ = x[Exit-50]
	_ = x[Array-51]
	_ = x[Isset-52]
	_ = x[Empty-53]
	_ = x[Eval-54]
	_ = x[Echo-55]
	_ = x[Try-56]
	_ = x[Catch-57]
	_ = x[Unset-58]
	_ = x[Stmts-59]
	_ = x[VarList-60]
	_ = x[ConstList-61]
	_ = x[NameList-62]
	_ = x[ParamList-63]
	_ = x[ModifierList-64]
	_ = x[ArrayPairList-65]
	_ = x[CaseListStart-66]
	_ = x[CaseListEnd-67]
	_ = x[ArgumentList-68]
	_ = x[PropertyList-69]
	_ = x[ParameterList-70]
	_ = x[AdaptationList-71]
	_ = x[LexicalVarList-72]
	_ = x[UseDeclarationList-73]
	_ = x[OpenParenthesisToken-74]
	_ = x[CloseParenthesisToken-75]
}

const _Position_name = "StartEndSlashColonSemiColonAltEndDollarAmpersandNamePrefixKeyVarUseTypeReturnTypeOptionalTypeCaseSeparatorLexicalVarsParamsRefCastExprInitExprCondExprIncExprTrueCondHaltCompillerNamespaceStaticClassUseWhileForSwitchBreakForeachDeclareLabelFinallyListDefaultIfElseIfElseVariadicFunctionDoubleArrowAliasAsEqualExitArrayIssetEmptyEvalEchoTryCatchUnsetStmtsVarListConstListNameListParamListModifierListArrayPairListCaseListStartCaseListEndArgumentListPropertyListParameterListAdaptationListLexicalVarListUseDeclarationListOpenParenthesisTokenCloseParenthesisToken"

var _Position_index = [...]uint16{0, 5, 8, 13, 18, 27, 33, 39, 48, 52, 58, 61, 64, 71, 81, 93, 106, 117, 123, 126, 130, 134, 142, 150, 157, 161, 165, 178, 187, 193, 198, 201, 206, 209, 215, 220, 227, 234, 239, 246, 250, 257, 259, 265, 269, 277, 285, 296, 301, 303, 308, 312, 317, 322, 327, 331, 335, 338, 343, 348, 353, 360, 369, 377, 386, 398, 411, 424, 435, 447, 459, 472, 486, 500, 518, 538, 559}

func (i Position) String() string {
	if i < 0 || i >= Position(len(_Position_index)-1) {
		return "Position(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Position_name[_Position_index[i]:_Position_index[i+1]]
}
