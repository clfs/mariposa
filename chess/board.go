package chess

type Board struct {
	Friends      Bitboard
	Enemies      Bitboard
	Pawns        Bitboard
	Knights      Bitboard
	Cardinals    Bitboard
	Ordinals     Bitboard
	FriendlyKing Square
	EnemyKing    Square
	EnPassant    Square
	Castling     Castling
}
