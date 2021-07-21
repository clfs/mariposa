package core_test

import (
	"testing"

	. "github.com/clfs/mariposa/engine/core"
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
	}
	for _, c := range cases {
		if got := c.r.IsValid(); got != c.b {
			t.Errorf("IsValid(%s) = %v; want %v", c.r, got, c.b)
		}
	}
}
