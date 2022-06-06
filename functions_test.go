package higherorder_test

import (
	"strconv"
	"testing"

	"github.com/churmd/higherorder"
	"github.com/stretchr/testify/assert"
)

func TestIdentity(t *testing.T) {
	myInt := 1
	ids := []interface{}{
		10,
		"a",
		10.0,
		'a',
		[]int{1, 2, 3},
		&myInt,
	}

	for _, val := range ids {
		actualOutput := higherorder.Identity(val)

		assert.Equal(t, val, actualOutput)
	}
}

func TestCompose(t *testing.T) {
	f := func(x int) string { return strconv.Itoa(x) }
	g := func(y string) bool { return y == "10" }

	actualOutput := higherorder.Compose(g, f, 10)

	assert.True(t, actualOutput)
}

func TestReverse(t *testing.T) {
	inputs := [][]int{
		{},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{1, 2, 3, 4, 5},
	}

	expectedOutputs := [][]int{
		{},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		{5, 4, 3, 2, 1},
	}

	for i, input := range inputs {
		originalInput := make([]int, len(input))
		copy(originalInput, input)
		expectedOutput := expectedOutputs[i]
		actualOutput := higherorder.Reverse(input)

		assert.Equal(t, expectedOutput, actualOutput)
		assert.Equal(t, originalInput, input, "orignal list is not preserved")
	}

}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toString := func(x int) string { return strconv.Itoa(x) }
	expectedOutput := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	actualOutput := higherorder.Map(toString, input)

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, input, "orignal list is not preserved")
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isEven := func(x int) bool { return x%2 == 0 }
	expectedOutput := []int{2, 4, 6, 8, 10}

	actualOutput := higherorder.Filter(isEven, input)

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, input, "orignal list is not preserved")
}

func TestFilterNoElemsPass(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isNegavtive := func(x int) bool { return x < 0 }
	expectedOutput := []int{}

	actualOutput := higherorder.Filter(isNegavtive, input)

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, input, "orignal list is not preserved")
}

func TestAll(t *testing.T) {
	inputs := [][]int{
		{},
		{0, 2, 4, 6, 8},
		{0, 1, 2, 3, 4},
	}

	expectedOutputs := []bool{
		true,
		true,
		false,
	}

	for i, input := range inputs {
		isEven := func(x int) bool { return x%2 == 0 }
		actualOutput := higherorder.All(isEven, input)

		assert.Equal(t, expectedOutputs[i], actualOutput)
	}
}

func TestAny(t *testing.T) {
	inputs := [][]int{
		{},
		{1, 2, 3, 4},
		{1, 3, 5, 7},
	}

	expectedOutputs := []bool{
		false,
		true,
		false,
	}

	for i, input := range inputs {
		isEven := func(x int) bool { return x%2 == 0 }
		actualOutput := higherorder.Any(isEven, input)

		assert.Equal(t, expectedOutputs[i], actualOutput)
	}
}

func TestFirst(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	predicates := []higherorder.Predicate[int]{
		func(x int) bool { return x%2 != 0 },
		func(x int) bool { return x%2 == 0 },
		func(x int) bool { return x > 9 },
	}
	expectedOutputs := []int{
		1,
		2,
		10,
	}

	for i := 0; i < len(predicates); i++ {
		actualOutput, err := higherorder.First(predicates[i], input)

		assert.NoError(t, err)
		assert.Equal(t, expectedOutputs[i], actualOutput)
	}
}

func TestFirstWithNoMatches(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	predicate := func(x int) bool { return x > 10 }
	expectedError := "first failed to find an element that satisfied the given predicate"

	actualOutput, err := higherorder.First(predicate, input)

	assert.EqualError(t, err, expectedError)
	assert.Equal(t, 0, actualOutput)

}

func TestFoldl(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	identity := 0
	sum := func(x, y int) int { return x + y }
	expectedOutput := 55

	actualOutput := higherorder.Foldl(sum, identity, input)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestFoldlNonAssociativeFunc(t *testing.T) {
	input := []int{1, 2, 3}
	identity := 4
	f := func(x, y int) int { return 2*x + y }
	expectedOutput := 43

	actualOutput := higherorder.Foldl(f, identity, input)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestFoldlChangeType(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	identity := ""
	toStringAndConcat := func(x string, y int) string { return x + strconv.Itoa(y) }
	expectedOutput := "12345678910"

	actualOutput := higherorder.Foldl(toStringAndConcat, identity, input)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestFoldr(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	identity := 0
	sum := func(x, y int) int { return x + y }
	expectedOutput := 55

	actualOutput := higherorder.Foldr(sum, identity, input)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestFoldrNonAssociativeFunc(t *testing.T) {
	input := []int{1, 2, 3}
	identity := 4
	f := func(x, y int) int { return 2*x + y }
	expectedOutput := 16

	actualOutput := higherorder.Foldr(f, identity, input)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestFoldrChangeType(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	identity := ""
	toStringAndConcat := func(x int, y string) string { return strconv.Itoa(x) + y }
	expectedOutput := "12345678910"

	actualOutput := higherorder.Foldr(toStringAndConcat, identity, input)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestSort(t *testing.T) {
	input := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	lessThan := func(x, y int) bool { return x < y }
	expectedOutput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	actualOutput := higherorder.Sort(lessThan, input)

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, input)
}

func TestSortRandomInput(t *testing.T) {
	input := []int{1, 9, 6, 7, 6, 3, 4, 5, 10, 2, 1, 8}
	lessThan := func(x, y int) bool { return x < y }
	expectedOutput := []int{1, 1, 2, 3, 4, 5, 6, 6, 7, 8, 9, 10}

	actualOutput := higherorder.Sort(lessThan, input)

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, []int{1, 9, 6, 7, 6, 3, 4, 5, 10, 2, 1, 8}, input)
}

func TestSortCustomType(t *testing.T) {
	type MyStruct struct {
		s   string
		num int
	}

	input := []MyStruct{{"10", 10}, {"9", 9}, {"8", 8}, {"7", 7}, {"6", 6}, {"5", 5}, {"4", 4}, {"3", 3}, {"2", 2}, {"1", 1}}
	lessThan := func(x, y MyStruct) bool { return x.num < y.num }
	expectedOutput := []MyStruct{{"1", 1}, {"2", 2}, {"3", 3}, {"4", 4}, {"5", 5}, {"6", 6}, {"7", 7}, {"8", 8}, {"9", 9}, {"10", 10}}

	actualOutput := higherorder.Sort(lessThan, input)

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, []MyStruct{{"10", 10}, {"9", 9}, {"8", 8}, {"7", 7}, {"6", 6}, {"5", 5}, {"4", 4}, {"3", 3}, {"2", 2}, {"1", 1}}, input)
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

func BenchmarkReverse(b *testing.B) {
	input := largeList()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		higherorder.Reverse(input)
	}
}

func BenchmarkMap(b *testing.B) {
	input := largeList()
	double := func(x int) int { return x * 2 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		higherorder.Map(double, input)
	}
}

func BenchmarkFilter(b *testing.B) {
	input := largeList()
	isEven := func(x int) bool { return x%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		higherorder.Filter(isEven, input)
	}
}

func BenchmarkAll(b *testing.B) {
	input := largeList()
	isPositive := func(x int) bool { return x >= 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		higherorder.All(isPositive, input)
	}
}

func BenchmarkAny(b *testing.B) {
	input := largeList()
	isNegative := func(x int) bool { return x < 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		higherorder.Any(isNegative, input)
	}
}

func BenchmarkFirst(b *testing.B) {
	input := largeList()
	over500 := func(x int) bool { return x > 500 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		higherorder.First(over500, input)
	}
}

func BenchmarkFoldl(b *testing.B) {
	input := largeList()
	sum := func(x, y int) int { return x + y }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		higherorder.Foldl(sum, 0, input)
	}
}

func BenchmarkFoldr(b *testing.B) {
	input := largeList()
	sum := func(x, y int) int { return x + y }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		higherorder.Foldr(sum, 0, input)
	}
}

func BenchmarkSort(b *testing.B) {
	input := largeList()
	lessThan := func(x, y int) bool { return x < y }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		higherorder.Sort(lessThan, input)
	}
}
