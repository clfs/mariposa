package chess

type Bitboard uint64

// Value returns the bitboard's integer value.
func (b *Bitboard) Value() uint64 {
	return uint64(*b)
}

// Get returns the bit at the given square.
func (b *Bitboard) Get(s Square) bool {
	return (*b & (1 << s.Value())) > 0
}

// Set sets the bit for the given square and returns b.
func (b *Bitboard) Set(s Square) *Bitboard {
	*b |= (1 << s.Value())
	return b
}

// Clear clears the bit for the given square and returns b.
func (b *Bitboard) Clear(s Square) *Bitboard {
	*b &^= (1 << s.Value())
	return b
}

// Toggle toggles the bit for the given square and returns b.
func (b *Bitboard) Toggle(s Square) *Bitboard {
	*b ^= (1 << s.Value())
	return b
}

// Reset sets all bits to 0 and returns b.
func (b *Bitboard) Reset() *Bitboard {
	*b = 0
	return b
}
