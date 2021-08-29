package chess

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Square uint8

const NumSquares = 64

const (
	A1 Square = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

// NewSquare returns a new square from coordinates.
func NewSquare(f File, r Rank) Square {
	return Square(uint8(r)*8 + uint8(f))
}

func ParseSquareFEN(s string) (Square, error) {
	if len(s) != 2 {
		return 0, fmt.Errorf("invalid square: %s", s)
	}
	f := File(s[0] - 'a')
	r := Rank(s[1] - '1')
	sq := NewSquare(f, r)
	// TODO: validate square
	return sq, nil
}

// Valid returns whether the square is valid.
func (s Square) Valid() bool {
	return s < NumSquares
}

// File returns the file the square is on.
func (s Square) File() File {
	return File(s % 8)
}

// Rank returns the rank the square is on.
func (s Square) Rank() Rank {
	return Rank(s / 8)
}

// Flipped returns the vertically opposite square.
func (s Square) Flipped() Square {
	return s ^ 56
}

// EnPassantRight returns an en passant right that targets the square.
func (s Square) EnPassantRight() EnPassantRight {
	return EnPassantRight(enPassantAllowedMask | s)
}

// Bitboard returns a bitboard with just the square set.
func (s Square) Bitboard() Bitboard {
	return Bitboard(1 << s)
}

// Generate lets Square satisfy testing/quick.Generator.
func (Square) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Square(rand.Intn(NumSquares)))
}

func (s Square) FEN() string {
	switch s {
	case A1:
		return "a1"
	case B1:
		return "b1"
	case C1:
		return "c1"
	case D1:
		return "d1"
	case E1:
		return "e1"
	case F1:
		return "f1"
	case G1:
		return "g1"
	case H1:
		return "h1"
	case A2:
		return "a2"
	case B2:
		return "b2"
	case C2:
		return "c2"
	case D2:
		return "d2"
	case E2:
		return "e2"
	case F2:
		return "f2"
	case G2:
		return "g2"
	case H2:
		return "h2"
	case A3:
		return "a3"
	case B3:
		return "b3"
	case C3:
		return "c3"
	case D3:
		return "d3"
	case E3:
		return "e3"
	case F3:
		return "f3"
	case G3:
		return "g3"
	case H3:
		return "h3"
	case A4:
		return "a4"
	case B4:
		return "b4"
	case C4:
		return "c4"
	case D4:
		return "d4"
	case E4:
		return "e4"
	case F4:
		return "f4"
	case G4:
		return "g4"
	case H4:
		return "h4"
	case A5:
		return "a5"
	case B5:
		return "b5"
	case C5:
		return "c5"
	case D5:
		return "d5"
	case E5:
		return "e5"
	case F5:
		return "f5"
	case G5:
		return "g5"
	case H5:
		return "h5"
	case A6:
		return "a6"
	case B6:
		return "b6"
	case C6:
		return "c6"
	case D6:
		return "d6"
	case E6:
		return "e6"
	case F6:
		return "f6"
	case G6:
		return "g6"
	case H6:
		return "h6"
	case A7:
		return "a7"
	case B7:
		return "b7"
	case C7:
		return "c7"
	case D7:
		return "d7"
	case E7:
		return "e7"
	case F7:
		return "f7"
	case G7:
		return "g7"
	case H7:
		return "h7"
	case A8:
		return "a8"
	case B8:
		return "b8"
	case C8:
		return "c8"
	case D8:
		return "d8"
	case E8:
		return "e8"
	case F8:
		return "f8"
	case G8:
		return "g8"
	case H8:
		return "h8"
	default:
		return ""
	}
}
