package chess

import "fmt"

//go:generate stringer -type Piece -linecomment=true
type Piece uint8

const (
	WhitePawn   Piece = iota + 1 // P
	WhiteKnight                  // N
	WhiteBishop                  // B
	WhiteRook                    // R
	WhiteQueen                   // Q
	WhiteKing                    // K
	_
	_
	BlackPawn   // p
	BlackKnight // n
	BlackBishop // b
	BlackRook   // r
	BlackQueen  // q
	BlackKing   // k
)

func (p Piece) Value() uint8 {
	return uint8(p)
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
		return 0, fmt.Errorf("TODO")
	}
}
