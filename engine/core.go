package engine

// Colors (1 bit).
const (
	White = iota
	Black
)

// Roles (3 bits).
const (
	Pawn = iota + 1
	Knight
	Bishop
	Rook
	Queen
	King
)

// Pieces (4 bits).
const (
	WhitePawn   = (White << 3) ^ Pawn
	WhiteKnight = (White << 3) ^ Knight
	WhiteBishop = (White << 3) ^ Bishop
	WhiteRook   = (White << 3) ^ Rook
	WhiteQueen  = (White << 3) ^ Queen
	WhiteKing   = (White << 3) ^ King
	BlackPawn   = (Black << 3) ^ Pawn
	BlackKnight = (Black << 3) ^ Knight
	BlackBishop = (Black << 3) ^ Bishop
	BlackRook   = (Black << 3) ^ Rook
	BlackQueen  = (Black << 3) ^ Queen
	BlackKing   = (Black << 3) ^ King
)

// Squares (6 bits).
const (
	A1 = iota
	B1
	C1
	D1
	E1
	F1
	G1
	H1
	A2
	B2
	C2
	D2
	E2
	F2
	G2
	H2
	A3
	B3
	C3
	D3
	E3
	F3
	G3
	H3
	A4
	B4
	C4
	D4
	E4
	F4
	G4
	H4
	A5
	B5
	C5
	D5
	E5
	F5
	G5
	H5
	A6
	B6
	C6
	D6
	E6
	F6
	G6
	H6
	A7
	B7
	C7
	D7
	E7
	F7
	G7
	H7
	A8
	B8
	C8
	D8
	E8
	F8
	G8
	H8
)

// Files (3 bits).
const (
	FileA = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

// Ranks (3 bits). Note that Rank1 equals 0.
const (
	Rank1 = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

// Directions.
const (
	North     = 8
	East      = 1
	South     = -North
	West      = -East
	Northeast = North + East
	Northwest = North + West
	Southeast = South + East
	Southwest = South + West
)

func PieceToRole(piece uint8) uint8 {
	return piece & 0b111
}

func PieceToColor(piece uint8) uint8 {
	return piece & (1 << 3)
}

func SquareToFile(square uint8) uint8 {
	return square & 0b111
}

func SquareToRank(square uint8) uint8 {
	return square & (0b111 << 3)
}

func SquareFromFileRank(file uint8, rank uint8) uint8 {
	return (file << 3) | rank
}
