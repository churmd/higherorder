package higherorder_test

import (
	"strconv"
	"testing"

	"github.com/churmd/higherorder"
	"github.com/stretchr/testify/assert"
)

func TestIdentity(t *testing.T) {
	ids := []interface{}{
		10,
		"a",
		10.0,
		'a',
		[]int{1, 2, 3},
	}

	for _, val := range ids {
		actualOutput := higherorder.Identity(val)

		assert.Equal(t, val, actualOutput)
	}
}

func TestCompose(t *testing.T) {
	f := func(x int) string {return strconv.Itoa(x)}
	g := func(y string) bool {return y == "10"}

	actualOutput := higherorder.Compose(10, f, g)

	assert.True(t, actualOutput)
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toString := func(x int) string { return strconv.Itoa(x) }
	expectedOutput := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	actualOutput := higherorder.Map(input, toString)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isEven := func(x int) bool { return x%2 == 0 }
	expectedOutput := []int{2, 4, 6, 8, 10}

	actualOutput := higherorder.Filter(input, isEven)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestFilterNoElemsPass(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isNegavtive := func(x int) bool { return x < 0 }
	expectedOutput := []int{}

	actualOutput := higherorder.Filter(input, isNegavtive)

	assert.Equal(t, expectedOutput, actualOutput)
}

// Benchmarks

func largeList() []int {
	n := 1000
	l := make([]int, n)

	for i := 0; i < n; i++ {
		l[i] = i
	}

	return l
}

func BenchmarkMap(b *testing.B) {
	input := largeList()
	double := func(x int) int { return x * 2 }

	for i := 0; i < b.N; i++ {
		higherorder.Map(input, double)
	}
}

func BenchmarkFilter(b *testing.B) {
	input := largeList()
	isEven := func(x int) bool { return x%2 == 0 }

	for i := 0; i < b.N; i++ {
		higherorder.Filter(input, isEven)
	}
}
