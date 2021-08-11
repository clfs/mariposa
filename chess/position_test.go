package chess_test

import (
	"testing"
	"testing/quick"

	"github.com/google/go-cmp/cmp"

	. "github.com/clfs/mariposa/chess"
)

func TestPosition_SymmetricFEN(t *testing.T) {
	f := func(p Position) bool {
		fen, err := p.FEN()
		if err != nil {
			return false
		}
		q, err := NewPosition(fen)
		if err != nil {
			return false
		}
		return cmp.Equal(p, q)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
