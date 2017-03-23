package algebra

import "reflect"

type semigroup interface {
	Concat()
}

// Semigroup algebra
type Semigroup struct {
	List []interface{}
}

// Add Elements To Semigroup
func (s *Semigroup) Add(args ...interface{}) *Semigroup {
	if len(args) == 0 {
		return s
	}

	// check if all elements of same type
	t := reflect.TypeOf(args[0])
	for _, el := range args {
		if reflect.TypeOf(el) != t {
			panic("types are not monotonous")
		}
	}

	s.List = append(s.List, args...)
	return s
}

// Push Element To Semigroup
// func (s *Semigroup) Push(el interface{}) *Semigroup {
// 	if len(s.List) == 0 {
// 		s.List = append(s.List, el)
// 		return s
// 	}

// 	listtype := reflect.TypeOf(s.List[0])
// 	elemtype := reflect.TypeOf(el)
// 	if listtype == elemtype {
// 		s.List = append(s.List, el)
// 	}

// 	return s
// }

// Concat :: Semigroup a => a ~> a -> a
func (s *Semigroup) Concat(z *Semigroup) *Semigroup {
	x := new(Semigroup)
	x.Add(s.List...)
	x.Add(z.List...)
	return x
}
