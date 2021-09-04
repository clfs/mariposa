package movegen

import "github.com/clfs/mariposa/internal/chess"

type MoveGen interface {
	LegalMoves(p chess.Position) []chess.Move
	MaybeLegalMoves(p chess.Position) []chess.Move
	IsLegalMove(p chess.Position, m chess.Move) bool
}
