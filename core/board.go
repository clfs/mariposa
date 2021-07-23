package core

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Pieces         [NumSquares]Piece
	SideToMove     Color
	WhiteOOO       bool
	WhiteOO        bool
	BlackOOO       bool
	BlackOO        bool
	EPTarget       Square
	HalfMoveClock  int
	FullMoveNumber int
}

// NewBoard returns a new board from a FEN.
func NewBoard(fen string) (*Board, error) {
	board := Board{}

	fields := strings.Split(fen, " ")
	if len(fields) != 6 {
		return nil, &InvalidFENError{fen}
	}

	// 1. Piece placement.
	square := A8
	for _, ru := range fields[0] {
		switch ru {
		case 'P', 'p', 'N', 'n', 'B', 'b', 'R', 'r', 'Q', 'q', 'K', 'k':
			board.Pieces[square] = PieceFromString(string(ru))
			square += 1
		case '1', '2', '3', '4', '5', '6', '7', '8':
			square += Square(ru - '0')
		case '/':
			square -= 16
		default:
			return nil, &InvalidFENError{fen}
		}
	}

	// 2. Side to move.
	switch fields[1] {
	case "w":
		board.SideToMove = White
	case "b":
		board.SideToMove = Black
	default:
		return nil, &InvalidFENError{fen}
	}

	// 3. Castling availability.
	if fields[2] != "-" {
		for _, ru := range fields[2] {
			switch ru {
			case 'K':
				board.WhiteOO = true
			case 'Q':
				board.WhiteOOO = true
			case 'k':
				board.BlackOO = true
			case 'q':
				board.BlackOOO = true
			default:
				return nil, &InvalidFENError{fen}
			}
		}
	}

	// 4. En passant target square.
	board.EPTarget = SquareFromString(fields[3])

	// 5. Half move clock.
	halfMoveClock, err := strconv.Atoi(fields[4])
	if err != nil {
		return nil, &InvalidFENError{fen}
	}
	board.HalfMoveClock = halfMoveClock

	// 6. Full move number.
	fullMoveNumber, err := strconv.Atoi(fields[5])
	if err != nil {
		return nil, &InvalidFENError{fen}
	}
	board.FullMoveNumber = fullMoveNumber

	return &board, nil
}

func (b Board) FEN() string {
	return "not implemented"
}

func (b Board) Pretty() string {
	var sb strings.Builder

	for r := Rank8; r <= Rank8; r-- {
		for f := FileA; f <= FileH; f++ {
			sq := SquareFromFileRank(f, r)
			fmt.Fprintf(&sb, "%s ", b.Pieces[sq])
		}
		sb.WriteString("\n")
	}

	fmt.Fprintf(&sb, "Side to move: %s\n", b.SideToMove)
	fmt.Fprintf(&sb, "En passant target: %s\n", b.EPTarget)
	fmt.Fprintf(&sb, "Half move clock: %d\n", b.HalfMoveClock)
	fmt.Fprintf(&sb, "Full move number: %d\n", b.FullMoveNumber)

	fmt.Fprintf(&sb, "White O-O-O: %v\n", b.WhiteOOO)
	fmt.Fprintf(&sb, "White O-O: %v\n", b.WhiteOO)
	fmt.Fprintf(&sb, "Black O-O-O: %v\n", b.BlackOOO)
	fmt.Fprintf(&sb, "Black O-O: %v\n", b.BlackOO)

	fmt.Fprintf(&sb, "FEN: %s", b.FEN())

	return sb.String()
}
