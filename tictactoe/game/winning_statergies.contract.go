package game

type IWinningStatergies interface {
	CheckWinner(board *BoardDTO, move *MoveDTO) bool
	Undo(board *BoardDTO, move *MoveDTO)
}
