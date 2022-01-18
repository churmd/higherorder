package higherorder

func Identity[X any](x X) X {
	return x
}

func Compose[X, Y, Z any](val X, f func(X) Y, g func(Y) Z) Z {
	return g(f(val))
}

func Reverse[X any](xs []X) []X {
	numElems := len(xs)

	for i := 0; i < numElems/2; i++ {
		tmp := xs[i]
		rightIndex := numElems - 1 - i
		xs[i] = xs[rightIndex]
		xs[rightIndex] = tmp
	}

	return xs
}

func Map[X, Y any](xs []X, f func(X) Y) []Y {
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

func Foldl[X, Y any](f func(y Y, x X) Y, identity Y, values []X) Y {
	acc := identity

	for _, v := range values {
		acc = f(acc, v)
	}

	return acc
}

func Foldr[X, Y any](f func(x X, y Y) Y, identity Y, values []X) Y {
	acc := identity

	for i := len(values) - 1; i >= 0; i-- {
		acc = f(values[i], acc)
	}

	return acc
}
