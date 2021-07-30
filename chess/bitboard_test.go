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

func TestBitboard_Clear(t *testing.T) {
	t.Parallel()
	f := func(b Bitboard) bool {
		b.Clear()
		return b.Value() == 0
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
