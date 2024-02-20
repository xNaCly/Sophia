package token

const (
	UNKNOWN = iota + 1
	// constants
	FLOAT
	STRING
	TEMPLATE_STRING
	IDENT
	BOOL

	// symbols
	ADD
	SUB
	DIV
	MUL
	MOD
	HASHTAG // #

	// structure
	LEFT_CURLY
	RIGHT_CURLY
	COLON
	DOT
	DOUBLE_COLON
	LEFT_BRACE
	RIGHT_BRACE
	LEFT_BRACKET
	RIGHT_BRACKET

	// keywords
	LET
	FUNC
	IF
	EQUAL
	OR
	AND
	NEG
	FOR
	LT
	GT
	MATCH
	LOAD
	MERGE
	RETURN
	MODULE
	LAMBDA

	EOF
)

var TOKEN_NAME_MAP = map[int]string{
	UNKNOWN:         "UNKNOWN",
	FLOAT:           "float",
	STRING:          "string",
	TEMPLATE_STRING: "TEMPLATE_STRING",
	IDENT:           "ident",
	BOOL:            "bool",
	ADD:             "+",
	SUB:             "-",
	DIV:             "/",
	MUL:             "*",
	MOD:             "%",
	LEFT_CURLY:      "{",
	RIGHT_CURLY:     "}",
	COLON:           ":",
	DOUBLE_COLON:    "::",
	DOT:             ".",
	LEFT_BRACE:      "(",
	RIGHT_BRACE:     ")",
	LEFT_BRACKET:    "[",
	RIGHT_BRACKET:   "]",
	LET:             "let",
	FUNC:            "fun",
	IF:              "if",
	EQUAL:           "eq",
	OR:              "or",
	AND:             "and",
	NEG:             "not",
	FOR:             "for",
	LT:              "lt",
	GT:              "gt",
	MATCH:           "match",
	LOAD:            "load",
	MERGE:           "++",
	EOF:             "EOF",
	RETURN:          "return",
	MODULE:          "module",
	LAMBDA:          "lambda",
	HASHTAG:         "#",
}
