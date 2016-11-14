package linkedlist

type LinkedList struct {
  rootNode listNode
}

func New(num int) LinkedList {
  
  list := new(LinkedList)
  list.rootNode.value = num
  
  return list
}

type listNode struct {
  value int
  nextNode listNode
}

func (list *linkedList) AddToTail(num int) {
  
  return list.rootNode.addToTail(num)
}

func (node *listNode) addToTail(num int) {
  
  if node.nextNode == nil {
    node.nextNode := new(listNode)
    node.nextNode.value = num
    
    return
  }
  
  return node.nextNode.addToTail(num)
}

func (list *LinkedList) AddToHead(num int) {
  
  newRootNode := new(listNode)
  newRootNode.value = num
  
  newRootNode.nextNode = list.rootNode
  list.rootNode = newRootNode
}
