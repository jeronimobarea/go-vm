package execution

import (
	"math/big"

	"github.com/jeronimobarea/go-vm/internal/memory"
	"github.com/jeronimobarea/go-vm/internal/stack"
)

type Execution struct {
	code           [8]uint
	stack          stack.Stack
	memory         memory.Memory
	programCounter uint
	stopped        bool
	output         *big.Int
}
