package chess

type Board struct {
	whites  Bitboard
	blacks  Bitboard
	pawns   Bitboard
	knights Bitboard
	bishops Bitboard
	rooks   Bitboard
	queens  Bitboard
	kings   Bitboard
}

func (b *Board) Whites() Bitboard {
	return b.whites
}

func (b *Board) Blacks() Bitboard {
	return b.blacks
}

func (b *Board) Pawns() Bitboard {
	return b.pawns
}

func (b *Board) Knights() Bitboard {
	return b.knights
}

func (b *Board) Bishops() Bitboard {
	return b.bishops
}

func (b *Board) Rooks() Bitboard {
	return b.rooks
}

func (b *Board) Queens() Bitboard {
	return b.queens
}

func (b *Board) Put(p Piece, s Square) {
	switch p.Color() {
	case White:
		b.whites.Set(s)
	case Black:
		b.blacks.Set(s)
	}
	switch p.Role() {
	case Pawn:
		b.pawns.Set(s)
	case Knight:
		b.knights.Set(s)
	case Bishop:
		b.bishops.Set(s)
	case Rook:
		b.rooks.Set(s)
	case Queen:
		b.queens.Set(s)
	case King:
		b.kings.Set(s)
	}
}

func (b *Board) Get(s Square) (Piece, bool) {
	// todo
	return 0, false
}
