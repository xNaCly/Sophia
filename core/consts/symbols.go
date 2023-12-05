package consts

// contains all global objects
var SYMBOL_TABLE = make(map[uint32]any, 64)

// contains scope local objects
var SCOPE_TABLE = make(map[uint32]any, 64)
