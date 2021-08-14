package common

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/clfs/mariposa/internal/bitboard"
)

//go:generate stringer -type=Square -linecomment=true
type Square uint8

const NumSquares = 64

const (
	A1 Square = iota // a1
	B1               // b1
	C1               // c1
	D1               // d1
	E1               // e1
	F1               // f1
	G1               // g1
	H1               // h1
	A2               // a2
	B2               // b2
	C2               // c2
	D2               // d2
	E2               // e2
	F2               // f2
	G2               // g2
	H2               // h2
	A3               // a3
	B3               // b3
	C3               // c3
	D3               // d3
	E3               // e3
	F3               // f3
	G3               // g3
	H3               // h3
	A4               // a4
	B4               // b4
	C4               // c4
	D4               // d4
	E4               // e4
	F4               // f4
	G4               // g4
	H4               // h4
	A5               // a5
	B5               // b5
	C5               // c5
	D5               // d5
	E5               // e5
	F5               // f5
	G5               // g5
	H5               // h5
	A6               // a6
	B6               // b6
	C6               // c6
	D6               // d6
	E6               // e6
	F6               // f6
	G6               // g6
	H6               // h6
	A7               // a7
	B7               // b7
	C7               // c7
	D7               // d7
	E7               // e7
	F7               // f7
	G7               // g7
	H7               // h7
	A8               // a8
	B8               // b8
	C8               // c8
	D8               // d8
	E8               // e8
	F8               // f8
	G8               // g8
	H8               // h8
)

func (s Square) Value() uint8 {
	return uint8(s)
}

func (s Square) Valid() bool {
	return s <= H8
}

func (s Square) Invalid() bool {
	return !s.Valid()
}

func (s Square) File() File {
	return File(s % 8)
}

func (s Square) Rank() Rank {
	return Rank(s / 8)
}

func (s Square) Bitboard() bitboard.B {
	return bitboard.B(1 << s.Value())
}

func (s Square) Equal(a Square) bool {
	return s == a || (s.Invalid() && a.Invalid())
}

func (Square) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Square(rand.Intn(64)))
}

func SquareAt(f File, r Rank) (Square, error) {
	if f.Invalid() || r.Invalid() {
		return 0, fmt.Errorf("TODO")
	}
	return Square(r.Value()*8 + f.Value()), nil
}

func ParseSquare(s string) (Square, error) {
	if len(s) != 2 {
		return 0, fmt.Errorf("TODO")
	}
	f := File(s[0] - 'a')
	r := Rank(s[1] - '1')
	return SquareAt(f, r)
}
