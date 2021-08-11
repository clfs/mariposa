package chess

type State struct {
	EnPassantAllowed bool
	EnPassantTarget  Square
	SideToMove       Color
	Castling         Castling
	PlyCount         uint64
	MoveCount        uint64
}
