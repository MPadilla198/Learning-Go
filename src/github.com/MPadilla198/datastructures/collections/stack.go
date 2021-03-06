package collections

type Stack struct {
	list *linkedList
}

func (stack *Stack) Clear() {

	stack.list = new(linkedList)
}

func (stack *Stack) IsEmpty() bool {

	return stack.list.IsEmpty()
}

func (stack *Stack) Has(val int) bool {

	return stack.list.Has(val)
}

func (stack *Stack) IndexOf(val int) int {

	return stack.list.IndexOf(val)
}

func (stack *Stack) Peek() int {

	return stack.list.Peek()
}

func (stack *Stack) Pop() int {

	val := stack.list.Peek()
	stack.list.Remove(val)

	return val
}

func (stack *Stack) Push(val int) {

	stack.list.AddToHead(val)
}
