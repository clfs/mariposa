package main

import (
	"fmt"
	"log"

	"github.com/clfs/mariposa/core"
)

func main() {

	b, err := core.NewBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b.Pretty())

}
