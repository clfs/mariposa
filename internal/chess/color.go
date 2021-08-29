package chess

import (
	"fmt"
	"math/rand"
	"reflect"
)

// Color represents a color - either white or black.
type Color uint8

const (
	White Color = iota
	Black
)

func ParseColorFEN(s string) (Color, error) {
	switch s {
	case "w":
		return White, nil
	case "b":
		return Black, nil
	default:
		return 0, fmt.Errorf("invalid color %s", s)
	}
}

// Flipped returns the opposite color.
func (c Color) Flipped() Color {
	return c ^ 1
}

// FEN returns the FEN representation of a color. It returns an empty string if
// the color is invalid.
func (c *Color) FEN() string {
	switch *c {
	case White:
		return "w"
	case Black:
		return "b"
	default:
		return ""
	}
}

// Generate lets Color satisfy testing/quick.Generator.
func (Color) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Color(rand.Intn(2)))
}
