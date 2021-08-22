package common_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/common"
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

func TestEnPassantRight_NewEnPassantRightNotAllowed(t *testing.T) {
	t.Parallel()
	e := NewEnPassantRightNotAllowed()
	if _, ok := e.Get(); ok {
		t.Errorf("en passant was allowed")
	}
}
