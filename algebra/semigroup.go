package algebra

import "reflect"

// Isemigroup interface
type Isemigroup interface {
	Concat(Isemigroup) Isemigroup
	List() []interface{}
}

// Semigroup algebra
type Semigroup struct {
	L []interface{}
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

	s.L = append(s.L, args...)
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
func (s Semigroup) Concat(s1 Isemigroup) Isemigroup {
	x := new(Semigroup)
	x.Add(s.List()...)
	x.Add(s1.List()...)
	return x
}

// List returns internal slice
func (s Semigroup) List() []interface{} {
	return s.L
}
