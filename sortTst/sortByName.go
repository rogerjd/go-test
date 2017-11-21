package sortTst

import (
	"fmt"
	"sort"
)

//https://golang.org/pkg/sort/#Search

type (
	//Person is struct for sort test
	Person struct {
		Name string
		Num  int
	}

	//ByNum slice of Person sorted by Num
	//	ByNum []Person

	ByName []Person
)

/*
func (a ByNum) Len() int {
	return len(a)
}

func (a ByNum) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByNum) Less(i, j int) bool {
	return a[i].Num < a[j].Num
}
*/

func (a ByName) Len() int {
	return len(a)
}

func (a ByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

func NameTst() {
	a := ByName{
		{Name: "joe", Num: 12},
		{Name: "abc", Num: 34},
		{Name: "def", Num: 56},
		{Name: "gji", Num: 78},
	}
	sort.Sort(a)
	fmt.Println("\nBy Name", a)
}
