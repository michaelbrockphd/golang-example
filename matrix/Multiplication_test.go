package matrix

import "testing"

const matrix_size = (1024 * 1)

// Benchmark where both matrixes are row aligned.
func Benchmark_Multiply_RowAlignedMatrices(t *testing.B) {
	size := matrix_size

	a := NewRowAlignedMatrix(size, size)

	b := NewRowAlignedMatrix(size, size)

	Multiply(a, b)
}

// Benchmark where matrix b is column aligned.
func Benchmark_Multiply_Mixed(t *testing.B) {
	size := matrix_size

	a := NewRowAlignedMatrix(size, size)

	b := NewColumnAlignedMatrix(size, size)

	Multiply(a, b)
}

// Benchmark where matrix a is column aligned and matrix b is row aligned.
func Benchmark_Multiply_WrongAlignment(t *testing.B) {
	size := matrix_size

	a := NewColumnAlignedMatrix(size, size)

	b := NewRowAlignedMatrix(size, size)

	Multiply(a, b)
}

// Benchmark where matrix a is column aligned and matrix b is row aligned.
func Benchmark_Multiply_RowAlignedAndComputed(t *testing.B) {
	size := matrix_size

	a := NewRowAlignedMatrix(size, size)

	b := NewComputedIdentityMatrix(size, 2)

	Multiply(a, b)
}

// Benchmark where matrix a is column aligned and matrix b is row aligned.
func Benchmark_Multiply_ColumnAlignedAndComputed(t *testing.B) {
	size := matrix_size

	a := NewColumnAlignedMatrix(size, size)

	b := NewComputedIdentityMatrix(size, 2)

	Multiply(a, b)
}
