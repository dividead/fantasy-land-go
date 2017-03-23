package algebra

import (
	"fmt"
	"testing"
)

func TestConcat(t *testing.T) {
	s := new(Semigroup)
	s.Add(1, 2, 3)
	fmt.Println(s.List)
	b := new(Semigroup)
	b.Add(4, 5, 6)
	fmt.Println(b.List)

	c := new(Semigroup)
	c.Add(7, 8, 9)
	fmt.Println(c.List)

	//a.concat(b).concat(c) is equivalent to a.concat(b.concat(c)) (associativity)

	x := s.Concat(b).Concat(c)
	fmt.Println(x.List)

	y := s.Concat(b.Concat(c))
	fmt.Println(y.List)
}
