package higherorder_test

import (
	"strconv"
	"testing"

	"github.com/churmd/higherorder"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	double := func(x int) int { return x * 2 }
	expectedOutput := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	actualOutput := higherorder.Map(input, double)

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestMapTypeChange(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toString := func(x int) string { return strconv.Itoa(x) }
	expectedOutput := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

	actualOutput := higherorder.Map(input, toString)

	assert.Equal(t, expectedOutput, actualOutput)
}
