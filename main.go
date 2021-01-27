package main

import (
	"fmt"
)

var words []string

func main() {
	var q string
	fmt.Print("query >>")
	fmt.Scanln(&q)
	fmt.Println(sortSlice(words, q)[:20])
}
