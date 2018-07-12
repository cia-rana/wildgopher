package mat

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

var globalRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewSymmetricMatrix(a *mat.Dense) *mat.Dense {
	r, c := a.Dims()

	result := mat.NewDense(r, c, nil)

	result.Add(a, a.T())

	for i := 0; i < r; i++ {
		result.Set(i, i, result.At(i, i)-a.At(i, i))
	}

	return result
}

func NewRandomSymmetricMatrix(size int) *mat.Dense {
	if size < 0 {
		size = 100
	}

	m := mat.NewDense(size, size, nil)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			m.Set(i, j, globalRand.Float64()*2-1)
		}
	}

	return NewSymmetricMatrix(m)
}

func QuadraticForm(vector *mat.VecDense, matrix *mat.Dense) float64 {
	tmp := mat.NewDense(0, 0, nil)
	tmp.Mul(vector.T(), matrix)
	return mat.Dot(tmp.RowView(0), vector)
}

func HamiltonianEnergy(spins *mat.VecDense, hamiltonian *mat.Dense) float64 {
	l, _ := hamiltonian.Dims()

	j := mat.NewDense(0, 0, nil)
	j.Copy(hamiltonian)
	for i := 0; i < l; i++ {
		j.Set(i, i, 0)
	}

	jSum := -QuadraticForm(spins, j) / 2
	hSum := 0
	for i := 0; i < l; i++ {
		hSum -= hamiltonian.At(i, i) * spins.AtVec(i)
	}
	return jSum + hSum
}

func QuadraticEnergy(vector *mat.VecDense, matrix *mat.Dense) float64 {
	product := QuadraticForm(vector, matrix)
	for i, _ := matrix.Dims(); i != 0; i-- {
		product += matrix.At(i, i) * vector.AtVec(i)
	}
	return 0.5 * product
}

func String(matrix mat.Matrix) string {
	return fmt.Sprintf("%.5g", mat.Formatted(matrix, mat.Squeeze()))
}
