package chess

import (
	"fmt"
	"strconv"
	"strings"
)

// Starting is the FEN for a starting position.
const Starting = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func Parse(fen string) (*Position, error) {
	p := New()

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

func ParseBoard(s string) (*Board, error) {
	var b = new(Board)
	square := A8
	for _, r := range s {
		switch r {
		case 'P', 'N', 'B', 'R', 'Q', 'K', 'p', 'n', 'b', 'r', 'q', 'k':
			piece, err := ParsePiece(string(r))
			if err != nil {
				return nil, fmt.Errorf("invalid board: %s", s)
			}
			b.Put(piece, square)
		case '1', '2', '3', '4', '5', '6', '7', '8':
			offset := r - '0'
			square += Square(offset)
		case '/':
			square -= 16
		default:
			return nil, fmt.Errorf("invalid board: %s", s)
		}
	}
	return b, nil
}

func ParseColor(s string) (Color, error) {
	switch s {
	case "w":
		return White, nil
	case "b":
		return Black, nil
	default:
		return 0, fmt.Errorf("invalid color: %s", s)
	}
}

func ParseSquare(s string) (Square, error) {
	if len(s) != 2 {
		return 0, fmt.Errorf("invalid square: %s", s)
	}
	f := File(s[0] - 'a')
	r := Rank(s[1] - '1')
	sq := SquareFromCoordinates(f, r)
	// TODO: validate square
	return sq, nil
}

func ParsePiece(s string) (Piece, error) {
	switch s {
	case "P":
		return WhitePawn, nil
	case "p":
		return BlackPawn, nil
	case "N":
		return WhiteKnight, nil
	case "n":
		return BlackKnight, nil
	case "B":
		return WhiteBishop, nil
	case "b":
		return BlackBishop, nil
	case "R":
		return WhiteRook, nil
	case "r":
		return BlackRook, nil
	case "Q":
		return WhiteQueen, nil
	case "q":
		return BlackQueen, nil
	case "K":
		return WhiteKing, nil
	case "k":
		return BlackKing, nil
	default:
		return 0, fmt.Errorf("invalid piece: %s", s)
	}
}

func ParseCastlingRights(s string) (CastlingRights, error) {
	var cr CastlingRights
	for _, r := range s {
		switch r {
		case 'K':
			cr.Enable(FriendOO)
		case 'Q':
			cr.Enable(FriendOOO)
		case 'k':
			cr.Enable(EnemyOO)
		case 'q':
			cr.Enable(EnemyOOO)
		case '-':
			return 0, nil // early exit!
		default:
			return 0, fmt.Errorf("invalid castling: %s", s)
		}
	}
	return cr, nil
}

func ParseEnPassantRight(s string) (EnPassantRight, error) {
	e := NewEnPassantRightNotAllowed()
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
