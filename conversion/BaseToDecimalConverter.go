package conversion

import "math"

// Index where lowercase letters start on the ASCII chart.
const _index_letters_lower_case = 97

// Index where uppercase letters start on the ASCII chart.
const _index_letters_uppers_case = 65

// Offset so ints from lowercase letters are shifted towards 10.
const _offset_letters_lower_case = 87

// Offset so ints from uppercase letters are shifted towards 10.
const _offset_letters_upper_case = 55

// Offset so ints from ASCII numbers are shifted towards 0
const _offset_numbers = 48

// Interface describing all methods a base-to-decimal converter must implement.
type BaseToDecimalConverter interface {
	Convert(src string) int
}

// Private, simple implementation of a base-to-decimal converter.
type baseToDecimalConverter struct {
	sourceBase int
}

// Convert the provided string back into a Integer value.
func (btod *baseToDecimalConverter) Convert(src string) int {
	rtn := 0

	srcLen := len(src) - 1

	i := 0

	for i <= srcLen {
		currVal := int(src[i])

		// Adjust where needed.

		if _index_letters_lower_case <= currVal {
			currVal -= _offset_letters_lower_case

		} else if _index_letters_uppers_case <= currVal {
			currVal -= _offset_letters_upper_case

		} else {
			currVal -= _offset_numbers
		}

		// Calculate the exponent and power.

		exponent := (srcLen - i)

		power := math.Pow(float64(btod.sourceBase), float64(exponent))

		// Increment the return value and indexer.

		rtn += (currVal * int(power))

		i += 1
	}

	return (rtn)

}

// Create a new base-to-decimal converter for the provided source base.
func NewBaseToDecimalConverter(b int) BaseToDecimalConverter {
	rtn := new(baseToDecimalConverter)

	rtn.sourceBase = b

	return (rtn)
}
