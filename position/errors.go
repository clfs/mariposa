package position

import "fmt"

type ParseFENError struct {
	FEN string
}

func (e *ParseFENError) Error() string {
	return fmt.Sprintf("invalid FEN: %v", e.FEN)
}

type ParseMoveError struct {
	Move string
}

func (e *ParseMoveError) Error() string {
	return fmt.Sprintf("invalid move: %v", e.Move)
}
