package algebra

import (
	"reflect"
)

//Setoid algebra
type Setoid struct {
	Value interface{}
}

type setoid interface {
	Equals()
}

//Equals :: Setoid a => a ~> a -> Boolean
func (s Setoid) Equals(b Setoid) bool {
	return reflect.TypeOf(s.Value) == reflect.TypeOf(b.Value) && s.Value == b.Value
}
