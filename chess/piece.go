package chess

//go:generate stringer -type Piece -linecomment=true
type Piece uint8

const (
	WhitePawn   Piece = iota + 1 // P
	WhiteKnight                  // N
	WhiteBishop                  // B
	WhiteRook                    // R
	WhiteQueen                   // Q
	WhiteKing                    // K
	_
	_
	BlackPawn   // p
	BlackKnight // n
	BlackBishop // b
	BlackRook   // r
	BlackQueen  // q
	BlackKing   // k
)

func (p Piece) Value() uint8 {
	return uint8(p)
}

func (p Piece) Valid() bool {
	return (WhitePawn <= p && p <= WhiteKing) || (BlackPawn <= p && p <= BlackKing)
}

func (p Piece) Invalid() bool {
	return !p.Valid()
}

func (p Piece) Color() Color {
	return Color(p & 8 >> 3)
}

func (p Piece) Role() Role {
	return Role(p & 7)
}
