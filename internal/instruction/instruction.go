package instruction

import "context"

type Instruction struct {
	opCode uint
	name   string
	exec   func(context.Context)
}
