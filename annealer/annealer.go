package annealer

import (
	"gonum.org/v1/gonum/mat"
)

type Annealer interface {
	Anneal(*mat.Dense) *mat.Dense
}
