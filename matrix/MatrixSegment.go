package matrix

import (
	"fmt"
	"strings"
)

// Small wrapper for holding an array.
//
// This is more for convenience and to avoid accidentally confusing a matrix
// row with a matrix column when performing math operations.
type matrixSegment struct {
	elements []int64
}

// Create a new matrix part with existing elements
func createSegmentWithElements(elements []int64) *matrixSegment {
	mp := new(matrixSegment)

	mp.elements = elements

	return mp
}

// Create a new matrix part instance.
func createSegment(n int) *matrixSegment {
	elements := make([]int64, n)

	return createSegmentWithElements(elements)
}

// Return the number of elements in the row/column.
func (ms *matrixSegment) Length() int {
	return len(ms.elements)
}

// Get the element in the row/column
func (ms *matrixSegment) Element(i int) int64 {
	return ms.elements[i]
}

// Set the element in the row/column
func (ms *matrixSegment) SetElement(i int, v int64) {
	ms.elements[i] = v
}

// Return a string of the row/column.
func (ms *matrixSegment) String() string {
	str := fmt.Sprint(ms.elements)

	elements := strings.Fields(str)

	joined := strings.Join(elements, ", ")

	return joined
}
