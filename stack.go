package stack

import (
	"github.com/intdxdt/list"
)

//Stack data structure - LIFO
type Stack struct {
	stack *list.List
}

//Create a new stack
func NewStack() *Stack {
	return &Stack{stack: list.NewList()}
}

//Size of the stack
func (o *Stack) Size() int {
	return o.stack.Len()
}

//Push item to the stack
func (o *Stack) Push(item ...interface{}) *Stack {
	for _, itm := range item {
		o.stack.Append(itm)
	}
	return o
}

//Pop item from the stack
func (o *Stack) Pop() interface{} {
	return o.stack.Pop()
}

//First item in front
func (o *Stack) First() interface{} {
	return o.stack.Last()
}

//Top at the item on stack
func (o *Stack) Top() interface{} {
	return o.First()
}

//Last item at the end of the stack
func (o *Stack) Last() interface{} {
	return o.stack.First()
}

//Empty the stack
func (o *Stack) Empty() *Stack {
	o.stack.Clear()
	return o
}

//IsEmpty checks if stack is empty.
func (o *Stack) IsEmpty() bool {
	return o.stack.Len() == 0
}

//Stringer interface
func (o *Stack) String() string {
	return o.stack.String()
}
