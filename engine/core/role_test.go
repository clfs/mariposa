package core_test

import (
	"testing"

	. "github.com/clfs/mariposa/engine/core"
)

func TestRole_IsValid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r Role
		b bool
	}{
		{0, false},
		{Pawn, true},
		{Knight, true},
		{Bishop, true},
		{Rook, true},
		{Queen, true},
		{King, true},
		{King + 1, false},
	}
	for _, c := range cases {
		if got := c.r.IsValid(); got != c.b {
			t.Errorf("Role.IsValid(%s) = %v, want %v", c.r, got, c.b)
		}
	}
}
