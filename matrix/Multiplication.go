package matrix

func MultiplyLinear(a Matrix, b Matrix) Matrix {
	var rtn Matrix = nil

	return rtn
}

func (a *matrixSubset) Multiply(b *matrixSubset) int64 {
	result := int64(0)

	l := a.Length()

	for i := 0; i < l; i++ {
		result += (a.Element(i) * b.Element(i))
	}

	return result
}

/*func (a *Matrix) Multiply(b *Matrix) *Matrix {
	result := NewMatrix(a.Rows(), b.Columns())

	numRows := result.Rows()
	numCols := result.Columns()

	for i := 0; i < numRows; i++ {
		aRow := a.Row(i)

		for j := 0; j < numCols; j++ {
			bCol := b.Column(j)

			v := aRow.Multiply(bCol)

			result.SetElement(i, j, v)
		}
	}

	return result
}*/

func (a *Matrix) Multiply(b *Matrix) *Matrix {
	result := NewMatrix(a.Rows(), b.Columns())

	numRows := result.Rows()
	numCols := result.Columns()

	len := a.Columns()

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			v := int64(0)

			for k := 0; k < len; k++ {
				v += a.Element(i, k) * b.Element(k, j)
			}

			result.SetElement(i, j, v)
		}
	}

	return result
}
