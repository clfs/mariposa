package chess

import (
	"math/rand"
	"reflect"
)

// EnPassantRight represents the right to en passant.
type EnPassantRight uint8

const (
	enPassantAllowedMask = 0b0100_0000
	enPassantSquareMask  = 0b0011_1111
)

// NewEnPassantRight returns a new EnPassantRight.
func NewEnPassantRight() EnPassantRight {
	return 0
}

// Allowed returns whether en passant is allowed.
func (e EnPassantRight) Allowed() bool {
	return e&enPassantAllowedMask != 0
}

func (e EnPassantRight) square() Square {
	return Square(e & enPassantSquareMask)
}

// Get returns the target square (if any) and whether en passant is allowed.
func (e EnPassantRight) Get() (Square, bool) {
	return e.square(), e.Allowed()
}

// Flipped returns a new EnPassantRight with the target square vertically
// flipped (if any).
func (e EnPassantRight) Flipped() EnPassantRight {
	return EnPassantRight(uint8(e&enPassantAllowedMask) | uint8(e.square().Flipped()))
}

// FEN returns the FEN representation of the EnPassantRight.
func (e EnPassantRight) FEN() string {
	target, allowed := e.Get()
	if !allowed {
		return "-"
	}
	return target.FEN()
}

// Generate lets EnPassantRight satisfy testing/quick.Generator.
func (EnPassantRight) Generate(r *rand.Rand, size int) reflect.Value {
	n := (enPassantAllowedMask | enPassantSquareMask) + 1
	return reflect.ValueOf(EnPassantRight(rand.Intn(n)))
}

func ParseEnPassantRightFEN(s string) (EnPassantRight, error) {
	if s == "-" {
		return NewEnPassantRight(), nil
	}
	target, err := ParseSquareFEN(s)
	if err != nil {
		return 0, err
	}
	return target.EnPassantRight(), nil
}
