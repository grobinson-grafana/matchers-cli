# matchers-cli

`matchers-cli` is a simple command-line interface for Prometheus-like matchers.

## Installation

You can install `matchers-cli` with `go install github.com/grobinson-grafana/matchers-cli` or build it from source.

## Usage

```
Usage of matchers-cli:
  -just-lex
    	skip parsing and print the tokens from the lexer
```

`matchers-cli` reads input from stdin. You can use it as an interactive program:

```
$ matchers-cli
{foo="bar"}
{foo="bar"}
```

 or pass input from other programs using `|`:
 
 ```
 echo {foo="bar"} | matchers-cli
{foo="bar"}
```

If the input is invalid an error will be printed to stderr:

```
$ matchers-cli
{foo="bar",,}
9:10: unexpected ,: expected a matcher or close paren after comma
```

You can also use `matchers-cli` to print the tokens from the lexer and skip parsing altogether:

```
$ matchers-cli -just-lex
{foo=bar}
(OpenParen) '{'
(Ident) 'foo'
(Op) '='
(Ident) 'bar'
(CloseParen) '}'
```
