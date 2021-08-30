package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestBoard_Flip(t *testing.T) {
	t.Parallel()
	// Flipping the board twice must not change the board.
	f := func(b Board) bool {
		old := b
		return old == *b.Flip().Flip()
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
