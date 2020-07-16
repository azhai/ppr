package freefloating

import (
	"strings"

	"github.com/z7zmey/php-parser/position"
)

type StringType int

const (
	WhiteSpaceType StringType = iota
	CommentType
	TokenType
)

type Position int

//go:generate stringer -type=Position -output ./position_string.go
const (
	Start Position = iota
	End
	Slash
	Colon
	SemiColon
	AltEnd
	Dollar
	Ampersand
	Name
	Prefix
	Key
	Var
	UseType
	ReturnType
	OptionalType
	CaseSeparator
	LexicalVars
	Params
	Ref
	Cast
	Expr
	InitExpr
	CondExpr
	IncExpr
	True
	Cond

	HaltCompiller
	Namespace
	Static
	Class
	Use
	While
	For
	Switch
	Break
	Foreach
	Declare
	Label
	Finally
	List
	Default
	If
	ElseIf
	Else
	Variadic
	Function
	DoubleArrow
	Alias
	As
	Equal
	Exit
	Array
	Isset
	Empty
	Eval
	Echo
	Try
	Catch
	Unset

	Stmts
	VarList
	ConstList
	NameList
	ParamList
	ModifierList
	ArrayPairList
	CaseListStart
	CaseListEnd
	ArgumentList
	PropertyList
	ParameterList
	AdaptationList
	LexicalVarList
	UseDeclarationList

	OpenParenthesisToken
	CloseParenthesisToken
)

type String struct {
	StringType StringType
	Value      string
	Position   *position.Position
}

func (s String) IsEmpty() bool {
	return strings.TrimSpace(s.Value) == ""
}

// get spaces after the last "\n"
func (s String) GetLastSpaces() string {
	n := strings.LastIndex(s.Value, "\n")
	if n > 0 {
		return s.Value[n+1:]
	}
	return s.Value
}

type Collection map[Position][]String

func (c Collection) IsEmpty() bool {
	for _, v := range c {
		if len(v) > 0 {
			return false
		}
	}
	return true
}

func (c Collection) GetDocComment() (comment string) {
	var strs []String
	ok, has := false, false
	if strs, ok = c[Start]; !ok {
		return
	}
	for _, s := range strs {
		if s.StringType >= TokenType {
			continue
		}
		if s.StringType == CommentType {
			has = !s.IsEmpty()
			comment += s.Value
		} else if has {
			comment += s.Value
		}
	}
	if !has {
		comment = ""
	}
	return
}
