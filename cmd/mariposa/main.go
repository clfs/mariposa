package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/clfs/mariposa/engine/board"
)

func main() {
	b, err := board.New("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(b)

	fmt.Println(b.Pretty())
}
