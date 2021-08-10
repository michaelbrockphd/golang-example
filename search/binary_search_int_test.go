package search

import "testing"

func Test_BinarySearchInt(t *testing.T) {
	testCases := []struct {
		data     []int
		target   int
		expected int
	}{
		{
			[]int{},
			42,
			-1,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			1,
			1,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			9,
			8,
		},
	}

	result := 0

	for _, testCase := range testCases {
		result = BinarySearchInt(testCase.data, testCase.target)

		if result != testCase.expected {
			t.Errorf(
				"BinarySearchInt FAIL: expected %v, but got %v",
				testCase.expected,
				result,
			)
		}
	}
}
