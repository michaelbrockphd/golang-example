package conversion

import "testing"

func Test_DecimalToBaseConverter_Convert(t *testing.T) {
	testCases := []struct {
		base     int
		input    int
		expected string
	}{
		{2, 2, "10"},
		{2, 42, "101010"},
		{16, 8, "8"},
		{16, 10, "A"},
		{16, 15, "F"},
		{16, 16, "10"},
		{16, 42, "2A"},
	}

	for _, testCase := range testCases {
		subject := NewDecimalToBaseConverter(testCase.base)

		result := subject.Convert(testCase.input)

		if result != testCase.expected {
			t.Errorf(
				"Base %v conversion of %v failed: expected %s but got %s",
				testCase.base,
				testCase.input,
				testCase.expected,
				result)
		}
	}
}
