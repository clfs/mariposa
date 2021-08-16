package common

import (
	"math/rand"
	"reflect"
)

//go:generate stringer -type Piece -linecomment=true
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
	return Piece(c.Value()<<3 | r.Value())
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
