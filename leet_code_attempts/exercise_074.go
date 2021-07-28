package leet_code_attempts

// Detect if there is a matching element by literally starting in the bottom
// right element and working towards the top left.
func searchMatrix(matrix [][]int, target int) bool {
	found := false

	cLen := len(matrix[0])

	r := len(matrix)
	cMax := cLen - 1
	c := int(0)

	// Break out if we run out of rows or a match is found.
	for 0 < r && !found {
		r -= 1

		if matrix[r][0] <= target {
			// Reset c so we start at the end.
			c = cMax

			// Break out if we:
			// 1) Run out of elements
			// 2) Have found a match
			// 3) The current element is lower than the target.
			for 0 <= c && !found && !(matrix[r][c] < target) {
				found = matrix[r][c] == target

				c -= 1
			}
		}
	}

	return found
}
