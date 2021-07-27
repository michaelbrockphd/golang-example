package conversion

// Offset needed to reach the number region on the ASCII chart.
const _offset_ascii_to_numbers = 48

// Further offset from ASCII numbers to reach the uppercase letters.
const _offset_numbers_to_letters = 7

// Interface describing all methods a decimal-to-base converter must implement.
type DecimalToBaseConverter interface {
	Convert(src int) string
}

// Private implementation of decimal-to-base converter.
type decimalToBaseConverter struct {
	destinationBase int
}

// When given a source integer, return a string in the required base.
func (dtob *decimalToBaseConverter) Convert(src int) string {
	rtn := ""

	quotient := src

	for 0 < quotient {
		remainder := (quotient % dtob.destinationBase)

		if 9 < remainder {
			remainder += _offset_numbers_to_letters
		}

		c := rune(remainder + _offset_ascii_to_numbers)

		rtn = (string(c) + rtn)

		quotient = (quotient / dtob.destinationBase)
	}

	return (rtn)
}

// Create a new decimal-to-base converter using the specified base.
func NewDecimalToBaseConverter(b int) DecimalToBaseConverter {
	rtn := new(decimalToBaseConverter)

	rtn.destinationBase = b

	return rtn
}
