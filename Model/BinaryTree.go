package Model

import "fmt"

type BinaryTree struct {
	root *BinaryTreeNode
}

func NewBinaryTree(root *BinaryTreeNode) *BinaryTree {
	return &BinaryTree{root: root}
}

func (root *BinaryTreeNode) CreateBSTTree(data ...interface{}) {
	root.SetData(data[0])
	for i := 1; i < len(data); i++ {
		insertNode(root, data[i])
	}
}

func (tree *BinaryTree) InOrderWalk() {
	fmt.Println("中序遍历结果:")
	inOrderWalk(tree.root)
}

func TreeSearchRecursive(node *BinaryTreeNode, k interface{}) *BinaryTreeNode {
	if node == nil || k.(int) == node.GetData().(int) {
		return node
	}
	if k.(int) < node.GetData().(int) {
		return TreeSearchRecursive(node.GetLChild(), k)
	} else {
		return TreeSearchRecursive(node.GetRChild(), k)
	}
}

func TreeSearchIteration(node *BinaryTreeNode, k interface{}) *BinaryTreeNode {
	for node != nil && k.(int) != node.GetData().(int) {
		if k.(int) < node.GetData().(int) {
			node = node.GetLChild()
		} else {
			node = node.GetRChild()
		}
	}
	return node
}

func inOrderWalk(node *BinaryTreeNode)  {
	if node != nil {
		inOrderWalk(node.GetLChild())
		fmt.Printf("%v ", node.GetData())
		inOrderWalk(node.GetRChild())
	}
}

func insertNode(node *BinaryTreeNode, data interface{})  {
		if data.(int) < node.GetData().(int) {
			if node.HasLChild() {
				insertNode(node.GetLChild(), data)
			} else {
				child := NewBinaryTreeNode(data)
				child.SetParent(node)
				node.SetLChild(child)
			}
		} else {
			if node.HasRChild() {
				insertNode(node.GetRChild(), data)
			} else {
				child := NewBinaryTreeNode(data)
				child.SetParent(node)
				node.SetRChild(child)
			}
		}
}

func (tree *BinaryTree) Minimum() *BinaryTreeNode {
	node := tree.root
	for node.GetLChild() != nil {
		node = node.GetLChild()
	}
	return node
}

func (tree *BinaryTree) Maxmum() *BinaryTreeNode {
	node := tree.root
	for node.GetRChild() != nil {
		node = node.GetRChild()
	}
	return node
}