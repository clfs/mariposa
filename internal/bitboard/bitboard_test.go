package bitboard_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/bitboard"
	"github.com/clfs/mariposa/internal/common"
)

func TestB_Value(t *testing.T) {
	f := func(n uint64) bool {
		b := B(n)
		return b.Value() == n
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestB_Set(t *testing.T) {
	f := func(b B, s common.Square) bool {
		return b.Set(s).Get(s)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestB_Clear(t *testing.T) {
	f := func(b B, s common.Square) bool {
		return !b.Clear(s).Get(s)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestB_Toggle(t *testing.T) {
	f := func(b B, s common.Square) bool {
		old := b.Get(s)
		return old != b.Toggle(s).Get(s)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
