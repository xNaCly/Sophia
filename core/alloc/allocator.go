package alloc

var Default = Allocator{
	varCount:  0,
	funcCount: 0,
	Functions: map[string]uint32{},
	Variables: map[string]uint32{},
}

func NewFunc(name string) uint32 {
	return Default.NewFunc(name)
}

func NewVar(name string) uint32 {
	return Default.NewVar(name)
}

type Allocator struct {
	Functions map[string]uint32
	funcCount uint32
	Variables map[string]uint32
	varCount  uint32
}

func (a *Allocator) NewFunc(name string) uint32 {
	a.funcCount++
	a.Functions[name] = a.funcCount
	return a.funcCount
}

func (a *Allocator) NewVar(name string) uint32 {
	a.varCount++
	a.Variables[name] = a.varCount
	return a.varCount
}
