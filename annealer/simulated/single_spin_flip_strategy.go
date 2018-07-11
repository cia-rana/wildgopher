package simulated

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type SingleSpinFlipStrategy struct {
	repetition     int
	updateCallback func(*mat.VecDense)
}

func NewSingleSpinFlipStrategy(repetition int, updateCallback func(*mat.VecDense)) *SingleSpinFlipStrategy {
	if repetition < 0 {
		repetition = 2000
	}

	return &SingleSpinFlipStrategy{
		repetition:     repetition,
		updateCallback: updateCallback,
	}
}

func (s *SingleSpinFlipStrategy) Update(annealer *Annealer, t float64) {
	for i := 0; i < s.repetition; i++ {
		k := annealer.rand.Intn(annealer.dim)
		dE := 2.0 * (annealer.h.AtVec(k) + mat.Dot(annealer.j.RowView(k), annealer.q)) * annealer.q.AtVec(k)
		if dE < 0 || math.Exp(-dE/t) > annealer.rand.Float64() {
			annealer.q.SetVec(k, -annealer.q.AtVec(k))
			if s.updateCallback != nil {
				s.updateCallback(annealer.q)
			}
		}
	}
}
