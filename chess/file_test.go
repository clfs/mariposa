package chess_test

import (
	"testing"

	. "github.com/clfs/mariposa/chess"
)

func TestFile_Value(t *testing.T) {
	t.Parallel()
	cases := []struct {
		f    File
		want uint8
	}{
		{FileA, 0},
		{FileH, 7},
		{File(8), 8},
	}
	for _, c := range cases {
		if got := c.f.Value(); got != c.want {
			t.Errorf("%v.Value() = %v; want %v", c.f, got, c.want)
		}
	}
}

func TestFile_Valid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		f    File
		want bool
	}{
		{FileA, true},
		{FileH, true},
		{File(8), false},
	}
	for _, c := range cases {
		if got := c.f.Valid(); got != c.want {
			t.Errorf("%v.Valid() = %v; want %v", c.f, got, c.want)
		}
	}
}

func TestFile_Invalid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		f    File
		want bool
	}{
		{FileA, false},
		{FileH, false},
		{File(8), true},
	}
	for _, c := range cases {
		if got := c.f.Invalid(); got != c.want {
			t.Errorf("%v.Invalid() = %v; want %v", c.f, got, c.want)
		}
	}
}
