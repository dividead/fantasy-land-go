package main

import (
	"fmt"

	a "github.com/dividead/fantasy-land-go/algebra"
)

func main() {
	// hold dep
	_ = new(a.Setoid)
	fmt.Println("unicorns")

	m := new(a.Monoid)

	m1 := a.Empty().Concat(m)
	m2 := m.Concat(a.Empty())

}

// func pipe(fs ...any) {
// 	for f := range fs {

// 	}
// }

// func one() int {
// 	return 1
// }

// func addX(i int) func(int) int {
// 	return func(y int) int {
// 		return y + i
// 	}
// }

// var add2 = addX(2)

// pipe(add2, one) => 3
