package engine_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/engine"
)

func TestPieceFromColorRole(t *testing.T) {
	f := func(c Color, r Role) bool {
		p := PieceFromColorRole(c, r)
		return p.Color() == c && p.Role() == r
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSquareFromFileRank(t *testing.T) {
	f := func(f File, r Rank) bool {
		s := SquareFromFileRank(f, r)
		return s.File() == f && s.Rank() == r
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
