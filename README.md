# Sophia

My take on a small lisp like programming language named after my girlfriend.

View the docs containing an overview, an in depth overview and a lot of
Examples [here](https://xnacly.github.io/Sophia/)

```lisp
(fun
    square
    (_ n)
    (* n n)
)

(let n 12)
(let res (square a))

(put '{n}*{n} is {res}')
;; 12*12 is 144
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

Specify the desired compilation target from the list above using the
`-target`-flag when invoking sophia with a source expression:

```text
$ sophia -exp='(put "hello world")' -target=js
console.log("Hello World")
```

#### Compiling sophia to javascript

Sophia produces valid minified javascript, for example compiling the example
from the beginning:

```lisp
(fun
    square
    (_ n)
    (* n n)
)

(let n 12)
(let res (square a))

(put '{n}*{n} is {res}')
;; 12*12 is 144
```

```js
function square(t) {
  return t * t;
}
let a = 12;
let r = square(a);
console.log(`${a}*${a} is ${r}`);
```
