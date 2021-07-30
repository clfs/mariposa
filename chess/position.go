package chess

type Position struct {
	allyBoard  Board
	enemyBoard Board
	state      State
}

func (p *Position) WhiteBoard() Board {
	if p.allyBoard.IsFlipped() {
		return p.enemyBoard
	}
	return p.allyBoard
}
