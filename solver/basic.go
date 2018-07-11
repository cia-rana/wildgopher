package solver

import (
	"reflect"

	"gonum.org/v1/gonum/mat"

	"github.com/cia-rana/wildgopher/endpoint"
	"github.com/cia-rana/wildgopher/solver/solution"
)

type BasicSolver struct {
	Solution *solution.BasicSolution
	endpoint endpoint.Endpoint
}

func NewBasicSolver() *BasicSolver {
	return &BasicSolver{
		Solution: solution.NewBasicSolution(),
		endpoint: endpoint.NewRemoteEndpoint(),
	}
}

func (b *BasicSolver) Solve(endpoint endpoint.Endpoint, callback func(*mat.VecDense)) <-chan *mat.VecDense {
	if l, _ := b.Solution.IsingInteractions.Dims(); l == 0 {
		b.InitIsingInteractions()
	}
	if l, _ := b.Solution.Qubo.Dims(); l == 0 {
		b.InitQubo()
	}
	if !reflect.ValueOf(endpoint).IsNil() {
		b.endpoint = endpoint
	}
	return b.endpoint.Dispatch(b.Solution, callback)
}

func (b *BasicSolver) InitQubo() {
	b.Solution.Qubo.Scale(-4.0, b.Solution.IsingInteractions)

	l, _ := b.Solution.Qubo.Dims()
	for i := 0; i < l; i++ {
		b.Solution.Qubo.Set(
			i, i,
			2*(mat.Sum(b.Solution.IsingInteractions.RowView(i))+mat.Sum(b.Solution.IsingInteractions.ColView(i))-6*b.Solution.IsingInteractions.At(i, i)),
		)
	}
}

func (b *BasicSolver) InitIsingInteractions() {
	b.Solution.IsingInteractions.Scale(-0.25, b.Solution.Qubo)

	l, _ := b.Solution.IsingInteractions.Dims()
	for i := 0; i < l; i++ {
		b.Solution.IsingInteractions.Set(
			i, i,
			-0.25*(mat.Sum(b.Solution.Qubo.RowView(i))+mat.Sum(b.Solution.Qubo.ColView(i))),
		)
	}
}
