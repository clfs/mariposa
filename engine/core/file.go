package core

import (
	"fmt"
	"math/rand"
	"reflect"
)

type File uint8

const NumFiles = 8

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

func (f File) IsValid() bool {
	return f <= FileH
}

func (f File) String() string {
	if !f.IsValid() {
		return fmt.Sprintf("File(%d)", uint8(f))
	}
	return string('a' + uint8(f))
}

// Generate implements quick.Generator. It only generates pre-defined
// constants of type File.
func (File) Generate(rand *rand.Rand, size int) reflect.Value {
	_ = size
	return reflect.ValueOf(File(rand.Intn(NumFiles)))
}
