package main

import (
	"fmt"

	mymat "github.com/cia-rana/wildgopher/mat"

	"gonum.org/v1/gonum/mat"
)

func main() {
	m := mat.NewDense(2, 2, []float64{
		1, 2,
		3, 4,
	})
	v := mat.NewVecDense(2, []float64{11, 12})
	fmt.Println(mymat.QuadraticForm(v, m))
}
