package higherorder

import "sort"

// The identity function
// Returns the give value unchanged
func Identity[X any](x X) X {
	return x
}

// Composes 2 functions together, applies the given value and returns the result
func Compose[X, Y, Z any](g func(Y) Z, f func(X) Y, val X) Z {
	return g(f(val))
}

// Reverses a slice of nay type and returns the result
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

// Applies a function to each element of a slice and returns the resulting list
func Map[X, Y any](f func(X) Y, xs []X) []Y {
	ys := make([]Y, len(xs))

	for i, x := range xs {
		ys[i] = f(x)
	}

	return ys
}

// Given a list, returns a list of those elements that satisfy the given predicate
func Filter[X any](predicate func(X) bool, xs []X) []X {
	result := make([]X, 0)

	for _, x := range xs {
		if predicate(x) {
			result = append(result, x)
		}
	}

	return result
}

// Reduces the slice using the given binary function f from left to right
//
// The identity value is passed to f with the first element in the slice to start with
// the result of that is passed to f with the sencond element of the slice
// and so on
//
// Example:
// f := func(x, y int) int { return x + y }
// Foldl(f, 0, []int{1,2,3}]) == (((0 + 1) + 2) + 3)
func Foldl[X, Y any](f func(y Y, x X) Y, identity Y, values []X) Y {
	acc := identity

	for _, v := range values {
		acc = f(acc, v)
	}

	return acc
}

// Reduces the slice using the given binary function f from right to left
//
// The last element of the slice is passed to f with the identity value to start
// the second to last element of the slice is passed to f with the result of that
// and so on
//
// Example:
// f := func(x, y int) int { return x + y }
// Foldr(f, 0, []int{1,2,3}]) == 1 + (2 + (3 + 0))
func Foldr[X, Y any](f func(x X, y Y) Y, identity Y, values []X) Y {
	acc := identity

	for i := len(values) - 1; i >= 0; i-- {
		acc = f(values[i], acc)
	}

	return acc
}

// private struct that implements the sort.Interface
// used in the Sort function to leverage the std lib
type sortableSlice[X any] struct {
	values   []X
	lessThan func(a, b X) bool
}

func (ss sortableSlice[X]) Len() int {
	return len(ss.values)
}

func (ss sortableSlice[X]) Less(i, j int) bool {
	return ss.lessThan(ss.values[i], ss.values[j])
}

func (ss *sortableSlice[X]) Swap(i, j int) {
	tmp := ss.values[i]
	ss.values[i] = ss.values[j]
	ss.values[j] = tmp
}

// Given a list, returns the sorted list according to the given lessThan function
//
// lessThan function returns true
// when the first param a comes before the second param b
func Sort[X any](lessThan func(a, b X) bool, xs []X) []X {
	ss := sortableSlice[X]{
		values:   xs,
		lessThan: lessThan,
	}

	sort.Sort(&ss)
	return ss.values
}
