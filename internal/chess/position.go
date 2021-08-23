package chess

import (
	"fmt"
	"strconv"
	"strings"
)

// StartingPositionFEN is the FEN for the starting position.
const StartingPositionFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

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
	fields := strings.Fields(fen)
	if len(fields) != 6 {
		return fmt.Errorf("invalid FEN %s", fen)
	}

	if err := p.Board.setFromBoardFEN(fields[0]); err != nil {
		return fmt.Errorf("invalid FEN %s: %v", fen, err)
	}

	color, err := ParseColorFEN(fields[1])
	if err != nil {
		return fmt.Errorf("invalid FEN %s: %v", fen, err)
	}
	p.SideToMove = color

	castlingRights, err := ParseCastlingRightsFEN(fields[2])
	if err != nil {
		return fmt.Errorf("invalid FEN %s: %v", fen, err)
	}
	p.Castling = castlingRights

	enPassantRight, err := ParseEnPassantRightFEN(fields[3])
	if err != nil {
		return fmt.Errorf("invalid FEN %s: %v", fen, err)
	}
	p.EnPassant = enPassantRight

	halfMoveClock, err := strconv.ParseUint(fields[4], 10, 8)
	if err != nil {
		return fmt.Errorf("invalid FEN %s: %v", fen, err)
	}
	p.HalfMoveClock = uint8(halfMoveClock)

	fullMoveCount, err := strconv.ParseUint(fields[5], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid FEN %s: %v", fen, err)
	}
	p.FullMoveCount = uint16(fullMoveCount)

	return nil
}
