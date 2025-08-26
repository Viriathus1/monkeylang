package compiler

type SymbolScope string

const (
	GlobalScope SymbolScope = "Global"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	store          map[string]Symbol
	numDefinitions int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store: make(map[string]Symbol),
	}
}

func (st *SymbolTable) Define(name string) Symbol {
	sym := Symbol{Name: name, Scope: GlobalScope, Index: st.numDefinitions}
	st.store[name] = sym
	st.numDefinitions++
	return sym
}

func (st *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := st.store[name]
	return obj, ok
}
