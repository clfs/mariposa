package fen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/clfs/mariposa/internal/board"
	"github.com/clfs/mariposa/internal/common"
	"github.com/clfs/mariposa/internal/position"
)

func Parse(fen string) (*position.Position, error) {
	p := position.New()

	fields := strings.Fields(fen)
	if len(fields) != 6 {
		return nil, &ParseFENError{FEN: fen}
	}

	// Field 0: Piece placement.
	board, err := ParseBoard(fields[0])
	if err != nil {
		return nil, &ParseFENError{FEN: fen}
	}
	p.Board = *board

	// Field 1: Side to move.
	color, err := ParseColor(fields[1])
	if err != nil {
		return nil, &ParseFENError{FEN: fen}
	}
	p.SideToMove = color

	// Field 2: Castling rights.
	castlingRights, err := ParseCastlingRights(fields[2])
	if err != nil {
		return nil, &ParseFENError{FEN: fen}
	}
	p.Castling = castlingRights

	// Field 3: En passant right.
	enPassantRight, err := ParseEnPassantRight(fields[3])
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

func ParseBoard(s string) (*board.Board, error) {
	var b *board.Board
	square := common.A8
	for _, r := range s {
		switch r {
		case 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k':
			piece, err := common.ParsePiece(string(r))
			if err != nil {
				return nil, fmt.Errorf("invalid board: %s", s)
			}
			b.Put(piece, square)
		case '1', '2', '3', '4', '5', '6', '7', '8':
			offset := r - '0'
			square += common.Square(offset)
		case '/':
			square -= 16
		default:
			return nil, fmt.Errorf("invalid board: %s", s)
		}
	}
	return b, nil
}

func ParseColor(s string) (common.Color, error) {
	switch s {
	case "w":
		return common.White, nil
	case "b":
		return common.Black, nil
	default:
		return 0, fmt.Errorf("invalid color: %s", s)
	}
}

func ParseSquare(s string) (common.Square, error) {
	if len(s) != 2 {
		return 0, fmt.Errorf("invalid square: %s", s)
	}
	f := common.File(s[0] - 'a')
	r := common.Rank(s[1] - '1')
	sq := common.SquareAt(f, r)
	if sq.Invalid() {
		return 0, fmt.Errorf("invalid square: %s", s)
	}
	return sq, nil
}

func ParseCastlingRights(s string) (common.CastlingRights, error) {
	var cr common.CastlingRights
	for _, r := range s {
		switch r {
		case 'K':
			cr.Add(common.WhiteOO)
		case 'Q':
			cr.Add(common.WhiteOOO)
		case 'k':
			cr.Add(common.BlackOO)
		case 'q':
			cr.Add(common.BlackOOO)
		case '-':
			return 0, nil // early exit!
		default:
			return 0, fmt.Errorf("invalid castling: %s", s)
		}
	}
	return cr, nil
}

func ParseEnPassantRight(s string) (common.EnPassantRight, error) {
	var e common.EnPassantRight
	if s == "-" {
		return e, nil
	}
	sq, err := ParseSquare(s)
	if err != nil {
		return e, err
	}
	e.Set(sq)
	return e, nil
}

// todo: consider carrying another error inside?
type ParseFENError struct {
	FEN string
}

func (e *ParseFENError) Error() string {
	return fmt.Sprintf("invalid fen: %s", e.FEN)
}
