package leet_code_attempts

import "testing"

// Constraints
// 1) 0 <= s.length <= 200
// 2) s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.

func Test_008_myAtoi(t *testing.T) {
	testCases := []struct {
		src string
		exp int32
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
	}

	for _, testCase := range testCases {
		result := myAtoi(testCase.src)

		if result != testCase.exp {
			t.Errorf(
				"Failed to convert %v: expected %v but got %v",
				testCase.src,
				testCase.exp,
				result)
		}
	}
}
