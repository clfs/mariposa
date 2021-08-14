package common_test

import (
	"testing"

	. "github.com/clfs/mariposa/internal/common"
)

func TestRank_Value(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r    Rank
		want uint8
	}{
		{Rank1, 0},
		{Rank8, 7},
		{Rank(8), 8},
	}
	for _, c := range cases {
		if got := c.r.Value(); got != c.want {
			t.Errorf("%s: got %d, want %d", c.r, got, c.want)
		}
	}
}

func TestRank_Valid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r    Rank
		want bool
	}{
		{Rank1, true},
		{Rank8, true},
		{Rank(8), false},
	}
	for _, c := range cases {
		if got := c.r.Valid(); got != c.want {
			t.Errorf("%s: got %t, want %t", c.r, got, c.want)
		}
	}
}

func TestRank_Invalid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r    Rank
		want bool
	}{
		{Rank1, false},
		{Rank8, false},
		{Rank(8), true},
	}
	for _, c := range cases {
		if got := c.r.Invalid(); got != c.want {
			t.Errorf("%s: got %t, want %t", c.r, got, c.want)
		}
	}
}
