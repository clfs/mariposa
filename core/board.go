package core

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Board struct {
	Pieces         [NumSquares]Piece
	SideToMove     Color
	CastleRights   Castling
	EPTarget       Square
	HalfMoveClock  uint64
	FullMoveNumber uint64
}

// NewBoard returns a new board from a FEN.
func NewBoard(fen string) (Board, error) {
	board := Board{}

	fields := strings.Split(fen, " ")
	if len(fields) != 6 {
		return Board{}, &InvalidFENError{fen}
	}

	// 1. Piece placement.
	square := A8
	for _, ru := range fields[0] {
		switch ru {
		case 'P', 'p', 'N', 'n', 'B', 'b', 'R', 'r', 'Q', 'q', 'K', 'k':
			board.Pieces[square] = PieceFromRune(ru)
			square++
		case '1', '2', '3', '4', '5', '6', '7', '8':
			square += Square(ru - '0')
		case '/':
			square -= 16
		default:
			return Board{}, &InvalidFENError{fen}
		}
	}

	// 2. Side to move.
	switch fields[1] {
	case "w":
		board.SideToMove = White
	case "b":
		board.SideToMove = Black
	default:
		return Board{}, &InvalidFENError{fen}
	}

	// 3. Castling availability.
	board.CastleRights = Castling{}
	if fields[2] != "-" {
		for _, ru := range fields[2] {
			log.Println(ru)
			switch ru {
			case 'K':
				board.CastleRights.WhiteOO = true
			case 'Q':
				board.CastleRights.WhiteOOO = true
			case 'k':
				board.CastleRights.BlackOO = true
			case 'q':
				board.CastleRights.BlackOOO = true
			default:
				return Board{}, &InvalidFENError{fen}
			}
		}
	}

	// 4. En passant target square.
	board.EPTarget = SquareFromString(fields[3])

	// 5. Half move clock.
	halfMoveClock, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return Board{}, &InvalidFENError{fen}
	}
	board.HalfMoveClock = uint64(halfMoveClock)

	// 6. Full move number.
	fullMoveNumber, err := strconv.ParseUint(fields[5], 10, 64)
	if err != nil {
		return Board{}, &InvalidFENError{fen}
	}
	board.FullMoveNumber = uint64(fullMoveNumber)

	return board, nil
}

func (b Board) PieceAt(s Square) Piece {
	return b.Pieces[s]
}

func (b Board) FEN() string {
	// rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1
	var sb strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		num := 0
		for f := FileA; f <= FileH; f++ {
			piece := b.PieceAt(SquareAt(f, r))
			if piece.IsEmpty() {
				num += 1
			} else {
				if num != 0 {
					sb.WriteString(strconv.Itoa(num))
				}
				sb.WriteString(piece.String())
				num = 0
			}
		}
		if num != 0 {
			sb.WriteString(strconv.Itoa(num))
		}
		if r != Rank1 {
			sb.WriteString("/")
		}
	}

	sb.WriteString(" ")
	sb.WriteString(b.SideToMove.String())
	sb.WriteString(" ")
	sb.WriteString(b.CastleRights.String())
	sb.WriteString(" ")
	sb.WriteString(b.EPTarget.String())
	sb.WriteString(" ")
	sb.WriteString(strconv.FormatUint(b.HalfMoveClock, 10))
	sb.WriteString(" ")
	sb.WriteString(strconv.FormatUint(b.FullMoveNumber, 10))

	return sb.String()
}

func (b Board) Pretty() string {
	var sb strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		for f := FileA; f <= FileH; f++ {
			fmt.Fprintf(&sb, "%s ", b.PieceAt(SquareAt(f, r)))
		}
		sb.WriteString("\n")
	}
	fmt.Fprintf(&sb, "FEN: %s", b.FEN())

	return sb.String()
}
