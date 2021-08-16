package position

import (
	"fmt"

	"github.com/clfs/mariposa/internal/board"
	"github.com/clfs/mariposa/internal/common"
)

type Position struct {
	Board         board.Board
	EnPassant     common.EnPassantRight
	Castling      common.CastlingRights
	SideToMove    common.Color
	HalfMoveClock uint8  // Max half move clock is 75.
	FullMoveCount uint16 // Max full move count is ~4000 (?).
}

func New() *Position {
	return new(Position)
}

func (p *Position) FEN() string {
	return fmt.Sprintf(
		"%s %s %s %s %d %d",
		p.Board.FEN(),
		p.SideToMove.FEN(),
		p.Castling.FEN(),
		p.EnPassant.FEN(),
		p.HalfMoveClock,
		p.FullMoveCount,
	)
}
