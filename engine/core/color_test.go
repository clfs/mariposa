package core_test

import (
	"testing"

	. "github.com/clfs/mariposa/engine/core"
)

func TestColor_IsValid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		c Color
		b bool
	}{
		{White, true},
		{Black, true},
		{2, false},
	}
	for _, c := range cases {
		if got := c.c.IsValid(); c.b != c.c.IsValid() {
			t.Errorf("Color.IsValid(%s) = %t, want %t", c.c, got, c.b)
		}
	}
}
