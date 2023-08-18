# Interal Documentation

## Concept

The sophia language interpreter is build around my previous experience building
a markdown to html compiler. I made use of the visitor design pattern in this
implementation and made use of it in the sophia interpreter. The sophia
interpreter is inspired by my compiler design university course.

## Lexical analyisis

The first step for interpreting a given sophia expression is to scan the
characters for token. For the sake of simplicity we will use the most basic
expression i can think of:

```lisp
(put "hello world")
```

The above statement can be roughly translated into the following list of token:

```
BRACE       (
PUT         put
STRING      "hello world"
BRACE       )
```

This list of tokens is passed to the next interpretation step, the abstract syntax tree creation.

## Creating an AST

Before creating the ast the interpreter checks if the expression matches our
languages grammar, for instance in the above expression we could've put a
second keyword after `put`, which would not match our language grammar:

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

The resulting ast of our before defined expression expressed using json would look like:

```json
{
  "put": {
    "arguments": [
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

### Typechecking

<!-- TODO: -->
