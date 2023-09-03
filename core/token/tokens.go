package token

const (
	UNKNOWN = iota + 1
	FLOAT
	STRING
	ADD
	SUB
	DIV
	MUL
	PUT
	MOD
	LET
	LEFT_BRACE
	RIGHT_BRACE
	IF
	EQUAL
	OR
	AND
	NEG
	FUNC
	FOR
	LT
	GT
	MATCH
	TEMPLATE_STRING
	MERGE
	PARAM
	IDENT
	BOOL
	EOF
)

var TOKEN_NAME_MAP = map[int]string{
	UNKNOWN:         "UNKNOWN",
	FLOAT:           "FLOAT",
	STRING:          "STRING",
	BOOL:            "BOOL",
	LEFT_BRACE:      "LEFT_BRACE",
	RIGHT_BRACE:     "RIGHT_BRACE",
	ADD:             "PLUS",
	SUB:             "MINUS",
	DIV:             "DIV",
	MUL:             "MUL",
	PARAM:           "PARAM",
	MOD:             "MOD",
	MERGE:           "MERGE",
	PUT:             "PUT",
	LET:             "LET",
	IF:              "IF",
	EQUAL:           "EQ",
	OR:              "OR",
	NEG:             "NOT",
	AND:             "AND",
	FUNC:            "FUN",
	FOR:             "FOR",
	LT:              "LT",
	GT:              "GT",
	MATCH:           "MATCH",
	IDENT:           "IDENT",
	TEMPLATE_STRING: "TEMPLATE_STRING",
	EOF:             "EOF",
}
