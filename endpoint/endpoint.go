package endpoint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gonum.org/v1/gonum/mat"

	"github.com/cia-rana/wildgopher/annealer/simulated"
	mymat "github.com/cia-rana/wildgopher/mat"
	"github.com/cia-rana/wildgopher/solver/solution"
)

type Endpoint interface {
	Dispatch(*solution.BasicSolution) *mat.VecDense
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

func (le *LocalEndpoint) Dispatch(sltn *solution.BasicSolution) *mat.VecDense {
	if sltn.IsingInteractions == nil {
		panic("")
	}

	if r, _ := sltn.IsingInteractions.Dims(); r == 0 {
		panic("")
	}

	return le.anneal(sltn)
}

func (le *LocalEndpoint) anneal(sltn *solution.BasicSolution) *mat.VecDense {
	q := le.annealer.Anneal(sltn.IsingInteractions)
	return sltn.AdjustSolutionsFromIsingSpins(q)
}

type RemoteEndpoint struct {
	requestPrecisionFormat string
	isingSolverPath        string
}

const (
	defaultIsingSolverPath = "http://api.mdrft.com/apiv1/ising"
)

func NewRemoteEndpoint(isingSolverPath string) *RemoteEndpoint {
	if isingSolverPath == "" {
		isingSolverPath = defaultIsingSolverPath
	}

	return &RemoteEndpoint{
		requestPrecisionFormat: "%.3g",
		isingSolverPath:        isingSolverPath,
	}
}

func (re *RemoteEndpoint) Dispatch(sltn *solution.BasicSolution) *mat.VecDense {
	if r, _ := sltn.IsingInteractions.Dims(); r == 0 {
		panic("")
	}

	buff := bytes.NewBuffer(nil)
	{
		params := map[string][][]float64{
			"hami": mymat.BuildMatrixForParams(
				mymat.ApplyPrecisionMatrix(
					sltn.IsingInteractions,
					re.requestPrecisionFormat,
				),
				false,
			),
		}
		encoder := json.NewEncoder(buff)
		encoder.Encode(params)
	}
	resp, err := http.Post(re.isingSolverPath, "application/json", buff)
	if err != nil {
		// TODO
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
		return mat.NewVecDense(0, nil)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO
		return nil
	}

	v := mymat.BuildVectorFromVectorString(string(respBytes))
	return mat.NewVecDense(len(v), v)
}
