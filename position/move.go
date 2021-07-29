package position

import (
	"fmt"

	"github.com/clfs/mariposa/core"
)

type Move struct {
	From      core.Square
	To        core.Square
	Promotion core.Role
}

func MoveFromString(s string) (Move, error) {
	if l := len(s); l != 4 && l != 5 {
		return Move{}, fmt.Errorf("invalid move: %s", s)
	}
	return Move{}, nil
}

func (m Move) String() string {
	return fmt.Sprintf("%s%s%s", m.From, m.To, m.Promotion)
}
