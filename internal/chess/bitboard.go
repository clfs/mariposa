package chess

import (
	"math/bits"
	"strings"
)

// Bitboard is a 64-bit bitboard. From LSB to MSB, the bitboard goes A1, A2,
// ..., H8.
type Bitboard uint64

const (
	// BitboardAllOne is a bitboard with all bits set to 1.
	BitboardAllOne Bitboard = 0xffffffffffffffff
	// BitboardAllZero is a bitboard with all bits set to 0.
	BitboardAllZero Bitboard = 0
)

// Get returns the bit at s.
func (b *Bitboard) Get(s Square) bool {
	return (*b & (1 << uint8(s))) > 0
}

// Set sets the bit at s and returns b.
func (b *Bitboard) Set(s Square) *Bitboard {
	*b |= (1 << uint8(s))
	return b
}

// Clear clears the bit at s and returns b.
func (b *Bitboard) Clear(s Square) *Bitboard {
	*b &= ^(1 << uint8(s))
	return b
}

// Toggle toggles the bit at s and returns b.
func (b *Bitboard) Toggle(s Square) *Bitboard {
	*b ^= (1 << uint8(s))
	return b
}

// Flip flips the bitboard vertically and returns b.
func (b *Bitboard) Flip() *Bitboard {
	*b = Bitboard(bits.ReverseBytes64(uint64(*b)))
	return b
}

// Debug returns a multi-line debug string representing the bitboard. True is 1,
// and false is ".".
func (b *Bitboard) Debug() string {
	var sb strings.Builder
	for r := 7; r >= 0; r-- {
		for f := 0; f < 8; f++ {
			file := File(f)
			rank := Rank(r)
			square := SquareFromCoordinates(file, rank)
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
