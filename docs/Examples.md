# Sophia examples

## Hello world

```lisp
(put "hello world")
(put "hello" "world")
```

## Variables

```lisp
(put
    "pi:"
    (let pi 3.1514))

(put
    "pi*pi" (* pi pi))

(put
    "version:"
    (let v "v1.1.1")
    v
    (let v "1")
    v
)

(put
    "boolean variables:"
    (let bool true)
    (put bool)
)
```

## Functions

```lisp
(fun
    square
    (_ a)
    (* a a)
)

(put
    "function call:"
    (square 12))
```

## Lists

```lisp
(let list 0 1 2 3 4)
(let listEq1 0..4)
(let listEq2 ..4)
```

```lisp
(let list ..4)
(++ list list)
```

## Math

```lisp
(put
    "1+2=" (+ 1 2))
(put
    "1-2=" (- 1 2))
(put
    "1*2=" (* 1 2))
(put
    "1/2=" (/ 1 2))

(put
    "1%2=" (% 1 2))

(- 25
    (+ 1
        (* 5
            (/ 5 2))))

```

## Controlflow

```lisp
(if
    true
    (put "the condition is true")
)

(if
    false
    (put "?")
)

(eq 1 1)
(eq "equal" "equal")
(eq true false)
(lt 10 1)
(gt 1 10)
(lt 1 10)
(gt 10 1)

(let a 1)
(put
    "a == a -> "
    (eq a a))

(or false true false)

(let a false)
(put
    "a | a | true -> "
    (or a a true))

(and true false)

(let a true)
(put
    "a & a -> "
    (and a a))

(not
    (eq 1 2))

(put
    "not (and a) -> "
    (not
        (eq a)))
```

## Iteration

```lisp
(let arr 9)
(for (_ i) arr
     (put i))
```

```lisp
(let arr 1..9)
(for (_ i) arr
     (put i))
```

```lisp
(let sum 0)
(let arr 1..9)
(for (_ e) arr
    (let sum (+ e sum)))
(put sum)
```

## Fibonacci sequence

```lisp
(let beforeLast 0)
(let lst 1)
(for (_ i) 0..15
    (let t (+ beforeLast lst))
    (put t)
    (let beforeLast lst)
    (let lst t))
```

## Fizzbuzz

```lisp
(for (_ i) 1..15
    (let mod3 (eq 0 (% i 3)))
    (let mod5 (eq 0 (% i 5)))
    (match
        (if (and mod3 mod5) (put "FizzBuzz"))
        (if mod3 (put "Fizz"))
        (if mod5 (put "Buzz"))
        (put i)
))
```
