package core

var KEYWORDS = map[string]int{
	"add":  ADD,
	"sub":  SUB,
	"div":  DIV,
	"mul":  MUL,
	"putv": PUTV,
}

var EXPECTED_KEYWORDS = func() []int {
	words := make([]int, 0)
	for _, v := range KEYWORDS {
		words = append(words, v)
	}
	return words
}()

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
	PUTV
	LEFT_BRACE  // [
	RIGHT_BRACE // ]
	EOF
)

var TOKEN_NAME_MAP = map[int]string{
	UNKNOWN:     "UNKNOWN",
	FLOAT:       "FLOAT",
	STRING:      "STRING",
	ADD:         "ADD",
	SUB:         "SUB",
	DIV:         "DIV",
	MUL:         "MUL",
	PUTV:        "PUTV",
	LEFT_BRACE:  "LEFT_BRACE",
	RIGHT_BRACE: "RIGHT_BRACE",
	EOF:         "EOF",
}
