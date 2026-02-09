// Generated from Jack.g4 by ANTLR 4.7.

package parser // Jack

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JackListener is a complete listener for a parse tree produced by JackParser.
type JackListener interface {
	antlr.ParseTreeListener

	// EnterClassdef is called when entering the classdef production.
	EnterClassdef(c *ClassdefContext)

	// EnterClassvardec is called when entering the classvardec production.
	EnterClassvardec(c *ClassvardecContext)

	// EnterAtype is called when entering the atype production.
	EnterAtype(c *AtypeContext)

	// EnterSubrotinedec is called when entering the subrotinedec production.
	EnterSubrotinedec(c *SubrotinedecContext)

	// EnterSubroutinetype is called when entering the subroutinetype production.
	EnterSubroutinetype(c *SubroutinetypeContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterSubroutinebody is called when entering the subroutinebody production.
	EnterSubroutinebody(c *SubroutinebodyContext)

	// EnterVardecs is called when entering the vardecs production.
	EnterVardecs(c *VardecsContext)

	// EnterVardec is called when entering the vardec production.
	EnterVardec(c *VardecContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLetStatement is called when entering the letStatement production.
	EnterLetStatement(c *LetStatementContext)

	// EnterLvalue is called when entering the lvalue production.
	EnterLvalue(c *LvalueContext)

	// EnterRvalue is called when entering the rvalue production.
	EnterRvalue(c *RvalueContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterIfBlock is called when entering the ifBlock production.
	EnterIfBlock(c *IfBlockContext)

	// EnterElseBlock is called when entering the elseBlock production.
	EnterElseBlock(c *ElseBlockContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterWhileBlock is called when entering the whileBlock production.
	EnterWhileBlock(c *WhileBlockContext)

	// EnterDoStatement is called when entering the doStatement production.
	EnterDoStatement(c *DoStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterSubroutinecall is called when entering the subroutinecall production.
	EnterSubroutinecall(c *SubroutinecallContext)

	// EnterClassObject is called when entering the classObject production.
	EnterClassObject(c *ClassObjectContext)

	// EnterExpressionlist is called when entering the expressionlist production.
	EnterExpressionlist(c *ExpressionlistContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBinaryoperation is called when entering the binaryoperation production.
	EnterBinaryoperation(c *BinaryoperationContext)

	// EnterIntegerTerm is called when entering the IntegerTerm production.
	EnterIntegerTerm(c *IntegerTermContext)

	// EnterStringTerm is called when entering the StringTerm production.
	EnterStringTerm(c *StringTermContext)

	// EnterKeywordTerm is called when entering the KeywordTerm production.
	EnterKeywordTerm(c *KeywordTermContext)

	// EnterVarnameTerm is called when entering the VarnameTerm production.
	EnterVarnameTerm(c *VarnameTermContext)

	// EnterArrayTerm is called when entering the ArrayTerm production.
	EnterArrayTerm(c *ArrayTermContext)

	// EnterSubroutineTerm is called when entering the SubroutineTerm production.
	EnterSubroutineTerm(c *SubroutineTermContext)

	// EnterParsTerm is called when entering the ParsTerm production.
	EnterParsTerm(c *ParsTermContext)

	// EnterUnaryopTerm is called when entering the unaryopTerm production.
	EnterUnaryopTerm(c *UnaryopTermContext)

	// EnterClassname is called when entering the classname production.
	EnterClassname(c *ClassnameContext)

	// EnterVarname is called when entering the varname production.
	EnterVarname(c *VarnameContext)

	// EnterSubroutinename is called when entering the subroutinename production.
	EnterSubroutinename(c *SubroutinenameContext)

	// ExitClassdef is called when exiting the classdef production.
	ExitClassdef(c *ClassdefContext)

	// ExitClassvardec is called when exiting the classvardec production.
	ExitClassvardec(c *ClassvardecContext)

	// ExitAtype is called when exiting the atype production.
	ExitAtype(c *AtypeContext)

	// ExitSubrotinedec is called when exiting the subrotinedec production.
	ExitSubrotinedec(c *SubrotinedecContext)

	// ExitSubroutinetype is called when exiting the subroutinetype production.
	ExitSubroutinetype(c *SubroutinetypeContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitSubroutinebody is called when exiting the subroutinebody production.
	ExitSubroutinebody(c *SubroutinebodyContext)

	// ExitVardecs is called when exiting the vardecs production.
	ExitVardecs(c *VardecsContext)

	// ExitVardec is called when exiting the vardec production.
	ExitVardec(c *VardecContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLetStatement is called when exiting the letStatement production.
	ExitLetStatement(c *LetStatementContext)

	// ExitLvalue is called when exiting the lvalue production.
	ExitLvalue(c *LvalueContext)

	// ExitRvalue is called when exiting the rvalue production.
	ExitRvalue(c *RvalueContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitIfBlock is called when exiting the ifBlock production.
	ExitIfBlock(c *IfBlockContext)

	// ExitElseBlock is called when exiting the elseBlock production.
	ExitElseBlock(c *ElseBlockContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitWhileBlock is called when exiting the whileBlock production.
	ExitWhileBlock(c *WhileBlockContext)

	// ExitDoStatement is called when exiting the doStatement production.
	ExitDoStatement(c *DoStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitSubroutinecall is called when exiting the subroutinecall production.
	ExitSubroutinecall(c *SubroutinecallContext)

	// ExitClassObject is called when exiting the classObject production.
	ExitClassObject(c *ClassObjectContext)

	// ExitExpressionlist is called when exiting the expressionlist production.
	ExitExpressionlist(c *ExpressionlistContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBinaryoperation is called when exiting the binaryoperation production.
	ExitBinaryoperation(c *BinaryoperationContext)

	// ExitIntegerTerm is called when exiting the IntegerTerm production.
	ExitIntegerTerm(c *IntegerTermContext)

	// ExitStringTerm is called when exiting the StringTerm production.
	ExitStringTerm(c *StringTermContext)

	// ExitKeywordTerm is called when exiting the KeywordTerm production.
	ExitKeywordTerm(c *KeywordTermContext)

	// ExitVarnameTerm is called when exiting the VarnameTerm production.
	ExitVarnameTerm(c *VarnameTermContext)

	// ExitArrayTerm is called when exiting the ArrayTerm production.
	ExitArrayTerm(c *ArrayTermContext)

	// ExitSubroutineTerm is called when exiting the SubroutineTerm production.
	ExitSubroutineTerm(c *SubroutineTermContext)

	// ExitParsTerm is called when exiting the ParsTerm production.
	ExitParsTerm(c *ParsTermContext)

	// ExitUnaryopTerm is called when exiting the unaryopTerm production.
	ExitUnaryopTerm(c *UnaryopTermContext)

	// ExitClassname is called when exiting the classname production.
	ExitClassname(c *ClassnameContext)

	// ExitVarname is called when exiting the varname production.
	ExitVarname(c *VarnameContext)

	// ExitSubroutinename is called when exiting the subroutinename production.
	ExitSubroutinename(c *SubroutinenameContext)
}
