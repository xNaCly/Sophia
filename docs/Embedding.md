## Embedding

### Embedding the runtime

#### Installing the runtime

Install sophia as a project dependency:

```shell
$ go get github.com/xnacly/sophia
```

#### Initial embedding

Add skeleton:

```go
package main

import (
	"os"

	"github.com/xnacly/sophia/embed"
)

func main() {
	embed.Embed(embed.Configuration{})
	file, err := os.Open("config.phia")
	if err != nil {
		panic(err)
	}
	embed.Execute(file, nil)
}
```

#### Configuration script

Lets add some configuration we want to modify with our sophia script:

```go
package main

import (
	"fmt"
	"os"

	"github.com/xnacly/sophia/embed"
)

type Configuration struct {
	Port int
}

var config = &Configuration{}

func main() {
	embed.Embed(embed.Configuration{})
	file, err := os.Open("config.phia")
	if err != nil {
		panic(err)
	}
	embed.Execute(file, nil)
}
```

And write a sophia script:

```lisp
(let port 8080)
(set-port port)
```

#### Interfacing with go

Now lets create the function by adding the `Functions` field to the `embed.Configuration` structure:

```go
package main

import (
	"fmt"
	"os"

	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
	"github.com/xnacly/sophia/embed"
)

type Configuration struct {
	Port int
}

var config = &Configuration{}

func main() {
	embed.Embed(embed.Configuration{
		Functions: map[string]types.KnownFunctionInterface{
			"set-port": func(t *token.Token, n ...types.Node) any {
				return nil
			},
		},
	})
	file, err := os.Open("config.phia")
	if err != nil {
		panic(err)
	}
	embed.Execute(file, nil)
}
```

#### Input validation

And do some input validation for the arguments we passed to `set-port`:

```go
package main

import (
	"fmt"
	"os"

	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
	"github.com/xnacly/sophia/embed"
)

type Configuration struct {
	Port int
}

var config = &Configuration{}

func main() {
	embed.Embed(embed.Configuration{
		Functions: map[string]types.KnownFunctionInterface{
			"set-port": func(t *token.Token, n ...types.Node) any {
				if len(n) > 1 {
					serror.Add(n[1].GetToken(), "Too many arguments", "Expected 1 argument for set-port, got %d", len(n))
					serror.Panic()
				}
				return nil
			},
		},
	})
	file, err := os.Open("config.phia")
	if err != nil {
		panic(err)
	}
	embed.Execute(file, nil)
	fmt.Println("port:", config.Port)
}
```

If we pass two ports to our `set-port` function we will get the following error message:

```shell
$ cat config.phia
(let port 8080)
(set-port port port)
$ go run .
error: Too many arguments

        at: /home/teo/programming/embedding_sophia/config.phia:3:16:

            1| ;; vim: syntax=lisp
            2| (let port 8080)
            3| (set-port port port)
             |                ^^^^

Expected 1 argument for set-port, got 2
```

Lets evaluate the result of the argument passed to our function, cast it to a float64 and assign it to `config.Port`:

```go
package main

import (
	"fmt"
	"os"

	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
	"github.com/xnacly/sophia/embed"
)

type Configuration struct {
	Port int
}

var config = &Configuration{}

func main() {
	embed.Embed(embed.Configuration{
		Functions: map[string]types.KnownFunctionInterface{
			"set-port": func(t *token.Token, n ...types.Node) any {
				if len(n) > 1 {
					serror.Add(n[1].GetToken(), "Too many arguments", "Expected 1 argument for set-port, got %d", len(n))
					serror.Panic()
				}
				res := n[0].Eval()
				port, ok := res.(float64)
				if !ok {
					serror.Add(n[0].GetToken(), "Type error", "Expected float64 for port, got %T", res)
					serror.Panic()
				}

				config.Port = int(port)

				return nil
			},
		},
	})
	file, err := os.Open("config.phia")
	if err != nil {
		panic(err)
	}
	embed.Execute(file, nil)
	fmt.Println("port:", config.Port)
}
```

Again, lets check the error handling:

```shell
$ cat config.phia
(let port "8080")
(set-port port)
$ go run .
error: Type error

        at: /home/teo/programming/embedding_sophia/config.phia:3:11:

            1| ;; vim: syntax=lisp
            2| (let port "8080")
            3| (set-port port)
             |           ^^^^

Expected float64 for port, got string
```

#### Resulting embedding of Sophia

Simply running our script with valid inputs according to our previous checks will result in the following output:

```shell
$ cat config.phia
(let port 8080)
(set-port port)
$ go run .
port: 8080
```

### KFI - Known function interface

> KFI is a pun on FFI, because we know our functions and they must be defined
> in the same binary the sophia language runtime is embedded in.

The sophia language includes capabilities for exposing go functions to use
inside of the sophia language, for example see the following function
definition included in `core/builtin/builtin.go`:

#### Example: Linking strings.Split

```go
func init() {
	// [...]
	consts.FUNC_TABLE[alloc.NewFunc("strings-split")] = func(tok *token.Token, n ...types.Node) any {
		if len(n) != 2 {
			serror.Add(tok, "Argument error", "Expected exactly 2 argument for strings-split built-in")
			serror.Panic()
		}
		v := n[0].Eval()
		str, ok := v.(string)
		if !ok {
			serror.Add(tok, "Error", "Can't split target of type %T, use a string", v)
			serror.Panic()
		}

		v = n[1].Eval()
		sep, ok := v.(string)
		if !ok {
			serror.Add(tok, "Error", "Can't split string with anything other than a string (%T)", v)
			serror.Panic()
		}

		out := strings.Split(str, sep)

		// sophia lang runtime only sees arrays containing
		// elements whose types were erased as an array.
		r := make([]any, len(out))
		for i, e := range out {
			r[i] = e
		}

		return r
	}
}
```

This maps the `strings.Split` function from the go standard library to the
`strings-split` sophia function. All functions defined with the KFI have access
to the callees token and all its arguments, for instance:

```lisp
(strings-split "Hello World" "")
;; token: strings-split
;; n: "Hello World", " "
```

The `token` parameter points to `strings-split`, `n` contains 0 or more
arguments to the call, here its `["Hello World", " "]`.

### Example: `typeof`

We can do whatever go and the sophia lang type system allow. You can print an
expressions type without evaluating it:

```go
consts.FUNC_TABLE[alloc.NewFunc("typeof")] = func(tok *token.Token, n ...types.Node) any {
    if len(n) != 1 {
        serror.Add(tok, "Argument error", "Expected exactly 1 argument for typeof built-in")
        serror.Panic()
    }
    return fmt.Sprintf("%T", n[0])
}
```

And call this function from sophia:

```shell
$ cat test.phia; echo "------"; sophia test.phia
(println (typeof [1 "test" test 25.0]))
(println (typeof true))
(println (typeof "test"))
(println (typeof 12))
(println (typeof { key: "value" }))
------
*expr.Array
*expr.Boolean
*expr.String
*expr.Float
*expr.Object
```
