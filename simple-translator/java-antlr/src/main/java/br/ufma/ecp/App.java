package br.ufma.ecp;


import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CharStreams;
import org.antlr.v4.runtime.CommonTokenStream;

import br.ufma.ecp.parser.SimpleLexer;
import br.ufma.ecp.parser.SimpleParser;


/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
        try {
            SimpleLexer lexer;
            SimpleParser parser;
            lexer = new SimpleLexer(CharStreams.fromFileName("input.si"));
            CommonTokenStream commonTokenStream = new CommonTokenStream(lexer);
            parser = new SimpleParser( commonTokenStream);
            parser.program();
            System.out.println("compilado com sucesso");
        } catch (Exception ex) {
            System.out.println("erro");
        }
    }
}
