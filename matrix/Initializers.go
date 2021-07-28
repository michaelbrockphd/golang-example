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
	}
}
