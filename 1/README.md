# Parsing Text using DSL Grammar

source for inspiration for this approach taken from this [article](https://neil.computer/notes/parsing-text-using-dsl-grammar/)

gramma file is `names.peg` [PEG](https://en.wikipedia.org/wiki/Parsing_expression_grammar)

## how to use

install `pigeon`

`go install github.com/mna/pigeon@latest`

generate the parser

`pigeon -o parser.go names.peg`

build and run

`go build . && ./dsl`