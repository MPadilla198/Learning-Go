package main

import {
  "fmt"
  "MPadilla198/datastructures/binarytree"
}

func main() {
  testBinaryTree()
}

func testBinaryTree() {
  
  bTree := binarytree.New(42)
  bTree.Add(38)
  bTree.Add(69)
  bTree.Add(99)
  bTree.Add(78)
  bTree.Add(22)
  bTree.Add(1)
  bTree.Add(13)
  bTree.Add(47)
  bTree.Add(52)
  bTree.Add(50)
  bTree.Add(31)
  bTree.Add(82)
  bTree.Add(24)
  bTree.Add(15)
  bTree.Add(20)
  
  fmt.Printf("Tree size:" + bTree.Size())
  fmt.Printf("Tree depth:" + bTree.TreeDepth())
  
  fmt.Printf("Pre-Order Traversal:" + bTree.PreorderTraversal())
  fmt.Printf("In-Order Traversal:" + bTree.InorderTraversal())
  fmt.Printf("Post-Order Traversal:" + bTree.PostorderTraversal())
  
  fmt.Printf("Has 34?" + bTree.Has(34))
  fmt.Printf("Has 13?" + bTree.Has(13))
}
