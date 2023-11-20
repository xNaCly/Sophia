// walks the ast and performs several optimisations, such as:
//   - dead code elimination for empty if, for, unused functions and variables
//   - replacing variable names with integers for faster map access
//   - precomputing constants if leafs are of type expr.Float or expr.Boolean
package optimizer

import (
	"math/rand"
	"sophia/core/debug"
	"sophia/core/expr"
	"strings"
)

var alphabet = []rune("0123456789ABCDEF")
var alphabetlen = len(alphabet)

// TODO: Replace variable names with integers -> should reduce time spend in
// runtime.mapassign_faststr and aeshashbody (watch out for error handling,
// etc)

// TODO: precompute constants

// TODO: dead code elim, empty if, match, put, for, fun and all references to them

// Optimisations
//   - Dead code elimination
//   - Precomputed constants, less load on the evaluation stage especially in
//     loops of repeated function calls
//   - Replacing variable and function names with integers for faster map access
//
// Dead code elimination:
//
//   - Variables and functions are subject to removal if the optimizer walked the
//     tree, populated the counters and determined that a variable or function is
//     defined but not used.
//
//   - Functions, if, match and put statements as well as for loops are
//     furthermore removed if their body contain no expressions, rendering them
//     useless and unnecessary pressure on the evaluation stage.
//
//   - All statements referencing empty functions are removed, such as
//     variables or expressions calling these functions
type Optimiser struct {
	variables   map[string]Node // counter for keeping track of functions
	functions   map[string]Node // counter for keeping track of variables
	builder     strings.Builder
	didOptimise bool
}

type Node struct {
	Used   bool
	Parent expr.Node
	Child  expr.Node
}

func New() *Optimiser {
	return &Optimiser{
		variables: map[string]Node{},
		functions: map[string]Node{},
		builder:   strings.Builder{},
	}
}

func (o *Optimiser) Start(ast []expr.Node) []expr.Node {
	astHolder := &expr.Root{
		Children: make([]expr.Node, len(ast)),
	}
	copy(astHolder.Children, ast)

	// walk ast and populate counters for unused variables and functions
	for _, node := range ast {
		o.walkAst(astHolder, node)
	}

	// unused variables
	for k, v := range o.variables {
		if v.Used {
			continue
		}
		if v.Parent == nil {
			continue
		}
		ch := v.Parent.GetChildren()
		if ch == nil {
			continue
		}
		for i, c := range ch {
			if c == v.Child {
				ch = append(ch[:i], ch[i+1:]...)
				v.Parent.SetChildren(ch)
				t := v.Child.GetToken()
				debug.Logf("removed: %T(%s) [%d:%d]\n", v.Child, k, t.Line+1, t.LinePos)
				delete(o.variables, k)
				o.didOptimise = true
				break
			}
		}
	}

	// unused functions
	for k, v := range o.functions {
		if v.Used {
			continue
		}
		if v.Parent == nil {
			continue
		}
		ch := v.Parent.GetChildren()
		if ch == nil {
			continue
		}
		for i, c := range ch {
			if c == v.Child {
				ch = append(ch[:i], ch[i+1:]...)
				v.Parent.SetChildren(ch)
				t := v.Child.GetToken()
				debug.Logf("removed: %T(%s) [%d:%d]\n", v.Child, k, t.Line+1, t.LinePos)
				delete(o.functions, k)
				o.didOptimise = true
				break
			}
		}
	}

	if o.didOptimise {
		o.didOptimise = false
		return o.Start(astHolder.Children)
	}

	return astHolder.Children
}

// postFix appends a random id of length 5 to val, returns result
func (o *Optimiser) postFix(val string) string {
	o.builder.WriteString(val)
	o.builder.WriteRune('#')
	for i := 0; i < 5; i++ {
		o.builder.WriteRune(alphabet[rand.Intn(alphabetlen)])
	}
	defer o.builder.Reset()
	return o.builder.String()
}

func (o *Optimiser) walkAst(parent, node expr.Node) {
	if node == nil {
		return
	}

	switch v := node.(type) {
	case *expr.Func:
		name := v.Name.GetToken().Raw
		if fun, ok := o.functions[name]; ok {
			o.functions[o.postFix(name)] = fun
		}
		o.functions[name] = Node{Used: false, Parent: parent, Child: v}
	case *expr.Call:
		name := v.Token.Raw
		// detects a function usage, updates counter
		if val, ok := o.functions[name]; ok && !val.Used {
			c := val
			c.Used = true
			o.functions[name] = c
		}
	case *expr.Var:
		// detect a variable definition
		name := v.Ident.GetToken().Raw
		if variable, ok := o.variables[name]; ok {
			o.variables[o.postFix(name)] = variable
		}
		o.variables[name] = Node{Used: false, Parent: parent, Child: v}
	case *expr.Ident:
		// detects a variable usage, updates counter
		name := v.Name
		if val, ok := o.variables[name]; ok && !val.Used {
			c := val
			c.Used = true
			o.variables[name] = c
			// !ok impossible codepath
		}
	}

	children := node.GetChildren()
	for _, c := range children {
		o.walkAst(node, c)
	}
}
