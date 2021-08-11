package chess

type Bitboard uint64

func NewBitboard() (b Bitboard) {
	return 0
}

func (b *Bitboard) Value() uint64 {
	return uint64(*b)
}

func (b *Bitboard) Clear() {
	*b = 0
}

func (b *Bitboard) Set(s Square) {
	*b |= (1 << s.Value())
}
