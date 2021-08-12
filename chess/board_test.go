package chess_test

import (
	"testing"
	"testing/quick"

	. "github.com/clfs/mariposa/chess"
)

func TestBoard_PutAndGet(t *testing.T) {
	t.Parallel()
	f := func(p Piece, s Square) bool {
		b := new(Board) // todo improve board generation
		b.Put(p, s)
		got, ok := b.Get(s)
		return ok && got == p
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
