package main

import (
	"algorithm/Model"
	"fmt"
)
/**
 * @description: Binary Search Tree & AVL Tree Exercise
 * @Note: BST sort & QuickSort make the same comparisons,
 *        but in a different order.
 */
func main() {
	bst_root := Model.NewBinaryTreeNode(0)
	bst_root.CreateBSTTree(6,5,4,3,2,1)
	bst_tree := Model.NewBinaryTree(bst_root)
	bst_tree.InOrderWalk()
	fmt.Println("\nBST树的根结点为：",bst_root)
	//创建AVL树
	avl_root := Model.NewBinaryTreeNode(0)
	avl_root = avl_root.CreateAVLTree(1,2,3,4,5,6)
	avl_tree := Model.NewBinaryTree(avl_root)
	avl_tree.InOrderWalk()
	fmt.Println("\nAVL树的根结点为:", avl_root)
}
