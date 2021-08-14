package board

import (
	"math/rand"
	"reflect"
)

type Color uint8

//go:generate stringer -type=Color -linecomment=true
const (
	White Color = iota // w
	Black              // b
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

func (Color) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	return reflect.ValueOf(Color(rand.Intn(2)))
}
