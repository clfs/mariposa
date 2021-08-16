package common

import (
	"fmt"
	"strings"
)

type (
	CastlingRights uint8
	CastlingRight  uint8
)

const (
	WhiteOO CastlingRight = 1 << iota
	WhiteOOO
	BlackOO
	BlackOOO
)

func (c *CastlingRights) Add(r CastlingRight) *CastlingRights {
	*c |= CastlingRights(r)
	return c
}

func (c *CastlingRights) Remove(r CastlingRight) *CastlingRights {
	*c &^= CastlingRights(r)
	return c
}

func (c *CastlingRights) IsAllowed(r CastlingRight) bool {
	return (*c & CastlingRights(r)) != 0
}

func (c *CastlingRights) FEN() string {
	var b strings.Builder
	if c.IsAllowed(WhiteOO) {
		b.WriteString("K") // todo: use stringer for these
	}
	if c.IsAllowed(WhiteOOO) {
		b.WriteString("Q")
	}
	if c.IsAllowed(BlackOO) {
		b.WriteString("k")
	}
	if c.IsAllowed(BlackOOO) {
		b.WriteString("q")
	}
	s := b.String()
	if s == "" {
		return "-"
	}
	return s
}

// todo: move this to fen package
func ParseCastlingRights(s string) (CastlingRights, error) {
	var rights CastlingRights
	for _, r := range s {
		switch r {
		case 'K':
			rights.Add(WhiteOO)
		case 'Q':
			rights.Add(WhiteOOO)
		case 'k':
			rights.Add(BlackOO)
		case 'q':
			rights.Add(BlackOOO)
		case '-':
			return 0, nil // early exit!
		default:
			return 0, &ParseCastlingRightsError{Rights: s}
		}
	}
	return rights, nil
}

type ParseCastlingRightsError struct {
	Rights string
}

func (e *ParseCastlingRightsError) Error() string {
	return fmt.Sprintf("invalid castling rights: %s", e.Rights)
}
