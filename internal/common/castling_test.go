package common_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/common"
)

func TestCastlingRights_Set(t *testing.T) {
	t.Parallel()
	f := func(c CastlingRights, r CastlingRight) bool {
		return c.Set(r).Get(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastlingRights_Clear(t *testing.T) {
	t.Parallel()
	f := func(c CastlingRights, r CastlingRight) bool {
		return !c.Clear(r).Get(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastlingRights_FEN(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in  *CastlingRights
		out string
	}{
		{NewCastlingRights(), "-"},
		{NewCastlingRights(WhiteOO), "K"},
		{NewCastlingRights(WhiteOOO), "Q"},
		{NewCastlingRights(BlackOO), "k"},
		{NewCastlingRights(BlackOOO), "q"},
		{NewCastlingRights(WhiteOO, WhiteOOO), "KQ"},
		{NewCastlingRights(WhiteOO, WhiteOOO, BlackOO), "KQk"},
		{NewCastlingRights(WhiteOO, WhiteOOO, BlackOO, BlackOOO), "KQkq"},
	}
	for _, c := range cases {
		if got := c.in.FEN(); got != c.out {
			t.Errorf("%v.FEN() = %q, want %q", c.in, got, c.out)
		}
	}
}
