package core_test

import (
	"testing"

	. "github.com/clfs/mariposa/core"
)

func TestFile_IsValid(t *testing.T) {
	cases := []struct {
		f File
		b bool
	}{
		{FileA, true},
		{FileB, true},
		{FileC, true},
		{FileD, true},
		{FileE, true},
		{FileF, true},
		{FileG, true},
		{FileH, true},
		{9, false},
	}
	for _, c := range cases {
		if got := c.f.IsValid(); got != c.b {
			t.Errorf("%s.IsValid() = %v, want %v", c.f, got, c.b)
		}
	}
}

func TestFile_String(t *testing.T) {
	cases := []struct {
		f File
		s string
	}{
		{FileA, "a"},
		{FileB, "b"},
		{FileC, "c"},
		{FileD, "d"},
		{FileE, "e"},
		{FileF, "f"},
		{FileG, "g"},
		{FileH, "h"},
		{9, "File(9)"},
	}
	for _, c := range cases {
		if got := c.f.String(); got != c.s {
			t.Errorf("File(%d).String() = %s, want %s", c.f, got, c.s)
		}
	}
}
