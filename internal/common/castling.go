package common

import (
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

func (c *CastlingRights) Add(r CastlingRight) *CastlingRights {
	*c |= CastlingRights(r)
	return c
}

func (c *CastlingRights) Remove(r CastlingRight) *CastlingRights {
	*c &^= CastlingRights(r)
	return c
}

func (c *CastlingRights) IsAllowed(r CastlingRight) bool {
	return (*c & CastlingRights(r)) != 0
}

func (c *CastlingRights) FEN() string {
	var b strings.Builder
	if c.IsAllowed(WhiteOO) {
		b.WriteString("K")
	}
	if c.IsAllowed(WhiteOOO) {
		b.WriteString("Q")
	}
	if c.IsAllowed(BlackOO) {
		b.WriteString("k")
	}
	if c.IsAllowed(BlackOOO) {
		b.WriteString("q")
	}
	s := b.String()
	if s == "" {
		return "-"
	}
	return s
}
