# Sophia

My take on a small lisp like programming language named after my girlfriend.

View the docs containing an overview, an in depth overview and a lot of
Examples [here](https://xnacly.github.io/Sophia/)

```lisp
(fun square [n]
    (* n n))

(let n 12)
(let res (square n))

(println '{n}*{n} is {res}')
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
$ sophia -exp '(println "Hello World")'

Hello World!
```

```
$ echo '(println "Hello World")' | sophia

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
sophia> (let person { name: "user" })
= [user]
sophia> (println "Hello World," person#["name"] ":)")
Hello World, user :)
= [<nil>]
sophia>
```
