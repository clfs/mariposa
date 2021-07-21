package board

import "fmt"

type InvalidFENError struct {
	FEN string
}

func (e *InvalidFENError) Error() string {
	return fmt.Sprintf("invalid FEN: %v", e.FEN)
}
