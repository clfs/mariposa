package bitboard

import (
	"math/bits"
	"strings"

	"github.com/clfs/mariposa/internal/common"
)

// B is a bitboard.
type B uint64

// Get returns the bit at s and returns b.
func (b *B) Get(s common.Square) bool {
	return (*b & (1 << uint8(s))) > 0
}

// Set sets the bit at s and returns b.
func (b *B) Set(s common.Square) *B {
	*b |= (1 << uint8(s))
	return b
}

// Clear clears the bit at s and returns b.
func (b *B) Clear(s common.Square) *B {
	*b &= ^(1 << uint8(s))
	return b
}

// Toggle toggles the bit at s and returns b.
func (b *B) Toggle(s common.Square) *B {
	*b ^= (1 << uint8(s))
	return b
}

// Flip flips the bitboard vertically and returns b. For example, A1 becomes A8.
func (b *B) Flip() *B {
	*b = B(bits.ReverseBytes64(uint64(*b)))
	return b
}

func (b *B) Debug() string {
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
