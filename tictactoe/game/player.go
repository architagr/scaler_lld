package game

type BasePlayerDTO struct {
	symbol     *SymbolDTO
	name       string
	playerType PlayerType
}

func (player *BasePlayerDTO) GetSymbol() *SymbolDTO {
	return player.symbol
}

func (player *BasePlayerDTO) GetName() string {
	return player.name
}

func (player *BasePlayerDTO) GetPlayerType() PlayerType {
	return player.playerType
}

type IPlayer interface {
	MakeMove(board *BoardDTO) (row int, col int)
	GetPlayerType() PlayerType
	GetName() string
	GetSymbol() *SymbolDTO
}
