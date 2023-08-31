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

### Running

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

```
$ echo '(put "Hello World")' | sophia

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

### Compiling

#### Targets

- [x] Javascript
- [ ] C
- [ ] Go
- [ ] Python

#### Compiling sophia to a target

Specify the desired compilation target from the list above using the `-target`-flag when invoking sophia with a source expression:

```text
$ sophia -exp='(put "hello world")' -target=js
console.log("Hello World")
```
