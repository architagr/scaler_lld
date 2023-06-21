package models

import "tictactoe/contracts"

type MoveDTO struct {
	player contracts.IPlayer
	cell   contracts.ICell
}

func (move *MoveDTO) GetPlayer() contracts.IPlayer {
	return move.player
}
func (move *MoveDTO) GetCell() contracts.ICell {
	return move.cell
}

func InitMove(cell contracts.ICell, player contracts.IPlayer) contracts.IMove {
	return &MoveDTO{
		cell:   cell,
		player: player,
	}
}
