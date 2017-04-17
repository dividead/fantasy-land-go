package algebra

type functor interface {
	Map()
}

// Functor struct
type Functor struct {
}

// Map :: Functor f => f a ~> (a -> b) -> f b
// must return a value of the same Functor
func (f Functor) Map(fn func()) Functor {
	return f
}
