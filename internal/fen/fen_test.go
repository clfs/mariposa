package fen_test

import (
	"bufio"
	"os"
	"testing"
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
