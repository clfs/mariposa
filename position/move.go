package position

import (
	"fmt"

	"github.com/clfs/mariposa/chess"
)

type Move struct {
	From      chess.Square
	To        chess.Square
	Promotion chess.Role
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
