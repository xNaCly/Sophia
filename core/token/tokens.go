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
	MERGE
	PARAM // _
	IDENT // ([a-z]|_)+
	BOOL  // true | false
	EOF
)

var TOKEN_NAME_MAP = map[int]string{
	UNKNOWN:     "UNKNOWN",
	FLOAT:       "float",
	STRING:      "string",
	BOOL:        "bool",
	LEFT_BRACE:  "(",
	RIGHT_BRACE: ")",
	ADD:         "+",
	SUB:         "-",
	DIV:         "/",
	MUL:         "*",
	PARAM:       "_",
	MOD:         "%",
	MERGE:       "++",
	PUT:         "put",
	LET:         "let",
	IF:          "if",
	EQUAL:       "eq",
	OR:          "or",
	NEG:         "not",
	AND:         "and",
	FUNC:        "fun",
	FOR:         "for",
	LT:          "lt",
	GT:          "gt",
	IDENT:       "identifier",
	EOF:         "End of file",
}
