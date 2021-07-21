package core

import (
	"fmt"
	"strconv"
)

func (c Color) String() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	default:
		return fmt.Sprintf("Color(%d)", uint8(c))
	}
}

func (c Color) IsValid() bool {
	return c == White || c == Black
}

func (r Role) String() string {
	switch r {
	case Pawn:
		return "Pawn"
	case Knight:
		return "Knight"
	case Bishop:
		return "Bishop"
	case Rook:
		return "Rook"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return fmt.Sprintf("Role(%d)", uint8(r))
	}
}

func (r Role) IsValid() bool {
	return r <= King
}

func (p Piece) String() string {
	switch p {
	case WhitePawn:
		return "P"
	case WhiteKnight:
		return "N"
	case WhiteBishop:
		return "B"
	case WhiteRook:
		return "R"
	case WhiteQueen:
		return "Q"
	case WhiteKing:
		return "K"
	case BlackPawn:
		return "p"
	case BlackKnight:
		return "n"
	case BlackBishop:
		return "b"
	case BlackRook:
		return "r"
	case BlackQueen:
		return "q"
	case BlackKing:
		return "k"
	default:
		return fmt.Sprintf("Piece(%d)", uint8(p))
	}
}

func (p Piece) IsValid() bool {
	return p <= BlackKing
}

func (p Piece) Color() Color {
	switch p {
	case WhitePawn, WhiteKnight, WhiteBishop, WhiteRook, WhiteQueen, WhiteKing:
		return White
	case BlackPawn, BlackKnight, BlackBishop, BlackRook, BlackQueen, BlackKing:
		return Black
	default:
		return NoColor
	}
}

func (p Piece) Role() Role {
	switch p {
	case WhitePawn, BlackPawn:
		return Pawn
	case WhiteKnight, BlackKnight:
		return Knight
	case WhiteBishop, BlackBishop:
		return Bishop
	case WhiteRook, BlackRook:
		return Rook
	case WhiteQueen, BlackQueen:
		return Queen
	case WhiteKing, BlackKing:
		return King
	default:
		return NoRole
	}
}

func (f File) String() string {
	if !f.IsValid() {
		return fmt.Sprintf("File(%d)", uint8(f))
	}
	return string('A' + uint8(f))
}

func (f File) IsValid() bool {
	return f <= FileH
}

func (r Rank) String() string {
	if !r.IsValid() {
		return fmt.Sprintf("Rank(%d)", uint8(r))
	}
	return string('1' + uint8(r))
}

func (r Rank) IsValid() bool {
	return r <= Rank8
}

func RankFromString(s string) Rank {
	if len(s) != 1 {
		return NoRank
	}
	r, err := strconv.Atoi(s)
	if err != nil {
		return NoRank
	}
	return Rank(r)
}
