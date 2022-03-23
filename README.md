[![GoDoc](https://godoc.org/github.com/churmd/higherorder?status.svg)](https://pkg.go.dev/github.com/churmd/higherorder)

# HigherOrder

Higher order functions written with golang 1.18 generics

Functions included are:

-   All
-   Compose
-   Filter
-   Foldl
-   Foldr
-   Identity
-   Map
-   Reverse
-   Sort

These functions will complete all their work before returning and will return a variable that can be used immediately in any other function. It does not take an iterable approach to increase performance on large slices or when only the first n elements are needed.

Example

```go
package main

import (
	"fmt"

	"github.com/churmd/higherorder"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	double := func(x int) int { return x * 2 }
	lessThan10 := func(x int) bool { return x < 10 }

	// output == []int{2, 4, 6, 8}
	output := higherorder.Filter(lessThan10, higherorder.Map(double, input))
	fmt.Printf("%v\n", output)
}

```
