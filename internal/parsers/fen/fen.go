package fen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/clfs/mariposa/internal/common"
	"github.com/clfs/mariposa/internal/position"
)

// FromPosition returns the FEN for a position.
func FromPosition(p position.Position) (string, error) {
	var b strings.Builder
	fmt.Fprintf(&b, "%s", p.Board.FEN())
	fmt.Fprintf(&b, " %s", p.SideToMove)
	fmt.Fprintf(&b, " %s", p.Castling.FEN())
	fmt.Fprintf(&b, " %s", p.EnPassant.String()) // todo: why is there a string call?
	fmt.Fprintf(&b, " %d", p.HalfMoveClock)
	fmt.Fprintf(&b, " %d", p.FullMoveCount)
	return b.String(), nil
}

// ToPosition returns the position for a FEN.
func ToPosition(fen string) (*position.Position, error) {
	p := position.New()

	fields := strings.Fields(fen)
	if len(fields) != 6 {
		return nil, &ParseFENError{FEN: fen}
	}

	// Field 0: Piece placement.
	// todo: move this to a separate function.
	square := common.A8
	for _, r := range fields[0] {
		switch r {
		case 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k':
			piece, err := common.ParsePiece(string(r))
			if err != nil {
				return nil, &ParseFENError{FEN: fen}
			}
			p.Board.Put(piece, square)
		case '1', '2', '3', '4', '5', '6', '7', '8':
			offset := r - '0'
			square += common.Square(offset) // TODO: Is there a better way to express this?
		case '/':
			square -= 16
		default:
			return nil, &ParseFENError{FEN: fen}
		}
	}

	// Field 1: Side to move.
	color, err := common.ParseColor(fields[1])
	if err != nil {
		return nil, &ParseFENError{FEN: fen}
	}
	p.SideToMove = color

	// Field 2: Castling rights.
	castlingRights, err := common.ParseCastlingRights(fields[2])
	if err != nil {
		return nil, &ParseFENError{FEN: fen}
	}
	p.Castling = castlingRights

	// Field 3: En passant right.
	enPassantRight, err := common.ParseEnPassantRight(fields[3])
	if err != nil {
		return nil, &ParseFENError{FEN: fen}
	}
	p.EnPassant = enPassantRight

	// Field 4: Half-move clock.
	halfMoveClock, err := strconv.ParseUint(fields[4], 10, 8)
	if err != nil {
		return nil, &ParseFENError{FEN: fen}
	}
	p.HalfMoveClock = uint8(halfMoveClock)

	// Field 5: Full-move count.
	fullMoveCount, err := strconv.ParseUint(fields[5], 10, 16)
	if err != nil {
		return nil, &ParseFENError{FEN: fen}
	}
	p.FullMoveCount = uint16(fullMoveCount)

	return p, nil
}

// todo: consider carrying another error inside?
type ParseFENError struct {
	FEN string
}

func (e *ParseFENError) Error() string {
	return fmt.Sprintf("invalid fen: %s", e.FEN)
}
