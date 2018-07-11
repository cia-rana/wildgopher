package main

import (
	"fmt"

	"github.com/cia-rana/wildgopher/annealer/simulated"
	mymat "github.com/cia-rana/wildgopher/mat"
)

func main() {
	a := simulated.NewAnnealer(nil, nil)

	m := mymat.NewRandomSymmetricMatrix(10)

	fmt.Println(mymat.String(m))
	fmt.Println(a.Anneal(m))
}
