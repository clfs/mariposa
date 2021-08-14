package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/chess"
)

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
		return !(b.Clear(s).Get(s))
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Toggle(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard, s Square) bool {
		tmp := b.Get(s)
		return tmp != b.Toggle(s).Get(s)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestBitboard_Reset(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard, s Square) bool {
		return !(b.Reset().Get(s))
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
