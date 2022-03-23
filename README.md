[![GoDoc](https://godoc.org/github.com/churmd/higherorder?status.svg)](https://pkg.go.dev/github.com/churmd/higherorder)

# HigherOrder

Higher order functions written with golang 1.18 beta generics

Functions included are:

-   Identity
-   Compose
-   Reverse
-   Map
-   Filter
-   Foldl
-   Foldr
-   Sort

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

	output := higherorder.Filter(lessThan10, higherorder.Map(double, input)) // output == []int{2, 4, 6, 8}
	fmt.Printf("%v\n", output)
}

```

These functions will complete all their work before returning and will return a variable that can be used immediately in any other function. It does not take an iterable approach to increase performance on large slices or when only the first n elements are needed.
