# Documentation - Sophia

The sophia langauge is a minimal lisp like programming language. It features
common programming language characteristics such as data types (float, string,
bool, containers), controlflow (if, match, for), comparisons (eq, lt, gt), bool
comparisons (and, or, not) arithmetics (+,-,/,\*,%), declarations (let, fun),
merging lists (++) as well as printing to stdout.

It's implementation can be fed expressions from stdin, the repl, a file or a
flag. The Sophia language is implemented with a tree walk interpreter.

```sophia
(println "Hello World")
```

The Sophia language documentation is split up into several chapters.
Each chapter contains the block for learning the concepts in the next
chapter. The documentation is designed for readers which are considered
experienced programmers and are able to create rudimentary programs in
programming language such as c or python.

## News

For the latest and greatest on sophia lang development, see my blog
[here](https://xnacly.me/tags/sophia/).

## Overview

The documentation is split up into the following files:

| Page                        | Description                                                                                   |
| --------------------------- | --------------------------------------------------------------------------------------------- |
| [Language](./Language.md)   | Overview over the sophia programming language                                                 |
| [Examples](./Examples.md)   | Examples with in depth explanations                                                           |
| [Embedding](./Embedding.md) | Guide for embedding sophia into applications and registering go functions for usage in sophia |
| [Internal](./Internal.md)   | Overview over the inner workings and the code of the interpreter                              |

## Motivation

I was motivated to create an interpreter for a programming language ever since
I have started reading [crafting
interpreters](https://craftinginterpreters.com/). The first project I made
which featured a compiler was [fleck](https://github.com/xnacly/fleck) - a
markdown to html compiler. I kept being intrigued by the different stages of
compiling and still wanted to create a programming language by myself. My main
issue was the parsing, I couldn't wrap my head around operator precedence and
how to implement it. Therefore i chose the lisp inspired syntax for ease of
expression parsing.

## Inspiration

As i said before the syntax is based around
[S-Expressions](https://en.wikipedia.org/wiki/S-expression), which is clearly
lisp inspired. I let myself get inspired by several concepts of the rust
programming language, such as the [range
notation](https://doc.rust-lang.org/rust-by-example/flow_control/for.html) or
the very much less powerful implementation of the [match
statement](https://doc.rust-lang.org/rust-by-example/flow_control/match.html) I
like the list [concatenation
operator](https://zvon.org/other/haskell/Outputprelude/HH_o.html) from Haskell,
so i added it for string and list concatenation. For the [match
statement](https://book.realworldhaskell.org/read/defining-types-streamlining-functions.html#deftypes.guardhttps://book.realworldhaskell.org/read/defining-types-streamlining-functions.html#deftypes.guard)
i was also inspired by Haskell, therefore it uses similar syntax. I liked the
way Python uses [keywords for
operations](https://docs.python.org/3/reference/expressions.html#and) other
programming language use symbols, so i implemented `and` instead of `&&`, `eq`
instead of `==`, you get the gist.

List of features and their inspiration

| Feature                     | Inspiration   | Description                                                  |
| --------------------------- | ------------- | ------------------------------------------------------------ |
| match statement             | rust, haskell | switch statement on speed                                    |
| keywords instead of symbols | python        | `&&`->`and`, `\|\|`->`or`, `==`->`eq`,`<`->`lt`              |
| S-Expressions               | lisp          | `(keyword arguments)`, some people dislike them, i love them |
| list concatenation          | haskell       | merging lists and strings is incredibly useful               |
| loop                        | go,rust       | iterate over containers, such as lists                       |
| fancy lexer error messages  | rust          | context display and extensive debugging information          |

Books and resources i consider inspiring:

- [Writing An Interpreter in Go by Thorsten Ball](https://interpreterbook.com/) (go)
- [Crafting Interpreters by Robert Nystrom](https://craftinginterpreters.com/) (java & c)
- [Compilers: Principles, Techniques, and Tools](https://en.wikipedia.org/wiki/Compilers:_Principles,_Techniques,_and_Tools)

I wrote a blog post about why you should write a programming language which
included the baby steps of this project, read it
[here](https://xnacly.me/posts/2023/write-your-own-programming-language/).

## Performance

The performance is not too bad and not too great. The implementation is
currently using the visitor pattern for evaluation, which means the interpreter
is not fast. I am thinking about a bytecode interpreter rewrite, but I am
nowhere near experienced enough for that yet. The lexer and parser itself
aren't doing much work and are pretty fast - the evaluation is the slowest part
of the interpreter.

For recent optimisiation work, see:

- [Improving Programming Language Performance](https://xnacly.me/posts/2023/language-performance/)
