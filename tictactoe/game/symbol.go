package game

type SymbolDTO struct {
	char byte
}

func (symbol *SymbolDTO) GetChar() byte {
	return symbol.char
}

func InitSymbol(ch byte) *SymbolDTO {
	return &SymbolDTO{
		char: ch,
	}
}
