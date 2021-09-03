package conversion

import (
	"fmt"
	"testing"
)

func Test_BaseToDecimalConverter_Convert(t *testing.T) {
	testCases := []struct {
		base     int
		input    string
		expected int
	}{
		{2, "10", 2},
		{2, "101010", 42},
		{16, "8", 8},
		{16, "A", 10},
		{16, "F", 15},
		{16, "10", 16},
		{16, "2A", 42},
	}

	var subTestName string = ""

	for _, testCase := range testCases {
		subTestName = fmt.Sprintf("%v:%v", testCase.base, testCase.input)

		subTestFunc := func(st *testing.T) {
			subject := NewBaseToDecimalConverter(testCase.base)

			result := subject.Convert(testCase.input)

			if result != testCase.expected {
				st.Errorf(
					"Base %v conversion of %s failed: expected %v but got %v",
					testCase.base,
					testCase.input,
					testCase.expected,
					result)
			}
		}

		t.Run(subTestName, subTestFunc)
	}
}
