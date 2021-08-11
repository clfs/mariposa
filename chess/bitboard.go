package chess

type Bitboard uint64

func NewBitboard() (b *Bitboard) {
	x := Bitboard(0)
	return &x
}

func (b *Bitboard) Value() uint64 {
	return uint64(*b)
}

func (b *Bitboard) Zero() {
	*b = 0
}

func (b *Bitboard) Set(s Square) {
	*b |= (1 << s.Value())
}

func (b *Bitboard) Clear(s Square) {
	*b &^= (1 << s.Value())
}

func (b *Bitboard) At(s Square) bool {
	return (*b & (1 << s.Value())) != 0
}

func (b *Bitboard) Invert() {
	*b = ^*b
}
