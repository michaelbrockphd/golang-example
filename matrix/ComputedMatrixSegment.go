package matrix

import (
	"fmt"
	"strings"
)

// Structure declaration for a fully computed matrix segment.
type computedMatrixSegment struct {
	// How long the segment is meant to be
	length int

	// The row/column index the computed segment is meant to represent.
	index int

	// The only, identity value that appears in the segment at the required index.
	value int64
}

// Return the number of elements in the row/column.
func (ms *computedMatrixSegment) Length() int {
	return ms.length
}

// Get the element in the row/column
func (ms *computedMatrixSegment) Element(i int) int64 {
	rtn := int64(0)

	if i == ms.index {
		rtn = ms.value
	}

	return rtn
}

// Set the element in the row/column
func (ms *computedMatrixSegment) SetElement(i int, v int64) {
	// Throw an error!
}

// Return a string of the row/column.
func (ms *computedMatrixSegment) String() string {
	tmp := make([]int64, ms.length)

	tmp[ms.index] = ms.value

	str := fmt.Sprint(tmp)

	elements := strings.Fields(str)

	joined := strings.Join(elements, ", ")

	return joined
}

// Create a new computed matrix segment.
func createComputedMatrixSegment(l int, i int, v int64) *computedMatrixSegment {
	rtn := new(computedMatrixSegment)

	rtn.length = l
	rtn.index = i
	rtn.value = v

	return rtn
}
