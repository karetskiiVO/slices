package slices

import "fmt"

func Map[S ~[]FromT, FromT, ToT any](s S, f func(FromT) ToT) []ToT {
	res := make([]ToT, len(s))
	for i, si := range s {
		res[i] = f(si)
	}

	return res
}

func Join[S1 ~[]T1, S2 ~[]T2, T1, T2 any](s1 S1, s2 S2) []struct {
	fst T1
	snd T2
} {
	if len(s1) != len(s2) {
		panic(fmt.Sprintf("in join: slices have different length s1[%v] s2[%v]", len(s1), len(s2)))
	}
	res := make([]struct {
		fst T1
		snd T2
	}, len(s1))

	for i := range res {
		res[i] = struct {
			fst T1
			snd T2
		}{
			fst: s1[i],
			snd: s2[i],
		}
	}

	return res
}
