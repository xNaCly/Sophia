package token

var CONSTANTS = []int{
	FLOAT,
	STRING,
	IDENT,
	BOOL,
	HASHTAG,    // array
	LEFT_CURLY, // object
}

// operators
var EXPECTED_KEYWORDS = []int{
	ADD,
	SUB,
	DIV,
	MUL,
	MOD,
	LET,
	IF,
	EQUAL,
	OR,
	AND,
	NEG,
	MERGE,
	FUNC,
	IDENT,
	FOR,
	LT,
	GT,
	MATCH,
	LOAD,
	RETURN,
	MODULE,
	LAMBDA,
}

var KEYWORD_MAP = map[string]int{
	"return": RETURN,
	"let":    LET,
	"if":     IF,
	"or":     OR,
	"and":    AND,
	"not":    NEG,
	"fun":    FUNC,
	"for":    FOR,
	"ident":  IDENT,
	"match":  MATCH,
	"load":   LOAD,
	"lambda": LAMBDA,
	"module": MODULE,
}
