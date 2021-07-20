package core_test

import (
	"testing"

	. "github.com/clfs/mariposa/engine/core"
)

func TestColor_String(t *testing.T) {
	t.Parallel()
	cases := []struct {
		c Color
		s string
	}{
		{White, "White"},
		{Black, "Black"},
		{NoColor, "Color(255)"},
	}
	for _, c := range cases {
		if got := c.c.String(); c.s != got {
			t.Errorf("want %v, got %v", c.s, got)
		}
	}
}

func TestColor_IsValid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		c Color
		b bool
	}{
		{White, true},
		{Black, true},
		{NoColor, false},
	}
	for _, c := range cases {
		if got := c.c.IsValid(); c.b != got {
			t.Errorf("Color.IsValid(%s): want %v, got %v", c.c, c.b, got)
		}
	}
}

func TestRole_String(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r Role
		s string
	}{
		{Pawn, "Pawn"},
		{Knight, "Knight"},
		{Bishop, "Bishop"},
		{Rook, "Rook"},
		{Queen, "Queen"},
		{King, "King"},
		{NoRole, "Role(255)"},
	}
	for _, c := range cases {
		if got := c.r.String(); c.s != got {
			t.Errorf("want %v, got %v", c.s, got)
		}
	}
}

func TestRole_IsValid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r Role
		b bool
	}{
		{Pawn, true},
		{Knight, true},
		{Bishop, true},
		{Rook, true},
		{Queen, true},
		{King, true},
		{NoRole, false},
	}
	for _, c := range cases {
		if got := c.r.IsValid(); c.b != got {
			t.Errorf("Role.IsValid(%s): want %v, got %v", c.r, c.b, got)
		}
	}
}

func TestPiece_String(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p Piece
		s string
	}{
		{WhitePawn, "P"},
		{BlackPawn, "p"},
		{WhiteKnight, "N"},
		{BlackKnight, "n"},
		{WhiteBishop, "B"},
		{BlackBishop, "b"},
		{WhiteRook, "R"},
		{BlackRook, "r"},
		{WhiteQueen, "Q"},
		{BlackQueen, "q"},
		{WhiteKing, "K"},
		{BlackKing, "k"},
		{NoPiece, "Piece(255)"},
	}
	for _, c := range cases {
		if got := c.p.String(); c.s != got {
			t.Errorf("want %v, got %v", c.s, got)
		}
	}
}

func TestPiece_IsValid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p Piece
		b bool
	}{
		{WhitePawn, true},
		{BlackPawn, true},
		{WhiteKnight, true},
		{BlackKnight, true},
		{WhiteBishop, true},
		{BlackBishop, true},
		{WhiteRook, true},
		{BlackRook, true},
		{WhiteQueen, true},
		{BlackQueen, true},
		{WhiteKing, true},
		{BlackKing, true},
		{NoPiece, false},
	}
	for _, c := range cases {
		if got := c.p.IsValid(); c.b != got {
			t.Errorf("Piece.IsValid(%s): want %v, got %v", c.p, c.b, got)
		}
	}
}

func TestPiece_Color(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p Piece
		c Color
	}{
		{WhitePawn, White},
		{BlackPawn, Black},
		{WhiteKnight, White},
		{BlackKnight, Black},
		{WhiteBishop, White},
		{BlackBishop, Black},
		{WhiteRook, White},
		{BlackRook, Black},
		{WhiteQueen, White},
		{BlackQueen, Black},
		{WhiteKing, White},
		{BlackKing, Black},
		{NoPiece, NoColor},
	}
	for _, c := range cases {
		if got := c.p.Color(); c.c != got {
			t.Errorf("Piece.Color(%s): want %v, got %v", c.p, c.c, got)
		}
	}
}
