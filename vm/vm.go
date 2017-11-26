package vm

import (
	"github.com/Zac-Garby/lang/bytecode"
	"github.com/Zac-Garby/lang/object"
	"github.com/Zac-Garby/lang/store"
)

// A VM interprets and executes bytecode.
type VM struct {
	frames    []*Frame
	frame     *Frame
	returnVal object.Object
	err       error
}

// New creates a new virtual machine.
func New() *VM {
	return &VM{
		frames:    make([]*Frame, 0),
		frame:     nil,
		returnVal: nil,
		err:       nil,
	}
}

// Run executes some bytecode in a new frame.
func (vm *VM) Run(code bytecode.Code, s *store.Store, constants []object.Object) {
	frame := vm.makeFrame(code, store.New(), s, constants)

	vm.runFrame(frame)
}

// RunDefault executes some bytecode in a new frame
// with empty globals and locals.
func (vm *VM) RunDefault(code bytecode.Code, constants []object.Object) {
	vm.Run(code, store.New(), constants)
}

func (vm *VM) makeFrame(code bytecode.Code, args, s *store.Store, constants []object.Object) *Frame {
	frame := &Frame{
		code:      code,
		store:     s,
		offset:    0,
		stack:     newStack(),
		constants: constants,
		vm:        vm,
	}

	for k, v := range args.Data {
		frame.store.Set(k, v.Value, true)
	}

	return frame
}

func (vm *VM) pushFrame(frame *Frame) {
	vm.frames = append(vm.frames, frame)
}

func (vm *VM) popFrame() *Frame {
	f := vm.frames[len(vm.frames)-1]
	vm.frames = vm.frames[:len(vm.frames)-1]
	return f
}

func (vm *VM) runFrame(frame *Frame) {
	vm.pushFrame(frame)
	vm.popFrame().execute()
}

// ExtractValue returns the top value from the top frame
func (vm *VM) ExtractValue() object.Object {
	if len(vm.frames) < 1 || len(vm.frames[0].stack.objects) < 1 {
		return nil
	}

	return vm.frames[0].stack.top()
}
