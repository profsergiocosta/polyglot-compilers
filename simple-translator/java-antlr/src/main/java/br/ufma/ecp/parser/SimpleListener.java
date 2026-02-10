// Generated from Simple.g4 by ANTLR 4.9.3
package br.ufma.ecp.parser;
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SimpleParser}.
 */
public interface SimpleListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SimpleParser#exp}.
	 * @param ctx the parse tree
	 */
	void enterExp(SimpleParser.ExpContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleParser#exp}.
	 * @param ctx the parse tree
	 */
	void exitExp(SimpleParser.ExpContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleParser#oper}.
	 * @param ctx the parse tree
	 */
	void enterOper(SimpleParser.OperContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleParser#oper}.
	 * @param ctx the parse tree
	 */
	void exitOper(SimpleParser.OperContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleParser#term}.
	 * @param ctx the parse tree
	 */
	void enterTerm(SimpleParser.TermContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleParser#term}.
	 * @param ctx the parse tree
	 */
	void exitTerm(SimpleParser.TermContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(SimpleParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(SimpleParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleParser#statements}.
	 * @param ctx the parse tree
	 */
	void enterStatements(SimpleParser.StatementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleParser#statements}.
	 * @param ctx the parse tree
	 */
	void exitStatements(SimpleParser.StatementsContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(SimpleParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(SimpleParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleParser#letStatement}.
	 * @param ctx the parse tree
	 */
	void enterLetStatement(SimpleParser.LetStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleParser#letStatement}.
	 * @param ctx the parse tree
	 */
	void exitLetStatement(SimpleParser.LetStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link SimpleParser#printStament}.
	 * @param ctx the parse tree
	 */
	void enterPrintStament(SimpleParser.PrintStamentContext ctx);
	/**
	 * Exit a parse tree produced by {@link SimpleParser#printStament}.
	 * @param ctx the parse tree
	 */
	void exitPrintStament(SimpleParser.PrintStamentContext ctx);
}