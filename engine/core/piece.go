package core

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Piece uint8

const NumPieces = 12

const (
	NoPiece     Piece = 0
	WhitePawn   Piece = Piece(uint8(White)<<3 | uint8(Pawn))
	WhiteKnight Piece = Piece(uint8(White)<<3 | uint8(Knight))
	WhiteBishop Piece = Piece(uint8(White)<<3 | uint8(Bishop))
	WhiteRook   Piece = Piece(uint8(White)<<3 | uint8(Rook))
	WhiteQueen  Piece = Piece(uint8(White)<<3 | uint8(Queen))
	WhiteKing   Piece = Piece(uint8(White)<<3 | uint8(King))
	BlackPawn   Piece = Piece(uint8(Black)<<3 | uint8(Pawn))
	BlackKnight Piece = Piece(uint8(Black)<<3 | uint8(Knight))
	BlackBishop Piece = Piece(uint8(Black)<<3 | uint8(Bishop))
	BlackRook   Piece = Piece(uint8(Black)<<3 | uint8(Rook))
	BlackQueen  Piece = Piece(uint8(Black)<<3 | uint8(Queen))
	BlackKing   Piece = Piece(uint8(Black)<<3 | uint8(King))
)

func (p Piece) IsValid() bool {
	return p.Color().IsValid() && p.Role().IsValid()
}

func (p Piece) Color() Color {
	return Color(p >> 3)
}

func (p Piece) Role() Role {
	return Role(p & 7)
}

func (p Piece) String() string {
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
		return fmt.Sprintf("Piece(%d)", uint8(p))
	}
}

// Generate implements quick.Generator. It only generates pre-defined
// constants of type Piece.
func (Piece) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	x := []Piece{
		WhitePawn, WhiteKnight, WhiteBishop, WhiteRook, WhiteQueen, WhiteKing,
		BlackPawn, BlackKnight, BlackBishop, BlackRook, BlackQueen, BlackKing,
		NoPiece,
	}
	return reflect.ValueOf(Piece(x[rand.Intn(len(x))]))
}

func PieceFromString(s string) Piece {
	switch s {
	case "P":
		return WhitePawn
	case "N":
		return WhiteKnight
	case "B":
		return WhiteBishop
	case "R":
		return WhiteRook
	case "Q":
		return WhiteQueen
	case "K":
		return WhiteKing
	case "p":
		return BlackPawn
	case "n":
		return BlackKnight
	case "b":
		return BlackBishop
	case "r":
		return BlackRook
	case "q":
		return BlackQueen
	case "k":
		return BlackKing
	default:
		return NoPiece
	}
}

func PieceFromColorRole(c Color, r Role) Piece {
	return Piece(uint8(c)<<3 | uint8(r))
}
