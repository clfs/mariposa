package chess_test

import (
	"testing"

	. "github.com/clfs/mariposa/chess"
)

func TestPiece_Value(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p    Piece
		want uint8
	}{
		{WhitePawn, 1},
		{BlackPawn, 9},
		{Piece(30), 30},
	}
	for _, c := range cases {
		got := c.p.Value()
		if got != c.want {
			t.Errorf("%v.Value() = %v, want %v", c.p, got, c.want)
		}
	}
}

func TestPiece_Valid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p    Piece
		want bool
	}{
		{WhitePawn, true},
		{BlackPawn, true},
		{Piece(30), false},
	}
	for _, c := range cases {
		got := c.p.Valid()
		if got != c.want {
			t.Errorf("%v.Valid() = %v, want %v", c.p, got, c.want)
		}
	}
}

func TestPiece_Invalid(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p    Piece
		want bool
	}{
		{WhitePawn, false},
		{BlackPawn, false},
		{Piece(30), true},
	}
	for _, c := range cases {
		got := c.p.Invalid()
		if got != c.want {
			t.Errorf("%v.Invalid() = %v, want %v", c.p, got, c.want)
		}
	}
}

func TestPiece_Color(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p    Piece
		want Color
	}{
		{WhitePawn, White},
		{BlackPawn, Black},
	}
	for _, c := range cases {
		got := c.p.Color()
		if got != c.want {
			t.Errorf("%v.Color() = %v, want %v", c.p, got, c.want)
		}
	}
}

func TestPiece_Role(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p    Piece
		want Role
	}{
		{WhitePawn, Pawn},
		{WhiteKing, King},
		{BlackPawn, Pawn},
		{BlackKing, King},
	}
	for _, c := range cases {
		got := c.p.Role()
		if got != c.want {
			t.Errorf("%v.Role() = %v, want %v", c.p, got, c.want)
		}
	}
}
