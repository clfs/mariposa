package core

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Color uint8

const NumColors = 2

const (
	White Color = iota
	Black
)

func (c Color) IsValid() bool {
	return c <= Black
}

func (c Color) String() string {
	switch c {
	case White:
		return "w"
	case Black:
		return "b"
	default:
		return fmt.Sprintf("Color(%d)", uint8(c))
	}
}

// Generate implements quick.Generator. It only generates pre-defined
// constants of type Color.
func (Color) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Color(rand.Intn(NumColors)))
}
