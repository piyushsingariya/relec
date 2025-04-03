# !/bin/bash

# alias antlr4='java -jar $PWD/antlr-4.13.1-complete.jar'
wget http://www.antlr.org/download/antlr-4.13.1-complete.jar
alias antlr4='java -Xmx500M -cp "./antlr-4.13.1-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -visitor -package engine *.g4
rm antlr-4.13.1-complete.jar
