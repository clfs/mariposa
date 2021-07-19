package engine

import (
	"math/rand"
	"reflect"
)

type (
	Color     uint8
	Role      uint8
	Piece     uint8
	Square    uint8
	File      uint8
	Rank      uint8
	Direction int8
)

const (
	NumColors     = 2
	NumRoles      = 6
	NumPieces     = 12
	NumSquares    = 64
	NumFiles      = 8
	NumRanks      = 8
	NumDirections = 8
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

// Pieces have a color (bit 3) and a role (bits 2-0).
const (
	WhitePawn   Piece = Piece(uint8(White<<3) | uint8(Pawn))
	WhiteKnight Piece = Piece(uint8(White<<3) | uint8(Knight))
	WhiteBishop Piece = Piece(uint8(White<<3) | uint8(Bishop))
	WhiteRook   Piece = Piece(uint8(White<<3) | uint8(Rook))
	WhiteQueen  Piece = Piece(uint8(White<<3) | uint8(Queen))
	WhiteKing   Piece = Piece(uint8(White<<3) | uint8(King))
	BlackPawn   Piece = Piece(uint8(Black<<3) | uint8(Pawn))
	BlackKnight Piece = Piece(uint8(Black<<3) | uint8(Knight))
	BlackBishop Piece = Piece(uint8(Black<<3) | uint8(Bishop))
	BlackRook   Piece = Piece(uint8(Black<<3) | uint8(Rook))
	BlackQueen  Piece = Piece(uint8(Black<<3) | uint8(Queen))
	BlackKing   Piece = Piece(uint8(Black<<3) | uint8(King))
)

// Squares start at A1 = 0. They increase left to right, and then bottom to top.
// Squares have a rank (bits 5-3) and a file (bits 2-0).
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

// Directions.
const (
	North     Direction = 8
	East                = 1
	South               = -North
	West                = -East
	Northeast           = North + East
	Northwest           = North + West
	Southeast           = South + East
	Southwest           = South + West
)

func (p Piece) Role() Role {
	return Role(p & 0b111)
}

func (p Piece) Color() Color {
	return Color(p >> 3 & 1)
}

func (s Square) File() File {
	return File(s & 0b111)
}

func (s Square) Rank() Rank {
	return Rank(s >> 3 & 0b111)
}

func PieceFromColorRole(c Color, r Role) Piece {
	return Piece(uint8(c)<<3 | uint8(r))
}

func SquareFromFileRank(f File, r Rank) Square {
	return Square(uint8(r)<<3 | uint8(f))
}

// Generate lets Color implement testing/quick.Generator.
func (Color) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	c := Color(rand.Intn(NumColors))
	return reflect.ValueOf(c)
}

// Generate lets Role implement testing/quick.Generator.
func (Role) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	r := Role(rand.Intn(NumRoles))
	return reflect.ValueOf(r)
}

// Generate lets Piece implement testing/quick.Generator.
func (Piece) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	p := Piece(rand.Intn(NumPieces))
	return reflect.ValueOf(p)
}

// Generate lets File implement testing/quick.Generator.
func (File) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	f := File(rand.Intn(NumFiles))
	return reflect.ValueOf(f)
}

// Generate lets Rank implement testing/quick.Generator.
func (Rank) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	r := Rank(rand.Intn(NumRanks))
	return reflect.ValueOf(r)
}

// Generate lets Square implement testing/quick.Generator.
func (Square) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	s := Square(rand.Intn(NumSquares))
	return reflect.ValueOf(s)
}
