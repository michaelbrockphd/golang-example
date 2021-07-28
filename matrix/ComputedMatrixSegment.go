package matrix

import (
	"fmt"
	"strings"
)

type computedMatrixSegment struct {
	length int
	index  int
	value  int64
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

func createComputedMatrixSegment(l int, i int, v int64) *computedMatrixSegment {
	rtn := new(computedMatrixSegment)

	rtn.length = l
	rtn.index = i
	rtn.value = v

	return rtn
}
