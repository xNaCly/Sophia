package expr

import (
	"github.com/xnacly/sophia/core/consts"
	"github.com/xnacly/sophia/core/serror"
	"github.com/xnacly/sophia/core/token"
	"github.com/xnacly/sophia/core/types"
)

type Use struct {
	Token *token.Token
	Name  types.Node
}

func (u *Use) GetChildren() []types.Node {
	return nil
}

func (u *Use) SetChildren(c []types.Node) {}

func (u *Use) GetToken() *token.Token {
	return u.Token
}

func (u *Use) Eval() any {
	ident, _ := u.Name.(*Ident)
	module, ok := consts.MODULE_TABLE[ident.Name]
	if !ok {
		serror.Add(ident.Token, "Undefined Module", "Can't find a module named %q", ident.Name)
		serror.Panic()
	}
	m := module.(*Module)
	for _, c := range m.Children {
		function, ok := c.(*Func)
		if !ok {
			serror.Add(c.GetToken(), "Type Error", "Expected a function inside a module, got %T", c)
			serror.Panic()
		}
		fName := function.Name.(*Ident)
		fName.Name = ident.Name + "::" + fName.Name
		function.Name = fName

	}
	return nil
}
