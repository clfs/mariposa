package movegen

import (
	"github.com/clfs/mariposa/internal/common"
	"github.com/clfs/mariposa/internal/position"
)

// MaybeLegalMoves returns potentially legal moves for a position. It does not
// check for legality. For example, it might return a bishop move that causes
// a king to remain in check.
func MaybeLegalMoves(p *position.Position) []common.Move {
	return nil
}

// LegalMoves returns legal moves for a position.
func LegalMoves(p *position.Position) []common.Move {
	var res []common.Move
	for _, m := range MaybeLegalMoves(p) {
		if IsLegalMove(p, m) {
			res = append(res, m)
		}
	}
	return res
}

// IsLegalMove returns true if the move is legal in the position.
func IsLegalMove(p *position.Position, m common.Move) bool {
	return false
}
