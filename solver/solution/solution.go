package solution

import (
	"gonum.org/v1/gonum/mat"

	mymat "github.com/cia-rana/wildgopher/mat"
)

type Solution interface {
	AdjustSolutionsFromIsingSpins(int) int
	HamiltonianEnergy(int)
	QuboEnergy(int)
}

type BasicSolution struct {
	IsingInteractions *mat.Dense
	Qubo              *mat.Dense
}

func NewBasicSolution() *BasicSolution {
	return &BasicSolution{
		IsingInteractions: mat.NewDense(0, 0, nil),
		Qubo:              mat.NewDense(0, 0, nil),
	}
}

func (bs BasicSolution) AdjustSolutionsFromIsingSpins(solutions *mat.VecDense) *mat.VecDense {
	return solutions
}

func (bs *BasicSolution) HamiltonianEnergy(spins *mat.VecDense) float64 {
	return mymat.HamiltonianEnergy(spins, bs.IsingInteractions)
}

func (bs BasicSolution) QuboEnergy(vector *mat.VecDense) float64 {
	return mymat.QuadraticEnergy(vector, bs.Qubo)
}
