package core

import (
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
		case '1', '2', '3', '4', '5', '6', '7', '8':
			square += Square(ru - '1')
		case '/':
			square -= 16
		default:
			return nil, &InvalidFENError{fen}
		}
		square++
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
	return "fen: not implemented"
}

func (b Board) Pretty() string {
	var sb strings.Builder

	sb.WriteString(" +---+---+---+---+---+---+---+---+\n")

	for i, piece := range b.Pieces {
		sb.WriteString(" | ")
		if piece.IsValid() {
			sb.WriteString(string(piece))
		} else {
			sb.WriteString(" ")
		}
		if i%8 == 7 {
			sb.WriteString(" |\n")
		}
	}

	sb.WriteString(" +---+---+---+---+---+---+---+---+\n")

	sb.WriteString("\n")
	sb.WriteString(b.FEN())

	return sb.String()
}
