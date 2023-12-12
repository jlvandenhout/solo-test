package test

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
)

func TestFibonacci(t *testing.T) {
	env := solo.New(t)
	chain := env.NewChain()
	chain.DeployEVMContract(nil, FibonacciContractABI, FibonacciContractBytecode)
}
