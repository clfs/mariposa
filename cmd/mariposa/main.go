package main

import (
	"log"
	"os"

	"github.com/clfs/mariposa/uci"
)

func main() {
	if err := uci.New(os.Stdin, os.Stdout).Run(); err != nil {
		log.Fatal(err)
	}
}
