package main

import (
	"fmt"

	"github.com/clfs/mariposa/position"
)

func main() {
	s := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	p, err := position.NewPosition(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(p.Pretty())
}
