# Interal Documentation

## Concept

The sophia language interpreter is build around my previous experience building
a markdown to html compiler. I made use of the visitor design pattern in this
implementation and made use of it in the sophia interpreter. The sophia
interpreter is inspired by my compiler design university course.

## Interpretation

The interpretation is split up in the following three abstract steps:

| Step                  | Description                                                                      |
| --------------------- | -------------------------------------------------------------------------------- |
| Lexical analyisis     | Converts the character input to a stream of token with meanings                  |
| Syntactical analyisis | Converts the tokenstream to an abstract syntax tree, includes semantic analyisis |
| Evaluation            | Walks the tree and evaluates each node                                           |

The following chapters will explain each step with an example.

### Lexical analysis

The first step for interpreting a given sophia expression is to scan the
characters for token. For the sake of simplicity we will use the most basic
expression i can think of:

```lisp
(println "hello world")
```

The above statement can be roughly translated into the following list of token:

| Token type    | Line index | Line | Raw value     |
| ------------- | ---------- | ---- | ------------- |
| `LEFT_BRACE`  | 0          | 0    | `(`           |
| `IDENT`       | 1          | 0    | `println`     |
| `STRING`      | 7          | 0    | `hello world` |
| `RIGHT_BRACE` | 18         | 0    | `)`           |

The lexer specifically distinguishes between strings, operators, keywords,
booleans, identifiers and floating point integers.

This list is passed to the next interpretation step - the abstract syntax tree creation.

### Syntactical analysis

The parser checks if the expression starts with a token of type `LEFT_BRACE`,
if it does the parser proceeds to check for an operator or a keyword, such as
`put` or `+-*/%` or `if`. If the token matches the parser checks if the
following token are valid arguments for the previously matched operator. If
so the last step for an expression to be valid is the `RIGHT_BRACE` token.

With our example:

```
LEFT_BRACE      -> valid expression start, proceed
println         -> valid function, proceed
"hello world"   -> valid argument for operator, proceed
RIGHT_BRACE     -> valid expression end, return ast
```

The resulting abstract syntax tree of our before defined expression expressed using JSON would look like:

```json
{
  "println": {
    "children": [
      {
        "string": {
          "value": "hello world"
        }
      }
    ]
  }
}
```

## Evaluation

As said before the evaluation step is realised using the visitor pattern, which
enables the evaluation of each node. Consider the following expression definition for a sophia string:

```go
// core/expr/str.go
package expr

import "sophia/core/token"

type String struct {
	Token token.Token
}

func (s *String) GetToken() token.Token {
	return s.Token
}

func (s *String) Eval() any {
	return s.Token.Raw
}
```

By attaching the `GetToken` and the `Eval` methods to the `String` structure,
the `Node` structure is implemented:

```go
// core/expr/node.go
package expr

import "sophia/core/token"

type Node interface {
	GetToken() token.Token
	Eval() any
}
```

Therefore a string can be accepted as a value for all other expressions which
accept structures of type `Node` as values for their attributes.

### Typechecking

Some Nodes accept children of all types, other operations, such as arithmetics
should and can not accept arguments of all types. To fix this the `expr`
package contains the `expr.castPanicIfNotType` function, which does exactly
what its name implies - it panics if the given argument is not of the generic
type specified. It's used extensively in functions such as `expr.Add.Eval()`:

```go
// core/expr/add.go
package expr

import "sophia/core/token"

type Add struct {
	Token    token.Token
	Children []Node
}

func (a *Add) GetToken() token.Token {
	return a.Token
}

func (a *Add) Eval() any {
	if len(a.Children) == 0 {
		return 0.0
	}
	res := 0.0
	for i, c := range a.Children {
		if i == 0 {
			res = castPanicIfNotType[float64](c.Eval(), a.Token)
		} else {
			res += castPanicIfNotType[float64](c.Eval(), a.Token)
		}
	}
	return res
}
```

The first children is the starting point to add all other children onto. The
`expr.castPanicIfNotType` function is used to check if the return value of
evaluating an expression can be added onto an other value. The second argument
is used for indicating the operation which caused the runtime panic in the
error message.

## Error handling

Error handling is a big topic for the developer experience. The error handling
of the three different interpreter steps is different, therefore each stage has
its own following chapter.

### While lexing

If the lexer is invoked with an empty input it exits and returns the following message:

```
$ echo " " | sophia
error: Unexpected end of file

	at: /home/teo/stdin:1:1:

	   1|
	    | ^

Source possibly empty
Syntax errors found, skipping remaining interpreter stages. (parsing and evaluation)
```

If the lexer encounters an unknown character it tries to scan the resulting input to provide the user with all possible errors:

```
$ sophia -exp '(= ? "unclosed string )'

error: Unknown character

	at: cli:1:2:

	   1| (= ? "unclosed string )
	    |  ^

Unexpected "="

error: Unknown character

	at: cli:1:4:

	   1| (= ? "unclosed string )
	    |    ^

Unexpected "?"

error: Unterminated string

	at: cli:1:6:

	   1| (= ? "unclosed string )
	    |      ^^^^^^^^^^^^^^^^^^

Consider closing the string via "
Syntax errors found, skipping remaining interpreter stages. (parsing and evaluation)
```

Providing the lexer with input containing more than one line, the lexer shows context for the currently found error:

```
$ cat err.phia
;; this input contains two errors
(println " ?)
(? "test")
$ sophia err.phia
error: Unterminated string

	at: /home/teo/err.phia:2:7:

	   1| ;; this input contains two errors
	   2| (println " ?)
	    |          ^^^^
	   3| (? "test")
	   4|

Consider closing the string via "

error: Unknown character

	at: /home/teo/err.phia:3:2:

	   1| ;; this input contains two errors
	   2| (println " ?)
	   3| (? "test")
	    |  ^
	   4|

Unexpected "?"

Syntax errors found, skipping remaining interpreter stages. (parsing and evaluation)
```

### While parsing

Before creating the ast the interpreter checks if the expression matches our
languages grammar, for instance in the aforementioned expression we could've
put a second keyword after `put`, which would not match our language grammar:

```lisp
(println put)
```

```
$ sophia -exp '(println put)'
error: Undefined variable

        at: cli:1:10:

            1| (println put)
             |          ^^^

Variable "put" is not defined.
```

Trying to define a function without an identifier as its name results in the following error:

```
$ sophia -exp '(fun (_) (println))'
error: Type error

	at: cli:1:7:

	   1| (fun (_) (println))
	    |       ^

Expected the first argument for function definition to be of type IDENT, got "PARAM".
Semantic errors found, skipping remaining interpreter stages. (evaluation)
```

Omitting the parameters out, results in:

```
$ sophia -exp '(fun square (* n n))'
error: Type error

	at: cli:1:14:

	   1| (fun square (* n n))
	    |              ^

Expected the second argument for function definition to be of type PARAM, got "MUL".
Semantic errors found, skipping remaining interpreter stages. (evaluation)
```

### While evaluating

Using an undefined variable or function:

```
$ sophia -exp '(println a)'
rror: Undefined variable

	at: cli:1:6:

	   1| (println a)
	    |      ^

Variable "a" is not defined.
$ sophia -exp '(square 12)'
error: Undefined function

	at: cli:1:2:

	   1| (square 12)
	    |  ^^^^^^

Function "square" not defined
```

Iterating of anything which is not a container:

```
$ sophia -exp '(for (_ i) "test" (println "iteration" i))'
error: Invalid iterator

	at: cli:1:13:

	   1| (for (_ i) "test" (println "iteration" i))
	    |             ^^^^

expected container or upper bound for iteration, got: string
```

Trying to use a string in addition:

```
$ sophia -exp '(+ "test" 1)'
error: Type error

	at: cli:1:2:

	   1| (+ "test" 1)
	    |  ^

Incompatiable types string and float64
```

Sophia contains a massive amount of error handling and edge cases, the above are the most prominent cases.
