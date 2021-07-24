package position

import "strings"

type Castling struct {
	WhiteOO  bool
	WhiteOOO bool
	BlackOO  bool
	BlackOOO bool
}

func (c Castling) String() string {
	var sb strings.Builder
	if c.WhiteOO {
		sb.WriteString("K")
	}
	if c.WhiteOOO {
		sb.WriteString("Q")
	}
	if c.BlackOO {
		sb.WriteString("k")
	}
	if c.BlackOOO {
		sb.WriteString("q")
	}
	s := sb.String()
	if s == "" {
		s = "-"
	}
	return s
}
