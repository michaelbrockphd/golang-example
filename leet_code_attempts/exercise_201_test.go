package leet_code_attempts

import "testing"

func Test_201_rangeBitwiseAnd(t *testing.T) {
	testCases := []struct {
		left     int
		right    int
		expected int
	}{
		{5, 7, 4},
		{0, 0, 0},
		{1, 2147483647, 0},
		{30000, 2147483643, 0},
	}

	for _, testCase := range testCases {
		result := rangeBitwiseAnd(testCase.left, testCase.right)

		if result != testCase.expected {
			t.Errorf(
				"Search failed: expected %v but got %v",
				testCase.expected,
				result)
		}
	}
}
