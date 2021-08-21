package position

import (
	"fmt"
	"strings"

	"github.com/clfs/mariposa/internal/bitboard"
	"github.com/clfs/mariposa/internal/common"
)

type Board struct {
	whites  bitboard.B
	blacks  bitboard.B
	pawns   bitboard.B
	knights bitboard.B
	bishops bitboard.B
	rooks   bitboard.B
	queens  bitboard.B
	kings   bitboard.B
}

func (b *Board) FEN() string {
	var sb strings.Builder

	for r := common.Rank8; r <= common.Rank8; r-- {
		skip := 0
		for f := common.FileA; f <= common.FileH; f++ {
			piece, ok := b.Get(common.SquareAt(f, r))
			if !ok {
				skip++
				continue
			}
			if skip > 0 {
				fmt.Fprintf(&sb, "%d", skip)
				skip = 0
			}
			fmt.Fprintf(&sb, "%s", piece)
		}
		if skip > 0 {
			fmt.Fprintf(&sb, "%d", skip)
		}
		if r != common.Rank1 {
			fmt.Fprintf(&sb, "/")
		}
	}

	return sb.String()
}

// Put places p at s and returns b.
func (b *Board) Put(p common.Piece, s common.Square) *Board {
	switch p.Color() {
	case common.White:
		b.whites.Set(s)
	case common.Black:
		b.blacks.Set(s)
	}
	switch p.Role() {
	case common.Pawn:
		b.pawns.Set(s)
	case common.Knight:
		b.knights.Set(s)
	case common.Bishop:
		b.bishops.Set(s)
	case common.Rook:
		b.rooks.Set(s)
	case common.Queen:
		b.queens.Set(s)
	case common.King:
		b.kings.Set(s)
	}
	return b
}

func (b *Board) Get(s common.Square) (common.Piece, bool) {
	var (
		c common.Color
		r common.Role
	)
	switch {
	case b.whites.Get(s):
		c = common.White
	case b.blacks.Get(s):
		c = common.Black
	default:
		return 0, false
	}
	switch {
	case b.pawns.Get(s):
		r = common.Pawn
	case b.knights.Get(s):
		r = common.Knight
	case b.bishops.Get(s):
		r = common.Bishop
	case b.rooks.Get(s):
		r = common.Rook
	case b.queens.Get(s):
		r = common.Queen
	case b.kings.Get(s):
		r = common.King
	default:
		return 0, false
	}
	return common.NewPiece(c, r), true
}
