package core_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/core"
)

func TestPiece_IsValid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p Piece
		b bool
	}{
		{NoPiece, false},
		{WhitePawn, true},
		{WhiteKnight, true},
		{WhiteBishop, true},
		{WhiteRook, true},
		{WhiteQueen, true},
		{WhiteKing, true},
		{BlackPawn, true},
		{BlackKnight, true},
		{BlackBishop, true},
		{BlackRook, true},
		{BlackQueen, true},
		{BlackKing, true},
	}
	for _, c := range cases {
		if got := c.p.IsValid(); got != c.b {
			t.Errorf("%s.IsValid() = %v, want %v", c.p, got, c.b)
		}
	}
}

func TestPiece_String(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p Piece
		s string
	}{
		{NoPiece, "."},
		{WhitePawn, "P"},
		{WhiteKnight, "N"},
		{WhiteBishop, "B"},
		{WhiteRook, "R"},
		{WhiteQueen, "Q"},
		{WhiteKing, "K"},
		{BlackPawn, "p"},
		{BlackKnight, "n"},
		{BlackBishop, "b"},
		{BlackRook, "r"},
		{BlackQueen, "q"},
		{BlackKing, "k"},
		{Piece(200), "."},
	}
	for _, c := range cases {
		if got := c.p.String(); got != c.s {
			t.Errorf("%s.String() = %s, want %s", c.p, got, c.s)
		}
	}
}

func TestPiece_DebugString(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p Piece
		s string
	}{
		{NoPiece, "NoPiece"},
		{WhitePawn, "WhitePawn"},
		{WhiteKnight, "WhiteKnight"},
		{WhiteBishop, "WhiteBishop"},
		{WhiteRook, "WhiteRook"},
		{WhiteQueen, "WhiteQueen"},
		{WhiteKing, "WhiteKing"},
		{BlackPawn, "BlackPawn"},
		{BlackKnight, "BlackKnight"},
		{BlackBishop, "BlackBishop"},
		{BlackRook, "BlackRook"},
		{BlackQueen, "BlackQueen"},
		{BlackKing, "BlackKing"},
		{Piece(200), "Piece(200)"},
	}
	for _, c := range cases {
		if got := c.p.DebugString(); got != c.s {
			t.Errorf("%s.DebugString() = %s, want %s", c.p, got, c.s)
		}
	}
}

func TestPieceFromString(t *testing.T) {
	t.Parallel()
	cases := []struct {
		s string
		p Piece
	}{
		{"P", WhitePawn},
		{"N", WhiteKnight},
		{"B", WhiteBishop},
		{"R", WhiteRook},
		{"Q", WhiteQueen},
		{"K", WhiteKing},
		{"p", BlackPawn},
		{"n", BlackKnight},
		{"b", BlackBishop},
		{"r", BlackRook},
		{"q", BlackQueen},
		{"k", BlackKing},
		{"", NoPiece},
		{"x", NoPiece},
		{"xxx", NoPiece},
	}
	for _, c := range cases {
		if got := PieceFromString(c.s); got != c.p {
			t.Errorf("PieceFromString(%q) = %q, want %q", c.s, got, c.p)
		}
	}
}

func TestPieceFromColorRole(t *testing.T) {
	t.Parallel()
	f := func(c Color, r Role) bool {
		if !c.IsValid() || !r.IsValid() {
			return true // skip this run
		}
		p := PieceFromColorRole(c, r)
		return p.IsValid() && p.Color() == c && p.Role() == r
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
