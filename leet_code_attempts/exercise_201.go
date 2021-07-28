package leet_code_attempts

func rangeBitwiseAnd(left int, right int) int {
	tally := left

	i := left

	for i < right {
		i += 1

		tally &= i
	}

	return tally
}
