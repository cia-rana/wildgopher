package solver

import (
	"gonum.org/v1/gonum/mat"

	"github.com/cia-rana/wildgopher/endpoint"
)

type Solver interface {
	Solve(endpoint endpoint.Endpoint) *mat.VecDense
	InitQubo()
	InitIsingInteractions()
}
