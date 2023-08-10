# Sophia

My take on a small lisp like programming language with single characters as
keywords and operators which is named after my girlfriend.

View [examples/so.phia](examples/so.phia) for sophia by example and [this blog
article](https://xnacly.me/posts/2023/write-your-own-programming-language/) for
a short overview.

```lisp
($
    square
    (_ t)
    (* t t)
)
(: a 12)
(.
    "a*a is"
    (square 12))
```

## Try

```bash
git clone https://github.com/xnacly/sophia
go build
```

With a file:

```text
$ sophia ./examples/helloworld.phia

Hello World!
```

With an expression:

```
$ sophia -exp '(. "Hello World")'

Hello World!
```

As a repl:

```
$ sophia

Welcome to the Sophia repl - press <CTRL-D> or <CTRL-C> to quit...
ß :: (. "Hi!")
Hi!
= []
ß ::
```

## Reference

Sophia currently supports:

- the o in io :^)
- floats
- strings
- booleans
- arithmetics

### Hello world:

```bash
sophia -exp '(. "Hello World!")'
# Hello World!
```

### Keyword reference:

| keyword | description                                                                                                         |
| ------- | ------------------------------------------------------------------------------------------------------------------- |
| `.`     | prints all arguments to stdout                                                                                      |
| `+`     | adds all arguments together                                                                                         |
| `-`     | subtracts all arguments                                                                                             |
| `*`     | multiplies all arguments                                                                                            |
| `/`     | divides all arguments                                                                                               |
| `%`     | modulo all arguments                                                                                                |
| `,`     | concatinate strings, returns combination of arguments as string                                                     |
| `:`     | defines a variable, with the first argument as the name and the remaining argument as the value                     |
| `?`     | defines a condition, evaluates the first argument, evaluates all following argument if the first argument is truthy |
| `=`     | returns true if all arguments are equal to each other                                                               |
| `&`     | returns true if all arguments are true                                                                              |
| `\|`    | returns true if one of the arguments is true                                                                        |
| `!`     | negates the first argument, returns the first argument                                                              |
| `$`     | defines a function, first param name, second param parameters, rest function body                                   |
| `_`     | defines the parameters of a function                                                                                |

### Execution direction

All execution is done left to right, meaning:

```lisp
(+ 2 3 4) ;; results in 2+3+4 -> 9
(- 2 3 4) ;; results in 2-3-4 -> -5
(* 2 3 4) ;; results in 2*3*4 -> 24
(/ 2 3 4) ;; results in 2/3/4 -> 0.166667
(% 2 3 4) ;; results in 2%3%4 -> 2
```

## Progress

- [x] lexing
- [x] parsing
- [x] evaluation
- [x] variables
- [x] controlflow
- [x] functions
