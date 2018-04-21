package object

import (
	"fmt"

	"github.com/Zac-Garby/radon/bytecode"
)

// A Function is a piece of code which can be called from somewhere else,
// pushing a frame to the VM's frame stack. A Function is usually referred
// to as a Method if .Self != nil.
type Function struct {
	Parameters []string
	Code       bytecode.Code
	Constants  []Object
	Names      []string
	Self       *Map
}

func (f *Function) String() string {
	if f.IsMethod() {
		return fmt.Sprintf("<method (%d) %p>", len(f.Parameters), f.Self)
	} else {
		return fmt.Sprintf("<function (%d)>", len(f.Parameters))
	}
}

// Type returns the type of an Object.
func (f *Function) Type() Type {
	return FunctionType
}

// Equals checks whether or not two objects are equal to each other.
func (f *Function) Equals(other Object) bool {
	return false
}

// IsMethod checks whether on not a function is a method - i.e., a
// function is a method if Self != nil.
func (f *Function) IsMethod() bool {
	return f.Self != nil
}
