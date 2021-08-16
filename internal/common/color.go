package common

import (
	"math/rand"
	"reflect"
)

type Color uint8

//go:generate stringer -type=Color
const (
	White Color = iota
	Black
)

func (c Color) Value() uint8 {
	return uint8(c)
}

func (c Color) Valid() bool {
	return c <= Black
}

func (c Color) Invalid() bool {
	return !c.Valid()
}

func (c Color) FEN() string {
	switch c {
	case White:
		return "w"
	case Black:
		return "b"
	default:
		return "-"
	}
}

func (Color) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	return reflect.ValueOf(Color(rand.Intn(2)))
}
