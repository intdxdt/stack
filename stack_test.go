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
		var obj = NewStack[int]().Push(array[0], array[1])

		for _, v := range array[2:] {
			obj.Push(v)
		}
		//print
		fmt.Println(obj)

		g.It("should test items in the queue", func() {
			var panics = 0
			var panicFn = func() {
				if err := recover(); err != nil {
					panics += 1
				}
			}
			g.Assert(obj.IsEmpty()).IsFalse()
			g.Assert(obj.Size()).Equal(len(array))
			g.Assert(obj.First()).Equal(3)
			g.Assert(obj.Top()).Equal(3)
			g.Assert(obj.Pop()).Equal(3)
			g.Assert(obj.Pop()).Equal(2)
			g.Assert(obj.Pop()).Equal(1)
			g.Assert(obj.Pop()).Equal(8)

			g.Assert(obj.First()).Equal(2)
			g.Assert(obj.Top()).Equal(2)
			g.Assert(obj.Last()).Equal(9)

			fmt.Println(obj)

			obj.Empty()
			g.Assert(obj.IsEmpty()).IsTrue()
			func() {
				defer panicFn()
				obj.Pop() //panic
			}()
			g.Assert(panics == 1).IsTrue()
			//print
			fmt.Println(obj)
		})
	})
}
