# Sophia

## Datatypes

Sophia features four data types:

| Datatype | Description                                   | Examples                         |
| -------- | --------------------------------------------- | -------------------------------- |
| float    | 64Bit floating point number                   | `.1`, `1e-3`, `1.1`, `1_000_000` |
| string   | text, multiple and single characters          | `"Hello world"`, `"t"`, `"!!!"`  |
| bool     | boolean                                       | `true`, `false`                  |
| array    | list that is able to contain all of the above | `[1 2 3]`, `[1 "test" true]`     |

## Printing

> When talking about operators all build in keywords and symbols are meant.
> Keywords are textual words with multiple characters, operators are solely
> symbols. Both words are used synonymously in the documentation.

The most useful keyword is the `put` keyword. It prints all arguments to the
standard output. Our current knowledge about Sophia can be applied to create
the famous `Hello World` example:

```lisp
(put "Hello World")
```

The following chapter explains how to execute the expression.

## Executing expressions

When invoking the interpreter without any arguments the repl is started:

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
ß ::
```

The REPL accepts all valid Sophia expressions as well as several commands (read more [here](#repl-commands))

Running the aforementioned `Hello World` can therefore be simply typed into the REPL and executed with pressing `ENTER`:

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
ß :: (put "Hello World")
Hello World
= [<nil>]
```

Noticeable the string we wanted to print got printed but we notice the line
below it, which in the repl indicates the return value of the previously
executed expression. We printed to stdout, therefore our expression does not
return anything (indicated with the `nil` type).

## Comments

Sophia supports single line comments:

```lisp
;; this is a comment
```

Notice the prefix `;;`, similar how in c single line comments start with `//`.

## Artithmetics

Sophia supports the following mathematical operators:

| Operator | Description                                                     |
| -------- | --------------------------------------------------------------- |
| `+`      | Sum of all arguments                                            |
| `-`      | Difference of all arguments                                     |
| `*`      | Product of all arguments                                        |
| `/`      | Quotient of all arguments                                       |
| `%`      | [Modulo](https://en.wikipedia.org/wiki/Modulo) of all arguments |

Examples for all of the above:

```lisp
;; 1+2+3 = 6
(+ 1 2 3)

;; 1-2-3 = -4
(- 1 2 3)

;; 1*2*3 = 6
(* 1 2 3)

;; 1/2/3 = 0.16666666666666666
(/ 1 2 3)

;; 1%2%3 = 1
(% 1 2 3)
```

## Variables

Sophia enables variable definition with the `let`-keyword:

```lisp
(let pi 3.1415)
(let moneyInTheBank 12)
(let myName "anon")
(let iLikeCheeseCake true)

;; assigned to nil
(let thisIsAnEmptyVariable)

;; array of numbers
(let arrayOfNumbers 3.1415 3.1415)
```

> Using the `let` keyword without specifying any arguments after the variable
> name causes the variable to have the `nil` value.

## Comparison

Comparing values is elementary for a programming language, so Sophia too allows it:

### Comparing booleans

```lisp
(and true false)
(or true false)
(not true)
```

`and` evaluates to true if all arguments are true. `or` evaluates to true if
one of its arguments is true. `not` negates the arguments value:
`true`->`false`, etc.

### Comparing numbers

```lisp
(eq 1 2 1)
(lt 1 10)
(gt 20 1)
```

`eq` evaluates to true if all arguments are the same. `lt` evaluates to true if
the first argument is smaller than the second. `gt` evaluates to true if the
first argument is bigger than the second.

## Controlflow

Sophia features conditional evaluation as well as iterating over containers:

### Conditional evaluation

`if` evaluates the first argument, if it returns true all following arguments
are executed.

```lisp
(if
    true                ;; if-head
    (put "true!")       ;; if-body
    (put "true!")       ;; if-body
    (put "true!")       ;; if-body
)

(if
    (and true false)
    (put "true!")
    (put "true!")
    (put "true!")
)
```

### Iteration

Lets take a look at iteration over containers:

```lisp
(let array 1 2 3 4)
(for
    (_ i)               ;; loop variable
    array               ;; container to iterate over
    (put i))            ;; for body
```

`for` assigns the current value of the container iteration to the defined loop
variable, here `i`. As shown above, the first argument is the loop variable,
the second variable the container to iterate over and the following arguments
are, similar to the `if` keyword, evaluated.

### Match

> This is currently not implemented

## Functions

### Definition

Functions are a big part of programming, lets see how they are defined and
called in Sophia:

```lisp
(fun
    square              ;; function name
    (_ n)               ;; parameters
    (* a a))            ;; function body
```

Functions are defined using the `fun` keyword. The first argument is the name
of the function, here `square`. The second argument are, similar to the `for`
keyword, the parameters the function accepts. The following expressions are
again executed.

Notice the missing `return` or anything similar? Functions in Sophia always
return the last expression as their value.

### Usage

Functions can be used as keywords in Sophia. Calling the `square` function from
above can be done as follows:

```lisp
(square 2)
```

Specifying more or less arguments than the function accepts will cause a
runtime error.

Functions with multiple arguments, such as summing two values can be expressed as follows:

```lisp
(fun
    sum
    (_ a b)
    (+ a b))

(sum 1 2)
```

## Repl commands

All repl commands are prefixed with the tilde (`~`).

### Syms command

Defining two variables and printing the symbol table using the `~syms` command:

```
ß :: (let a "hello world")(let b 1 2 3)
= [hello world [1 2 3]]
ß :: ~syms
map[string]interface {}{"a":"hello world", "b":[]interface {}{1, 2, 3}}
ß ::
```

### Funs command

Defining a function and printing the function table using the `~funs` command:

```
ß :: (fun square (_ a) (* a a))
= [<nil>]
ß :: ~funs
map[string]interface {}{"square":(*expr.Func)(0xc0000ac000)}
ß ::
```

### Debug command

The debug command enables debug logging in the repl, which will log:

- timing information
  - lexer starting, took x ns
  - parser starting, took x ns
- lexed tokens
- ast

```
ß ::~debug
12:22:32.748067 toggled debug logging to='true'
ß ::
```
