package chess

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

type (
	// CastlingRights represents all available castling rights.
	CastlingRights uint8
	// CastlingFlag represents a single castling right.
	CastlingFlag uint8
)

const (
	// FriendOO represents the friendly king-side castling right.
	FriendOO CastlingFlag = 1 << iota
	// FriendOOO represents the friendly queen-side castling right.
	FriendOOO
	// EnemyOO represents the enemy king-side castling right.
	EnemyOO
	// EnemyOOO represents the enemy queen-side castling right.
	EnemyOOO
)

// NewCastlingRights returns a new CastlingRights with the provided rights set
// to true.
func NewCastlingRights(flags ...CastlingFlag) CastlingRights {
	var n CastlingRights
	for _, f := range flags {
		n.Enable(f)
	}
	return CastlingRights(n)
}

func ParseCastlingRightsFEN(s string) (CastlingRights, error) {
	var flags []CastlingFlag
	for _, r := range s {
		switch r {
		case 'K':
			flags = append(flags, FriendOO)
		case 'Q':
			flags = append(flags, FriendOOO)
		case 'k':
			flags = append(flags, EnemyOO)
		case 'q':
			flags = append(flags, EnemyOOO)
		case '-':
			continue
		default:
			return 0, fmt.Errorf("invalid castling rights %s", s)
		}
	}
	return NewCastlingRights(flags...), nil
}

// Enable sets the provided castling right to true.
func (c *CastlingRights) Enable(f CastlingFlag) {
	*c |= CastlingRights(f)
}

// Disable sets the provided castling right to false.
func (c *CastlingRights) Disable(f CastlingFlag) {
	*c &^= CastlingRights(f)
}

// Get returns whether the provided castling right is available.
func (c *CastlingRights) Get(f CastlingFlag) bool {
	return (*c & CastlingRights(f)) != 0
}

// Mirror mirrors the castling rights by color. For example, the friendly
// king-side castling right becomes the enemy king-side castling right.
func (c *CastlingRights) Mirror() {
	*c = (*c << 2) | (*c >> 2)
}

// FEN returns the FEN representation of the castling rights. It assumes the
// castling rights are not mirrored. For example, the friendly king-side
// castling right is represented by K.
func (c *CastlingRights) FEN() string {
	var b strings.Builder
	if c.Get(FriendOO) {
		b.WriteString("K")
	}
	if c.Get(FriendOOO) {
		b.WriteString("Q")
	}
	if c.Get(EnemyOO) {
		b.WriteString("k")
	}
	if c.Get(EnemyOOO) {
		b.WriteString("q")
	}
	s := b.String()
	if s == "" {
		return "-"
	}
	return s
}

// Equal checks if two CastlingRights represent equivalent castling rights.
func (c CastlingRights) Equal(o CastlingRights) bool {
	return c&0xf == o&0xf
}

// Generate lets CastlingRights satisify testing/quick.Generator.
func (CastlingRights) Generate(rand *rand.Rand, size int) reflect.Value {
	return reflect.ValueOf(CastlingRights(rand.Intn(1 << 4)))
}

// Generate lets CastlingFlag satisfy testing/quick.Generator.
func (CastlingFlag) Generate(r *rand.Rand, size int) reflect.Value {
	all := []CastlingFlag{FriendOO, FriendOOO, EnemyOO, EnemyOOO}
	return reflect.ValueOf(all[r.Intn(len(all))])
}
