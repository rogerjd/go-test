package sortTst

import (
	"fmt"
	"sort"
)

type (
	by func(i, j Person) bool

	personSorter struct {
		person []Person
		by     by
	}
)

func (ps personSorter) Len() int {
	return len(ps.person)
}

func (ps personSorter) Swap(i, j int) {
	ps.person[i], ps.person[j] = ps.person[j], ps.person[i]
}

func (ps personSorter) Less(i, j int) bool {
	return ps.by(ps.person[i], ps.person[j])
}

func (a by) Sort(p []Person) {
	ps := personSorter{
		person: p,
		by:     a,
	}

	sort.Sort(ps)
}

//NameOrNum is entry point
func NameOrNum() {
	p := []Person{
		{Name: "Zqw", Num: 208},
		{Name: "Akku", Num: 178},
		{Name: "Wawa", Num: 015},
		{Name: "Sheetz", Num: -3},
	}
	fmt.Println(p)

	name := func(p1, p2 Person) bool {
		return p1.Name < p2.Name
	}

	num := func(p1, p2 Person) bool {
		return p1.Num < p2.Num
	}

	by(name).Sort(p)
	fmt.Println("\nBy Name", p)

	by(num).Sort(p)
	fmt.Println("\nBy Num", p)
}
