package game

type MoveDTO struct {
	player IPlayer
	cell   *CellDTO
}

func (move *MoveDTO) GetPlayer() IPlayer {
	return move.player
}
func (move *MoveDTO) GetCell() *CellDTO {
	return move.cell
}

func InitMove(cell *CellDTO, player IPlayer) *MoveDTO {
	return &MoveDTO{
		cell:   cell,
		player: player,
	}
}
