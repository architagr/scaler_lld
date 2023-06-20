package models

type MoveDTO struct {
	player *PlayerDTO
	cell   *CellDTO
}

func (move *MoveDTO) GetPlayer() *PlayerDTO {
	return move.player
}
func (move *MoveDTO) GetCell() *CellDTO {
	return move.cell
}

func InitMove(cell *CellDTO, player *PlayerDTO) *MoveDTO {
	return &MoveDTO{
		cell:   cell,
		player: player,
	}
}
