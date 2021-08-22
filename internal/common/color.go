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

// Mirror sets c to the opposite color.
func (c *Color) Mirror() {
	*c = *c ^ 1
}

// FEN returns the FEN representation of a color. If the color is invalid, it
// returns an empty string.
func (c *Color) FEN() string {
	switch *c {
	case White:
		return "w"
	default:
		return "b"
	}
}

// Generate lets Color satisfy testing/quick.Generator.
func (Color) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(Color(rand.Intn(2)))
}
