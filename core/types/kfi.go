package types

import "sophia/core/token"

type KnownFunctionInterface func(*token.Token, ...Node) any
