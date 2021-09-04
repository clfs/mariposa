package chess

type Role uint8

const (
	Pawn Role = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

// Valid returns whether the role is valid.
func (r Role) Valid() bool {
	return Pawn <= r && r <= King
}
