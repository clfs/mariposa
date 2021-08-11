package chess

type Board struct {
	Whites  Bitboard
	Blacks  Bitboard
	Pawns   Bitboard
	Knights Bitboard
	Bishops Bitboard
	Rooks   Bitboard
	Queens  Bitboard
	Kings   Bitboard
}

func (b *Board) Put(p Piece, s Square) {
	switch p.Color() {
	case White:
		b.Whites.Set(s)
	case Black:
		b.Blacks.Set(s)
	}
	switch p.Role() {
	case Pawn:
		b.Pawns.Set(s)
	case Knight:
		b.Knights.Set(s)
	case Bishop:
		b.Bishops.Set(s)
	case Rook:
		b.Rooks.Set(s)
	case Queen:
		b.Queens.Set(s)
	case King:
		b.Kings.Set(s)
	}
}

func (b *Board) Get(s Square) (Piece, bool) {
	// todo
	return 0, false
}
