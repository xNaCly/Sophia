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
$ tisp -e '[add "Hello World"]'

~ [Hello World!]
```

as a repl:

```
$ tisp

Welcome to the Tisp repl - press <CTRL-D> or <CTRL-C> to quit...
τ :: [putv "Hi!"]
~ [Hi!]
= []
τ ::
```

## Reference

### Hello world:

tisp currently supports arithmetics and the o in io :^)

```bash
tisp -e '[putv "Hello World!"]'
# [Hello World!]
```

### Keyword reference:

| keyword | description                                                        |
| ------- | ------------------------------------------------------------------ |
| `putv`  | prints all arguments to stdout, supports floats, bools and strings |
| `add`   | adds all arguments together                                        |
| `sub`   | subtracts all arguments                                            |
| `mul`   | multiplies all arguments                                           |
| `div`   | divides all arguments                                              |

### Execution direction

All execution is done left to right, meaning:

```lisp
[add 2 3 4] ;; results in 2+3+4 -> 9
[sub 2 3 4] ;; results in 2-3-4 -> -5
[mul 2 3 4] ;; results in 2*3*4 -> 24
[div 2 3 4] ;; results in 2/3/4 -> 0.166667
```

Tips supports strings and 64 bit floating point integers.

## Progress

- [x] lexing
- [x] parsing
- [x] evaluation
- [ ] variables
- [ ] controlflow
- [ ] functions
