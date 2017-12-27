package main

import (
	"algorithm/Model"
	"fmt"
)

func main()  {
	 avl := Model.NewAVLTree(50,40,60)
	 avl.InsertAVLNode(55)
	 avl.InsertAVLNode(70)
	 avl.InsertAVLNode(53)
	 avl.InOrderWalk()
	 fmt.Println("\n根结点:", avl.GetRoot())
}
