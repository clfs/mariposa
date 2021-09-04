package fen

import (
	"github.com/clfs/mariposa/internal/chess"
)

func FromColor(c chess.Color) string {
	if c == chess.White {
		return "w"
	}
	return "b"
}
