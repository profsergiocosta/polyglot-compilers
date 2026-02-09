// Generated from Jack.g4 by ANTLR 4.7.

package parser // Jack

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJackListener is a complete listener for a parse tree produced by JackParser.
type BaseJackListener struct{}

var _ JackListener = &BaseJackListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJackListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJackListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJackListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJackListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterClassdef is called when production classdef is entered.
func (s *BaseJackListener) EnterClassdef(ctx *ClassdefContext) {}

// ExitClassdef is called when production classdef is exited.
func (s *BaseJackListener) ExitClassdef(ctx *ClassdefContext) {}

// EnterClassvardec is called when production classvardec is entered.
func (s *BaseJackListener) EnterClassvardec(ctx *ClassvardecContext) {}

// ExitClassvardec is called when production classvardec is exited.
func (s *BaseJackListener) ExitClassvardec(ctx *ClassvardecContext) {}

// EnterAtype is called when production atype is entered.
func (s *BaseJackListener) EnterAtype(ctx *AtypeContext) {}

// ExitAtype is called when production atype is exited.
func (s *BaseJackListener) ExitAtype(ctx *AtypeContext) {}

// EnterSubrotinedec is called when production subrotinedec is entered.
func (s *BaseJackListener) EnterSubrotinedec(ctx *SubrotinedecContext) {}

// ExitSubrotinedec is called when production subrotinedec is exited.
func (s *BaseJackListener) ExitSubrotinedec(ctx *SubrotinedecContext) {}

// EnterSubroutinetype is called when production subroutinetype is entered.
func (s *BaseJackListener) EnterSubroutinetype(ctx *SubroutinetypeContext) {}

// ExitSubroutinetype is called when production subroutinetype is exited.
func (s *BaseJackListener) ExitSubroutinetype(ctx *SubroutinetypeContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseJackListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseJackListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseJackListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseJackListener) ExitParameter(ctx *ParameterContext) {}

// EnterSubroutinebody is called when production subroutinebody is entered.
func (s *BaseJackListener) EnterSubroutinebody(ctx *SubroutinebodyContext) {}

// ExitSubroutinebody is called when production subroutinebody is exited.
func (s *BaseJackListener) ExitSubroutinebody(ctx *SubroutinebodyContext) {}

// EnterVardecs is called when production vardecs is entered.
func (s *BaseJackListener) EnterVardecs(ctx *VardecsContext) {}

// ExitVardecs is called when production vardecs is exited.
func (s *BaseJackListener) ExitVardecs(ctx *VardecsContext) {}

// EnterVardec is called when production vardec is entered.
func (s *BaseJackListener) EnterVardec(ctx *VardecContext) {}

// ExitVardec is called when production vardec is exited.
func (s *BaseJackListener) ExitVardec(ctx *VardecContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseJackListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseJackListener) ExitStatement(ctx *StatementContext) {}

// EnterLetStatement is called when production letStatement is entered.
func (s *BaseJackListener) EnterLetStatement(ctx *LetStatementContext) {}

// ExitLetStatement is called when production letStatement is exited.
func (s *BaseJackListener) ExitLetStatement(ctx *LetStatementContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BaseJackListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BaseJackListener) ExitLvalue(ctx *LvalueContext) {}

// EnterRvalue is called when production rvalue is entered.
func (s *BaseJackListener) EnterRvalue(ctx *RvalueContext) {}

// ExitRvalue is called when production rvalue is exited.
func (s *BaseJackListener) ExitRvalue(ctx *RvalueContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseJackListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseJackListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterIfBlock is called when production ifBlock is entered.
func (s *BaseJackListener) EnterIfBlock(ctx *IfBlockContext) {}

// ExitIfBlock is called when production ifBlock is exited.
func (s *BaseJackListener) ExitIfBlock(ctx *IfBlockContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseJackListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseJackListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseJackListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseJackListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterWhileBlock is called when production whileBlock is entered.
func (s *BaseJackListener) EnterWhileBlock(ctx *WhileBlockContext) {}

// ExitWhileBlock is called when production whileBlock is exited.
func (s *BaseJackListener) ExitWhileBlock(ctx *WhileBlockContext) {}

// EnterDoStatement is called when production doStatement is entered.
func (s *BaseJackListener) EnterDoStatement(ctx *DoStatementContext) {}

// ExitDoStatement is called when production doStatement is exited.
func (s *BaseJackListener) ExitDoStatement(ctx *DoStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseJackListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseJackListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterSubroutinecall is called when production subroutinecall is entered.
func (s *BaseJackListener) EnterSubroutinecall(ctx *SubroutinecallContext) {}

// ExitSubroutinecall is called when production subroutinecall is exited.
func (s *BaseJackListener) ExitSubroutinecall(ctx *SubroutinecallContext) {}

// EnterClassObject is called when production classObject is entered.
func (s *BaseJackListener) EnterClassObject(ctx *ClassObjectContext) {}

// ExitClassObject is called when production classObject is exited.
func (s *BaseJackListener) ExitClassObject(ctx *ClassObjectContext) {}

// EnterExpressionlist is called when production expressionlist is entered.
func (s *BaseJackListener) EnterExpressionlist(ctx *ExpressionlistContext) {}

// ExitExpressionlist is called when production expressionlist is exited.
func (s *BaseJackListener) ExitExpressionlist(ctx *ExpressionlistContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJackListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJackListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBinaryoperation is called when production binaryoperation is entered.
func (s *BaseJackListener) EnterBinaryoperation(ctx *BinaryoperationContext) {}

// ExitBinaryoperation is called when production binaryoperation is exited.
func (s *BaseJackListener) ExitBinaryoperation(ctx *BinaryoperationContext) {}

// EnterIntegerTerm is called when production IntegerTerm is entered.
func (s *BaseJackListener) EnterIntegerTerm(ctx *IntegerTermContext) {}

// ExitIntegerTerm is called when production IntegerTerm is exited.
func (s *BaseJackListener) ExitIntegerTerm(ctx *IntegerTermContext) {}

// EnterStringTerm is called when production StringTerm is entered.
func (s *BaseJackListener) EnterStringTerm(ctx *StringTermContext) {}

// ExitStringTerm is called when production StringTerm is exited.
func (s *BaseJackListener) ExitStringTerm(ctx *StringTermContext) {}

// EnterKeywordTerm is called when production KeywordTerm is entered.
func (s *BaseJackListener) EnterKeywordTerm(ctx *KeywordTermContext) {}

// ExitKeywordTerm is called when production KeywordTerm is exited.
func (s *BaseJackListener) ExitKeywordTerm(ctx *KeywordTermContext) {}

// EnterVarnameTerm is called when production VarnameTerm is entered.
func (s *BaseJackListener) EnterVarnameTerm(ctx *VarnameTermContext) {}

// ExitVarnameTerm is called when production VarnameTerm is exited.
func (s *BaseJackListener) ExitVarnameTerm(ctx *VarnameTermContext) {}

// EnterArrayTerm is called when production ArrayTerm is entered.
func (s *BaseJackListener) EnterArrayTerm(ctx *ArrayTermContext) {}

// ExitArrayTerm is called when production ArrayTerm is exited.
func (s *BaseJackListener) ExitArrayTerm(ctx *ArrayTermContext) {}

// EnterSubroutineTerm is called when production SubroutineTerm is entered.
func (s *BaseJackListener) EnterSubroutineTerm(ctx *SubroutineTermContext) {}

// ExitSubroutineTerm is called when production SubroutineTerm is exited.
func (s *BaseJackListener) ExitSubroutineTerm(ctx *SubroutineTermContext) {}

// EnterParsTerm is called when production ParsTerm is entered.
func (s *BaseJackListener) EnterParsTerm(ctx *ParsTermContext) {}

// ExitParsTerm is called when production ParsTerm is exited.
func (s *BaseJackListener) ExitParsTerm(ctx *ParsTermContext) {}

// EnterUnaryopTerm is called when production unaryopTerm is entered.
func (s *BaseJackListener) EnterUnaryopTerm(ctx *UnaryopTermContext) {}

// ExitUnaryopTerm is called when production unaryopTerm is exited.
func (s *BaseJackListener) ExitUnaryopTerm(ctx *UnaryopTermContext) {}

// EnterClassname is called when production classname is entered.
func (s *BaseJackListener) EnterClassname(ctx *ClassnameContext) {}

// ExitClassname is called when production classname is exited.
func (s *BaseJackListener) ExitClassname(ctx *ClassnameContext) {}

// EnterVarname is called when production varname is entered.
func (s *BaseJackListener) EnterVarname(ctx *VarnameContext) {}

// ExitVarname is called when production varname is exited.
func (s *BaseJackListener) ExitVarname(ctx *VarnameContext) {}

// EnterSubroutinename is called when production subroutinename is entered.
func (s *BaseJackListener) EnterSubroutinename(ctx *SubroutinenameContext) {}

// ExitSubroutinename is called when production subroutinename is exited.
func (s *BaseJackListener) ExitSubroutinename(ctx *SubroutinenameContext) {}
