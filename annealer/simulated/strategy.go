package simulated

type Strategy interface {
	Update(*Annealer, float64)
}
