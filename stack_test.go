package stack

import (
	"fmt"
	"github.com/franela/goblin"
	"testing"
)

func TestStack(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Stack", func() {
		var array = []int{9, 5, 3, 2, 8, 1, 2, 3}
		var obj = NewStack().Push(array[0], array[1])

		for _, v := range array[2:] {
			obj.Push(v)
		}
		//print
		fmt.Println(obj)

		g.It("should test items in the queue", func() {
			g.Assert(obj.IsEmpty()).IsFalse()
			g.Assert(obj.Size()).Equal(len(array))
			g.Assert(obj.First().(int)).Equal(3)
			g.Assert(obj.Top().(int)).Equal(3)
			g.Assert(obj.Pop().(int)).Equal(3)
			g.Assert(obj.Pop().(int)).Equal(2)
			g.Assert(obj.Pop().(int)).Equal(1)
			g.Assert(obj.Pop().(int)).Equal(8)

			g.Assert(obj.First().(int)).Equal(2)
			g.Assert(obj.Top().(int)).Equal(2)
			g.Assert(obj.Last().(int)).Equal(9)

			fmt.Println(obj)

			obj.Empty()
			g.Assert(obj.Pop()).Equal(nil)
			g.Assert(obj.IsEmpty()).IsTrue()
			//print
			fmt.Println(obj)
		})
	})
}
