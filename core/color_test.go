package core_test

import (
	"testing"

	. "github.com/clfs/mariposa/core"
)

func TestColor_IsValid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		c Color
		b bool
	}{
		{White, true},
		{Black, true},
		{2, false},
	}
	for _, c := range cases {
		if got := c.c.IsValid(); c.b != c.c.IsValid() {
			t.Errorf("%s.IsValid() = %v, want %v", c.c, got, c.b)
		}
	}
}

func TestColor_String(t *testing.T) {
	t.Parallel()
	cases := []struct {
		c Color
		s string
	}{
		{White, "w"},
		{Black, "b"},
		{2, "-"},
	}
	for _, c := range cases {
		if got := c.c.String(); c.s != got {
			t.Errorf("Color(%d).String() = %s, want %s", c.c, got, c.s)
		}
	}
}

func TestColor_DebugString(t *testing.T) {
	t.Parallel()
	cases := []struct {
		c Color
		s string
	}{
		{White, "White"},
		{Black, "Black"},
		{2, "Color(2)"},
	}
	for _, c := range cases {
		if got := c.c.DebugString(); c.s != got {
			t.Errorf("Color(%d).DebugString() = %s, want %s", c.c, got, c.s)
		}
	}
}
