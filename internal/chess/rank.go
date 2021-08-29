package chess

type Rank uint8

// Note that Rank1 is 0.
const (
	Rank1 Rank = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
)

// Flipped returns a rank vertically flipped from the original.
func (r Rank) Flipped() Rank {
	return 7 - r
}
