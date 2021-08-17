package common_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/common"
)

func TestEnPassantRight_Equal(t *testing.T) {
	t.Parallel()
	cases := []struct {
		a    EnPassantRight
		b    EnPassantRight
		want bool
	}{
		{
			EnPassantRight{Target: A1, Allowed: true},
			EnPassantRight{Target: A1, Allowed: true},
			true,
		},
		{
			EnPassantRight{Target: A1, Allowed: true},
			EnPassantRight{Target: A2, Allowed: true},
			false,
		},
		{
			EnPassantRight{Target: A1, Allowed: false},
			EnPassantRight{Target: H8, Allowed: false},
			true,
		},
		{
			EnPassantRight{Target: A1, Allowed: false},
			EnPassantRight{Target: A1, Allowed: false},
			true,
		},
	}
	for _, c := range cases {
		if got := c.a.Equal(c.b); got != c.want {
			t.Errorf("%v.Equal(%v) = %v; want %v", c.a, c.b, got, c.want)
		}
	}
}

func TestEnPassantRight_Set(t *testing.T) {
	f := func(e EnPassantRight, s Square) bool {
		e.Set(s)
		target, ok := e.Get()
		return ok && target == s
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestEnPassantRight_Clear(t *testing.T) {
	f := func(e EnPassantRight) bool {
		e.Clear()
		_, ok := e.Get()
		return !ok
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
