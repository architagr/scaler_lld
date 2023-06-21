package game

type ColWinningStatergy struct {
	colData []map[*SymbolDTO]int
}

func InitColWinnigStatergy(players []IPlayer, size int) IWinningStatergies {
	colData := make([]map[*SymbolDTO]int, size)
	for i := 0; i < size; i++ {
		colData[i] = make(map[*SymbolDTO]int)
		for _, player := range players {
			colData[i][player.GetSymbol()] = 0
		}
	}
	return &ColWinningStatergy{
		colData: colData,
	}
}
func (winningStatergy *ColWinningStatergy) CheckWinner(board *BoardDTO, move *MoveDTO) bool {
	winningStatergy.colData[move.GetCell().GetCol()][move.GetPlayer().GetSymbol()]++
	return winningStatergy.colData[move.GetCell().GetCol()][move.GetPlayer().GetSymbol()] == board.GetSize()
}
func (winningStatergy *ColWinningStatergy) Undo(board *BoardDTO, move *MoveDTO) {
	winningStatergy.colData[move.GetCell().GetCol()][move.GetPlayer().GetSymbol()]--
}
