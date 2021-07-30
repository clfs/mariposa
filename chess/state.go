package chess

type State struct {
	EnPassantAllowed bool
	EnPassantTarget  Square
	Castling         Castling
	PlyCount         uint32
}
