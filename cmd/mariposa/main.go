package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/clfs/mariposa/engine/core"
)

func main() {
	b, err := core.NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		log.Fatal(err)
	}

	spew.Dump(b)

	fmt.Println(b.Pretty())
}
