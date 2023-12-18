## Embedding

### Embedding the runtime

**Work in Progress**

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
