package fen

import (
	"fmt"

	"github.com/clfs/mariposa/internal/chess"
)

var (
	ErrInvalidColor = fmt.Errorf("invalid color")
	ErrInvalidPiece = fmt.Errorf("invalid piece")
)

func ToColor(s string) (chess.Color, error) {
	switch s {
	case "w":
		return chess.White, nil
	case "b":
		return chess.Black, nil
	}
	return 0, ErrInvalidColor
}

func ToPiece(s string) (chess.Piece, error) {
	switch s {
	case "P":
		return chess.WhitePawn, nil
	case "p":
		return chess.BlackPawn, nil
	case "N":
		return chess.WhiteKnight, nil
	case "n":
		return chess.BlackKnight, nil
	case "B":
		return chess.WhiteBishop, nil
	case "b":
		return chess.BlackBishop, nil
	case "R":
		return chess.WhiteRook, nil
	case "r":
		return chess.BlackRook, nil
	case "Q":
		return chess.WhiteQueen, nil
	case "q":
		return chess.BlackQueen, nil
	case "K":
		return chess.WhiteKing, nil
	case "k":
		return chess.BlackKing, nil
	default:
		return 0, ErrInvalidPiece
	}
}
