package main

import (
	"fmt"
	"strconv"

	"github.com/MPadilla198/datastructures/collections"
)

func main() {
	testBinaryTree()
}

func testBinaryTree() {

	bTree := collections.NewBinaryTree(42)
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

	fmt.Println("Tree size:" + strconv.Itoa(bTree.Size()))
	fmt.Println("Tree depth:" + strconv.Itoa(bTree.TreeDepth()))

	fmt.Print("Pre-Order Traversal: ")
	for _, num := range bTree.PreorderTraversal() {
		fmt.Print(strconv.Itoa(num) + " ")
	}
	fmt.Println(" ")

	fmt.Print("In-Order Traversal: ")
	for _, num := range bTree.InorderTraversal() {
		fmt.Print(strconv.Itoa(num) + " ")
	}
	fmt.Println(" ")

	fmt.Print("Post-Order Traversal: ")
	for _, num := range bTree.PostorderTraversal() {
		fmt.Print(strconv.Itoa(num) + " ")
	}
	fmt.Println(" ")

	fmt.Println("Has 34? " + strconv.FormatBool(bTree.Has(34)))
	fmt.Println("Has 13? " + strconv.FormatBool(bTree.Has(13)))

	fmt.Println("Remove 1")
	var wasRemoved = bTree.Remove(1)

	if wasRemoved {
		fmt.Print("Pre-Order Traversal: ")
		for _, num := range bTree.PreorderTraversal() {
			fmt.Print(strconv.Itoa(num) + " ")
		}
		fmt.Println(" ")

		fmt.Print("In-Order Traversal: ")
		for _, num := range bTree.InorderTraversal() {
			fmt.Print(strconv.Itoa(num) + " ")
		}
		fmt.Println(" ")

		fmt.Print("Post-Order Traversal: ")
		for _, num := range bTree.PostorderTraversal() {
			fmt.Print(strconv.Itoa(num) + " ")
		}
		fmt.Println(" ")
	}
}
