package chess

//go:generate stringer -type=Color
type Color uint8

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
