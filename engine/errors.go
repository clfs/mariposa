package engine

import "fmt"

type ParsePieceError struct {
	InvalidPiece rune
}

func (e *ParsePieceError) Error() string {
	return fmt.Sprintf("invalid piece: %c", e.InvalidPiece)
}

type ParseSquareError struct {
	InvalidSquare string
}

func (e *ParseSquareError) Error() string {
	return fmt.Sprintf("invalid square: %s", e.InvalidSquare)
}

type ParseFileError struct {
	InvalidFile rune
}

func (e *ParseFileError) Error() string {
	return fmt.Sprintf("invalid file: %c", e.InvalidFile)
}

type ParseRankError struct {
	InvalidRank rune
}

func (e *ParseRankError) Error() string {
	return fmt.Sprintf("invalid rank: %c", e.InvalidRank)
}

type ParseFENError struct {
	FEN string
}

func (e *ParseFENError) Error() string {
	return fmt.Sprintf("invalid FEN: %s", e.FEN)
}
