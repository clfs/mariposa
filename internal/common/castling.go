package common

import (
	"math/rand"
	"reflect"
	"strings"
)

type (
	CastlingRights uint8
	CastlingRight  uint8
)

//go:generate stringer -type=CastlingRight
const (
	WhiteOO CastlingRight = 1 << iota
	WhiteOOO
	BlackOO
	BlackOOO
)

func NewCastlingRights(r ...CastlingRight) *CastlingRights {
	c := CastlingRights(0)
	for _, rr := range r {
		c.Set(rr)
	}
	return &c
}

func (c *CastlingRights) Set(r CastlingRight) *CastlingRights {
	*c |= CastlingRights(r)
	return c
}

func (c *CastlingRights) Clear(r CastlingRight) *CastlingRights {
	*c &^= CastlingRights(r)
	return c
}

func (c *CastlingRights) Get(r CastlingRight) bool {
	return (*c & CastlingRights(r)) != 0
}

func (c *CastlingRights) FEN() string {
	var b strings.Builder
	if c.Get(WhiteOO) {
		b.WriteString("K")
	}
	if c.Get(WhiteOOO) {
		b.WriteString("Q")
	}
	if c.Get(BlackOO) {
		b.WriteString("k")
	}
	if c.Get(BlackOOO) {
		b.WriteString("q")
	}
	s := b.String()
	if s == "" {
		return "-"
	}
	return s
}

func (CastlingRight) Generate(rand *rand.Rand, size int) reflect.Value {
	all := []CastlingRight{WhiteOO, WhiteOOO, BlackOO, BlackOOO}
	return reflect.ValueOf(all[rand.Intn(len(all))])
}

func (CastlingRights) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(CastlingRights(rand.Intn(0b10000)))
}
