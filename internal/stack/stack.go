package stack

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

var (
	uint256    = int64(math.Pow(2, 256))
	MaxUint256 = big.NewInt(uint256)
)

type Stack struct {
	maxDepth uint
	stack    []*big.Int
}

func NewStack(maxDepth uint) *Stack {
	return &Stack{
		maxDepth: maxDepth,
		stack:    make([]*big.Int, maxDepth),
	}
}

func (s *Stack) Push(v *big.Int) error {
	zeroValue := big.NewInt(0)

	if zeroValue.Cmp(v) == -1 || v.Cmp(MaxUint256) == 1 {
		return errors.New("stack overflow")
	}

	s.stack = append(s.stack, v)
	return nil
}

func (s *Stack) Pop() (*big.Int, error) {
	if s.empty() {
		return nil, errors.New("stack underflow")
	}
	element := s.top()
	idx := s.size() - 1
	s.stack = s.stack[:idx]
	return element, nil
}

func (s *Stack) empty() bool {
	if s.stack == nil {
		return true
	}
	return s.size() == 0
}

func (s Stack) size() int {
	return len(s.stack)
}

func (s Stack) top() *big.Int {
	size := s.size()
	if size == 0 {
		return nil
	}
	return s.stack[size-1]
}

func (s *Stack) Swap(idxA, idxB uint) error {
	a, err := s.getAtIndex(idxA)
	if err != nil {
		return err
	}

	b, err := s.getAtIndex(idxB)
	if err != nil {
		return err
	}

	s.setAtIndex(idxA, b)
	s.setAtIndex(idxB, a)
	return nil
}

func (s Stack) getAtIndex(idx uint) (*big.Int, error) {
	if s.empty() {
		return nil, errors.New("index out of bounds")
	}
	return s.stack[idx-1], nil
}

func (s *Stack) setAtIndex(idx uint, v *big.Int) {
	s.stack[idx-1] = v
}

func (s Stack) Print() {
	fmt.Printf("stack: \t %v", s.stack)
}
