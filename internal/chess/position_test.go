package chess_test

import (
	"bufio"
	"os"
	"testing"

	. "github.com/clfs/mariposa/internal/chess"
)

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

func TestNewPosition_NoPanic(t *testing.T) {
	t.Parallel()
	fens := getFENs(t, "testdata/fens.txt")
	for _, fen := range fens {
		_, err := NewPosition(fen)
		if err != nil {
			t.Errorf("Parse(%q) failed: %v", fen, err)
		}
	}
}
