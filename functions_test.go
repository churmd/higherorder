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
	f := func(x int) string { return strconv.Itoa(x) }
	g := func(y string) bool { return y == "10" }

	actualOutput := higherorder.Compose(10, f, g)

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
		expectedOutput := expectedOutputs[i]
		actualOutput := higherorder.Reverse(input)

		assert.Equal(t, expectedOutput, actualOutput, "failed reversing list: %v\nExpected:%v\nActual  :%v", input, expectedOutput, actualOutput)
	}

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

func TestFoldl(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	identity := 0
	sum := func(x, y int) int { return x + y }
	expectedOutput := 55

	actualOutput := higherorder.Foldl(sum, identity, input)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestFoldlNonAssociativeFunc(t *testing.T) {
	input := []int{4, 2, 4}
	identity := 64
	div := func(x, y int) int { return x / y }
	expectedOutput := 2

	actualOutput := higherorder.Foldl(div, identity, input)

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
	input := []int{8, 12, 24, 4}
	identity := 2
	div := func(x, y int) int { return x / y }
	expectedOutput := 8

	actualOutput := higherorder.Foldr(div, identity, input)

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
		higherorder.Map(input, double)
	}
}

func BenchmarkFilter(b *testing.B) {
	input := largeList()
	isEven := func(x int) bool { return x%2 == 0 }

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		higherorder.Filter(input, isEven)
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
