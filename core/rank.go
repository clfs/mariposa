package core

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Rank uint8

const NumRanks = 8

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

func (r Rank) IsValid() bool {
	return r <= Rank8
}

func (r Rank) String() string {
	if !r.IsValid() {
		return fmt.Sprintf("Rank(%d)", uint8(r))
	}
	return string('1' + uint8(r))
}

// Generate implements quick.Generator. It only generates pre-defined
// constants of type Rank.
func (Rank) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	return reflect.ValueOf(Rank(rand.Intn(NumRanks)))
}
