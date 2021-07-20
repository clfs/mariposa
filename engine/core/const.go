package core

type (
	Color  uint8
	Role   uint8
	Piece  uint8
	File   uint8
	Rank   uint8
	Square uint8
)

// These exclude values like NoColor and NoRole.
const (
	NumColors  = 2
	NumRoles   = 2
	NumPieces  = 12
	NumFiles   = 8
	NumRanks   = 8
	NumSquares = 64
)

// These represent "zero" values for each type.
// For optimization purposes, they're not actually zero.
const (
	NoColor  Color  = 0xff
	NoRole   Role   = 0xff
	NoPiece  Piece  = 0xff
	NoFile   File   = 0xff
	NoRank   Rank   = 0xff
	NoSquare Square = 0xff
)

const (
	White Color = iota
	Black
)

const (
	Pawn Role = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

const (
	WhitePawn Piece = iota
	WhiteKnight
	WhiteBishop
	WhiteRook
	WhiteQueen
	WhiteKing
	BlackPawn
	BlackKnight
	BlackBishop
	BlackRook
	BlackQueen
	BlackKing
)

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

// Note that Rank1 is 0.
const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

// Squares start at A1 = 0. They increase left to right, and then
// increase bottom to top.
const (
	A1 Square = iota
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
