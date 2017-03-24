package algebra

type monoid interface {
	Empty()
}

// Monoid algebra
type Monoid struct {
	*Semigroup
}

// Empty :: Monoid m => () -> m
func Empty() *Monoid {
	x := Semigroup.NewSemigroup(Semigroup{})
	return x
}
