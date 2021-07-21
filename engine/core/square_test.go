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
			t.Errorf("%s.IsValid() = %v, want %v", c.s, got, c.b)
		}
	}
}

func TestSquare_String(t *testing.T) {
	t.Parallel()
	cases := []struct {
		sq Square
		st string
	}{
		{A1, "a1"},
		{F5, "f5"},
		{H8, "h8"},
		{200, "Square(200)"},
	}
	for _, c := range cases {
		if got := c.sq.String(); got != c.st {
			t.Errorf("Square(%d).String() = %v, want %v", c.sq, got, c.st)
		}
	}
}

func TestSquareFromString(t *testing.T) {
	t.Parallel()
	cases := []struct {
		st string
		sq Square
	}{
		{"a1", A1},
		{"f5", F5},
		{"h8", H8},
		{"", NoSquare},
		{"a", NoSquare},
		{"a1a", NoSquare},
		{"a1a1", NoSquare},
		{"1", NoSquare},
		{"1A", NoSquare},
		{"A1", NoSquare},
		{"-", NoSquare},
	}
	for _, c := range cases {
		if got := SquareFromString(c.st); got != c.sq {
			t.Errorf("SquareFromString(%q) = %v, want %v", c.st, got, c.sq)
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
