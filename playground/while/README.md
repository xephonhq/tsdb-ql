# WHILE-lang

A full implementation of WHILE-lang including 

- lexer and parser using ANTLR
- interpreter
  - small step
  - big step
- translator

## Develop

- Run `go test -v -cover ./ast` for Test or `Ayi test`
- Run `antlr4 -Dlanguage=Go while.g4` in `parser` folder if you updated the `while.g4`