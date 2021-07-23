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

// String returns "w" for White and "b" for Black. It returns
// "-" if the color is invalid.
func (c Color) String() string {
	switch c {
	case White:
		return "w"
	case Black:
		return "b"
	default:
		return "-"
	}
}

func (c Color) DebugString() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	default:
		return fmt.Sprintf("Color(%d)", uint8(c))
	}
}

// Generate implements quick.Generator. It only generates pre-defined
// constants of type Color.
func (Color) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Color(rand.Intn(NumColors)))
}
