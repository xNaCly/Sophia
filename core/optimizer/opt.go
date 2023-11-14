// walks the ast and performs several optimisations
package optimizer

// Possible optimisations found by executing and benchmarking example/leetcode.phia for 100_000 executions:

// TODO:  1. Replace variables with integers -> should reduce time spend in runtime.mapassign_faststr and aeshashbody (watch out for error handling, etc)

// TODO:  2. Reuse variables in Node.Eval(), should reduce gc pressure and thus time spent in runtime.mallocgc

// INFO: 3. Fastpath for expr.castPanicIfNotType via expr.castFloatPanic & expr.castBoolPanic to skip a heap allocation - done

// TODO:  4. introduce token pointers instead of copies, could be faster because less memory usage

// INFO: 5. move float64 parsing from Node.Eval() to the parser - done

// TODO:  6. reduce function calls in hot paths and the interpreter
