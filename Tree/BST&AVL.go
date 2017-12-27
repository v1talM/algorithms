package main

import (
	"fmt"
	"algorithm/Model"
)

func main() {
	avl := Model.NewAVLTree(1,2,3,4,5,6,7)
	bst_root := Model.NewBinaryTreeNode(0)
	bst_root.CreateBSTTree(1,2,3,4,5,6,7)
	bst_tree := Model.NewBinaryTree(bst_root)
	fmt.Println("------------------BST树------------------")
	bst_tree.InOrderWalk()
	fmt.Println("\n------------------AVL树------------------")
	avl.InOrderWalk()
	fmt.Println("\n-----------------------------------------")
	fmt.Println("\nBST树的根结点为:",bst_root.GetData())
	fmt.Println("\nAVL树的根结点为:", avl.GetRoot().GetData())
	fmt.Println("\n---------------BST树打印-----------------")
	Model.PrintBinaryTree(bst_root, 1)
	fmt.Println("\n---------------AVL树打印-----------------")
	Model.PrintAVLTree(avl.GetRoot(), 1)
}