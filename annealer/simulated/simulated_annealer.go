package simulated

import (
	"math/rand"
	"reflect"
	"time"

	"gonum.org/v1/gonum/mat"

	mymat "github.com/cia-rana/wildgopher/mat"
)

type Annealer struct {
	temperatureSchedule *TemperatureSchedule
	strategy            Strategy
	dim                 int
	j                   *mat.Dense
	h                   *mat.VecDense
	q                   *mat.VecDense

	rand *rand.Rand
}

func NewAnnealer(sch *TemperatureSchedule, str Strategy) *Annealer {
	if sch == nil {
		sch = NewTemperatureSchedule(-1, -1, -1)
	}

	if str == nil || reflect.ValueOf(str).IsNil() {
		str = NewSingleSpinFlipStrategy(-1, nil)
	}

	return &Annealer{
		temperatureSchedule: sch,
		strategy:            str,
		rand:                rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (a *Annealer) Anneal(hamiltonian *mat.Dense) *mat.VecDense {
	a.initialAnnealing(hamiltonian)

	for t := a.temperatureSchedule.Next(); t > 0; t = a.temperatureSchedule.Next() {
		a.strategy.Update(a, t)
	}

	return a.q
}

func (a *Annealer) initialAnnealing(hamiltonian *mat.Dense) {
	a.dim, _ = hamiltonian.Dims()
	a.j = mymat.NewSymmetricMatrix(hamiltonian)
	a.h = mat.NewVecDense(a.dim, nil)
	a.q = mat.NewVecDense(a.dim, nil)

	for i := 0; i < a.dim; i++ {
		a.j.Set(i, i, 0)
		a.h.SetVec(i, hamiltonian.At(i, i))

		if a.rand.Float64() > 0.5 {
			a.q.SetVec(i, 1)
		} else {
			a.q.SetVec(i, -1)
		}
	}
}
