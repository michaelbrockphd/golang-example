package leet_code_attempts

import "testing"

func Test_007_reverse(t *testing.T) {
	testCases := []struct {
		src int32
		exp int32
	}{
		{123, 321},
		{-123, -321},
		{-1234, -4321},
		{120, 21},
		{0, 0},
		{2147483641, 1463847412},
		{1463847412, 2147483641},
		{2147483647, 0},
		{-2147483648, 0},
		{1534236469, 0},
	}

	for _, testCase := range testCases {
		result := reverse(testCase.src)

		if result != testCase.exp {
			t.Errorf(
				"Failed to reverse %v: expected %v but got %v",
				testCase.src,
				testCase.exp,
				result)
		}
	}
}
