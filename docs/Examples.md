# Sophia examples

## Hello world

```lisp
(println "hello world")
(println "hello" "world")
```

## Variables

```lisp
(println
    "pi:"
    (let pi 3.1514))

(println
    "pi*pi" (* pi pi))

(println
    "version:"
    (let v "v1.1.1")
    v
    (let v "1")
    v
)

(println
    "boolean variables:"
    (let bool true)
    (println bool)
)
```

## Functions

```lisp
(fun
    square
    [a]
    (* a a)
)

(println
    "function call:"
    (square 12))
```

## Lists

```lisp
(let list 0 1 2 3 4)
```

```lisp
(let list 5)
(++ list list)
```

## Math

```lisp
(println
    "1+2=" (+ 1 2))
(println
    "1-2=" (- 1 2))
(println
    "1*2=" (* 1 2))
(println
    "1/2=" (/ 1 2))

(println
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
    (println "the condition is true")
)

(if
    false
    (println "?")
)

(= 1 1)
(= "equal" "equal")
(= true false)
(< 10 1)
(> 1 10)
(< 1 10)
(> 10 1)

(let a 1)
(println
    "a == a -> "
    (= a a))

(or false true false)

(let a false)
(println
    "a | a | true -> "
    (or a a true))

(and true false)

(let a true)
(println
    "a & a -> "
    (and a a))

(not
    (= 1 2))

(println
    "not (and a) -> "
    (not
        (= a)))
```

## Iteration

```lisp
(let arr 9)
(for [i] arr
     (println i))
```

```lisp
(for [i] 9
     (println i))
```

```lisp
(let sum 0)
(let arr 9)
(for [i] arr
    (let sum (+ e sum)))
(println sum)
```

## Fibonacci sequence

```lisp
(let beforeLast 0)
(let lst 1)
(for [i] 15
    (let t (+ beforeLast lst))
    (println t)
    (let beforeLast lst)
    (let lst t))
```

## Fizzbuzz

```lisp
(for [i] 15
    (let mod3 (eq 0 (% i 3)))
    (let mod5 (eq 0 (% i 5)))
    (match
        (if (and mod3 mod5) (println "FizzBuzz"))
        (if mod3 (println "Fizz"))
        (if mod5 (println "Buzz"))
        (println i)
))
```
