package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestColor_Mirror(t *testing.T) {
	c := White
	c.Mirror()
	if c != Black {
		t.Errorf("expected %v, got %v", Black, c)
	}
	c.Mirror()
	if c != White {
		t.Errorf("expected %v, got %v", White, c)
	}
}

func TestColor_FEN(t *testing.T) {
	t.Parallel()
	f := func(c Color) bool {
		c2, err := ParseColor(c.FEN())
		return err == nil && c == c2
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
