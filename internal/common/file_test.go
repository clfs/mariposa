package common_test

import (
	"testing"

	. "github.com/clfs/mariposa/internal/common"
)

func TestFile_Mirror(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in, want File
	}{
		{FileA, FileH},
		{FileB, FileG},
		{FileC, FileF},
		{FileD, FileE},
		{FileE, FileD},
		{FileF, FileC},
		{FileG, FileB},
		{FileH, FileA},
	}
	for _, c := range cases {
		old := c.in
		c.in.Mirror()
		if got := c.in; got != c.want {
			t.Errorf("%v.Mirror() = %v, want %v", old, got, c.want)
		}
	}
}
