package core

import "fmt"

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
	return string('A' + uint8(f))
}
