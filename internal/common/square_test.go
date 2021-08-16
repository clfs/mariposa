package common_test

import (
	"testing"

	. "github.com/clfs/mariposa/internal/common"
)

func TestSquare_Value(t *testing.T) {
	cases := []struct {
		s    Square
		want uint8
	}{
		{A1, 0},
		{B1, 1},
		{A2, 8},
		{H8, 63},
		{Square(100), 100},
	}
	for _, c := range cases {
		if got := c.s.Value(); got != c.want {
			t.Errorf("%v.Value() = %v; want %v", c.s, got, c.want)
		}
	}
}

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

func TestSquareAt(t *testing.T) {
	t.Parallel()
	cases := []struct {
		f    File
		r    Rank
		want Square
	}{
		{FileA, Rank1, A1},
		{FileF, Rank5, F5},
		{FileH, Rank8, H8},
	}
	for _, c := range cases {
		if got := SquareAt(c.f, c.r); got != c.want {
			t.Errorf("SquareAt(%s, %s) = %v; want %v", c.f, c.r, got, c.want)
		}
	}
}
