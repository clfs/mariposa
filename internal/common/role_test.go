package common_test

import (
	"testing"

	. "github.com/clfs/mariposa/internal/common"
)

func TestRole_Value(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r    Role
		want uint8
	}{
		{Pawn, 1},
		{King, 6},
	}
	for _, c := range cases {
		if got := c.r.Value(); got != c.want {
			t.Errorf("%v.Value() = %d; want %d", c.r, got, c.want)
		}
	}
}

func TestRole_Valid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r    Role
		want bool
	}{
		{Pawn, true},
		{King, true},
		{Role(30), false},
	}
	for _, c := range cases {
		if got := c.r.Valid(); got != c.want {
			t.Errorf("%v.Valid() = %v; want %v", c.r, got, c.want)
		}
	}
}

func TestRole_Invalid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		r    Role
		want bool
	}{
		{Pawn, false},
		{King, false},
		{Role(30), true},
	}
	for _, c := range cases {
		if got := c.r.Invalid(); got != c.want {
			t.Errorf("%v.Invalid() = %v; want %v", c.r, got, c.want)
		}
	}
}
