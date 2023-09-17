package memory

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/jeronimobarea/go-vm/internal/stack"
)

type Memory struct {
	memory []*big.Int
}

func NewMemory() *Memory {
	return &Memory{}
}

func (m *Memory) Store(offset, value *big.Int) error {
	zeroValue := big.NewInt(0)

	if zeroValue.Cmp(offset) == -1 || offset.Cmp(stack.MaxUint256) == 1 {
		return errors.New("invalid memory offset")
	}

	if zeroValue.Cmp(offset) == -1 || offset.Cmp(stack.MaxUint256) == 1 {
		return errors.New("invalid memory value")
	}

	m.memory[offset.Int64()] = value
	return nil
}

func (m *Memory) Load(offset *big.Int) (*big.Int, error) {
	zeroValue := big.NewInt(0)

	if zeroValue.Cmp(offset) == -1 || offset.Cmp(stack.MaxUint256) == 1 {
		return nil, errors.New("invalid memory offset")
	}

	memSize := big.NewInt(int64(len(m.memory)))
	if offset.Cmp(memSize) == 0 || offset.Cmp(memSize) > 1 {
		return big.NewInt(0), nil
	}

	return m.memory[offset.Int64()], nil
}

func (m *Memory) Print() {
	fmt.Printf("Memory:\t %v", m.memory)
}
