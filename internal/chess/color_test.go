package chess_test

import (
	"testing"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestColor_Flipped(t *testing.T) {
	if got := White.Flipped(); got != Black {
		t.Errorf("White.Flipped() should be Black, got %v", got)
	}
	if got := Black.Flipped(); got != White {
		t.Errorf("Black.Flipped() should be White, got %v", got)
	}
}
