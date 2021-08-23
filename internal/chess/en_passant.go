package chess

import (
	"math/rand"
	"reflect"
)

type EnPassantRight uint8

const (
	enPassantAllowedMask = 0b0100_0000
	enPassantSquareMask  = 0b0011_1111
)

func NewEnPassantRight(s Square) EnPassantRight {
	return EnPassantRight(s | enPassantAllowedMask)
}

func NewEnPassantRightNotAllowed() EnPassantRight {
	return 0
}

func (e *EnPassantRight) Allowed() bool {
	return *e&enPassantAllowedMask != 0
}

func (e *EnPassantRight) square() Square {
	return Square(*e & enPassantSquareMask)
}

func (e *EnPassantRight) Get() (Square, bool) {
	return e.square(), e.Allowed()
}

func (e *EnPassantRight) Set(s Square) {
	*e = EnPassantRight(s | enPassantAllowedMask)
}

func (e *EnPassantRight) Clear() {
	*e = 0
}

func (e *EnPassantRight) Flip() {
	s := e.square()
	s.Flip()
	e.Set(s)
}

func (e *EnPassantRight) FEN() string {
	target, allowed := e.Get()
	if !allowed {
		return "-"
	}
	return target.FEN()
}

// Equal checks if two EnPassantRight represent equivalent en passant rights.
func (e EnPassantRight) Equal(o EnPassantRight) bool {
	eTarget, eAllowed := e.Get()
	oTarget, oAllowed := o.Get()
	return !(eAllowed || oAllowed) || (eTarget == oTarget)
}

// Generate lets EnPassantRight satisfy testing/quick.Generator.
func (EnPassantRight) Generate(r *rand.Rand, size int) reflect.Value {
	n := (enPassantAllowedMask | enPassantSquareMask) + 1
	return reflect.ValueOf(EnPassantRight(rand.Intn(n)))
}
