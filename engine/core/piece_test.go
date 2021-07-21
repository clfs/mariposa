package core_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/engine/core"
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
			t.Errorf("IsValid(%v) = %v, want %v", c.p, got, c.b)
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
