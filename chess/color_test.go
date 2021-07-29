package chess_test

import (
	"testing"

	. "github.com/clfs/mariposa/chess"
)

func TestColor_Value(t *testing.T) {
	t.Parallel()
	cases := []struct {
		c    Color
		want uint8
	}{
		{White, 0},
		{Black, 1},
		{Color(2), 2},
	}
	for _, c := range cases {
		if got := c.c.Value(); got != c.want {
			t.Errorf("%v.Value() = %v; want %v", c.c, got, c.want)
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
			t.Errorf("%v.Valid() = %v; want %v", c.c, got, c.want)
		}
	}
}

func TestColor_Invalid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		c    Color
		want bool
	}{
		{White, false},
		{Black, false},
		{Color(2), true},
	}
	for _, c := range cases {
		if got := c.c.Invalid(); got != c.want {
			t.Errorf("%v.Invalid() = %v; want %v", c.c, got, c.want)
		}
	}
}
