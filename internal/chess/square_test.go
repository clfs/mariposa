package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestSquare_Flipped(t *testing.T) {
	t.Parallel()
	f := func(s Square) bool {
		fs := s.Flipped()
		// File stays the same, rank is flipped.
		return fs.File() == s.File() && fs.Rank() == s.Rank().Flipped()
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestSquare_Flipped_Involutary(t *testing.T) {
	t.Parallel()
	f := func(s Square) bool {
		return s.Flipped().Flipped() == s
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestNewSquare(t *testing.T) {
	t.Parallel()
	f := func(s Square) bool {
		return NewSquare(s.File(), s.Rank()) == s
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
