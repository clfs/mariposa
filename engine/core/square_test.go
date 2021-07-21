package core_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/engine/core"
)

func TestSquare_IsValid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		s Square
		b bool
	}{
		{A1, true},
		{F5, true},
		{H8, true},
		{H8 + 1, false},
		{0xff, false},
	}
	for _, c := range cases {
		if got := c.s.IsValid(); got != c.b {
			t.Errorf("Square.IsValid(%s) = %t, want %t", c.s, got, c.b)
		}
	}
}

func TestSquareFromFileRank(t *testing.T) {
	t.Parallel()
	f := func(f File, r Rank) bool {
		if !f.IsValid() || !r.IsValid() {
			return true // skip this run
		}
		s := SquareFromFileRank(f, r)
		return s.IsValid() && s.File() == f && s.Rank() == r
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
