package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestBitboard_Set(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard, s Square, x bool) bool {
		return b.Set(s, x).Get(s) == x
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Toggle(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard, s Square) bool {
		old := b.Get(s)
		return old != b.Toggle(s).Get(s)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Flip(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard, s Square) bool {
		x := b.Get(s)
		y := b.Flip().Get(s.Flipped())
		z := b.Flip().Get(s)
		return x == y && y == z
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Debug(t *testing.T) {
	t.Parallel()
	b := Bitboard(0x1e22_2212_0e0a_1222)
	want := `.1111...
.1...1..
.1...1..
.1..1...
.111....
.1.1....
.1..1...
.1...1..`
	if got := b.Debug(); got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}
