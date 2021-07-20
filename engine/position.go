package engine

import (
	"strconv"
	"strings"
)

type Position struct {
	Board           [64]Piece
	SideToMove      Color
	WhiteOOO        bool
	WhiteOO         bool
	BlackOOO        bool
	BlackOO         bool
	EPTarget        Square
	HalfMoveClock   uint8
	FullMoveCounter uint16
}

func NewPosition() Position {
	return Position{}
}

func (p *Position) FromFEN(fen string) error {
	fields := strings.Split(fen, " ")

	// 1. Piece placement
	square := A8
	for _, r := range fields[0] {
		switch r {
		case 'P', 'p', 'N', 'n', 'B', 'b', 'R', 'r', 'Q', 'q', 'K', 'k':
			piece, err := PieceFromRune(r)
			if err != nil {
				return &ParseFENError{fen}
			}
			p.Board[square] = piece
		case '1', '2', '3', '4', '5', '6', '7', '8':
			square += Square(r - '0')
		case '/':
			square -= 16
		default:
			return &ParseFENError{fen}
		}
	}

	// 2. Side to move
	switch fields[1] {
	case "w":
		p.SideToMove = White
	case "b":
		p.SideToMove = Black
	default:
		return &ParseFENError{fen}
	}

	// 3. Castling availability
	if fields[2] != "-" {
		for _, r := range fields[2] {
			switch r {
			case 'K':
				p.WhiteOO = true
			case 'Q':
				p.WhiteOOO = true
			case 'k':
				p.BlackOO = true
			case 'q':
				p.BlackOOO = true
			default:
				return &ParseFENError{fen}
			}
		}
	}

	// 4. En passant target square
	square, err := SquareFromString(fields[3])
	if err != nil {
		return &ParseFENError{fen}
	}
	p.EPTarget = square

	// 5. Half move clock
	halfMoveClock, err := strconv.ParseUint(fields[4], 10, 8)
	if err != nil {
		return &ParseFENError{fen}
	}
	p.HalfMoveClock = uint8(halfMoveClock)

	// 6. Full move number
	fullMoveCounter, err := strconv.ParseUint(fields[5], 10, 16)
	if err != nil {
		return &ParseFENError{fen}
	}
	p.FullMoveCounter = uint16(fullMoveCounter)

	return nil
}

func (p *Position) FEN() string {
	var fen strings.Builder

	return fen.String()
}
