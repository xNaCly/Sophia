package token

const (
	UNKNOWN     = iota + 1
	FLOAT       // 0.0
	STRING      // "text"
	ADD         // +
	SUB         // -
	DIV         // /
	MUL         // *
	PUT         // .
	MOD         // %
	COLON       // :
	LEFT_BRACE  // (
	RIGHT_BRACE // )
	IF          // ?
	EQUAL       // =
	OR          // |
	AND         // &
	NEG         // !
	CONCAT      // ,
	FUNC        // #
	PARAM       // _
	IDENT       // ([a-z]|_)+
	BOOL        // true | false
	EOF
)

var TOKEN_NAME_MAP = map[int]string{
	UNKNOWN:     "UNKNOWN",
	FLOAT:       "FLOAT",
	STRING:      "STRING",
	ADD:         "+",
	SUB:         "-",
	DIV:         "/",
	MUL:         "*",
	PUT:         ".",
	MOD:         "%",
	COLON:       ":",
	IF:          "?",
	EQUAL:       "=",
	OR:          "|",
	NEG:         "!",
	AND:         "&",
	CONCAT:      ",",
	FUNC:        "$",
	PARAM:       "_",
	IDENT:       "IDENT",
	LEFT_BRACE:  "(",
	RIGHT_BRACE: ")",
	BOOL:        "BOOL",
	EOF:         "EOF",
}
