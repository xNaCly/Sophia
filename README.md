# Sophia

My take on a small lisp like programming language with single characters as
keywords and operators.

## Try

```bash
git clone https://github.com/xnacly/sophia
go build
```

With a file:

```text
$ sophia -f ./examples/helloworld.sophia

~ [Hello World!]
```

With an expression:

```
$ sophia -e '[. "Hello World"]'

~ [Hello World!]
```

As a repl:

```
$ sophia

Welcome to the Sophia repl - press <CTRL-D> or <CTRL-C> to quit...
ß :: [. "Hi!"]
~ [Hi!]
= []
ß ::
```

## Reference

Sophia currently supports arithmetics, strings and the o in io :^)

### Hello world:

```bash
sophia -e '[. "Hello World!"]'
# [Hello World!]
```

### Keyword reference:

| keyword | description                                                        |
| ------- | ------------------------------------------------------------------ |
| `.`     | prints all arguments to stdout, supports floats, bools and strings |
| `+`     | adds all arguments together                                        |
| `-`     | subtracts all arguments                                            |
| `*`     | multiplies all arguments                                           |
| `/`     | divides all arguments                                              |
| `^`     | raises all arguments to the following power                        |
| `%`     | modulo for current argument with the following argument            |

#### Planned keywords:

| keyword | description                                                                                                      |
| ------- | ---------------------------------------------------------------------------------------------------------------- |
| `#`     | defines a function, with the first argument as the name, the second as the argument and the third as the body    |
| `:`     | defines a variable, with the first argument as the name and the second argument as the value                     |
| `?`     | defines a condition, evaluates the first argument, evaluates the second argument if the first argument is truthy |
| `&`     | returns true if all arguments are true                                                                           |
| `\|`    | returns true if one of the arguments is true                                                                     |

### Execution direction

All execution is done left to right, meaning:

```lisp
[+ 2 3 4] ;; results in 2+3+4 -> 9
[- 2 3 4] ;; results in 2-3-4 -> -5
[* 2 3 4] ;; results in 2*3*4 -> 24
[/ 2 3 4] ;; results in 2/3/4 -> 0.166667
[^ 2 3 4] ;; results in 2^3^4 -> 4096
[% 2 3 4] ;; results in 2%3%4 -> 2
```

Tips supports strings and 64 bit floating point integers.

## Progress

- [x] lexing
- [x] parsing
- [x] evaluation
- [ ] variables
- [ ] controlflow
- [ ] functions
