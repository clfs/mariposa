package position_test

import (
	"testing"
	"testing/quick"

	"github.com/google/go-cmp/cmp"

	. "github.com/clfs/mariposa/position"
)

func TestPosition_FEN(t *testing.T) {
	t.Parallel()
	// Converting a position to FEN and back shouldn't change anything.
	f := func(p Position) bool {
		fen, err := p.FEN()
		if err != nil {
			return true // skip
		}
		pNew, err := NewPosition(fen)
		return err == nil && cmp.Equal(p, pNew)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
