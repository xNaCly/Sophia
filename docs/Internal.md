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
(put "hello world")
```

The above statement can be roughly translated into the following list of token:

| Token type    | Line | Line index | Raw value     |
| ------------- | ---- | ---------- | ------------- |
| `LEFT_BRACE`  | 0    | 0          | `(`           |
| `PUT`         | 1    | 0          | `put`         |
| `STRING`      | 7    | 0          | `hello world` |
| `RIGHT_BRACE` | 18   | 0          | `)`           |

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
PUT             -> valid operator, proceed
"hello world"   -> valid argument for operator, proceed
RIGHT_BRACE     -> valid expression end, return ast
```

The resulting abstract syntax tree of our before defined expression expressed using JSON would look like:

```json
{
  "put": {
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

Taking a look at the `Put` expression definition we notice the `Children`
attribute of type `[]Node`.

```go
package expr

import (
	"fmt"
	"sophia/core/token"
	"strings"
)

type Put struct {
	Token    token.Token
	Children []Node
}

func (p *Put) GetToken() token.Token {
	return p.Token
}

func (p *Put) Eval() any {
	b := strings.Builder{}
	for i, c := range p.Children {
		if i != 0 {
			b.WriteRune(' ')
		}
		b.WriteString(fmt.Sprint(c.Eval()))
	}
	fmt.Printf("%s\n", b.String())
	return nil
}
```

This definition allows the interpreter to call the `Node.Eval()` function on
the `Put` structure, which will call the `Node.Eval()` function and return its
values on all its children, which will do the same if they contain children.
This makes for a nice, easy but slow recursive evaluation.

### Typechecking

The `Put` structure accepts children which can contain all values, other
operations, such as arithmetics should and can not accept arguments of all
types. To fix this the `expr` package contains the `expr.castPanicIfNotType`
function, which does exactly what its name implies - it panics if the given
argument is not of the generic type specified. It's used extensively in functions such as `expr.Add.Eval()`:

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
			res = castPanicIfNotType[float64](c.Eval(), token.ADD)
		} else {
			res += castPanicIfNotType[float64](c.Eval(), token.ADD)
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
$ sophia empty.phia
14:25:59.319979 err: input is empty, stopping
14:25:59.320032 lexer error
14:25:59.320038 error in source file 'empty.phia' detected, stopping...
exit status 1
```

If the lexer encounters an unknown character it tries to scan the resulting input to provide the user with all possible errors:

```
$ sophia -exp '(= ? "unclosed string )'
14:28:09.839304 err: Unknown token '=' at [l 1:1]

001 |   (= ? "unclosed string )
         ^

14:28:09.839355 err: Unknown token '?' at [l 1:3]

001 |   (= ? "unclosed string )
           ^

14:28:09.839363 err: Unterminated String at [l 1:5]

001 |   (= ? "unclosed string )
             ^^^^^^^^^^^^^^^^^

14:28:09.839373 lexer error
exit status 1
```

Providing the lexer with input containing more than one line, the lexer shows context for the currently found error:

```lisp
$ cat err.phia
;; this input contains two errors
(put " ?)
(? "test")
$ sophia err.phia
14:30:25.523710 err: Unterminated String at [l 2:7]

001 |   ;; this input contains two errors
002 |   (put " ?)
             ^^^

14:30:25.523761 err: Unknown token '?' at [l 3:2]

002 |   (put " ?)
003 |   (? "test")
         ^

14:30:25.523773 lexer error
14:30:25.523777 error in source file 'err.phia' detected, stopping...
exit status 1
```

### While parsing

Before creating the ast the interpreter checks if the expression matches our
languages grammar, for instance in the aforementioned expression we could've
put a second keyword after `put`, which would not match our language grammar:

```lisp
(put put)
```

Interpreter error:

```
$ sophia -exp "(put put)"
16:10:28.093639 err: Missing or unknown argument - Expected any of: 'float,string,identifier,bool', got 'put' [l: 0:5]
16:10:28.093692 parser error
exit status 1
```

Trying to define a function without an identifier as its name results in the following error:

```
$ sophia -exp '(fun (_) (put))'
14:35:48.123879 err: expected the first argument for function definition to be of type IDENT, got "_"
14:35:48.123929 err: Missing statement start - Expected Token '(' got ')' [l: 0:14]
14:35:48.123934 err: Missing or unknown operator - Expected any of: '+,-,/,*,%,put,let,if,eq,or,and,not,++,fun,_,identifier,for,lt,gt', got 'End of file' [l: 0:0]
14:35:48.123939 err: Missing statement end - Expected Token ')' got 'End of file' [l: 0:0]
14:35:48.123946 parser error
exit status 1
```

Omitting the parameters out, results in:

```
$ sophia -exp '(fun square (* n n))'
14:37:15.507972 err: expected the second argument for function definition to be of type PARAM, got "identifier"
14:37:15.508022 err: Missing statement start - Expected Token '(' got ')' [l: 0:19]
14:37:15.508026 err: Missing or unknown operator - Expected any of: '+,-,/,*,%,put,let,if,eq,or,and,not,++,fun,_,identifier,for,lt,gt', got 'End of file' [l: 0:0]
14:37:15.508032 err: Missing statement end - Expected Token ')' got 'End of file' [l: 0:0]
14:37:15.508038 parser error
exit status 1
```

### While evaluating

Using an undefined variable or function:

```
$ sophia -exp '(put a)'
14:38:25.630180 err: variable 'a' is not defined!
14:38:25.630219 runtime error
exit status 1
$ sophia -exp '(square 12)'
14:39:01.729986 err: function "square" not defined
14:39:01.730025 runtime error
exit status 1
```

Iterating of anything which is not a container:

```
$ sophia -exp '(for (_ i) 1 (put "iteration" i))'
14:40:07.581402 err: can not use variable of type float64 in current operation (for), expected []interface {} for value 1
14:40:07.581440 runtime error
exit status 1
```

Trying to use a string in addition:

```
$ sophia -exp '(+ "test" 1)'
14:41:07.836950 err: can not use variable of type string in current operation (+), expected float64 for value test
14:41:07.836999 runtime error
exit status
```

Sophia contains a massive amount of error handling and edge cases, the above are the most prominent cases.
