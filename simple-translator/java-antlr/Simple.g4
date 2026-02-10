grammar Simple;

@members {
    private String _id;
}

exp : term {System.out.println("push "+ _input.LT(-1).getText());} oper*;

oper: 
    '+' term {System.out.println("push "+ _input.LT(-1).getText() + "\n" + "add");    }
    | '-' term {System.out.println("push "+ _input.LT(-1).getText() + "\n" + "sub");    }
    ;


term: NUMBER | ID ;

program: statements;

statements : statement* ;

statement : letStatement | printStament;

letStatement : 
            'let' ID {_id = _input.LT(-1).getText();} 
             '=' exp ';' {System.out.println("pop "+ _id);  };


printStament : 'print' exp ';' {System.out.println("print");} ; 


// terminais
NUMBER : [0-9]+;

ID: ([a-z] | [A-Z]) ([a-z] | [A-Z] | [0-9])*;

WHITESPACE: [ \r\n\t]+ -> skip;