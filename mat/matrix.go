package mat

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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
	hSum := 0.0
	for i := 0; i < l; i++ {
		hSum -= hamiltonian.At(i, i) * spins.AtVec(i)
	}
	return jSum + hSum
}

func QuadraticEnergy(vector *mat.VecDense, matrix *mat.Dense) float64 {
	product := QuadraticForm(vector, matrix)
	{
		l, _ := matrix.Dims()
		for i := 0; i < l; i++ {
			product += matrix.At(i, i) * vector.AtVec(i)
		}
	}
	return 0.5 * product
}

func ApplyPrecisionMatrix(matrix *mat.Dense, precisionFormat string) *mat.Dense {
	f := func(i, j int, v float64) float64 {
		w, err := strconv.ParseFloat(fmt.Sprintf(precisionFormat, v), 64)
		if err != nil {
			return 0
		}
		return w
	}
	r, c := matrix.Dims()
	result := mat.NewDense(r, c, nil)
	result.Apply(f, matrix)
	return result
}

func BuildMatrixForParams(matrix *mat.Dense, strip bool) [][]float64 {
	l, _ := matrix.Dims()
	list := make([][]float64, l)
	for i := 0; i < l; i++ {
		list[i] = matrix.RawRowView(i)
		if strip {
			list[i] = list[i][i:]
		}
	}
	return list
}

func BuildVectorFromVectorString(vectorString string) []float64 {
	vectorString = strings.TrimSpace(vectorString)
	if vectorString[0] != '[' || vectorString[len(vectorString)-1] != ']' {
		return []float64{}
	}

	vectorStringElements := strings.Split(vectorString[1:len(vectorString)-1], ",")
	result := make([]float64, len(vectorStringElements))
	for i, e := range vectorStringElements {
		f, err := strconv.ParseFloat(strings.TrimSpace(e), 64)
		if err != nil {
			return []float64{}
		}
		result[i] = f
	}

	return result
}

func String(matrix mat.Matrix) string {
	return fmt.Sprintf("%.5g", mat.Formatted(matrix, mat.Squeeze()))
}
