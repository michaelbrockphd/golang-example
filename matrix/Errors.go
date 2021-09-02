package matrix

import (
	"fmt"
)

// Error thrown by computed segments and matrixes.
type ComputedError struct {
	errorMessage string
}

func (ce *ComputedError) Error() string {
	return fmt.Sprintf("ComputedError: %v", ce.errorMessage)
}

// Create a new computed error.
func NewComputedError(message string) error {
	rtn := new(ComputedError)

	rtn.errorMessage = message

	return rtn
}

// Error thrown when trying to use a out-of-bounds index.
type MatrixSegmentIndexError struct {
	index int
}

func (msie *MatrixSegmentIndexError) Error() string {
	return fmt.Sprintf("MatrixSegmentIndexError: Index out of range, %v", msie.index)
}

// Create a new matrix segment index error.
func NewMatrixSegmentIndexError(i int) error {
	rtn := new(MatrixSegmentIndexError)

	rtn.index = i

	return rtn
}

// Error thrown when an index is out of bounds.
type MatrixIndexError struct {
	x int
	y int
}

func (mie *MatrixIndexError) Error() string {
	return fmt.Sprintf("MatrixIndexError: Index out of range, [%v, %v]", mie.x, mie.y)
}

// Create a new matrix index error.
func NewMatrixIndexError(row int, col int) error {
	rtn := new(MatrixIndexError)

	rtn.x = row
	rtn.y = col

	return rtn
}
