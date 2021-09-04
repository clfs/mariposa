package chess

import (
	"math/bits"
	"strings"
)

// Bitboard is a 64-bit bitboard. From LSB to MSB, the bitboard goes A1, A2,
// ..., H8.
type Bitboard uint64

// Get returns the bit at s.
func (b *Bitboard) Get(s Square) bool {
	return (*b & s.Bitboard()) != 0
}

// Set sets the bit at s to x.
func (b *Bitboard) Set(s Square, x bool) {
	if x {
		b.Set1(s)
	} else {
		b.Set0(s)
	}
}

// Set1 sets the bit at s to 1.
func (b *Bitboard) Set1(s Square) {
	*b |= s.Bitboard()
}

// Set0 sets the bit at s to 0.
func (b *Bitboard) Set0(s Square) {
	*b &^= s.Bitboard()
}

// Toggle toggles the bit at s.
func (b *Bitboard) Toggle(s Square) {
	*b ^= s.Bitboard()
}

// Flip vertically mirrors the bitboard. For example, A1 and A8 swap places.
func (b *Bitboard) Flip() {
	*b = Bitboard(bits.ReverseBytes64(uint64(*b)))
}

// Debug returns a multi-line debug string representing the bitboard. True is
// "1", and false is ".".
func (b *Bitboard) Debug() string {
	var sb strings.Builder
	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			square := NewSquare(File(f), Rank(r))
			if b.Get(square) {
				sb.WriteString("1")
			} else {
				sb.WriteString(".")
			}
		}
		if r != 0 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}
