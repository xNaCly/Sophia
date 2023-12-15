# Sophia

## Datatypes

Sophia features four data types:

| Datatype | Description                                              | Examples                         |
| -------- | -------------------------------------------------------- | -------------------------------- |
| float    | 64Bit floating point number                              | `.1`, `1e-3`, `1.1`, `1_000_000` |
| string   | text, multiple and single characters                     | `"Hello world"`, `"t"`, `"!!!"`  |
| bool     | boolean                                                  | `true`, `false`                  |
| array    | list that is able to contain all of the above            | `[1 2 3]`, `[1 "test" true]`     |
| objects  | key value pairs that is able to contain all of the above | `{}`, `{ name: "anon" age: 25 }` |

## Printing

> When talking about operators all build in keywords and symbols are meant.
> Keywords are textual words with multiple characters, operators are solely
> symbols. Both words are used synonymously in the documentation.

The most useful keyword is the `println` keyword. It prints all arguments to the
standard out stream. Our current knowledge about Sophia can be applied to create
the famous `Hello World` example:

```lisp
(println "Hello World")
```

The following chapter explains how to execute the expression.

## Executing expressions

When invoking the interpreter without any arguments the repl is started:

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
sophia>
```

The REPL accepts all valid Sophia expressions as well as several commands (read more [here](#repl-commands))

Running the aforementioned `Hello World` can therefore be simply typed into the REPL and executed with pressing `ENTER`:

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
sophia> (println "Hello World!")
Hello World!
= [<nil>]
sophia>
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

## Template strings

Sophia supports interpolation similar to rust or javascript via the following syntax:

```
(let name "anon")
(let money 500_912.99)
(println 'Hi "{money}", you have {money}€ in the bank!')
;; Hi "anon", you have 500912.99€ in the bank!
```

## Merging lists and strings

The `++` operator can be applied to lists, strings, booleans, floats:

```lisp
(++ "hello" "world")
;; => "hello world"

(++ 0.1 12e1)
;; => [0.1 120]

(++ true false "t")
;; => [true false "t"]

(let arr 1 2 3)
;; add 1 to 'arr' and return it as a new list
(++ arr 1)
;; [1 2 3 1]
(++ 1 arr)
;; [1 1 2 3]

#[1 2 3 5] ;; defining a list
```

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
(= 1 2 1)
(> 1 10)
(< 20 1)
```

`=` evaluates to true if all arguments are the same. `<` evaluates to true if
the first argument is smaller than the second. `>` evaluates to true if the
first argument is bigger than the second.

## Controlflow

Sophia features conditional evaluation as well as iterating over containers:

### Conditional evaluation

`if` evaluates the first argument, if it returns true all following arguments
are executed.

```lisp
(if
    true                ;; if-head
    (println "true!")       ;; if-body
    (println "true!")       ;; if-body
    (println "true!")       ;; if-body
)

(if
    (and true false)
    (println "true!")
    (println "true!")
    (println "true!")
)
```

### Iteration

Lets take a look at iteration over containers:

```lisp
(let array 1 2 3 4)
(for
    (_ i)               ;; loop variable
    array               ;; container to iterate over
    (println i))            ;; for body
```

`for` assigns the current value of the container iteration to the defined loop
variable, here `i`. As shown above, the first argument is the loop variable,
the second variable the container to iterate over and the following arguments
are, similar to the `if` keyword, evaluated.

Sophia supports array range syntax:

```lisp
(for
    (_ i)
    100
    (println i))
```

### Match

```lisp
(let a true)
(let b false)
(match
    (if a (println "a true"))
    (if b (println "b true"))
    (println i)
```

Match creates an environment which contains guards (`if`), similar to a switch
or a match in languages like c or rust. The guards are evaluated top to bottom,
the first matched guard exits the match statement.

The statement which is not a guard, is executed if all other guards do not match.

## Including sources

Sophia supports the modularisation of source code into multiple files. To do
this simply create the following files and compile the entry point
(`main.phia`) using `sophia main.phia`:

```lisp
;; square.phia
(fun square (_ n)
    (* n n))
;; main.phia
(load "square.phia")
(println
    (square 12))
```

The same expression can be used in the repl to archive the same effect of importing expressions from an other file:

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
sophia> (load "square.phia")
= [<nil>]
sophia> (square 12)
= [144]
sophia>
```

## Objects

Objects are created using the `{}` notation. In an object a key value pair is
denoted as `key: value`, see below for an example:

```lisp
(let person { name: "anon" age: 25})
```

## Accessing object and list elements

Objects are only as useful as the data they contain, therefore the syntax for
accessing said data is as intuitive as possible:

```lisp
(let person { name: "anon" age: 25})
;; accessing the 'name' value from the 'person' object
(println person["name"])
```

The same can be applied to lists:

```lisp
(let list 1 2 3 4)
;; accessing the first element in the 'l'-list
(println list[0])
```

<!-- Updating the values at either the key or the index is possible via the same syntax: -->

<!-- ```lisp -->
<!-- (let list.0 5) -->
<!-- (let person.name "unknown") -->
<!-- (println list person) -->
<!-- [5 2 3 4] map[age:25 name:unknown] -->
<!-- ``` -->

<!-- Objects support one more feature, which is the addition of new keys also via the same syntax: -->

<!-- ```lisp -->
<!-- (let person[newKey] "thisIsNew") -->
<!-- (println person) -->
<!--  map[age:25 name:unknown newKey:thisisNew] -->
<!-- ``` -->

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
