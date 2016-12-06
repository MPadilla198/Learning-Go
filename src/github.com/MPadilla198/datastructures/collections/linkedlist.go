package collections

type linkedList struct {
  rootNode *listNode
}

func NewLinkedList(num int) linkedList {

  return &linkedList {
    rootNode: &listNode {
      value: num,
    },
  }
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
	  
	  node.nextNode = &listNode {
		  value: num,
	  }

    return
  }

  node.nextNode.addToTail(num)
}

func (list *linkedList) AddToHead(num int) {
	
	list.rootNode = &listNode {
		value: num,
		nextNode: list.rootNode,
	}
}

func (list *linkedList) AddAfter(num int) {
	
	list.rootNode.addAfter(num)
}

func (node *listNode) addAfter(num int) {
	
	if node.value != num {
		
		return node.nextNode.addAfter(num)
	}
	
	if node.nextNode == nil {
		
		node.nextNode = &listNode {
			value: num,
		}
		
		return
	}
	
	node.nextNode = &listNode {
		value: num,
		nextNode: node.nextNode,
	}
}

func (list *linkedList) Remove(num int) bool {
	
	wasRemoved, _ := list.rootNode.remove(num)
	
	return wasRemoved
}

func (node *listNode) remove(num int) bool, bool {
	
	if node.value == num {
		
		return true, true
	}
	
	wasRemoved, wasLast := node.nextNode.remove(num)
	
	if wasLast {
		
		node.nextNode = &listNode {
			value: num,
			nextNode: node.nextNode.nextNode,
		}
	}
	
	return wasRemoved, false
}
