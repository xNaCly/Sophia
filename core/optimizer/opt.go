// walks the ast and performs several optimisations, such as:
//   - dead code elimination for empty if, for, unused functions and variables
//   - replacing variable names with integers for faster map access
//   - precomputing constants if leafs are of type expr.Float or expr.Boolean
package optimizer

import (
	"sophia/core/debug"
	"sophia/core/expr"
)

// TODO: Replace variable names with integers -> should reduce time spend in
// runtime.mapassign_faststr and aeshashbody (watch out for error handling,
// etc)

// TODO: precompute constants

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
	nodes       []NodeTuple // stores variables and functions that are possible defined but not usnot used
	emptyNodes  []NodeTuple // stores expressions that are possible empty
	didOptimise bool
}

type NodeTuple struct {
	Name   string
	Parent expr.Node
	Child  expr.Node
}

func New() *Optimiser {
	return &Optimiser{
		nodes:      []NodeTuple{},
		emptyNodes: []NodeTuple{},
	}
}

func (o *Optimiser) Start(ast []expr.Node) []expr.Node {
	astHolder := &expr.Root{
		Children: ast,
	}

	// walk ast and populate counters for unused variables and functions
	for _, node := range ast {
		o.walkAst(astHolder, node)
	}

	// unused variables and functions
	for i := 0; i < len(o.nodes); i++ {
		tuple := o.nodes[i]
		if tuple.Parent == nil {
			continue
		}
		ch := tuple.Parent.GetChildren()
		if ch == nil {
			continue
		}
		for i, c := range ch {
			if c == tuple.Child {
				ch[i] = ch[len(ch)-1]
				ch = ch[:len(ch)-1]
				tuple.Parent.SetChildren(ch)
				debug.Logf("removed: %T(%s) [%d:%d]\n", tuple.Child, tuple.Name, tuple.Child.GetToken().Line+1, tuple.Child.GetToken().LinePos)
				o.didOptimise = true
				break
			}
		}
	}

	// dead code removal
	for i := 0; i < len(o.emptyNodes); i++ {
		tuple := o.emptyNodes[i]
		if tuple.Parent == nil {
			continue
		}
		ch := tuple.Parent.GetChildren()
		if ch == nil {
			continue
		}
		for i, c := range ch {
			if c == tuple.Child {
				ch[i] = ch[len(ch)-1]
				ch = ch[:len(ch)-1]
				tuple.Parent.SetChildren(ch)
				debug.Logf("removed: %T(%s) [%d:%d]\n", tuple.Child, tuple.Name, tuple.Child.GetToken().Line+1, tuple.Child.GetToken().LinePos)
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

func (o *Optimiser) removeNodeByName(nodes []NodeTuple, name string) []NodeTuple {
	for i, k := range nodes {
		if k.Name == name {
			nodes[i] = nodes[len(nodes)-1]
			return nodes[:len(nodes)-1]
		}
	}
	return nodes
}

func (o *Optimiser) containsNode(nodes []NodeTuple, name string) bool {
	for _, k := range nodes {
		if k.Name == name {
			return true
		}
	}
	return false
}

func (o *Optimiser) isEmpty(node expr.Node) bool {
	if node == nil {
		return false
	} else if len(node.GetChildren()) == 0 {
		return true
	}
	return false
}

func (o *Optimiser) walkAst(parent, node expr.Node) {
	if node == nil {
		return
	}

	switch v := node.(type) {
	case *expr.If, *expr.Match, *expr.For, *expr.Put:
		// empty expressions are subject to removal
		if o.isEmpty(v) {
			o.emptyNodes = append(o.emptyNodes, NodeTuple{Parent: parent, Child: v})
		}
	case *expr.Func:
		// detect a function definition
		o.nodes = append(o.nodes, NodeTuple{Name: v.Name.GetToken().Raw, Parent: parent, Child: v})

		// empty node are subject to removal
		if o.isEmpty(v) {
			o.emptyNodes = append(o.emptyNodes, NodeTuple{Name: v.Name.GetToken().Raw, Parent: parent, Child: v})
		}
	case *expr.Call:
		// detects a function usage, removes the item from the unused functions tracker
		o.removeNodeByName(o.nodes, v.Token.Raw)

		// if a function with a matching name is subject to removal we want to remove the call as well
		if o.containsNode(o.emptyNodes, v.Token.Raw) {
			o.emptyNodes = append(o.emptyNodes, NodeTuple{Name: v.Token.Raw, Parent: parent, Child: v})
		}
	case *expr.Var:
		// detect a variable definition
		o.nodes = append(o.nodes, NodeTuple{Name: v.Ident.GetToken().Raw, Parent: parent, Child: v})

		if o.isEmpty(v) {
			o.emptyNodes = append(o.emptyNodes, NodeTuple{Parent: parent, Child: v})
		}
	case *expr.Ident:
		// if a variable with a matching name is subject to removal we want to remove its uses as well
		if o.containsNode(o.emptyNodes, v.Name) {
			o.emptyNodes = append(o.emptyNodes, NodeTuple{Name: v.Name, Parent: parent, Child: v})
		}

		// detects a variable usage, removes the item from the tracker
		o.removeNodeByName(o.nodes, v.Name)

	}

	children := node.GetChildren()
	for _, c := range children {
		o.walkAst(node, c)
	}
}
