package common

//go:generate stringer -type=Rank
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

func (r Rank) Value() uint8 {
	return uint8(r)
}

func (r Rank) Valid() bool {
	return r <= Rank8
}

func (r Rank) Invalid() bool {
	return !r.Valid()
}
