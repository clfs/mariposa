package chess

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

const StartingBoardFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"

type Board struct {
	friends   Bitboard
	enemies   Bitboard
	pawns     Bitboard
	knights   Bitboard
	bishops   Bitboard
	rooks     Bitboard
	queens    Bitboard
	kings     Bitboard
	isFlipped bool
}

// NewBoard returns a new board with the given board-FEN representation.
func NewBoard(boardFEN string) (Board, error) {
	b := new(Board)
	err := b.setFromBoardFEN(boardFEN)
	return *b, err
}

func (b *Board) setFromBoardFEN(boardFEN string) error {
	square := A8
	for _, r := range boardFEN {
		switch r {
		case '1', '2', '3', '4', '5', '6', '7', '8':
			square += Square(r - '0')
		case '/':
			square -= 16
		default:
			piece, err := ParsePieceFEN(string(r))
			if err != nil {
				return fmt.Errorf("invalid board: %s", boardFEN)
			}
			b.put(piece, square)
		}
		square++
	}
	return nil
}

// Flip flips the board vertically and also returns the board.
func (b *Board) Flip() *Board {
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
	return b
}

// FEN returns the FEN representation of the board.
func (b *Board) FEN() string {
	var sb strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		skip := 0
		for f := FileA; f <= FileH; f++ {
			piece, ok := b.get(NewSquare(f, r))
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
		if r != Rank1 {
			fmt.Fprintf(&sb, "/")
		}
	}

	return sb.String()
}

func (b *Board) put(p Piece, s Square) *Board {
	switch p.Color() {
	case White:
		b.friends.Set1(s)
	case Black:
		b.enemies.Set1(s)
	}
	switch p.Role() {
	case Pawn:
		b.pawns.Set1(s)
	case Knight:
		b.knights.Set1(s)
	case Bishop:
		b.bishops.Set1(s)
	case Rook:
		b.rooks.Set1(s)
	case Queen:
		b.queens.Set1(s)
	case King:
		b.kings.Set1(s)
	}
	return b
}

func (b *Board) get(s Square) (Piece, bool) {
	var (
		c Color
		r Role
	)
	switch {
	case b.friends.Get(s):
		c = White
	case b.enemies.Get(s):
		c = Black
	default:
		return 0, false
	}
	switch {
	case b.pawns.Get(s):
		r = Pawn
	case b.knights.Get(s):
		r = Knight
	case b.bishops.Get(s):
		r = Bishop
	case b.rooks.Get(s):
		r = Rook
	case b.queens.Get(s):
		r = Queen
	case b.kings.Get(s):
		r = King
	default:
		return 0, false
	}
	return NewPiece(c, r), true
}

// Generate lets Board satisfy testing/quick.Generator.
func (Board) Generate(r *rand.Rand, size int) reflect.Value {
	// TODO: improve this once move generation is implemented.
	b, err := NewBoard(StartingBoardFEN)
	if err != nil {
		panic(err)
	}
	return reflect.ValueOf(b)
}
