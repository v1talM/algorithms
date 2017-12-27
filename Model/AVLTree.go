package Model

import "fmt"

type AVLTree struct {
	root *AVLTreeNode
}

func NewAVLTree(data... interface{}) *AVLTree {
	root := NewAVLTreeNode(data[0])
	for i := 1; i < len(data); i++ {
		root = insertAVLNode(root, data[i])
	}
	return &AVLTree{root: root}
}

func (this *AVLTree) GetRoot() *AVLTreeNode {
	return this.root
}

func (this *AVLTree) InOrderWalk()  {
	fmt.Println("中序遍历结果：")
	avlInOrderWalk(this.root)
}

func (this *AVLTree) InsertAVLNode(data interface{}) {
	fmt.Println("\n插入新结点:", data)
	this.root = insertAVLNode(this.root, data)
}

func avlInOrderWalk(root *AVLTreeNode)  {
	if root != nil {
		avlInOrderWalk(root.GetLChild())
		fmt.Printf("%v ", root.GetData())
		avlInOrderWalk(root.GetRChild())
	}
}

func insertAVLNode(node *AVLTreeNode, data interface{}) *AVLTreeNode {
	if node == nil {
		node = NewAVLTreeNode(data)
	} else if data.(int) < node.GetData().(int) {
		node.SetLChild(insertAVLNode(node.GetLChild(), data))
		//插入之后失衡
		if height(node.GetLChild()) - height(node.GetRChild()) == 2 {
			if data.(int) < node.GetLChild().GetData().(int) {
				//右单旋
				node = rightRotation(node)
			} else {
				//先左旋再右旋
			}
		}
	} else {
		node.rChild = insertAVLNode(node.GetRChild(), data)
		//插入之后失衡
		if height(node.GetRChild()) - height(node.GetLChild()) == 2 {
			if data.(int) > node.GetRChild().GetData().(int) {
				//左单旋
				node = leftRotation(node)
			} else {
				//先右旋再左旋
			}
		}
	}
	node.height = max(height(node.GetLChild()), height(node.GetRChild())) + 1
	return node
}

func height(node *AVLTreeNode) int {
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

func leftRotation(node *AVLTreeNode) *AVLTreeNode {
	rchild := node.GetRChild() 				//待旋转的右子结点
	node.SetRChild(rchild.GetLChild())		//待旋转结点的左子结点设置到当前结点的右子结点
	rchild.SetLChild(node)					//当前结点转移到待旋转结点的左子结点下
	return rchild
}

func rightRotation(node *AVLTreeNode) *AVLTreeNode {
	lchild := node.GetLChild()
	node.SetLChild(lchild.GetRChild())
	lchild.SetRChild(node)
	return lchild
}