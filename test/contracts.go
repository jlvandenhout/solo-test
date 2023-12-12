package test

import (
	_ "embed"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

//go:generate solc --abi --bin --overwrite ../contracts/Fibonacci.sol -o build
var (
	//go:embed build/Fibonacci.abi
	FibonacciContractABI string
	//go:embed build/Fibonacci.bin
	fibonacciContractBytecodeHex string
	FibonacciContractBytecode    = common.FromHex(strings.TrimSpace(fibonacciContractBytecodeHex))
)
