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
	MERGE // ++
	PARAM // _
	IDENT // ([a-z]|_)+
	BOOL  // true | false
	EOF
)

var TOKEN_NAME_MAP = map[int]string{
	UNKNOWN:     "UNKNOWN",
	FLOAT:       "FLOAT",
	STRING:      "STRING",
	BOOL:        "BOOL",
	LEFT_BRACE:  "LEFT_BRACE",
	RIGHT_BRACE: "RIGHT_BRACE",
	ADD:         "PLUS",
	SUB:         "MINUS",
	DIV:         "DIV",
	MUL:         "MUL",
	PARAM:       "PARAM",
	MOD:         "MOD",
	MERGE:       "MERGE",
	PUT:         "PUT",
	LET:         "LET",
	IF:          "IF",
	EQUAL:       "EQ",
	OR:          "OR",
	NEG:         "NOT",
	AND:         "AND",
	FUNC:        "FUN",
	FOR:         "FOR",
	LT:          "LT",
	GT:          "GT",
	MATCH:       "MATCH",
	IDENT:       "IDENT",
	EOF:         "EOF",
}
