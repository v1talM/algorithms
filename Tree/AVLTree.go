package main

import (
	"algorithm/Model"
	"fmt"
)

func main()  {
	 avl := Model.NewAVLTree(50,40,60)
	 avl.InOrderWalk()
	 avl.InsertAVLNode(30)
	 avl.InsertAVLNode(45)
	 avl.InsertAVLNode(47)
	 avl.InOrderWalk()
	 fmt.Println("\n根结点:", avl.GetRoot())
}
