package main

import (
	"algorithm/Model"
	"fmt"
)

func main()  {
	 avl := Model.NewAVLTree(6,5,4,3,2,1)
	 avl.InOrderWalk()
	 avl.InsertAVLNode(7)
	 avl.InOrderWalk()
	 fmt.Println("\n根结点:", avl.GetRoot())
}
