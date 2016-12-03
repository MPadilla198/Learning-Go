package collections

type linkedList struct {
  rootNode *listNode
}

func NewLinkedList(num int) linkedList {

  var list linkedList
  list.rootNode.value = num

  return list
}

type listNode struct {
  value int
  nextNode *listNode
}

func (list *linkedList) AddToTail(num int) {

  list.rootNode.addToTail(num)
}

func (node *listNode) addToTail(num int) {

  if node.nextNode == nil {
    node.nextNode = new(listNode)
    node.nextNode.value = num

    return
  }

  node.nextNode.addToTail(num)
}

func (list *linkedList) AddToHead(num int) {

  newRootNode := new(listNode)
  newRootNode.value = num

  newRootNode.nextNode = list.rootNode
  list.rootNode = newRootNode
}
