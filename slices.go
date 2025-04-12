package slices

func Map[S ~[]FromT, FromT, ToT any](s S, f func(FromT) ToT) []ToT {
	res := make([]ToT, len(s))
	for i, si := range s {
		res[i] = f(si)
	}

	return res
}
