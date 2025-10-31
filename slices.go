package slices

import "fmt"

func Map[S ~[]FromT, FromT, ToT any](s S, f func(FromT) ToT) []ToT {
	res := make([]ToT, len(s))
	for i, si := range s {
		res[i] = f(si)
	}

	return res
}

func FilterMap[S ~[]FromT, FromT, ToT any](s S, f func(FromT) (ToT, bool)) []ToT {
	res := make([]ToT, 0, len(s))
	for _, si := range s {
		buf, ok := f(si)
		if !ok {
			continue
		}

		res = append(res, buf)
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

func Filter[S ~[]T, T any](s S, filter func(T) bool) S {
	res := make(S, 0)

	for _, si := range s {
		if filter(si) {
			res = append(res, si)
		}
	}

	return res
}

func CollectMap[M ~map[K]V, K comparable, V any](m M) []Pair[K, V] {
	res := make([]Pair[K, V], 0, len(m))

	for key, val := range m {
		res = append(res, Pair[K, V]{key, val})
	}

	return res
}
