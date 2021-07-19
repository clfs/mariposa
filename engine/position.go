package engine

/*
type Position struct {
	Board         [64]uint8
	ColorToMove   uint8
	WhiteOOO      bool
	WhiteOO       bool
	BlackOOO      bool
	BlackOO       bool
	EPTarget      uint8
	HalfMoveClock uint8
	FullMoveNum   uint16
}

// NewPosition returns a new Position.
func NewPosition() *Position {
	// I'll need to fill this out when I implement hash tables.
	return &Position{}
}

func (p *Position) FromFEN(fen string) error {
	fields := strings.Split(fen, " ")

	// 1. Piece placement
	square := A8
	for _, c := range fields[0] {
		switch c {
		case 'P', 'p', 'N', 'n', 'B', 'b', 'R', 'r', 'Q', 'q', 'K', 'k':
			p.Board[square] = PieceFromRune(c)
		case '1', '2', '3', '4', '5', '6', '7', '8':
			square += int(c) * East
		case '/':
			square += South
		}
		square++
	}

	// 2. Active color
	// 3. Castling availability
	// 4. En passant target square
	// 5. Halfmove clock
	// 6. Fullmove number
	return nil
}

// String lets Position implement fmt.Stringer.
func (p *Position) String() string {
	return "TODO"
}

// FEN returns the FEN record for the position.
func (p *Position) FEN() string {
	return "TODO"
}
*/
