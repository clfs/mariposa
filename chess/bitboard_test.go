package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/chess"
)

func TestNewBitboard(t *testing.T) {
	t.Parallel()
	if got := NewBitboard().Value(); got != 0 {
		t.Errorf("NewBitboard() = %v, want 0", got)
	}
}

func TestBitboard_Set(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard, s Square) bool {
		b.Set(s)
		return b.At(s)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Clear(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard, s Square) bool {
		b.Clear(s)
		return !b.At(s)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Value(t *testing.T) {
	t.Parallel()
	f := func(n uint64) bool {
		b := Bitboard(n)
		return b.Value() == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Invert(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard) bool {
		old := b
		b.Invert()
		b.Invert()
		return old == b
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Zero(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard) bool {
		b.Zero()
		return b.Value() == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
