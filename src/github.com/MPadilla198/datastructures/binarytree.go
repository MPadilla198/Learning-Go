package binarytree

import "strconv"

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
    return cnt
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
    return currentDepth
  }
  
  currentDepth++
  
  leftDepth = node.leftNode.fTreeDepth(currentDepth)
  rightDepth = node.rightNode.fTreeDepth(currentDepth)
  
  return (leftDepth>rightDepth) ? leftDepth : rightDepth
}

func (node *BinaryNode) PreorderTraversal() string {
  return node.fPreorderTraversal("")
}

func (node *BinaryNode) fPreorderTraversal(path *string) string {
  
  if !node.hasValue {
    return path
  }
  
  path = path + strconv.FormatInt(int64(node.value), 10)
  
  node.leftNode.fPreorderTraversal(path)
  node.rightNode.fPreorderTraversal(path)
  
  return path
}

func (node *BinaryNode) InorderTraversal() string {
  return node.fInorderTraversal("")
}

func (node *BinaryNode) fInorderTraversal(path *string) string {
  
  if !node.hasValue {
    return path
  }
  
  node.leftNode.fInorderTraversal(path)
  
  path = path + strconv.FormatInt(int64(node.value), 10)
  
  node.rightNode.fInorderTraversal(path)
  
  return path
}

func (node *BinaryNode) PostorderTraversal() string {
  return node.fPostorderTraversal("")
}

func (node *BinaryNode) fPostorderTraversal(path *string) string {
  
  if !node.hasValue {
    return path
  }
  
  node.leftNode.fPostorderTraversal(path)
  node.rightNode.fPostorderTraversal(path)
  
  path = path + strconvInt(int64(node.value), 10)
  
  return path
}
