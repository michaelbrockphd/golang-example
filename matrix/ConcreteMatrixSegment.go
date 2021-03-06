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
func (ms *concreteMatrixSegment) Element(i int) (int64, error) {
	rtn := int64(0)
	var err error = nil

	if i < 0 || len(ms.elements) <= i {
		err = NewMatrixSegmentIndexError(i)

	} else {
		rtn = ms.elements[i]
	}

	return rtn, err
}

// Set the element in the row/column
func (ms *concreteMatrixSegment) SetElement(i int, v int64) error {
	var err error = nil

	if i < 0 || len(ms.elements) <= i {
		err = NewMatrixSegmentIndexError(i)

	} else {
		ms.elements[i] = v
	}

	return err
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
