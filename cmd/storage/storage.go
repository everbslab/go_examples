package main

import (
	"github.com/everbslab/go_examples/pkg/storage"
	"log"
)

var stubData = []string{
	"hello",
	"world",
}

func main() {
	s := storage.NewFile("a.txt")

	err := s.Write(stubData)
	if err != nil {
		log.Fatal(err)
	}
}
