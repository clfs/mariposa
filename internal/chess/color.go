package chess

import (
	"math/rand"
	"reflect"
)

type Color uint8

const (
	White Color = iota
	Black
)

// Mirror sets c to the opposite color.
func (c *Color) Mirror() {
	*c = *c ^ 1
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
