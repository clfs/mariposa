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
	var (
		c Color
		r Role
	)
	switch {
	case b.Whites.At(s):
		c = White
	case b.Blacks.At(s):
		c = Black
	default:
		return 0, false
	}
	switch {
	case b.Pawns.At(s):
		r = Pawn
	case b.Knights.At(s):
		r = Knight
	case b.Bishops.At(s):
		r = Bishop
	case b.Rooks.At(s):
		r = Rook
	case b.Queens.At(s):
		r = Queen
	case b.Kings.At(s):
		r = King
	default:
		return 0, false
	}
	return NewPiece(c, r), false
}
