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

func (root *BinaryTreeNode) CreateAVLTree(data ...interface{}) *BinaryTreeNode {
	root.SetData(data[0])
	for i := 1; i < len(data); i++ {
		root = InsertAVLNode(root, data[i])
	}
	return root
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
	node.height = max(height(node.GetLChild()), height(node.GetRChild())) + 1
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

func height(node *BinaryTreeNode) int {
	if node == nil {
		return 0
	}
	lh, rh := 0, 0
	lChild := node.GetLChild()
	rChild := node.GetRChild()
	for lChild != nil {
		lh++
		lChild = lChild.GetLChild()
	}
	for rChild != nil{
		rh++
		rChild = rChild.GetRChild()
	}
	if lh > rh {
		return lh
	}
	return rh
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func InsertAVLNode(node *BinaryTreeNode, data interface{}) *BinaryTreeNode {
	if node == nil {
		node = NewBinaryTreeNode(data)
	} else if data.(int) < node.GetData().(int) {
		node.lChild = InsertAVLNode(node.GetLChild(), data)
		//插入后失衡
		if height(node.GetLChild()) - height(node.GetRChild()) == 1 {
			if data.(int) < node.GetLChild().GetData().(int) {
				//右旋
				node = rightRotation(node)
			} else {
				//先左旋、再右旋
				node = leftRightRotation(node)
			}
		}
	} else {
		node.rChild = InsertAVLNode(node.GetRChild(), data)
		//插入后失衡
		if height(node.GetRChild()) - height(node.GetLChild()) == 1 {
			if data.(int) > node.GetRChild().GetData().(int) {
				//左旋
				node = leftRotation(node)
			} else {
				//先右旋、再左旋
				node = rightLeftRotation(node)
			}
		}
	}
	node.height = max(height(node.GetLChild()), height(node.GetRChild())) + 1
	return node
}

func leftRotation(node *BinaryTreeNode) *BinaryTreeNode {
	prChild := node.GetRChild()
	node.rChild = prChild.GetLChild()
	if prChild.HasLChild() {
		prChild.GetLChild().CutOffParent()
		prChild.GetLChild().SetParent(node)
	}
	prChild.SetLChild(node)
	node.SetParent(prChild)
	prChild.CutOffParent()
	return prChild
}

func rightRotation(node *BinaryTreeNode) *BinaryTreeNode {
	plChild := node.GetLChild()
	node.lChild = plChild.GetRChild()
	if plChild.HasRChild() {
		plChild.GetRChild().CutOffParent()
		plChild.GetRChild().SetParent(node)
	}
	plChild.SetRChild(node)
	node.SetParent(plChild)
	plChild.CutOffParent()
	return plChild
}

func leftRightRotation(node *BinaryTreeNode) *BinaryTreeNode {

	return node
}

func rightLeftRotation(node *BinaryTreeNode) *BinaryTreeNode {
	return node
}