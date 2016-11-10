package binarytree

type BinaryNode struct {
  hasValue bool
  value int
  leftNode BinaryNode
  rightNode BinaryNode
}

func (node *BinaryNode) Add(num int) err {
  
  if !hasValue {
    node.value = num
    node.hasValue = true
    return nil
  } 
  
  if num == node.value {
    return nil
  } 
  
  if num < node.value {
    return node.leftNode.Add(num)
  }
  
  if num > node.value {
    return node.rightNode.Add(num)
  }
}

func (node *BinaryNode) Has(num int) bool {
  
  if !hasValue {
    return false
  }
  
  if num == node.Value {
    return true
  }
  
  if num < node.value {
    return node.leftNode.Has(num)
  }
  
  if num > node.value {
    return node.rightNode.Has(num)
  }
}

func (node *BinaryNode) Size() int {
  return node.fSize(0)
}

func (node *BinaryNode) fSize(cnt *int) int {
  if !hasValue {
    return 0
  }
  
  cnt++
  
  node.leftNode.fSize(cnt)
  node.rightNode.fSize(cnt)
  
  return cnt
}

func (node *BinaryNode) TreeDepth() int {
  return node.fTreeDepth(0)
}

func (node *BinaryNode) fTreeDepth(currentDepth int) int {
  
  if !hasValue {
    return 0
  }
  
  currentDepth++
  
  leftDepth = node.leftNode.fTreeDepth(currentDepth)
  rightDepth = node.rightNode.fTreeDepth(currentDepth)
  
  return (leftDepth>rightDepth) ? leftDepth : rightDepth
}
