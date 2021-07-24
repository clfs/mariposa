package core_test

import (
	"testing"
	"testing/quick"

	"github.com/google/go-cmp/cmp"

	. "github.com/clfs/mariposa/core"
)

func TestBoard_FEN(t *testing.T) {
	t.Parallel()
	// Converting the board to FEN and back shouldn't change anything important.
	f := func(b Board) bool {
		fen := b.FEN()
		bNew, err := NewBoard(fen)
		return err == nil && cmp.Equal(b, bNew)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
