package endpoint

import (
	"gonum.org/v1/gonum/mat"

	"github.com/cia-rana/wildgopher/annealer/simulated"
	"github.com/cia-rana/wildgopher/solver/solution"
)

type Endpoint interface {
	Dispatch(*solution.BasicSolution, func(*mat.VecDense)) <-chan *mat.VecDense
}

type LocalEndpoint struct {
	annealer *simulated.Annealer
}

func NewLocalEndpoint(annealer *simulated.Annealer) *LocalEndpoint {
	if annealer == nil {
		annealer = simulated.NewAnnealer(nil, nil)
	}

	return &LocalEndpoint{
		annealer: annealer,
	}
}

func (le *LocalEndpoint) Dispatch(sltn *solution.BasicSolution, callback func(*mat.VecDense)) <-chan *mat.VecDense {
	if sltn.IsingInteractions == nil {
		panic("")
	}

	if r, _ := sltn.IsingInteractions.Dims(); r == 0 {
		panic("")
	}

	result := make(chan *mat.VecDense)
	go func() {
		result <- le.anneal(sltn, callback)
	}()

	return result
}

func (le *LocalEndpoint) anneal(sltn *solution.BasicSolution, callback func(*mat.VecDense)) *mat.VecDense {
	q := le.annealer.Anneal(sltn.IsingInteractions)
	q = sltn.AdjustSolutionsFromIsingSpins(q)
	if callback != nil {
		callback(q)
	}
	return q
}

type RemoteEndpoint struct {
}

func NewRemoteEndpoint() *RemoteEndpoint {
	return &RemoteEndpoint{}
}

func (re *RemoteEndpoint) Dispatch(sltn *solution.BasicSolution, callback func(*mat.VecDense)) <-chan *mat.VecDense {
	result := make(chan *mat.VecDense)
	return result
}
