package fen_test

import (
	"bufio"
	"os"
	"testing"
	"testing/quick"

	"github.com/clfs/mariposa/internal/common"
	. "github.com/clfs/mariposa/internal/parsers/fen"
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

func TestParseColor_Invalid(t *testing.T) {
	t.Parallel()
	cases := []string{
		"",
		" ",
		"\n",
		"-",
		"a",
		"A",
		"W",
		"B",
		"wb",
		"bw",
		"White",
		"Black",
	}
	for _, c := range cases {
		if _, err := ParseColor(c); err == nil {
			t.Errorf("ParseColor(%q) did not fail", c)
		}
	}
}

func TestParseColor(t *testing.T) {
	t.Parallel()
	f := func(c common.Color) bool {
		got, err := ParseColor(c.FEN())
		return err == nil && got == c
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestParseCastlingRights(t *testing.T) {
	t.Parallel()
	f := func(c common.CastlingRights) bool {
		got, err := ParseCastlingRights(c.FEN())
		return err == nil && got == c
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestParseEnPassantRight(t *testing.T) {
	t.Parallel()
	f := func(e common.EnPassantRight) bool {
		got, err := ParseEnPassantRight(e.FEN())
		return err == nil && got.Equal(e)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestParse(t *testing.T) {
	t.Parallel()
	fens := getFENs(t, "testdata/fens.txt")
	for _, fen := range fens {
		_, err := Parse(fen)
		if err != nil {
			t.Errorf("Parse(%q) failed: %v", fen, err)
		}
	}
}
