package chess_test

import (
	"testing"
	"testing/quick"

	"github.com/google/go-cmp/cmp"

	. "github.com/clfs/mariposa/chess"
)

func TestStartingPosition(t *testing.T) {
	t.Parallel()
	p, err := StartingPosition()
	if err != nil {
		t.Errorf("StartingPosition() returned error: %v", err)
	}
	fen, err := p.FEN()
	if err != nil {
		t.Errorf("StartingPosition() to FEN returned error: %v", err)
	}
	if fen != StartingFEN {
		t.Errorf("StartingPosition() had fen %s, expected %s", fen, StartingFEN)
	}
}

func TestPosition_FENIsSymmetric(t *testing.T) {
	t.Parallel()
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
