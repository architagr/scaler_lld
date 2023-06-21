package models

import "tictactoe/contracts"

type SymbolDTO struct {
	char byte
}

func (symbol *SymbolDTO) GetChar() byte {
	return symbol.char
}

func InitSymbol(ch byte) contracts.ISymbol {
	return &SymbolDTO{
		char: ch,
	}
}
