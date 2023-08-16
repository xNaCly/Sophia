# Sophia

## Datatypes

Sophia features four data types:

| Datatype | Description                                   | Examples                         |
| -------- | --------------------------------------------- | -------------------------------- |
| float    | 64Bit floating point number                   | `.1`, `1e-3`, `1.1`, `1_000_000` |
| string   | text, multiple and single characters          | `"Hello world"`, `"t"`, `"!!!"`  |
| bool     | boolean                                       | `true`, `false`                  |
| array    | list that is able to contain all of the above | `[1 2 3]`, `[1 "test" true]`     |

## Keywords

> When talking about operators all build in keywords and symbols are meant.
> Keywords are textual words with multiple characters, operators are solely
> symbols. Both words are used synonymously in the documentation.

The most useful keyword is the `put` keyword. It prints all arguments to the
standard output. Our current knowledge about Sophia can be applied to create
the famous `Hello World` example:

```sophia
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
