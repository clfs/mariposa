package common_test

import (
	"errors"
	"math"
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/common"
	"github.com/clfs/mariposa/internal/parsers/fen"
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
		f, err := c.FEN()
		if err != nil {
			return false
		}
		c2, err := fen.ParseColor(f)
		return err == nil && c == c2
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestColor_FEN_Invalid(t *testing.T) {
	t.Parallel()
	for i := 2; i <= math.MaxUint8; i++ {
		c := Color(i)
		_, err := c.FEN()
		if !errors.As(err, &InvalidColorError{}) {
			t.Errorf("expected InvalidColorError, got %#v", err)
		}
	}
}
