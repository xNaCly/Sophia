# Tisp

> teo+lisp = tisp (τ)

My take on a small lisp like programming language.

## Try

```bash
git clone https://github.com/xnacly/tisp
go build
```

with a file:

```text
$ tisp -f ./examples/helloworld.tisp

~ [Hello World!]
```

with an expression:

```
$ tisp -e '[. "Hello World"]'

~ [Hello World!]
```

as a repl:

```
$ tisp

Welcome to the Tisp repl - press <CTRL-D> or <CTRL-C> to quit...
τ :: [. "Hi!"]
~ [Hi!]
= []
τ ::
```

## Reference

### Hello world:

tisp currently supports arithmetics and the o in io :^)

```bash
tisp -e '[. "Hello World!"]'
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
