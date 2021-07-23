package core_test

import (
	"testing"

	. "github.com/clfs/mariposa/core"
)

func TestRank_IsValid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r Rank
		b bool
	}{
		{Rank1, true},
		{Rank2, true},
		{Rank3, true},
		{Rank4, true},
		{Rank5, true},
		{Rank6, true},
		{Rank7, true},
		{Rank8, true},
		{8, false},
	}
	for _, c := range cases {
		if got := c.r.IsValid(); got != c.b {
			t.Errorf("%s.IsValid() = %v, want %v", c.r, got, c.b)
		}
	}
}

func TestRank_String(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r Rank
		s string
	}{
		{Rank1, "1"},
		{Rank2, "2"},
		{Rank3, "3"},
		{Rank4, "4"},
		{Rank5, "5"},
		{Rank6, "6"},
		{Rank7, "7"},
		{Rank8, "8"},
		{8, "Rank(8)"},
	}
	for _, c := range cases {
		if got := c.r.String(); got != c.s {
			t.Errorf("Rank(%d).String() = %s, want %s", c.r, got, c.s)
		}
	}
}
