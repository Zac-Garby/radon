package object

import (
	"fmt"
	"strings"
)

// A List is a dynamic mutable linked list.
type List struct {
	defaults
	Value []Object
}

func (l *List) String() string {
	var strs []string

	for _, item := range l.Value {
		strs = append(strs, item.String())
	}

	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}

func (l *List) Type() Type {
	return ListType
}

func (l *List) Equals(other Object) bool {
	switch o := other.(type) {
	case *List:
		left, right := l.Value, o.Value
		if len(left) != len(right) {
			return false
		}

		for i, item := range left {
			if !item.Equals(right[i]) {
				return false
			}
		}

		return true

	default:
		return false
	}
}

func (l *List) Infix(op string, right Object) (Object, bool) {
	if op == "+" {
		other, ok := right.(*List)
		if !ok {
			return nil, false
		}

		return &List{Value: append(l.Value, other.Value...)}, true
	} else if op == "[]" {
		num, ok := right.(*Number)
		if !ok {
			return nil, false
		}

		i := int(num.Value)
		if i < 0 || i >= len(l.Value) {
			return nil, false
		}
		return l.Value[i], true
	} else {
		return nil, false
	}
}

func (l *List) Items() ([]Object, bool) {
	return l.Value, true
}