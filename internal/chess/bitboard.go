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

// Set sets the bit at s to x and also returns the bitboard. If x is known
// beforehand, use Set1 or Set0 instead.
func (b *Bitboard) Set(s Square, x bool) *Bitboard {
	if x {
		return b.Set1(s)
	}
	return b.Set0(s)
}

// Set1 sets the bit at s to 1 and also returns the bitboard.
func (b *Bitboard) Set1(s Square) *Bitboard {
	*b |= s.Bitboard()
	return b
}

// Set0 sets the bit at s to 0 and also returns the bitboard.
func (b *Bitboard) Set0(s Square) *Bitboard {
	*b &^= s.Bitboard()
	return b
}

// Toggle toggles the bit at s and also returns the bitboard.
func (b *Bitboard) Toggle(s Square) *Bitboard {
	*b ^= s.Bitboard()
	return b
}

// Flip vertically mirrors the bitboard and also returns it. For example, the
// bit at A1 becomes the bit at A8.
func (b *Bitboard) Flip() *Bitboard {
	*b = Bitboard(bits.ReverseBytes64(uint64(*b)))
	return b
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
