alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.9.3-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 Simples.g4 -package br.ufma.ecp.parser -o ./src/main/java/br/ufma/ecp/parser/