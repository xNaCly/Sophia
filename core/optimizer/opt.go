// walks the ast and performs optimisations for a faster runtime
package optimizer

// Possible optimisations found by executing example/leetcode.phia:
// 1. Replace variables with integers -> should reduce time spend in
//    runtime.mapassign_faststr and aeshashbody (watch out for error handling,
//    etc)
// 2. Reuse variables in Node.Eval(), should reduce gc pressure and thus time spent in runtime.mallocgc
// 3. Fastpath for expr.castPanicIfNotType via expr.castFloatPanic &
//    expr.castBoolPanic to skip a heap allocation - done
// 4. introduce token pointers instead of copies, could be faster because less memory usage
// 5. move float64 parsing from Node.Eval() to the parser - done
