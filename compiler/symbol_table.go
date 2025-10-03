package compiler

type SymbolScope string

const (
	LocalScope   SymbolScope = "LOCAL"
	GlobalScope  SymbolScope = "GLOBAL"
	BuiltinScope SymbolScope = "BUILTIN"
	FreeScope    SymbolScope = "FREE"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	Outer *SymbolTable

	store          map[string]Symbol
	numDefinitions int
	FreeSymbols    []Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store:       make(map[string]Symbol),
		FreeSymbols: []Symbol{},
	}
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	s := NewSymbolTable()
	s.Outer = outer
	return s
}

func (st *SymbolTable) Define(name string) Symbol {
	sym := Symbol{Name: name, Index: st.numDefinitions}
	if st.Outer == nil {
		sym.Scope = GlobalScope
	} else {
		sym.Scope = LocalScope
	}

	st.store[name] = sym
	st.numDefinitions++
	return sym
}

func (st *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	sym := Symbol{Name: name, Index: index, Scope: BuiltinScope}
	st.store[name] = sym
	return sym
}

func (st *SymbolTable) DefineFree(original Symbol) Symbol {
	st.FreeSymbols = append(st.FreeSymbols, original)

	symbol := Symbol{Name: original.Name, Scope: FreeScope, Index: len(st.FreeSymbols) - 1}

	st.store[original.Name] = symbol
	return symbol
}

func (st *SymbolTable) Resolve(name string) (Symbol, bool) {
	obj, ok := st.store[name]
	if !ok && st.Outer != nil {
		obj, ok := st.Outer.Resolve(name)
		if !ok || obj.Scope == GlobalScope || obj.Scope == BuiltinScope {
			return obj, ok
		}

		free := st.DefineFree(obj)
		return free, true
	}
	return obj, ok
}
