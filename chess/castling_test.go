package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/chess"
)

func TestCastling_Set(t *testing.T) {
	t.Parallel()
	f := func(c Castling, r CastlingRight) bool {
		c.Set(r)
		return c.HasRight(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastling_Clear(t *testing.T) {
	t.Parallel()
	f := func(c Castling, r CastlingRight) bool {
		c.Clear(r)
		return !c.HasRight(r)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastling_NoRights(t *testing.T) {
	t.Parallel()
	f := func(c Castling) bool {
		c.Clear(WhiteOO)
		c.Clear(WhiteOOO)
		c.Clear(BlackOO)
		c.Clear(BlackOOO)
		return c.NoRights()
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastling_String(t *testing.T) {
	cases := []struct {
		c    Castling
		want string
	}{
		{Castling(0), "-"},
		{Castling(WhiteOO), "K"},
		{Castling(WhiteOOO), "Q"},
		{Castling(BlackOO), "k"},
		{Castling(BlackOOO), "q"},
		{Castling(WhiteOO | WhiteOOO), "KQ"},
		{Castling(WhiteOO | BlackOO), "Kk"},
		{Castling(WhiteOO | BlackOOO), "Kq"},
		{Castling(WhiteOOO | BlackOO), "Qk"},
		{Castling(WhiteOOO | BlackOOO), "Qq"},
		{Castling(BlackOO | BlackOOO), "kq"},
		{Castling(WhiteOO | WhiteOOO | BlackOO), "KQk"},
		{Castling(WhiteOO | WhiteOOO | BlackOOO), "KQq"},
		{Castling(WhiteOO | BlackOO | BlackOOO), "Kkq"},
		{Castling(WhiteOOO | BlackOO | BlackOOO), "Qkq"},
		{Castling(WhiteOO | WhiteOOO | BlackOO | BlackOOO), "KQkq"},
	}
	for _, c := range cases {
		if got := c.c.String(); got != c.want {
			t.Errorf("0x%x: got %q; want %q", c.c, got, c.want)
		}
	}
}
