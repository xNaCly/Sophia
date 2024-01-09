package types

import "github.com/xnacly/sophia/core/token"

type KnownFunctionInterface func(*token.Token, ...Node) any
