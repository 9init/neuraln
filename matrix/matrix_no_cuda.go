//go:build !cuda

package matrix

import (
	"math/rand/v2"
	"neuraln/errors"
)

// AddFromMatrixGPU adds another Matrix to the current Matrix using the CPU fallback.
func (m *Matrix) AddFromMatrix(sMatrix *Matrix) (*Matrix, error) {
	if m.Col != sMatrix.Col || m.Row != sMatrix.Row {
		return nil, errors.ErrMatricesDimensionsMustMatch
	}

	result := New(m.Row, m.Col)
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Col; j++ {
			result.Matrix[i][j] = m.Matrix[i][j] + sMatrix.Matrix[i][j]
		}
	}

	return result, nil
}

// SubtractMatrix subtracts another Matrix from the current Matrix using the CPU fallback.
func (m *Matrix) SubtractMatrix(sMatrix *Matrix) (*Matrix, error) {
	if m.Col != sMatrix.Col || m.Row != sMatrix.Row {
		return nil, errors.ErrMatricesDimensionsMustMatch
	}

	result := New(m.Row, m.Col)
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Col; j++ {
			result.Matrix[i][j] = m.Matrix[i][j] - sMatrix.Matrix[i][j]
		}
	}
	return result, nil
}

// DotProductGPU performs matrix multiplication using the CPU fallback.
func (m *Matrix) DotProduct(sMatrix *Matrix) (*Matrix, error) {
	if m.Col != sMatrix.Row {
		return nil, errors.ErrRowsMustEqualColumns
	}

	result := New(m.Row, sMatrix.Col)
	for i := 0; i < m.Row; i++ {
		for j := 0; j < sMatrix.Col; j++ {
			for k := 0; k < sMatrix.Row; k++ {
				result.Matrix[i][j] += m.Matrix[i][k] * sMatrix.Matrix[k][j]
			}
		}
	}

	return result, nil
}

// HadProduct performs element-wise multiplication (Hadamard product) and returns a new Matrix.
func (m *Matrix) HadProduct(sMatrix *Matrix) (*Matrix, error) {
	if m.Row != sMatrix.Row || m.Col != sMatrix.Col {
		return nil, errors.ErrRowsColsMustEqual
	}

	result := New(m.Row, m.Col)
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Col; j++ {
			result.Matrix[i][j] = m.Matrix[i][j] * sMatrix.Matrix[i][j]
		}
	}
	return result, nil
}

// Randomize fills the Matrix with random values between -1 and 1.
func (m *Matrix) Randomize() *Matrix {
	result := New(m.Row, m.Col)
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Col; j++ {
			result.Matrix[i][j] = rand.Float64()*(1-(-1)) - 1
		}
	}
	return result
}
