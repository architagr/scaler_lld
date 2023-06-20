package models

type symbolDTO struct {
	char byte
}

func (symbol *symbolDTO) GetChar() byte {
	return symbol.char
}
