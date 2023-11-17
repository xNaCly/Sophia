// walks the ast and performs several optimisations, such as:
//   - dead code elimination for empty if, for, unused functions and variables
//   - replacing variable names with integers for faster map access
//   - precomputing constants if leafs are of type expr.Float or expr.Boolean
package optimizer

import "sophia/core/expr"

// Possible optimisations found by executing and benchmarking example/leetcode.phia:
// Implemented:
// - move float64 parsing from Node.Eval() to the parser - done
// - Fastpath for expr.castPanicIfNotType via expr.castFloatPanic &
//   expr.castBoolPanic to skip a heap allocation - done
// - introduce token pointers instead of copies, could be faster because less
//   memory usage
// - Reuse variables in Node.Eval(), should reduce gc pressure and thus time
//   spent in runtime.mallocgc
// - reduce function calls in hot paths and the interpreter
// Planned:
// - Replace variable names with integers -> should reduce time spend in
//   runtime.mapassign_faststr and aeshashbody (watch out for error handling,
//   etc)
// - precompute constants
// - inline functions called once

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
//   - Functions, if and match statements as well as for loops are furthermore
//     removed if their body contain no expressions, rendering them useless and
//     unnecessary pressure on the evaluation stage.
//
//   - Expressions with no side effects such as being stored in a variable or
//     printed are also subject to removal
type Optimiser struct {
	functionDefinitions map[string]struct{} // keeps track of all defined functions
	functionUsage       map[string]struct{} // keeps track of all used functions
	variableDefinitions map[string]struct{} // keeps track of all defined variables
	variableUsage       map[string]struct{} // keeps track of all used variables
}

type Path struct {
	sideEffect bool // true if result of path is stored in a variable, is a function or printed
	constant   bool // true if all nodes in the path can be computed before the evaluation stage
}

func (o *Optimiser) Start(ast []expr.Node) {}
