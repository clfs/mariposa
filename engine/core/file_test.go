package core_test

import (
	"testing"

	. "github.com/clfs/mariposa/engine/core"
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
			t.Errorf("IsValid(%s) = %t, want %t", c.f, got, c.b)
		}
	}
}
