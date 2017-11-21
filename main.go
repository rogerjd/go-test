package main

import (
	"fmt"

	"github.com/rogerjd/go-test2/sortTst"
)

func main() {
	slice()
	sortTst.NameTst()
	sortTst.NameOrNum()
}

func slice() {
	fmt.Print("Hello Test Go Windows")
	fmt.Println("go test2")

	slc := make([]int, 3)
	slc[0] = 2
	slc = append(slc, 7)
	fmt.Printf("%v", slc)
}

func array() {
	var n [3]int
	n = [3]int{1, 2, 3}

	fmt.Printf("%v", n)
}
