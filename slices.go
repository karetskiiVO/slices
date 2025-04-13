package slices

import "fmt"

func Map[S ~[]FromT, FromT, ToT any](s S, f func(FromT) ToT) []ToT {
	res := make([]ToT, len(s))
	for i, si := range s {
		res[i] = f(si)
	}

	return res
}

type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

func Join[S1 ~[]T1, S2 ~[]T2, T1, T2 any](s1 S1, s2 S2) []Pair[T1, T2] {
	if len(s1) != len(s2) {
		panic(fmt.Sprintf("in join: slices have different length s1[%v] s2[%v]", len(s1), len(s2)))
	}
	res := make([]Pair[T1, T2], len(s1))

	for i := range res {
		res[i] = Pair[T1, T2]{
			First:  s1[i],
			Second: s2[i],
		}
	}

	return res
}
