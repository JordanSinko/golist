package main

import (
	"fmt"
	"log"

	"github.com/jordansinko/golist"
)

func main() {
	l, err := golist.NewFileList("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the list has %d items\n", l.Count())

	l.Randomize()

	item, err := l.Take()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("took %s\n", item)

}
