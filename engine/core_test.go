package engine_test

import (
	"testing"

	. "github.com/clfs/mariposa/engine"
)

func TestPieceToRole(t *testing.T) {
	cases := []struct {
		piece, role uint8
	}{
		{WhitePawn, Pawn},
		{WhiteKnight, Knight},
		{WhiteBishop, Bishop},
		{WhiteRook, Rook},
		{WhiteQueen, Queen},
		{WhiteKing, King},
		{BlackPawn, Pawn},
		{BlackKnight, Knight},
		{BlackBishop, Bishop},
		{BlackRook, Rook},
		{BlackQueen, Queen},
		{BlackKing, King},
	}
	for _, c := range cases {
		if role := PieceToRole(c.piece); role != c.role {
			t.Errorf("PieceToRole(%d) = %d, want %d", c.piece, role, c.role)
		}
	}
}

func TestPieceToColor(t *testing.T) {
	cases := []struct {
		piece, color uint8
	}{
		{WhitePawn, White},
		{WhiteKnight, White},
		{WhiteBishop, White},
		{WhiteRook, White},
		{WhiteQueen, White},
		{WhiteKing, White},
		{BlackPawn, Black},
		{BlackKnight, Black},
		{BlackBishop, Black},
		{BlackRook, Black},
		{BlackQueen, Black},
		{BlackKing, Black},
	}
	for _, c := range cases {
		if color := PieceToColor(c.piece); color != c.color {
			t.Errorf("PieceToColor(%d) = %d, want %d", c.piece, color, c.color)
		}
	}
}

func TestSquareToFile(t *testing.T) {
	cases := []struct {
		square, file uint8
	}{
		{A1, FileA},
		{B1, FileB},
		{C1, FileC},
		{D1, FileD},
		{E1, FileE},
		{F1, FileF},
		{G1, FileG},
		{H1, FileH},
		{A2, FileA},
		{B2, FileB},
		{C2, FileC},
		{D2, FileD},
		{E2, FileE},
		{F2, FileF},
		{G2, FileG},
		{H2, FileH},
		{A3, FileA},
		{B3, FileB},
		{C3, FileC},
		{D3, FileD},
		{E3, FileE},
		{F3, FileF},
		{G3, FileG},
		{H3, FileH},
		{A4, FileA},
		{B4, FileB},
		{C4, FileC},
		{D4, FileD},
		{E4, FileE},
		{F4, FileF},
		{G4, FileG},
		{H4, FileH},
		{A5, FileA},
		{B5, FileB},
		{C5, FileC},
		{D5, FileD},
		{E5, FileE},
		{F5, FileF},
		{G5, FileG},
		{H5, FileH},
		{A6, FileA},
		{B6, FileB},
		{C6, FileC},
		{D6, FileD},
		{E6, FileE},
		{F6, FileF},
		{G6, FileG},
		{H6, FileH},
		{A7, FileA},
		{B7, FileB},
		{C7, FileC},
		{D7, FileD},
		{E7, FileE},
		{F7, FileF},
		{G7, FileG},
		{H7, FileH},
		{A8, FileA},
		{B8, FileB},
		{C8, FileC},
		{D8, FileD},
		{E8, FileE},
		{F8, FileF},
		{G8, FileG},
		{H8, FileH},
	}
	for _, c := range cases {
		if file := SquareToFile(c.square); file != c.file {
			t.Errorf("SquareToFile(%d) = %d, want %d", c.square, file, c.file)
		}
	}
}

func TestSquareToRank(t *testing.T) {
	cases := []struct {
		square, rank uint8
	}{
		{A1, Rank1},
		{B1, Rank1},
		{C1, Rank1},
		{D1, Rank1},
		{E1, Rank1},
		{F1, Rank1},
		{G1, Rank1},
		{H1, Rank1},
		{A2, Rank2},
		{B2, Rank2},
		{C2, Rank2},
		{D2, Rank2},
		{E2, Rank2},
		{F2, Rank2},
		{G2, Rank2},
		{H2, Rank2},
		{A3, Rank3},
		{B3, Rank3},
		{C3, Rank3},
		{D3, Rank3},
		{E3, Rank3},
		{F3, Rank3},
		{G3, Rank3},
		{H3, Rank3},
		{A4, Rank4},
		{B4, Rank4},
		{C4, Rank4},
		{D4, Rank4},
		{E4, Rank4},
		{F4, Rank4},
		{G4, Rank4},
		{H4, Rank4},
		{A5, Rank5},
		{B5, Rank5},
		{C5, Rank5},
		{D5, Rank5},
		{E5, Rank5},
		{F5, Rank5},
		{G5, Rank5},
		{H5, Rank5},
		{A6, Rank6},
		{B6, Rank6},
		{C6, Rank6},
		{D6, Rank6},
		{E6, Rank6},
		{F6, Rank6},
		{G6, Rank6},
		{H6, Rank6},
		{A7, Rank7},
		{B7, Rank7},
		{C7, Rank7},
		{D7, Rank7},
		{E7, Rank7},
		{F7, Rank7},
		{G7, Rank7},
		{H7, Rank7},
		{A8, Rank8},
		{B8, Rank8},
		{C8, Rank8},
		{D8, Rank8},
		{E8, Rank8},
		{F8, Rank8},
		{G8, Rank8},
		{H8, Rank8},
	}
	for _, c := range cases {
		if rank := SquareToRank(c.square); rank != c.rank {
			t.Errorf("SquareToRank(%d) = %d, want %d", c.square, rank, c.rank)
		}
	}
}

func TestSquareFromFileRank(t *testing.T) {
	cases := []struct {
		file, rank, square uint8
	}{
		{FileA, Rank1, A1},
		{FileA, Rank2, A2},
		{FileA, Rank3, A3},
		{FileA, Rank4, A4},
		{FileA, Rank5, A5},
		{FileA, Rank6, A6},
		{FileA, Rank7, A7},
		{FileA, Rank8, A8},
		{FileB, Rank1, B1},
		{FileB, Rank2, B2},
		{FileB, Rank3, B3},
		{FileB, Rank4, B4},
		{FileB, Rank5, B5},
		{FileB, Rank6, B6},
		{FileB, Rank7, B7},
		{FileB, Rank8, B8},
		{FileC, Rank1, C1},
		{FileC, Rank2, C2},
		{FileC, Rank3, C3},
		{FileC, Rank4, C4},
		{FileC, Rank5, C5},
		{FileC, Rank6, C6},
		{FileC, Rank7, C7},
		{FileC, Rank8, C8},
		{FileD, Rank1, D1},
		{FileD, Rank2, D2},
		{FileD, Rank3, D3},
		{FileD, Rank4, D4},
		{FileD, Rank5, D5},
		{FileD, Rank6, D6},
		{FileD, Rank7, D7},
		{FileD, Rank8, D8},
		{FileE, Rank1, E1},
		{FileE, Rank2, E2},
		{FileE, Rank3, E3},
		{FileE, Rank4, E4},
		{FileE, Rank5, E5},
		{FileE, Rank6, E6},
		{FileE, Rank7, E7},
		{FileE, Rank8, E8},
		{FileF, Rank1, F1},
		{FileF, Rank2, F2},
		{FileF, Rank3, F3},
		{FileF, Rank4, F4},
		{FileF, Rank5, F5},
		{FileF, Rank6, F6},
		{FileF, Rank7, F7},
		{FileF, Rank8, F8},
		{FileG, Rank1, G1},
		{FileG, Rank2, G2},
		{FileG, Rank3, G3},
		{FileG, Rank4, G4},
		{FileG, Rank5, G5},
		{FileG, Rank6, G6},
		{FileG, Rank7, G7},
		{FileG, Rank8, G8},
		{FileH, Rank1, H1},
		{FileH, Rank2, H2},
		{FileH, Rank3, H3},
		{FileH, Rank4, H4},
		{FileH, Rank5, H5},
		{FileH, Rank6, H6},
		{FileH, Rank7, H7},
		{FileH, Rank8, H8},
	}
	for _, c := range cases {
		if square := SquareFromFileRank(c.file, c.rank); square != c.square {
			t.Errorf("SquareFromFileRank(%d, %d) = %d, want %d", c.file, c.rank, square, c.square)
		}
	}
}

func TestPieceToRune(t *testing.T) {
	cases := []struct {
		piece uint8
		r     rune
	}{
		{WhitePawn, 'P'},
		{WhiteKnight, 'N'},
		{WhiteBishop, 'B'},
		{WhiteRook, 'R'},
		{WhiteQueen, 'Q'},
		{WhiteKing, 'K'},
		{BlackPawn, 'p'},
		{BlackKnight, 'n'},
		{BlackBishop, 'b'},
		{BlackRook, 'r'},
		{BlackQueen, 'q'},
		{BlackKing, 'k'},
		{NoPiece, '-'},
	}
	for _, c := range cases {
		r := PieceToRune(c.piece)
		if r != c.r {
			t.Errorf("PieceToRune(%d) = %c, want %c", c.piece, r, c.r)
		}
	}
}

func TestPieceFromRune(t *testing.T) {
	cases := []struct {
		r rune
		p uint8
	}{
		{'P', WhitePawn},
		{'N', WhiteKnight},
		{'B', WhiteBishop},
		{'R', WhiteRook},
		{'Q', WhiteQueen},
		{'K', WhiteKing},
		{'p', BlackPawn},
		{'n', BlackKnight},
		{'b', BlackBishop},
		{'r', BlackRook},
		{'q', BlackQueen},
		{'k', BlackKing},
		{'-', NoPiece},
		{'x', NoPiece},
	}
	for _, c := range cases {
		p := PieceFromRune(c.r)
		if p != c.p {
			t.Errorf("PieceFromRune(%c) = %d, want %d", c.r, p, c.p)
		}
	}
}
