package collections

type Queue struct {
	list *linkedList
}

func (queue *Queue) Clear() {

	queue.list = new(linkedList)
}

func (queue *Queue) Has(val int) bool {

	return queue.list.Has(val)
}

func (queue *Queue) IndexOf(val int) int {

	return queue.list.IndexOf(val)
}

func (queue *Queue) IsEmpty() bool {

	return queue.list.IsEmpty()
}

func (queue *Queue) Peek() int {

	return queue.list.rootNode.value
}

func (queue *Queue) Pop() int {

	val := queue.list.Peek()
	queue.list.Remove(val)

	return val
}

func (queue *Queue) Push(val int) {

	queue.list.AddToTail(val)
}
