# Sophia

My take on a small lisp like programming language named after my girlfriend.

View [examples/so.phia](examples/so.phia) for sophia by example and [this blog
article](https://xnacly.me/posts/2023/write-your-own-programming-language/) for
a short overview.

```lisp
(fun
    square
    (_ t)
    (* t t)
)
(let a 12)
(put
    "a*a is"
    (square a))
;; puts a*a is 144
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
$ sophia -exp '(put "Hello World")'

Hello World!
```

As a repl:

```
$ sophia

  +####
 +\    #
+  \ ß  #
+   \   # <-> ß-calculus
+ ß  \  #
 +    \#
  ++++#

Welcome to the Sophia repl - press <CTRL-D> or <CTRL-C> to quit...
ß :: (put "Hi!")
Hi!
= [<nil>]
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
sophia -exp '(put "Hello World!")'
# Hello World!
```

### Keyword reference:

| keyword | description                                                                                                         |
| ------- | ------------------------------------------------------------------------------------------------------------------- |
| `put`   | prints all arguments to stdout                                                                                      |
| `+`     | adds all arguments together                                                                                         |
| `-`     | subtracts all arguments                                                                                             |
| `*`     | multiplies all arguments                                                                                            |
| `/`     | divides all arguments                                                                                               |
| `%`     | modulo all arguments                                                                                                |
| `++`    | merges and returns lists, concatinates strings                                                                      |
| `let`   | defines a variable, with the first argument as the name and the remaining argument as the value                     |
| `eq`    | returns true if all arguments are equal to each other                                                               |
| `lt`    | returns true if the first argument is less than the second                                                          |
| `gt`    | returns true if the first argument is greater than the second                                                       |
| `if`    | defines a condition, evaluates the first argument, evaluates all following argument if the first argument is truthy |
| `and`   | returns true if all arguments are true                                                                              |
| `or`    | returns true if one of the arguments is true                                                                        |
| `not`   | negates the first argument, returns the first argument                                                              |
| `_`     | defines the parameters of a function                                                                                |
| `fun`   | defines a function, first param name, second param parameters, rest function body                                   |
| `for`   | defines a loop, first param element, second arguemnt value to iterate over, rest loop body                          |

### Execution direction

All execution is done left to right, meaning:

```lisp
(+ 2 3 4) ;; results in 2+3+4 -> 9
(- 2 3 4) ;; results in 2-3-4 -> -5
(* 2 3 4) ;; results in 2*3*4 -> 24
(/ 2 3 4) ;; results in 2/3/4 -> 0.166667
(% 2 3 4) ;; results in 2%3%4 -> 2
```
