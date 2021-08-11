package chess

type State struct {
	EnPassantAllowed bool
	EnPassantTarget  Square
	SideToMove       Color
	Castling         Castling
	HalfMoveClock    uint64
	FullMoveCount    uint64
}
