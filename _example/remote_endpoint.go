package main

import (
	"fmt"

	"github.com/cia-rana/wildgopher/endpoint"
	mymat "github.com/cia-rana/wildgopher/mat"
	"github.com/cia-rana/wildgopher/solver"
	"github.com/cia-rana/wildgopher/solver/solution"
)

func main() {
	solution := solution.NewBasicSolution()
	solution.Qubo = mymat.NewRandomSymmetricMatrix(10)
	fmt.Println(mymat.String(solution.Qubo))

	solver := solver.NewBasicSolver()
	solver.Solution = solution

	result := solver.Solve(endpoint.NewRemoteEndpoint(""), nil)

	fmt.Println(mymat.String(result))
}
