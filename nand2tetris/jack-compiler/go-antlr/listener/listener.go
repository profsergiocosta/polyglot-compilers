package listener

import (
	"fmt"
	"strconv"

	"jackcompiler-antlr/parser"
	"jackcompiler-antlr/symboltable"
	"jackcompiler-antlr/vmwriter"
)

type JackListener struct {
	*parser.BaseJackListener

	st *symboltable.SymbolTable

	vm *vmwriter.VMWriter

	//context
	className    string
	functionName string

	// flow
	ifLabelNum    int
	whileLabelNum int

	whileLabelStack []int
	ifLabelStack    []int

	// call
	subroutineKind int
	numargs        int
	subroutineName string
}

func New(vmfilename string) *JackListener {

	s := &JackListener{}
	s.st = symboltable.NewSymbolTable()

	s.vm = vmwriter.New(vmfilename)

	s.ifLabelNum = 0
	s.whileLabelNum = 0

	return s
}

func (s *JackListener) EnterClassdef(ctx *parser.ClassdefContext) {
	s.className = ctx.Classname().GetText()
}

func (s *JackListener) EnterClassvardec(ctx *parser.ClassvardecContext) {

	var scope symboltable.SymbolScope

	scope = symboltable.STATIC

	if ctx.GetKind().GetTokenType() == parser.JackLexerFIELD {
		scope = symboltable.FIELD
	}

	ttype := ctx.Atype().GetText()

	varnames := ctx.AllVarname()
	for _, element := range varnames {
		s.st.Define(element.GetText(), ttype, scope)
	}

}

func (s *JackListener) EnterVardec(ctx *parser.VardecContext) {

	var scope symboltable.SymbolScope = symboltable.VAR

	ttype := ctx.Atype().GetText()

	varnames := ctx.AllVarname()
	for _, element := range varnames {
		s.st.Define(element.GetText(), ttype, scope)
	}
}

func (s *JackListener) EnterParameter(ctx *parser.ParameterContext) {
	var scope symboltable.SymbolScope = symboltable.ARG
	ttype := ctx.Atype().GetText()
	name := ctx.Varname().GetText()
	s.st.Define(name, ttype, scope)
}

func (s *JackListener) ExitVardecs(ctx *parser.VardecsContext) {
	nLocals := s.st.VarCount(symboltable.VAR)
	s.vm.WriteFunction(s.functionName, nLocals)

	switch s.subroutineKind {
	case parser.JackLexerMETHOD:
		s.vm.WritePush(vmwriter.ARG, 0)
		s.vm.WritePop(vmwriter.POINTER, 0)
	case parser.JackLexerCONSTRUCTOR:
		s.vm.WritePush(vmwriter.CONST, s.st.VarCount(symboltable.FIELD))
		s.vm.WriteCall("Memory.alloc", 1)
		s.vm.WritePop(vmwriter.POINTER, 0)
	}

}

func (s *JackListener) EnterSubrotinedec(ctx *parser.SubrotinedecContext) {

	s.st.StartSubroutine()

	s.functionName = s.className + "." + ctx.Subroutinename().GetText()
	s.subroutineKind = ctx.GetKind().GetTokenType()
	if ctx.GetKind().GetTokenType() == parser.JackLexerMETHOD {
		s.st.Define("this", s.className, symboltable.ARG)
	}

	s.ifLabelNum = 0
	s.whileLabelNum = 0

}

// ExitLvalue is called when production lvalue is exited.
func (s *JackListener) ExitLvalue(ctx *parser.LvalueContext) {
	sym := s.st.Resolve(ctx.GetIdent().GetText())
	if ctx.LBRACKET() != nil {
		s.vm.WritePush(scopeToSegment(sym.Scope), sym.Index)
		s.vm.WriteArithmetic(vmwriter.ADD)
	}
}

func (s *JackListener) ExitLetStatement(ctx *parser.LetStatementContext) {
	sym := s.st.Resolve(ctx.Lvalue().GetIdent().GetText())
	if ctx.Lvalue().GetChildCount() > 1 { // is array ?
		s.vm.WritePop(vmwriter.TEMP, 0)
		s.vm.WritePop(vmwriter.POINTER, 1)
		s.vm.WritePush(vmwriter.TEMP, 0)
		s.vm.WritePop(vmwriter.THAT, 0)
	} else {
		s.vm.WritePop(scopeToSegment(sym.Scope), sym.Index)
	}

}
func (s *JackListener) ExitDoStatement(ctx *parser.DoStatementContext) {
	s.vm.WritePop(vmwriter.TEMP, 0)
}

func (s *JackListener) ExitReturnStatement(ctx *parser.ReturnStatementContext) {
	if ctx.GetChildCount() > 2 { // has expression ?
		s.vm.WriteReturn()
	} else {
		s.vm.WritePush(vmwriter.CONST, 0)
		s.vm.WriteReturn()
	}
}

func (s *JackListener) ExitIfStatement(ctx *parser.IfStatementContext) {
	//s.ifLabelNum++
	s.ifLabelStack = s.ifLabelStack[:len(s.ifLabelStack)-1]
}

// EnterIfBlock is called when production ifBlock is entered.
func (s *JackListener) EnterIfBlock(ctx *parser.IfBlockContext) {
	labelTrue := fmt.Sprintf("IF_TRUE%d", s.ifLabelNum)
	labelFalse := fmt.Sprintf("IF_FALSE%d", s.ifLabelNum)

	s.vm.WriteIf(labelTrue)
	s.vm.WriteGoto(labelFalse)
	s.vm.WriteLabel(labelTrue)

	s.ifLabelStack = append(s.ifLabelStack, s.ifLabelNum)
	s.ifLabelNum++

}

// ExitIfBlock is called when production ifBlock is exited.
func (s *JackListener) ExitIfBlock(ctx *parser.IfBlockContext) {

	labelNum := s.ifLabelStack[len(s.ifLabelStack)-1]

	labelEnd := fmt.Sprintf("IF_END%d", labelNum)
	labelFalse := fmt.Sprintf("IF_FALSE%d", labelNum)

	if ctx.GetParent().GetChildCount() > 4 { // has else
		s.vm.WriteGoto(labelEnd)
	}

	s.vm.WriteLabel(labelFalse)
}

// ExitElseBlock is called when production elseBlock is exited.
func (s *JackListener) ExitElseBlock(ctx *parser.ElseBlockContext) {
	labelNum := s.ifLabelStack[len(s.ifLabelStack)-1]
	labelEnd := fmt.Sprintf("IF_END%d", labelNum)
	s.vm.WriteLabel(labelEnd)
}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *JackListener) EnterWhileStatement(ctx *parser.WhileStatementContext) {
	labelWhileExp := fmt.Sprintf("WHILE_EXP%d", s.whileLabelNum)
	s.vm.WriteLabel(labelWhileExp)
	s.whileLabelStack = append(s.whileLabelStack, s.whileLabelNum)
	s.whileLabelNum++
}

func (s *JackListener) ExitWhileStatement(ctx *parser.WhileStatementContext) {
	//s.whileLabelNum++
	// pop
	s.whileLabelStack = s.whileLabelStack[:len(s.whileLabelStack)-1]
}

// EnterWhileBlock is called when production whileBlock is entered.
func (s *JackListener) EnterWhileBlock(ctx *parser.WhileBlockContext) {
	labelNum := s.whileLabelStack[len(s.whileLabelStack)-1]
	labelWhileEnd := fmt.Sprintf("WHILE_END%d", labelNum)
	s.vm.WriteArithmetic(vmwriter.NOT)
	s.vm.WriteIf(labelWhileEnd)
}

// ExitWhileBlock is called when production whileBlock is exited.
func (s *JackListener) ExitWhileBlock(ctx *parser.WhileBlockContext) {
	labelNum := s.whileLabelStack[len(s.whileLabelStack)-1]
	labelWhileExp := fmt.Sprintf("WHILE_EXP%d", labelNum)
	labelWhileEnd := fmt.Sprintf("WHILE_END%d", labelNum)

	s.vm.WriteGoto(labelWhileExp)
	s.vm.WriteLabel(labelWhileEnd)
}

func (s *JackListener) EnterTerm(ctx *parser.TermContext) {

}

func (s *JackListener) EnterIntegerTerm(ctx *parser.IntegerTermContext) {
	s.vm.WritePush(vmwriter.CONST, asInt(ctx.GetText()))
}

// EnterKeywordTerm is called when production KeywordTerm is entered.
func (s *JackListener) EnterKeywordTerm(ctx *parser.KeywordTermContext) {
	switch ctx.GetKeywordConst().GetTokenType() {

	case parser.JackLexerTRUE:
		s.vm.WritePush(vmwriter.CONST, 0)
		s.vm.WriteArithmetic(vmwriter.NOT)

	case parser.JackLexerFALSE, parser.JackLexerNULL:
		s.vm.WritePush(vmwriter.CONST, 0)

	case parser.JackLexerTHIS:
		s.vm.WritePush(vmwriter.POINTER, 0)

	}

}

// EnterStringTerm is called when production StringTerm is entered.
func (s *JackListener) EnterStringTerm(ctx *parser.StringTermContext) {
	str := ctx.GetText()
	str = str[1 : len(str)-1]
	s.vm.WritePush(vmwriter.CONST, len(str))
	s.vm.WriteCall("String.new", 1)
	for i := 0; i < len(str); i++ {
		s.vm.WritePush(vmwriter.CONST, int(str[i]))
		s.vm.WriteCall("String.appendChar", 2)
	}
}

// EnterVarnameTerm is called when production VarnameTerm is entered.
func (s *JackListener) EnterVarnameTerm(ctx *parser.VarnameTermContext) {
	varname := ctx.GetText()
	sym := s.st.Resolve(varname)
	s.vm.WritePush(scopeToSegment(sym.Scope), sym.Index)
}

func (s *JackListener) ExitArrayTerm(ctx *parser.ArrayTermContext) {

	varname := ctx.GetIdent().GetText()

	sym := s.st.Resolve(varname)

	s.vm.WritePush(scopeToSegment(sym.Scope), sym.Index)
	s.vm.WriteArithmetic(vmwriter.ADD)

	s.vm.WritePop(vmwriter.POINTER, 1)
	s.vm.WritePush(vmwriter.THAT, 0)

}

func (s *JackListener) EnterSubroutinecall(ctx *parser.SubroutinecallContext) {
	s.numargs = 0
	s.subroutineName = ""
	if ctx.ClassObject() == nil { // é o metod da propria classe
		s.subroutineName = s.className + "."
		s.vm.WritePush(vmwriter.POINTER, 0)
		s.numargs = 1
	}
}

func (s *JackListener) ExitSubroutinecall(ctx *parser.SubroutinecallContext) {
	subName := s.subroutineName + ctx.Subroutinename().GetText()
	s.vm.WriteCall(subName, s.numargs)

}

// EnterClassObject is called when production classObject is entered.
func (s *JackListener) EnterClassObject(ctx *parser.ClassObjectContext) {
	sym, has := s.st.Lookup(ctx.ID().GetText())
	if has { // é um objeto

		s.vm.WritePush(scopeToSegment(sym.Scope), sym.Index)
		s.numargs = 1
		s.subroutineName = sym.Type + "."
	} else {
		s.subroutineName = ctx.ID().GetText() + "."
	}
}

// EnterExpressionlist is called when production expressionlist is entered.
func (s *JackListener) EnterExpressionlist(ctx *parser.ExpressionlistContext) {
	s.numargs += len(ctx.AllExpression())
}

// ExitUnaryopTerm is called when production unaryopTerm is exited.
func (s *JackListener) ExitUnaryopTerm(ctx *parser.UnaryopTermContext) {

	switch ctx.GetUnaryop().GetTokenType() {
	case parser.JackLexerMINUS:
		s.vm.WriteArithmetic(vmwriter.NEG)
	default:
		s.vm.WriteArithmetic(vmwriter.NOT)

	}
}

func (s *JackListener) ExitBinaryoperation(ctx *parser.BinaryoperationContext) {
	switch ctx.GetOperator().GetTokenType() {

	case parser.JackLexerPLUS:
		s.vm.WriteArithmetic(vmwriter.ADD)
	case parser.JackLexerMINUS:
		s.vm.WriteArithmetic(vmwriter.SUB)
	case parser.JackLexerASTERISK:
		s.vm.WriteCall("Math.multiply", 2)
	case parser.JackLexerSLASH:
		s.vm.WriteCall("Math.divide", 2)
	case parser.JackLexerAND:
		s.vm.WriteArithmetic(vmwriter.AND)
	case parser.JackLexerOR:
		s.vm.WriteArithmetic(vmwriter.OR)
	case parser.JackLexerLT:
		s.vm.WriteArithmetic(vmwriter.LT)
	case parser.JackLexerGT:
		s.vm.WriteArithmetic(vmwriter.GT)
	case parser.JackLexerEQ:
		s.vm.WriteArithmetic(vmwriter.EQ)
	case parser.JackLexerNOT:
		s.vm.WriteArithmetic(vmwriter.NOT)
	}

}

func asInt(s string) int {
	i1, err := strconv.Atoi(s)
	check(err)
	return i1
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func scopeToSegment(scope symboltable.SymbolScope) vmwriter.Segment {
	switch scope {
	case symboltable.STATIC:
		return vmwriter.STATIC
	case symboltable.FIELD:
		return vmwriter.THIS
	case symboltable.VAR:
		return vmwriter.LOCAL
	case symboltable.ARG:
		return vmwriter.ARG
	default:

	}
	panic("scope undefined")
}
