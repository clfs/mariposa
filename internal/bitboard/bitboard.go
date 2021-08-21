package bitboard

import (
	"strings"

	"github.com/clfs/mariposa/internal/common"
)

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

func (b *B) Pretty() string {
	var sb strings.Builder
	for r := 8; r >= 0; r-- {
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
		sb.WriteString("\n")
	}
	return sb.String()
}
