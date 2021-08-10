package search

// BinarySearchInt(arr []int, target int) int
//
// Uses the ol' binary search algorithm to find a match in a sorted array.
//
// Returns the index of the match or -1 if otherwise.
func BinarySearchInt(arr []int, target int) int {
	rtnIndex := -1

	minSearch := 0
	maxSearch := len(arr) - 1

	idxMiddle := 0
	curValue := 0

	// Loop through until maxSearch becomes less than minSearch, or if rtnIndex
	// is greater than -1.
	//
	// While this can be done recursively, a loop is better as it will not
	// consume the function stack if searching on an incredibly large array!
	for rtnIndex < 0 && minSearch <= maxSearch {
		idxMiddle = ((minSearch + maxSearch) / 2)

		curValue = arr[idxMiddle]

		if curValue < target {
			// If the current value is lower than the target, the match might be in
			// the right side of the array.
			minSearch = (idxMiddle + 1)

		} else if target < curValue {
			// If the current value is greater than the target, the match might be
			// in the left side of the array.
			maxSearch = (idxMiddle - 1)

		} else {
			// If we get here, were are an exact match.
			rtnIndex = idxMiddle
		}
	}

	// Return the index value.
	return rtnIndex
}
