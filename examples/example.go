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
