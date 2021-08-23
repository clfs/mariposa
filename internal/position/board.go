package position

import (
	"fmt"
	"strings"

	"github.com/clfs/mariposa/internal/bitboard"
	"github.com/clfs/mariposa/internal/common"
)

type Board struct {
	friends   bitboard.Bitboard
	enemies   bitboard.Bitboard
	pawns     bitboard.Bitboard
	knights   bitboard.Bitboard
	bishops   bitboard.Bitboard
	rooks     bitboard.Bitboard
	queens    bitboard.Bitboard
	kings     bitboard.Bitboard
	isFlipped bool
}

func (b *Board) Flip() {
	b.isFlipped = !b.isFlipped
	b.friends.Flip()
	b.enemies.Flip()
	b.pawns.Flip()
	b.knights.Flip()
	b.bishops.Flip()
	b.rooks.Flip()
	b.queens.Flip()
	b.kings.Flip()
	b.friends, b.enemies = b.enemies, b.friends
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
			fmt.Fprintf(&sb, "%s", piece.FEN())
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

// Put places p at s and returns b. This must not be used during gameplay.
func (b *Board) Put(p common.Piece, s common.Square) *Board {
	switch p.Color() {
	case common.White:
		b.friends.Set(s)
	case common.Black:
		b.enemies.Set(s)
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

// Get returns the piece at s and a boolean indicating whether one was found.
func (b *Board) Get(s common.Square) (common.Piece, bool) {
	var (
		c common.Color
		r common.Role
	)
	switch {
	case b.friends.Get(s):
		c = common.White
	case b.enemies.Get(s):
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
