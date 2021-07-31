# Matrix

A simple go module providing a collection of matrix class implementations.  While often used to teach upcoming distributed systems students how to program in a concurrent manner even sequential processing has its own challenges.

This simple module came about to see if simple changes to how matrixes are kept in memory have any impact on their performance - mainly when multiplying them.

---

## Matrix Class Types

This module has three implementations of a matrix.

1. **Row Aligned** - As the name implies, the matrix in memory is kept as a sequence of rows.  This makes it that a column is technically a "cross section" of a row aligned matrix.

2. **Column Aligned** - As the name implies, it is a matrix where in memory it is a sequence of columns making rows the cross sections of the matrix.

3. **Computed Identity Matrix** - An identity matrix where the value of a cell is calculated each time, rather than keeping a matrix that is mostly all zeros except for its diagnonal (identity) value in memory.

---

## Interfaces

While there are multiple matrix implementations, all implement the following two interfaces.

### Matrix

Describes all methods a matrix class (row/column aligned or computed) must implement.

- **Rows() int** - The number of rows in the matrix.

- **Columns() int** - The number of columns in the matrix.

- **Element(r int, c int) int64** - Returns the value of the given element in the matrix (effectively: M[r, c]).

- **SetElement(r int, c int, v int64)** - Assigns the provided value to the given element in the matrix (effectively M[r, c] &#8592; v)

- **Row() MatrixSegment** - Returns the given matrix row.

- **Column() MatrixSegment** - Returns the given column of the matrix.

### MatrixSegment

Describes all methods a matrix segment represents.  A matrix segment can be either a row or a column.

Originally, Row() and Column() was intended to just return a []int64 matrix, but this meant that computed matrixes would have to be create and return slices when the same methods were called on them.

Thus, MatrixSegments were created so that computed segments were possible for computed matrixes.

- **Length() int** - The number of elements in the segment.

- **Element(i int) int64** - Returns i-th, zero-based element from the segment.

- **SetElement(i int, v int)** - Assigns the value, v, to the i-th, zero-based element in the segment.

- **String() string** - Convenience method to allow the segment to be represented as a string.

---

## Classes

This section details the publically accessible matrix classes.  While there are classes for the required MatrixSegments, these are internal classes thus should not be accessed directly.

### RowAlignedMatrix

- As the name implies, the matrix is kept as a sequence of rows in memory.
- Should be used as the "B" matrix during matrix multiplications.
- Instantiated using **NewRowAlignedMatrix(rows int, columns int)**.

### ColumnAlignedMatrix

- As the name implies, the matrix is kept as a sequence of columns in memory.
- Should be used as the "A" mattrix during matrix multiplication.
- Instantiated using **NewColumnAlignMatrix(rows int, columns int)**

### ComputedIdentityMatrix

- A matrix that only tracks its dimension (it is the same on both sides) and its identity value.
- When accessing an element, the identity value is only returned when the row and column indices are the same.  Otherwise, zero is returned.
- The returned matrix segments for the row or column behave the same way - they track the index any represent and return the identity value when the element index matches.
- Can be used as either the "A" or "B" matrix in multiplications as there is no memory alignment to worry about.
- Instantiated using **NewComputedIdentityMatrix(n int, val int64)**.

---

## Initializers

For convenience, this module also provides functions that will initialize a given matrix with desired data.

### InitializeIdentityMatrix(m Matrix, v int64)

- Effectively sets the top-left to bottom-right diagonal of the provided matrix to the given value, v.
- Should **only** be used on newly instantiated matrixes.  This method only sets the diagonal and **does not** blank out (set to zero) the other matrix elements.
- Should be used when a full, non-computed identity matrix is needed.

### InitializeMatrix(m Matrix)

- Effectively initializes a matrix by populating each element (row by row) by an ever incrementing number, starting at 0.
- Maybe slow on column align matrixes due to the seeking needed to set each row element.

---

## Math Operations

At the time of writing, only multiplication has been implemented.  If more operations are implemented, they will be documented here.

### MultiplySegments(a MatrixSegment, b MatrixSegment) int64

- Calculates the (int64) product from the two provided segments.
- Made as its own function rather than part of the matrix multiplication function to ease maintenance.

### Multiply(a Matrix, b Matrix) Matrix

- Calculates the product of the provided two matrixes.
- Result is returned as a row aligned matrix.
- In future, a second version should be written to return the result in a given alignment.

---

Originally authored: July 31, 2021.