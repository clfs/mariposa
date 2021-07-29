package core_test

import (
	"testing"

	. "github.com/clfs/mariposa/core"
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
			t.Errorf("%s.IsValid() = %v, want %v", c.r, got, c.b)
		}
	}
}

func TestRole_DebugString(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r Role
		s string
	}{
		{0, "Role(0)"},
		{Pawn, "Pawn"},
		{Knight, "Knight"},
		{Bishop, "Bishop"},
		{Rook, "Rook"},
		{Queen, "Queen"},
		{King, "King"},
		{200, "Role(200)"},
	}
	for _, c := range cases {
		if got := c.r.DebugString(); got != c.s {
			t.Errorf("Role(%d).String() = %s, want %s", c.r, got, c.s)
		}
	}
}
