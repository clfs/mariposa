package chess

import "fmt"

type Move struct {
	From, To  Square
	Promotion Role
}

func ParseMove(s string) (Move, error) {
	return Move{}, fmt.Errorf("todo implement")
}
