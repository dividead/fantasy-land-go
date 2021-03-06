package algebra

type monoid interface {
	Empty() // Monoid
}

// Monoid algebra
type Monoid struct {
	Semigroup
}

// Empty :: Monoid m => () -> m
func Empty() *Monoid {
	x := new(Monoid)
	return x
}
