package chess

import (
	"math/rand"
	"reflect"
	"strings"
)

// todo better naming?
type (
	Castling      uint8
	CastlingRight uint8
)

const (
	WhiteOO CastlingRight = 1 << iota
	WhiteOOO
	BlackOO
	BlackOOO
)

func (c *Castling) Value() uint8 {
	return uint8(*c)
}

func (c *Castling) Set(r CastlingRight) {
	*c |= Castling(r)
}

func (c *Castling) Clear(r CastlingRight) {
	*c &= ^Castling(r)
}

func (c *Castling) HasRight(r CastlingRight) bool {
	return *c&Castling(r) != 0
}

func (c *Castling) NoRights() bool {
	return *c == 0
}

func (c *Castling) String() string {
	if c.NoRights() {
		return "-"
	}
	var b strings.Builder
	if c.HasRight(WhiteOO) {
		b.WriteString("K")
	}
	if c.HasRight(WhiteOOO) {
		b.WriteString("Q")
	}
	if c.HasRight(BlackOO) {
		b.WriteString("k")
	}
	if c.HasRight(BlackOOO) {
		b.WriteString("q")
	}
	return b.String()
}

func (Castling) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Castling(rand.Intn(1 << 4)))
}

func (CastlingRight) Generate(rand *rand.Rand, size int) reflect.Value {
	x := []CastlingRight{WhiteOO, WhiteOOO, BlackOO, BlackOOO}
	return reflect.ValueOf(CastlingRight(x[rand.Intn(len(x))]))
}
