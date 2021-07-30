package bitboard

type Bitboard uint64

func New(n uint64) (b Bitboard) {
	return 0
}

func (b *Bitboard) Value() uint64 {
	return uint64(*b)
}

func (b *Bitboard) Clear() {
	*b = 0
}
