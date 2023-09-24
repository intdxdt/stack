package stack

import (
	"github.com/intdxdt/list"
)

// Stack data structure - LIFO
type Stack[T any] struct {
	stack *list.List[T]
}

// NewStack creates a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{stack: list.NewList[T]()}
}

// Size of the stack
func (o *Stack[T]) Size() int {
	return o.stack.Len()
}

// Push item to the stack
func (o *Stack[T]) Push(item ...T) *Stack[T] {
	for _, itm := range item {
		o.stack.Append(itm)
	}
	return o
}

// Pop item from the stack
func (o *Stack[T]) Pop() T {
	return o.stack.Pop()
}

// First item in front
func (o *Stack[T]) First() T {
	return o.stack.Last()
}

// Top at the item on stack
func (o *Stack[T]) Top() T {
	return o.First()
}

// Last item at the end of the stack
func (o *Stack[T]) Last() T {
	return o.stack.First()
}

// Empty the stack
func (o *Stack[T]) Empty() *Stack[T] {
	o.stack.Clear()
	return o
}

// IsEmpty checks if stack is empty.
func (o *Stack[T]) IsEmpty() bool {
	return o.stack.Len() == 0
}

// Stringer interface
func (o *Stack[T]) String() string {
	return o.stack.String()
}
