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
	CONCAT
	FUNC
	FOR
	LT
	GT
	PARAM // _
	IDENT // ([a-z]|_)+
	BOOL  // true | false
	EOF
)

var TOKEN_NAME_MAP = map[int]string{
	UNKNOWN:     "UNKNOWN",
	FLOAT:       "float",
	STRING:      "string",
	ADD:         "+",
	SUB:         "-",
	DIV:         "/",
	MUL:         "*",
	PUT:         "put",
	MOD:         "%",
	LET:         "let",
	IF:          "if",
	EQUAL:       "eq",
	OR:          "or",
	NEG:         "not",
	AND:         "and",
	CONCAT:      "concat",
	FUNC:        "fun",
	FOR:         "for",
	LT:          "lt",
	GT:          "gt",
	PARAM:       "_",
	IDENT:       "identifier",
	LEFT_BRACE:  "(",
	RIGHT_BRACE: ")",
	BOOL:        "bool",
	EOF:         "End of file",
}
