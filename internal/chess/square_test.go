package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestSquare_Valid(t *testing.T) {
	cases := []struct {
		s    Square
		want bool
	}{
		{A1, true},
		{H8, true},
		{Square(100), false},
	}
	for _, c := range cases {
		if got := c.s.Valid(); got != c.want {
			t.Errorf("%v.Valid() = %t; want %t", c.s, got, c.want)
		}
	}
}

func TestSquare_Invalid(t *testing.T) {
	cases := []struct {
		s    Square
		want bool
	}{
		{A1, false},
		{H8, false},
		{Square(100), true},
	}
	for _, c := range cases {
		if got := c.s.Invalid(); got != c.want {
			t.Errorf("%v.Invalid() = %t; want %t", c.s, got, c.want)
		}
	}
}

func TestSquare_File(t *testing.T) {
	cases := []struct {
		s    Square
		want File
	}{
		{A1, FileA},
		{F5, FileF},
		{H8, FileH},
	}
	for _, c := range cases {
		if got := c.s.File(); got != c.want {
			t.Errorf("%v.File() = %v; want %v", c.s, got, c.want)
		}
	}
}

func TestSquare_Rank(t *testing.T) {
	cases := []struct {
		s    Square
		want Rank
	}{
		{A1, Rank1},
		{F5, Rank5},
		{H8, Rank8},
	}
	for _, c := range cases {
		if got := c.s.Rank(); got != c.want {
			t.Errorf("%v.Rank() = %v; want %v", c.s, got, c.want)
		}
	}
}

func TestSquare_Mirror(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in, want Square
	}{
		{A1, A8},
		{H8, H1},
		{E4, E5},
		{F2, F7},
		{C8, C1},
	}
	for _, c := range cases {
		old := c.in
		c.in.Flip()
		if got := c.in; got != c.want {
			t.Errorf("%v.Mirror() = %v; want %v", old.FEN(), got, c.want)
		}
	}
}

func TestSquare_Flip_Involutary(t *testing.T) {
	t.Parallel()
	f := func(s Square) bool {
		old := s
		s.Flip()
		s.Flip()
		return s == old
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSquareAt(t *testing.T) {
	t.Parallel()
	f := func(s Square) bool {
		file := s.File()
		rank := s.Rank()
		return SquareAt(file, rank) == s
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
