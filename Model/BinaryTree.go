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

func (tree *BinaryTree) Search(key interface{}) *BinaryTreeNode {
	return searchTree(tree.root, key)
}

func (tree *BinaryTree) InsertNode(key interface{}) bool {
	exist := tree.Search(key)
	if exist != nil {
		return false
	}
	return insertNode(tree.root, key)
}

func (tree *BinaryTree) DeleteNode(key interface{}) bool {
	exist := tree.Search(key)
	if exist == nil {
		return false
	}
	return deleteNode(exist)
}

func inOrderWalk(node *BinaryTreeNode)  {
	if node != nil {
		inOrderWalk(node.GetLChild())
		fmt.Printf("%v ", node.GetData())
		inOrderWalk(node.GetRChild())
	}
}

func insertNode(node *BinaryTreeNode, data interface{}) bool {
	if data.(int) < node.GetData().(int) {
		if node.HasLChild() {
			insertNode(node.GetLChild(), data)
		} else {
			newNode := NewBinaryTreeNode(data)
			newNode.SetParent(node)
			node.SetLChild(newNode)
		}
	} else {
		if node.HasRChild() {
			insertNode(node.GetRChild(), data)
		} else {
			newNode := NewBinaryTreeNode(data)
			newNode.SetParent(node)
			node.SetRChild(newNode)
		}
	}
	return true
}

func searchTree(node *BinaryTreeNode, key interface{}) *BinaryTreeNode {
	for node != nil && key.(int) != node.GetData().(int) {
		if key.(int) < node.GetData().(int) {
			node = node.GetLChild()
			return searchTree(node, key)
		} else {
			node = node.GetRChild()
			return searchTree(node, key)
		}
	}
	return node
}

func deleteNode(node *BinaryTreeNode) bool {
	if node.HasLChild() && node.HasRChild() {
		//寻找当前节点的后继元素
		iNode := node.GetRChild()
		for iNode.GetLChild() != nil {
			iNode = iNode.GetLChild()
		}
		node.SetData(iNode.GetData())
		return deleteNode(iNode)
	} else if node.HasLChild() {
		parent := node.GetParent()
		iNode := node.GetLChild()
		iNode.SetParent(parent)
		parent.SetLChild(iNode)
		node.CutOffParent()
	} else if node.HasRChild() {
		parent := node.GetParent()
		iNode := node.GetRChild()
		iNode.SetParent(parent)
		parent.SetRChild(iNode)
		node.CutOffParent()
	} else {
		node.CutOffParent()
	}
	return true
}


