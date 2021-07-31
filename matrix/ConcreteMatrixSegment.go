package matrix

import (
	"fmt"
	"strings"
)

// Small wrapper for holding an array.
//
// This is more for convenience and to avoid accidentally confusing a matrix
// row with a matrix column when performing math operations.
type concreteMatrixSegment struct {
	elements []int64
}

// Return the number of elements in the row/column.
func (ms *concreteMatrixSegment) Length() int {
	return len(ms.elements)
}

// Get the element in the row/column
func (ms *concreteMatrixSegment) Element(i int) int64 {
	return ms.elements[i]
}

// Set the element in the row/column
func (ms *concreteMatrixSegment) SetElement(i int, v int64) {
	ms.elements[i] = v
}

// Return a string of the row/column.
func (ms *concreteMatrixSegment) String() string {
	str := fmt.Sprint(ms.elements)

	elements := strings.Fields(str)

	joined := strings.Join(elements, ", ")

	return joined
}

// Create a new matrix part with existing elements
func createSegmentWithElements(elements []int64) *concreteMatrixSegment {
	mp := new(concreteMatrixSegment)

	mp.elements = elements

	return mp
}
