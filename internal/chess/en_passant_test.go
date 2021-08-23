package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/chess"
	"github.com/google/go-cmp/cmp"
)

func TestEnPassantRight_Set(t *testing.T) {
	t.Parallel()
	f := func(s Square) bool {
		e := NewEnPassantRight(s)
		e.Set(s)
		target, ok := e.Get()
		return ok && target == s
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestEnPassantRight_Clear(t *testing.T) {
	t.Parallel()
	f := func(e EnPassantRight) bool {
		e.Clear()
		_, ok := e.Get()
		return !ok
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestEnPassantRight_Flip(t *testing.T) {
	t.Parallel()
	f := func(e EnPassantRight) bool {
		old := e
		e.Flip()
		e.Flip()
		return cmp.Equal(e, old)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestEnPassantRight_NewEnPassantRightNone(t *testing.T) {
	t.Parallel()
	e := NewEnPassantRightNone()
	if _, ok := e.Get(); ok {
		t.Errorf("en passant was allowed")
	}
}
