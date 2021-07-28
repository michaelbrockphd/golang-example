package matrix

import (
	"fmt"
	"strings"
)

// Small wrapper for holding an array.
//
// This is more for convenience and to avoid accidentally confusing a matrix
// row with a matrix column when performing math operations.
type matrixPart struct {
	elements []int64
}

// Create a new matrix part with existing elements
func createMatrixPartWithElements(elements []int64) *matrixPart {
	mp := new(matrixPart)

	mp.elements = elements

	return mp
}

// Create a new matrix part instance.
func createMatrixPart(n int) *matrixPart {
	elements := make([]int64, n)

	return createMatrixPartWithElements(elements)
}

// Return the number of elements in the row/column.
func (mp *matrixPart) Length() int {
	return len(mp.elements)
}

// Get the element in the row/column
func (mp *matrixPart) Element(i int) int64 {
	return mp.elements[i]
}

// Set the element in the row/column
func (mp *matrixPart) SetElement(i int, v int64) {
	mp.elements[i] = v
}

// Return a string of the row/column.
func (mp *matrixPart) String() string {
	str := fmt.Sprint(mp.elements)

	elements := strings.Fields(str)

	joined := strings.Join(elements, ", ")

	return joined
}

// Alias so matrix rows are not confused for matrix columns.
type MatrixRow = matrixPart

// Create a new matrix row.
func NewMatrixRow(n int) *MatrixRow {
	return createMatrixPart(n)
}

// Create a new matrix row with existing elements.
func NewMatrixRowWithElements(elements []int64) *MatrixRow {
	return createMatrixPartWithElements(elements)
}

// Alias so matrix columns are not confused for matrix rows.
type MatrixColumn = matrixPart

// Create a new matrix column.
func NewMatrixColumn(n int) *MatrixColumn {
	return createMatrixPart(n)
}

// Create a new matrix column with existing elements.
func NewMatrixColumnWithElements(elements []int64) *MatrixColumn {
	return createMatrixPartWithElements(elements)
}
