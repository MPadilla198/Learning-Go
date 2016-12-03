package collections

type binaryTree struct {
	rootNode binaryNode
}

func NewBinaryTree(val int) binaryTree {

	var bTree binaryTree
	bTree.rootNode.value = val

	return bTree
}

type binaryNode struct {
	value               int
	leftNode, rightNode *binaryNode
}

func (bTree *binaryTree) Add(num int) {

	bTree.rootNode.add(num)
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

		node.leftNode.add(num)
	}

	if num > node.value {
		if node.rightNode == nil {
			node.rightNode = new(binaryNode)

			node.rightNode.value = num

			return
		}

		node.rightNode.add(num)
	}
}

func (bTree *binaryTree) Remove(val int) bool {

	return bTree.rootNode.remove(val)
}

func (node *binaryNode) remove(val int) bool {

	if node.value != val {

		if node.leftNode != nil {

			if node.rightNode != nil {

				return node.leftNode.remove(val) || node.rightNode.remove(val)
			}

			return node.leftNode.remove(val)
		}

		if node.rightNode != nil {

			return node.rightNode.remove(val)
		}

		return false
	}

	if node.leftNode == nil && node.rightNode == nil {

		node = new(binaryNode)
	} else if node.leftNode == nil {

		node = node.rightNode
	} else if node.rightNode == nil {

		node = node.leftNode
	} else {

		node.value = node.rightNode.findLowestNode()
	}

	return true
}

func (node *binaryNode) findLowestNode() int {

	if node.leftNode == nil && node.rightNode == nil {

		var val = node.value

		node = new(binaryNode)

		return val
	}

	if node.leftNode != nil {

		return node.leftNode.findLowestNode()
	}

	return node.rightNode.findLowestNode()
}

func (bTree *binaryTree) Has(num int) bool {

	return bTree.rootNode.has(num)
}

func (node *binaryNode) has(num int) bool {

	if num == node.value {
		return true
	}

	if num < node.value {
		if node.leftNode == nil {
			return false
		}

		return node.leftNode.has(num)
	} else {
		if node.rightNode == nil {
			return false
		}

		return node.rightNode.has(num)
	}
}

func (bTree *binaryTree) Size() int {

	return bTree.rootNode.size(0)
}

func (node *binaryNode) size(cnt int) int {

	cnt++

	if node.leftNode != nil {
		cnt = node.leftNode.size(cnt)
	}

	if node.rightNode != nil {
		cnt = node.rightNode.size(cnt)
	}

	return cnt
}

func (bTree *binaryTree) TreeDepth() int {

	return bTree.rootNode.treeDepth()
}

func (node *binaryNode) treeDepth() int {

	var leftDepth, rightDepth int

	if node.leftNode != nil {
		leftDepth = node.leftNode.treeDepth()
	}

	if node.rightNode != nil {
		rightDepth = node.rightNode.treeDepth()
	}

	if leftDepth < rightDepth {
		return rightDepth + 1
	}

	return leftDepth + 1
}

/*********************************************
//
// Traversal Functions
//
*********************************************/

func (bTree *binaryTree) PreorderTraversal() []int {

	arr := make([]int, bTree.Size())
	arr, _ = bTree.rootNode.preorderTraversal(arr, 0)

	return arr
}

func (node *binaryNode) preorderTraversal(path []int, cnt int) ([]int, int) {

	path[cnt] = node.value
	cnt++

	if node.leftNode != nil {
		path, cnt = node.leftNode.preorderTraversal(path, cnt)
	}

	if node.rightNode != nil {
		path, cnt = node.rightNode.preorderTraversal(path, cnt)
	}

	return path, cnt
}

func (bTree *binaryTree) InorderTraversal() []int {

	var arr = make([]int, bTree.Size())
	arr, _ = bTree.rootNode.inorderTraversal(arr, 0)

	return arr
}

func (node *binaryNode) inorderTraversal(path []int, cnt int) ([]int, int) {

	if node.leftNode != nil {
		path, cnt = node.leftNode.inorderTraversal(path, cnt)
	}

	path[cnt] = node.value
	cnt++

	if node.rightNode != nil {
		path, cnt = node.rightNode.inorderTraversal(path, cnt)
	}

	return path, cnt
}

func (bTree *binaryTree) PostorderTraversal() []int {

	var arr = make([]int, bTree.Size())
	arr, _ = bTree.rootNode.postorderTraversal(arr, 0)

	return arr
}

func (node *binaryNode) postorderTraversal(path []int, cnt int) ([]int, int) {

	if node.leftNode != nil {
		path, cnt = node.leftNode.postorderTraversal(path, cnt)
	}

	if node.rightNode != nil {
		path, cnt = node.rightNode.postorderTraversal(path, cnt)
	}

	path[cnt] = node.value
	cnt++

	return path, cnt
}

/*
func (bTree *binaryTree) LevelorderTraversal() []int {

  var arr = make([]int, bTree.Size())
  arr = bTree.rootNode.levelorderTraversal()

  return arr
}

func (node *binaryNode) levelorderTraversal() []int {


}*/
