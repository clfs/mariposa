package bitboard

import "github.com/clfs/mariposa/internal/common"

// B is a bitboard.
type B uint64

// Value returns the bitboard's integer value.
func (b *B) Value() uint64 {
	return uint64(*b)
}

// Get returns the bit at s and returns b.
func (b *B) Get(s common.Square) bool {
	return (*b & (1 << s.Value())) > 0
}

// Set sets the bit at s and returns b.
func (b *B) Set(s common.Square) *B {
	*b |= (1 << s.Value())
	return b
}

// Clear clears the bit at s and returns b.
func (b *B) Clear(s common.Square) *B {
	*b &= ^(1 << s.Value())
	return b
}

// Toggle toggles the bit at s and returns b.
func (b *B) Toggle(s common.Square) *B {
	*b ^= (1 << s.Value())
	return b
}
