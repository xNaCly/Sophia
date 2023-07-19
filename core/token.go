package core

var KEYWORDS = map[string]int{}

var EXPECTED_KEYWORDS = []int{
	ADD,
	SUB,
	DIV,
	MUL,
	PWR,
	MOD,
	PUT,
}

type Token struct {
	Pos   int
	Line  int
	Type  int
	Raw   string
	Float float64
}

const (
	UNKNOWN = iota + 1
	FLOAT   // 0.0
	STRING  // "text"
	ADD
	SUB
	DIV
	MUL
	PUT
	PWR
	MOD
	LEFT_BRACE  // [
	RIGHT_BRACE // ]
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
	PWR:         "^",
	MOD:         "%",
	LEFT_BRACE:  "[",
	RIGHT_BRACE: "]",
	EOF:         "EOF",
}
