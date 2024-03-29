package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestNewCastlingRights(t *testing.T) {
	t.Parallel()
	f := func(flags []CastlingFlag) bool {
		c := NewCastlingRights(flags...)
		for _, f := range flags {
			if !c.Get(f) {
				return false
			}
		}
		return true
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastlingRights_Enable(t *testing.T) {
	t.Parallel()
	f := func(c CastlingRights, f CastlingFlag) bool {
		c.Enable(f)
		return c.Get(f)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastlingRights_Disable(t *testing.T) {
	t.Parallel()
	f := func(c CastlingRights, f CastlingFlag) bool {
		c.Disable(f)
		return !c.Get(f)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCastlingRights_Flip(t *testing.T) {
	t.Parallel()
	f := func(c CastlingRights) bool {
		old := c
		c.Flip()
		c.Flip()
		return old == c
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
