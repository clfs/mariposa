package bitboard

import (
	"math/bits"
	"strings"

	"github.com/clfs/mariposa/internal/common"
)

// Bitboard is a 64-bit bitboard. From LSB to MSB, the bitboard goes A1, A2,
// ..., H8.
type Bitboard uint64

// Get returns the bit at s.
func (b *Bitboard) Get(s common.Square) bool {
	return (*b & (1 << uint8(s))) > 0
}

// Set sets the bit at s and returns b.
func (b *Bitboard) Set(s common.Square) *Bitboard {
	*b |= (1 << uint8(s))
	return b
}

// Clear clears the bit at s and returns b.
func (b *Bitboard) Clear(s common.Square) *Bitboard {
	*b &= ^(1 << uint8(s))
	return b
}

// Toggle toggles the bit at s and returns b.
func (b *Bitboard) Toggle(s common.Square) *Bitboard {
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
			file := common.File(f)
			rank := common.Rank(r)
			square := common.SquareAt(file, rank)
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
