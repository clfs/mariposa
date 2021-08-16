package fen_test

import (
	"bufio"
	"os"
	"testing"

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

func TestParseColor(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in   string
		want common.Color
	}{
		{"w", common.White},
		{"b", common.Black},
	}
	for _, c := range cases {
		got, err := ParseColor(c.in)
		if err != nil {
			t.Errorf("ParseColor(%q) failed: %v", c.in, err)
		}
		if got != c.want {
			t.Errorf("ParseColor(%q) = %v, want %v", c.in, got, c.want)
		}
	}
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
