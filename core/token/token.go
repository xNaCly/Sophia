package token

type Token struct {
	Pos     int
	Line    int
	LinePos int
	Type    int
	Raw     string
}
