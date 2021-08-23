package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestBoard_PutThenGet(t *testing.T) {
	f := func(p Piece, s Square) bool {
		got, ok := new(Board).Put(p, s).Get(s)
		return ok && got == p
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
