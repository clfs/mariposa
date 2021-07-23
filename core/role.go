package core

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Role uint8

const NumRoles = 6

const (
	Pawn Role = iota + 1
	Knight
	Bishop
	Rook
	Queen
	King
)

func (r Role) IsValid() bool {
	return r >= Pawn && r <= King
}

func (r Role) String() string {
	switch r {
	case Pawn:
		return "Pawn"
	case Knight:
		return "Knight"
	case Bishop:
		return "Bishop"
	case Rook:
		return "Rook"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return fmt.Sprintf("Role(%d)", uint8(r))
	}
}

// Generate implements quick.Generator. It only generates pre-defined
// constants of type Role.
func (Role) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	x := []Role{Pawn, Knight, Bishop, Rook, Queen, King}
	return reflect.ValueOf(x[rand.Intn(len(x))])
}
