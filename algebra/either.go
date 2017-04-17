package algebra

import "reflect"

func main() {
	e := &Either{&Left{"nothing"}, nil}
	e2 := &Either{&Left{"nothing"}, &Right{1}}
	maybePrint(e)
	maybePrint(e2)
}

// Either ->
type Either struct {
	L *Left
	R *Right
}

// Left <-
type Left struct {
	Error string
}

// Right ->
type Right struct {
	Value interface{}
}

func maybePrint(e *Either) {
	if e.R == nil {
		println(e.L.Error)
		return
	}

	println(reflect.ValueOf(e.R.Value).Int())
}
