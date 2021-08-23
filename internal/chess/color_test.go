package chess_test

import (
	"testing"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestColor_Flip(t *testing.T) {
	c := White
	c.Flip()
	if c != Black {
		t.Errorf("expected %v, got %v", Black, c)
	}
	c.Flip()
	if c != White {
		t.Errorf("expected %v, got %v", White, c)
	}
}
