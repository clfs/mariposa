package chess

type State struct {
	EnPassantAllowed bool
	EnPassantTarget  Square
	SideToMove       Color
	Castling         Castling
	HalfMoveClock    uint64
	FullMoveCount    uint64
}

func (s State) Equal(x State) bool {
	// If en passant isn't allowed, we don't care about the target square.
	if !s.EnPassantAllowed && !x.EnPassantAllowed {
		return s.SideToMove == x.SideToMove &&
			s.Castling == x.Castling &&
			s.HalfMoveClock == x.HalfMoveClock &&
			s.FullMoveCount == x.FullMoveCount
	} else {
		return s.EnPassantAllowed == x.EnPassantAllowed &&
			s.EnPassantTarget == x.EnPassantTarget &&
			s.SideToMove == x.SideToMove &&
			s.Castling == x.Castling &&
			s.HalfMoveClock == x.HalfMoveClock &&
			s.FullMoveCount == x.FullMoveCount
	}
}
