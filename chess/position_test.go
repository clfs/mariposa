package chess_test

import (
	"bufio"
	"os"
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

func getFENs(t *testing.T, filename string) []string {
	t.Helper()
	f, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var fens []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fens = append(fens, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
	return fens
}

func TestPosition_SetFEN(t *testing.T) {
	t.Parallel()
	fens := getFENs(t, "testdata/fens.txt")
	for _, want := range fens {
		p := new(Position)
		if err := p.SetFEN(want); err != nil {
			t.Error(err)
		}
		got, err := p.FEN()
		if err != nil {
			t.Error(err)
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("SetFEN() failed: (-got +want)\n%s", diff)
		}
	}
}

func TestPosition_FEN(t *testing.T) {
	t.Skip() // failing; todo implement Position.Generate
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
