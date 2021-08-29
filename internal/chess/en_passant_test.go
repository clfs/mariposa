package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestEnPassantRight_Allowed(t *testing.T) {
	f := func(s Square) bool {
		return s.EnPassantRight().Allowed()
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestEnPassantRight_Flipped(t *testing.T) {
	f := func(e EnPassantRight) bool {
		return e.Flipped().Flipped() == e
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
