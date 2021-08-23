package chess

import (
	"fmt"
)

type Position struct {
	Board         Board
	EnPassant     EnPassantRight
	Castling      CastlingRights
	SideToMove    Color
	HalfMoveClock uint8  // Max half move clock is 75.
	FullMoveCount uint16 // Max full move count is ~4000 (?).
}

func New() *Position {
	return new(Position)
}

func (p *Position) FEN() (string, error) {
	return fmt.Sprintf(
		"%s %s %s %s %d %d",
		p.Board.FEN(),
		p.SideToMove.FEN(),
		p.Castling.FEN(),
		p.EnPassant.FEN(),
		p.HalfMoveClock,
		p.FullMoveCount,
	), nil
}
