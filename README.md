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
(let res (square n))

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

  ██████  ▒█████   ██▓███   ██░ ██  ██▓ ▄▄▄
▒██    ▒ ▒██▒  ██▒▓██░  ██▒▓██░ ██▒▓██▒▒████▄
░ ▓██▄   ▒██░  ██▒▓██░ ██▓▒▒██▀▀██░▒██▒▒██  ▀█▄
  ▒   ██▒▒██   ██░▒██▄█▓▒ ▒░▓█ ░██ ░██░░██▄▄▄▄██
▒██████▒▒░ ████▓▒░▒██▒ ░  ░░▓█▒░██▓░██░ ▓█   ▓██▒
▒ ▒▓▒ ▒ ░░ ▒░▒░▒░ ▒▓▒░ ░  ░ ▒ ░░▒░▒░▓   ▒▒   ▓▒█░
░ ░▒  ░ ░  ░ ▒ ▒░ ░▒ ░      ▒ ░▒░ ░ ▒ ░  ▒   ▒▒ ░
░  ░  ░  ░ ░ ░ ▒  ░░        ░  ░░ ░ ▒ ░  ░   ▒
      ░      ░ ░            ░  ░  ░ ░        ░  ░

Welcome to the Sophia programming language repl - press <CTRL-D> or <CTRL-C> to quit...
sophia> (let name "user")
= [user]
sophia> (put 'Hello World, {name}!')
Hello World, user!
= [<nil>]
sophia>
```

### Compiling

By default sophia executes input with its tree walk interpreter.

#### Targets

Currently the following targets are supported:

- [x] Javascript (all language features supported)
- [ ] Python (planned)

Specify the desired compilation target from the list above using the
`-target`-flag when invoking sophia with a source expression:

```text
$ sophia -exp='(put "hello world")' -target=js
console.log("Hello World")
```

#### Compiling sophia to javascript

Sophia produces valid minified javascript, for example compiling the example
from the beginning or a more complex example:

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

```lisp
;; leetcode 327 Maximum Count of positive integer and negative integer
(let example1
    (not 2) ;; negated: (not 2) -> -2
    (not 1)
    (not 1) 1 2 3)
(let example2 (not 3) (not 2) (not 1) 0 0 1 2)
(let example3 5 20 66 1314)

;; returns a if bigger than b, otherwise b
;; max: float, float -> float
(fun max (_ a b)
    (let max) ;; created as nil
    (match
        (if (lt a b) (let max b))
        (let max a))
    max) ;; return without extra statement

;; counts negative and positve numbers in arr, returns the higher amount
;; solve: [float] -> float
(fun solve (_ arr)
    (let pos 0)
    (let neg 0)
    (for (_ i) arr
        (match
            (if (lt i 0) (let pos (+ pos 1)))
            (let neg (+ neg 1))))
    (max neg pos))

(put (solve example1)) ;; 3
(put (solve example3)) ;; 4
(put (solve example2)) ;; 4
```

```js
let example1 = [-2, -1, -1, 1, 2, 3];
let example2 = [-3, -2, -1, 0, 0, 1, 2];
let example3 = [5, 20, 66, 1314];
function max(a, b) {
  let max;
  if (a < b) {
    max = b;
  } else {
    max = a;
  }
  return max;
}
function solve(arr) {
  let pos = 0;
  let neg = 0;
  for (let i of arr) {
    if (i < 0) {
      pos = pos + 1;
    } else {
      neg = neg + 1;
    }
  }
  return max(neg, pos);
}
console.log(solve(example1));
console.log(solve(example3));
console.log(solve(example2));
```
