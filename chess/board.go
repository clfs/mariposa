package chess

type Board struct {
	allies       Bitboard
	enemies      Bitboard
	pawns        Bitboard
	knights      Bitboard
	cardinals    Bitboard // N, E, S, W
	ordinals     Bitboard // NW, NE, SE, SW
	friendlyKing Square
	enemyKing    Square
	isFlipped    bool
}

func (b *Board) Allies() Bitboard {
	return b.allies
}

func (b *Board) Enemies() Bitboard {
	return b.enemies
}

func (b *Board) Pawns() Bitboard {
	return b.pawns
}

func (b *Board) Knights() Bitboard {
	return b.knights
}

func (b *Board) Bishops() Bitboard {
	return b.ordinals &^ b.cardinals
}

func (b *Board) Rooks() Bitboard {
	return b.cardinals &^ b.ordinals
}

func (b *Board) Queens() Bitboard {
	return b.cardinals & b.ordinals
}

func (b *Board) IsFlipped() bool {
	return b.isFlipped
}

// todo
func (b *Board) Flip() {
}
