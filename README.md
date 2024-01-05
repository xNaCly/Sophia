# Sophia

My take on a small lisp like embeddable programming language with go
interoperability.

View the docs containing an overview, an in depth overview and a lot of
Examples [here](https://xnacly.github.io/Sophia/)

```lisp
(let arr 1 2 3 4 5)

(fun square [n] (* n n))
(map (square) arr) ;; [1 4 9 16 25]

(filter
    (lambda [n]
        (= (% n 2) 0)) arr) ;; [2 4]

(let person {
    array: [1 2 3]
    bank: {
        institute: {
            name: "western union"
        }
    }
})
(println person#["array"][0]) ;; 1
(println person#["bank"]["institute"]["name"]) ;; "western union"

(let name "anon")
(println 'Hello {name}!')
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
