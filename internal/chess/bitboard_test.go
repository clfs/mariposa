package chess_test

import (
	"math"
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestBitboard_Set(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard, s Square) bool {
		return b.Set(s).Get(s)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Clear(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard, s Square) bool {
		return !b.Clear(s).Get(s)
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
		y := b.Flip().Get(*s.Flip())
		return x == y
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Flip_Involutary(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard) bool {
		old := b
		return old == *b.Flip().Flip()
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

func BenchmarkBitboard_Get(b *testing.B) {
	bitboard := Bitboard(math.MaxUint64)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 64; j++ {
			bitboard.Get(Square(i))
		}
	}
}

func BenchmarkBitboard_Set(b *testing.B) {
	bitboard := BitboardAllZero
	for i := 0; i < b.N; i++ {
		for j := 0; j < 64; j++ {
			bitboard.Set(Square(i))
		}
	}
}

func BenchmarkBitboard_Clear(b *testing.B) {
	bitboard := BitboardAllOne
	for i := 0; i < b.N; i++ {
		for j := 0; j < 64; j++ {
			bitboard.Clear(Square(i))
		}
	}
}

func BenchmarkBitboard_Toggle(b *testing.B) {
	bitboard := BitboardAllZero
	for i := 0; i < b.N; i++ {
		for j := 0; j < 64; j++ {
			bitboard.Toggle(Square(i))
		}
	}
}

func BenchmarkBitboard_Flip(b *testing.B) {
	bitboard := BitboardAllOne
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			bitboard.Flip()
		}
	}
}
