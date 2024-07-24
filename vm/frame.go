package vm

import "monkey-lang/object"

type Frame struct {
	fn *object.CompiledFunction
	ip int
	basePointer int
}

func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	f := &Frame{fn: fn, ip: -1, basePointer: basePointer}
	return f
}
