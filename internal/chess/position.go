package chess

import (
	"fmt"
)

// Position represents a chess position.
type Position struct {
	Board         Board
	EnPassant     EnPassantRight
	Castling      CastlingRights
	SideToMove    Color
	HalfMoveClock uint8  // Max half move clock is 75.
	FullMoveCount uint16 // Max full move count is ~4000 (?).
}

// NewPosition returns a new position set to the given FEN.
func NewPosition(fen string) (*Position, error) {
	p := new(Position)
	err := p.setFromFEN(fen)
	return p, err
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

func (p *Position) setFromFEN(fen string) error {
	// TODO: implement.
	return nil
}
