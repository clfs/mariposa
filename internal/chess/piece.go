package chess

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Piece uint8

const (
	WhitePawn Piece = iota + 1
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	_
	_
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

func NewPiece(c Color, r Role) Piece {
	return Piece(uint8(c)<<3 | r.Value())
}

func ParsePieceFEN(s string) (Piece, error) {
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

func (p Piece) Value() uint8 {
	return uint8(p)
}

func (p Piece) FEN() string {
	switch p {
	case WhitePawn:
		return "P"
	case WhiteKnight:
		return "N"
	case WhiteBishop:
		return "B"
	case WhiteRook:
		return "R"
	case WhiteQueen:
		return "Q"
	case WhiteKing:
		return "K"
	case BlackPawn:
		return "p"
	case BlackKnight:
		return "n"
	case BlackBishop:
		return "b"
	case BlackRook:
		return "r"
	case BlackQueen:
		return "q"
	case BlackKing:
		return "k"
	default:
		return ""
	}
}

func (p Piece) Valid() bool {
	return (WhitePawn <= p && p <= WhiteKing) || (BlackPawn <= p && p <= BlackKing)
}

func (p Piece) Invalid() bool {
	return !p.Valid()
}

func (p Piece) Color() Color {
	return Color(p & 8 >> 3)
}

func (p Piece) Role() Role {
	return Role(p & 7)
}

func (Piece) Generate(rand *rand.Rand, size int) reflect.Value {
	allPieces := []Piece{WhitePawn, WhiteKnight, WhiteBishop, WhiteRook, WhiteQueen, WhiteKing, BlackPawn, BlackKnight, BlackBishop, BlackRook, BlackQueen, BlackKing}
	return reflect.ValueOf(allPieces[rand.Intn(len(allPieces))])
}
