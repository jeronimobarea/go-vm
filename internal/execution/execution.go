package execution

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/trie"

	"github.com/jeronimobarea/go-vm/internal/memory"
	"github.com/jeronimobarea/go-vm/internal/stack"
)

type ExecutionContext struct {
	code            [8]uint
	stack           *stack.Stack
	memory          *memory.Memory
	programCounter  uint
	stopped         bool
	output          *big.Int
	storage         *trie.Trie
	originalStorage *trie.Trie
	gas             *big.Int
}

func NewExecutionContext(
	code string,
	storage *trie.Trie,
	gas *big.Int,
) (*ExecutionContext, error) {
	decodedCode, err := hexutil.DecodeBig(code)
	if err != nil {
		return nil, err
	}
	return &ExecutionContext{
		code:            cod,
		stack:           stack.NewStack(1024),
		memory:          memory.NewMemory(),
		programCounter:  0,
		stopped:         false,
		storage:         storage,
		originalStorage: storage.Copy(),
		gas:             gas,
	}
}
