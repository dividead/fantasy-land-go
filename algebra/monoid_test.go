package algebra

import (
	"testing"
)

func TestMonoid(t *testing.T) {
	s1 := new(Semigroup)
	s2 := new(Semigroup)

	s1.Add(1, 2, 3)

	s3 := s2.Concat(*s1)

	t.Log(s1, s2, s3)

	m := new(Monoid)

	m1 := Empty().Concat(m)
	m2 := m.Concat(Empty())

	s4 := s3.Concat(m.Semigroup)

	s5 := s4.Concat(Empty()).Concat(Empty().Concat(m.Concat(s3)))

	t.Log(m1, m2, m, s4, s5)
}
