package core

import "fmt"

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
