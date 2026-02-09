grammar Jack;

// Rules
classdef:
	CLASS classname LBRACE classvardec* subrotinedec* RBRACE EOF;

classvardec:
	kind = (STATIC | FIELD) atype varname (COMMA varname)* SEMICOLON;

atype: (INT | CHAR | BOOLEAN | ID);

subrotinedec:
	kind = (CONSTRUCTOR | FUNCTION | METHOD) subroutinetype subroutinename LPAREN parameterList
		RPAREN subroutinebody;

subroutinetype: (VOID | atype);

parameterList: ( parameter (COMMA parameter)*)?;

parameter: atype varname;

subroutinebody: LBRACE vardecs statement* RBRACE;

vardecs: vardec*;

vardec: VAR atype varname (COMMA varname)* SEMICOLON;

statement:
	letStatement
	| ifStatement
	| whileStatement
	| doStatement
	| returnStatement;

letStatement: LET lvalue EQ rvalue SEMICOLON;

lvalue: ident = varname (LBRACKET expression RBRACKET)?;
rvalue: expression;

ifStatement: IF LPAREN expression ifBlock ( (elseBlock)?);

ifBlock: RPAREN LBRACE statement* RBRACE;

elseBlock: ELSE LBRACE statement* RBRACE;

whileStatement: WHILE LPAREN expression whileBlock;

whileBlock: RPAREN LBRACE statement* RBRACE;

doStatement: DO subroutinecall SEMICOLON;

returnStatement: RETURN expression? SEMICOLON;

subroutinecall: (classObject)? subroutinename LPAREN expressionlist RPAREN;

classObject: ID DOT;

expressionlist: (expression (COMMA expression)*)?;

expression: term binaryoperation*;

binaryoperation: (
		operator = (
			PLUS
			| MINUS
			| ASTERISK
			| SLASH
			| AND
			| OR
			| LT
			| GT
			| EQ
		) term
	);

term:
	INTEGER											# IntegerTerm
	| STRING										# StringTerm
	| keywordConst = (TRUE | FALSE | NULL | THIS)	# KeywordTerm
	| varname										# VarnameTerm
	| ident = varname LBRACKET expression RBRACKET	# ArrayTerm
	| subroutinecall								# SubroutineTerm
	| LPAREN expression RPAREN						# ParsTerm
	| unaryop = (MINUS | NOT) term					# unaryopTerm;

classname: ID;
varname: ID;
subroutinename: ID;

// Symbols

// operators

PLUS: '+';
MINUS: '-';
ASTERISK: '*';
SLASH: '/';
AND: '&';
OR: '|';
NOT: '~';
LT: '<';
GT: '>';
EQ: '=';

// Delimiters
DOT: '.';
COMMA: ',';
SEMICOLON: ';';
LPAREN: '(';
RPAREN: ')';
LBRACE: '{';
RBRACE: '}';
LBRACKET: '[';
RBRACKET: ']';

// Keywords
METHOD: 'method';
STATIC: 'static';
INT: 'int';
BOOLEAN: 'boolean';
TRUE: 'true';
NULL: 'null';
LET: 'let';
IF: 'if';
WHILE: 'while';
CONSTRUCTOR: 'constructor';
FIELD: 'field';
VAR: 'var';
CHAR: 'char';
VOID: 'void';
CLASS: 'class';
FALSE: 'false';
DO: 'do';
ELSE: 'else';
RETURN: 'return';
FUNCTION: 'function';
THIS: 'this';

ID: ([a-z] | [A-Z]) ([a-z] | [A-Z] | [0-9])*;

STRING: '"' .*? '"';

INTEGER: [0-9]+;

COMMENT: '/*' .*? '*/' -> skip;
LINE_COMMENT: '//' ~[\r\n]* -> skip;

WHITESPACE: [ \r\n\t]+ -> skip;