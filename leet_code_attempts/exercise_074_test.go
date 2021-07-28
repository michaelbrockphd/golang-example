package leet_code_attempts

import "testing"

// Constraints
// 1) 0 <= s.length <= 200
// 2) s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.

func Test_074_searchMatrix(t *testing.T) {
	testCases := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
		{[][]int{{1}}, 1, true},
	}

	for _, testCase := range testCases {
		result := searchMatrix(testCase.matrix, testCase.target)

		if result != testCase.expected {
			t.Errorf(
				"Search failed: expected %v but got %v",
				testCase.expected,
				result)
		}
	}
}
