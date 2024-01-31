package main

import (
	"fmt"

	"github.com/jordansinko/golist"
)

func main() {
	l := golist.NewList([]string{})
	fmt.Printf("the list has %d items", l.Count())
}
