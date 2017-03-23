package algebra

import (
	"testing"
)

func TestEquals(t *testing.T) {
	a := Setoid{1}
	b := Setoid{1}
	c := Setoid{1}

	t.Parallel()

	if !a.Equals(a) {
		t.Errorf("reflexivity -> expected %s to be equal %s", a, a)
	}
	if !(a.Equals(b) && b.Equals(a)) {
		t.Errorf("symmetry -> expected %s to be equal %s", a, b)
	}
	if !(a.Equals(b) && b.Equals(c) && a.Equals(c)) {
		t.Errorf("transitivity -> expected %s to be equal %s and %s", a, b, c)
	}
}
