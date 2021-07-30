package chess

import (
	"math/rand"
	"reflect"
	"strings"
)

type (
	Castling      uint8
	CastlingRight uint8
)

const (
	FriendOO CastlingRight = 1 << iota
	FriendOOO
	EnemyOO
	EnemyOOO
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

func (c *Castling) Swap() {
	*c = (*c >> 2) | (*c << 2)
}

func (c *Castling) String() string {
	if c.NoRights() {
		return "-"
	}
	var b strings.Builder
	if c.HasRight(FriendOO) {
		b.WriteString("K")
	}
	if c.HasRight(FriendOOO) {
		b.WriteString("Q")
	}
	if c.HasRight(EnemyOO) {
		b.WriteString("k")
	}
	if c.HasRight(EnemyOOO) {
		b.WriteString("q")
	}
	return b.String()
}

func (Castling) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Castling(rand.Intn(1 << 4)))
}

func (CastlingRight) Generate(rand *rand.Rand, size int) reflect.Value {
	x := []CastlingRight{FriendOO, FriendOOO, EnemyOO, EnemyOOO}
	return reflect.ValueOf(CastlingRight(x[rand.Intn(len(x))]))
}
