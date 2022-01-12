package higherorder

func Identity[X any](x X) X {
	return x
}

func Map[X any, Y any](xs []X, f func(X) Y) []Y {
	ys := make([]Y, len(xs))

	for i, x := range xs {
		ys[i] = f(x)
	}

	return ys
}

func Filter[X any](xs []X, predicate func(X) bool) []X {
	result := make([]X, 0)

	for _, x := range xs {
		if predicate(x) {
			result = append(result, x)
		}
	}

	return result
}
