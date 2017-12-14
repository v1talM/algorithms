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
				node.SetLChild(NewBinaryTreeNode(data))
			}
		} else {
			if node.HasRChild() {
				insertNode(node.GetRChild(), data)
			} else {
				node.SetRChild(NewBinaryTreeNode(data))
			}
		}
}