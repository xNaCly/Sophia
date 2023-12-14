package token

var CONSTANTS = []int{
	FLOAT,
	STRING,
	IDENT,
	BOOL,
	LEFT_CURLY, // object
}

// operators
var EXPECTED_KEYWORDS = []int{
	ADD,
	SUB,
	DIV,
	MUL,
	MOD,
	PUT,
	LET,
	IF,
	EQUAL,
	OR,
	AND,
	NEG,
	MERGE,
	FUNC,
	PARAM,
	IDENT,
	FOR,
	LT,
	GT,
	MATCH,
	LOAD,
	RETURN,
}

var KEYWORD_MAP = map[string]int{
	"return": RETURN,
	"put":    PUT,
	"let":    LET,
	"if":     IF,
	"eq":     EQUAL,
	"or":     OR,
	"and":    AND,
	"not":    NEG,
	"fun":    FUNC,
	"for":    FOR,
	"param":  PARAM,
	"ident":  IDENT,
	"gt":     GT,
	"lt":     LT,
	"match":  MATCH,
	"load":   LOAD,
}
