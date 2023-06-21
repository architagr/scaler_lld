package models

import (
	"tictactoe/contracts"
	"tictactoe/enums"
)

type BasePlayerDTO struct {
	symbol     contracts.ISymbol
	name       string
	playerType enums.PlayerType
}

func (player *BasePlayerDTO) GetSymbol() contracts.ISymbol {
	return player.symbol
}

func (player *BasePlayerDTO) GetName() string {
	return player.name
}

func (player *BasePlayerDTO) GetPlayerType() enums.PlayerType {
	return player.playerType
}
