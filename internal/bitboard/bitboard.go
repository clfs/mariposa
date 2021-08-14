package bitboard

// B is a bitboard.
type B uint64

// Value returns the bitboard's integer value.
func (b B) Value() uint64 {
	return uint64(b)
}
