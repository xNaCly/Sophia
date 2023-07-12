package main

var KEYWORDS = map[string]int{
	"add":  ADD,
	"sub":  SUB,
	"div":  DIV,
	"mul":  MUL,
	"putv": PUTV,
}

type Token struct {
	Pos   int
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
)
