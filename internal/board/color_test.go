package board_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/board"
)

func TestColor_Value(t *testing.T) {
	t.Parallel()
	cases := []struct {
		c    Color
		want uint8
	}{
		{White, 0},
		{Black, 1},
	}
	for _, c := range cases {
		if got := c.c.Value(); got != c.want {
			t.Errorf("%s.Value() = %v, want %v", c.c, got, c.want)
		}
	}
}

func TestColor_Valid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		c    Color
		want bool
	}{
		{White, true},
		{Black, true},
		{Color(2), false},
	}
	for _, c := range cases {
		if got := c.c.Valid(); got != c.want {
			t.Errorf("%s.Valid() = %v, want %v", c.c, got, c.want)
		}
	}
}

func TestColor_Invalid(t *testing.T) {
	t.Parallel()
	f := func(c Color) bool { return c.Valid() != c.Invalid() }
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
