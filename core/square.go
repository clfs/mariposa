package core

import "fmt"

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
	NoSquare
)

func (s Square) IsValid() bool {
	return s.File().IsValid() && s.Rank().IsValid()
}

// String returns strings like "e3" and "a1". It returns "-" if
// the square is invalid.
func (s Square) String() string {
	if !s.IsValid() {
		return "-"
	}
	return fmt.Sprintf("%s%s", s.File(), s.Rank())
}

func (s Square) DebugString() string {
	if !s.IsValid() {
		return fmt.Sprintf("Square(%d)", uint8(s))
	}
	return fmt.Sprintf("%s%s", s.File(), s.Rank())
}

func (s Square) File() File {
	return File(s % 8)
}

func (s Square) Rank() Rank {
	return Rank(s / 8)
}

func SquareFromString(s string) Square {
	if len(s) != 2 {
		return NoSquare
	}
	f := File(s[0] - 'a')
	r := Rank(s[1] - '1')
	if !f.IsValid() || !r.IsValid() {
		return NoSquare
	}
	return SquareFromFileRank(f, r)
}

func SquareFromFileRank(f File, r Rank) Square {
	return Square(uint8(r)<<3 | uint8(f))
}
