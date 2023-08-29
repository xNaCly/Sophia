# Sophia

My take on a small lisp like programming language named after my girlfriend.

View the docs containing an overview, an in depth overview and a lot of
Examples [here](https://xnacly.github.io/Sophia/)

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
