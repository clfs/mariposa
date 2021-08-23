package chess_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	. "github.com/clfs/mariposa/internal/chess"
)

func TestBoard_Flip(t *testing.T) {
	b, err := NewBoard(StartingBoardFEN)
	if err != nil {
		t.Fatal(err)
	}
	b.Flip()
	if diff := cmp.Diff(StartingBoardFEN, b.FEN()); diff != "" {
		t.Errorf("Flip() failed (-want +got): %s", diff)
	}
	b.Flip()
	if diff := cmp.Diff(StartingBoardFEN, b.FEN()); diff != "" {
		t.Errorf("Flip() failed (-want +got): %s", diff)
	}
}
