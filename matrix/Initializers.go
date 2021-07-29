package matrix

// Take the provided matrix and convert/initialize it to be an indentity matrix
//
// Note: this function assumes the  matrix is now and has no data.  It does not
// zero out any existing elements.  It only overwrites where r == c.
func InitializeIdentityMatrix(m Matrix, v int64) {
	numRows := m.Rows()

	i := 0

	for i < numRows {
		m.SetElement(i, i, v)

		i += 1
	}
}

// Take the provided matrix and set each element from left to right and top to
// bottom with an incremented number starting at 0
func InitializeMatrix(m Matrix) {
	numRows := m.Rows()

	numCols := m.Columns()

	v := int64(0)

	r := 0

	c := 0

	for r < numRows {
		c = 0

		for c < numCols {
			m.SetElement(r, c, v)

			v += 1
			c += 1
		}

		r += 1
	}
}
