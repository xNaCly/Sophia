package token

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
	CONCAT,
	FUNC,
	PARAM,
	IDENT,
	FOR,
	LT,
	GT,
}

var KEYWORD_MAP = map[string]int{
	"put":    PUT,
	"let":    LET,
	"if":     IF,
	"eq":     EQUAL,
	"or":     OR,
	"and":    AND,
	"not":    NEG,
	"concat": CONCAT,
	"fun":    FUNC,
	"for":    FOR,
	"param":  PARAM,
	"ident":  IDENT,
	"gt":     GT,
	"lt":     LT,
}
