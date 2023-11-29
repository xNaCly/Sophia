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
	PUT
	MOD
	PARAM

	// structure
	LEFT_CURLY
	RIGHT_CURLY
	COLON
	DOT
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

	EOF
)

var TOKEN_NAME_MAP = map[int]string{
	UNKNOWN:         "UNKNOWN",
	FLOAT:           "FLOAT",
	STRING:          "STRING",
	TEMPLATE_STRING: "TEMPLATE_STRING",
	IDENT:           "IDENT",
	BOOL:            "BOOL",
	ADD:             "ADD",
	SUB:             "SUB",
	DIV:             "DIV",
	MUL:             "MUL",
	PUT:             "PUT",
	MOD:             "MOD",
	PARAM:           "PARAM",
	LEFT_CURLY:      "LEFT_CURLY",
	RIGHT_CURLY:     "RIGHT_CURLY",
	COLON:           "COLON",
	DOT:             "DOT",
	LEFT_BRACE:      "LEFT_BRACE",
	RIGHT_BRACE:     "RIGHT_BRACE",
	LEFT_BRACKET:    "LEFT_BRACKET",
	RIGHT_BRACKET:   "RIGHT_BRACKET",
	LET:             "LET",
	FUNC:            "FUNC",
	IF:              "IF",
	EQUAL:           "EQUAL",
	OR:              "OR",
	AND:             "AND",
	NEG:             "NEG",
	FOR:             "FOR",
	LT:              "LT",
	GT:              "GT",
	MATCH:           "MATCH",
	LOAD:            "LOAD",
	MERGE:           "MERGE",
	EOF:             "EOF",
	RETURN:          "RETURN",
}
