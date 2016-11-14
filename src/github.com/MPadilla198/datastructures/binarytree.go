package binarytree

import {
  "strconv"
  //"errors"
}

type BinaryTree struct {
  rootNode binaryNode
}

type binaryNode struct {
  value int
  leftNode, rightNode binaryNode
}

func (bTree *BinaryTree) Add(num int) {
  
  return bTree.rootNode.add(num)
}

func (node *binaryNode) add(num int) {
  
  if num == node.value {
    return
  } 
  
  if num < node.value {
    if node.leftNode == nil {
      node.leftNode = new(binaryNode)
      
      node.leftNode.value = num
      
      return
    }
      
    return node.leftNode.add(num)
  }
  
  if num > node.value {
    if node.rightNode == nil {
      node.rightNode = new(binaryNode)
      
      node.rightNode.value = num
      
      return
    }
    
    return node.rightNode.add(num)
  }
}

func (bTree *BinaryTree) Has(num int) bool {
  
  return bTree.rootNode.has(num)
}

func (node *binaryNode) has(num int) bool {
  
  if num == node.Value {
    return true
  }
  
  if num < node.value {
    if node.leftNode = nil {
      return false
    }
    
    return node.leftNode.has(num)
  }
  
  if num > node.value {
    if node.rightNode = nil {
      return false
    }
    
    return node.rightNode.has(num)
  }
}

func (bTree *BinaryTree) Size() int {
  
  return bTree.rootNode.size(0)
}

func (node *binaryNode) size(cnt *int) int {
  
  cnt++
  
  if node.leftNode != nil {
    node.leftNode.size(cnt)
  }
  
  if node.rightNode != nil {
    node.rightNode.size(cnt)
  }
  
  return cnt
}

func (bTree *BinaryTree) TreeDepth() int {
  
  return bTree.rootNode.treeDepth(0)
}

func (node *binaryNode) treeDepth(currentDepth int) int {
  
  currentDepth++
  
  if node.leftNode != nil {
    leftDepth = node.leftNode.treeDepth(currentDepth)
  }
  
  if node.rightNode != nil {
    rightDepth = node.rightNode.treeDepth(currentDepth)
  }
  
  return currentDepth + (leftDepth>rightDepth) ? leftDepth : rightDepth
}

/*********************************************
//
// Traversal Functions
//
*********************************************/

func (bTree *BinaryTree) PreorderTraversal() []int {
  
  var arr [bTree.Size()]int
  
  return bTree.rootNode.preorderTraversal(arr, 0)
}

func (node *binaryNode) preorderTraversal(path *[]int, cnt *int) []int {
  
  path[cnt] = node.value
  cnt++
  
  if node.leftNode != nil {
    node.leftNode.preorderTraversal(path, cnt)
  }
  
  if node.rightNode != nil {
    node.rightNode.preorderTraversal(path, cnt)
  }
  
  return path
}

func (bTree *BinaryTree) InorderTraversal() []int {
  
  var arr [bTree.Size()]int
  
  return bTree.rootNode.inorderTraversal(arr, 0)
}

func (node *binaryNode) inorderTraversal(path *[]int, cnt *int) []int {
  
  if node.leftNode != nil {
    node.leftNode.inorderTraversal(path)
  }
  
  path[cnt] = node.value
  cnt++
  
  if node.rightNode != nil {
    node.rightNode.inorderTraversal(path)
  }
  
  return path
}

func (bTree *BinaryTree) PostorderTraversal() []int {
  
  var arr [bTree.Size()]int
  
  return bTree.rootNode.postorderTraversal(arr, 0)
}

func (node *binaryNode) postorderTraversal(path *[]int, cnt *int) []int {
  
  if node.leftNode != nil {
    node.leftNode.postorderTraversal(path)
  }
  
  if node.rightNode != nil {
    node.rightNode.postorderTraversal(path)
  }
  
  path[cnt] = node.value
  cnt++
  
  return path
}
