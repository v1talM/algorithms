package main

import (
	"algorithm/Model"
	"fmt"
)

func main()  {
	 avl := Model.NewAVLTree(50,40,60,30,45,42,23,33,41,13,7)
	 Model.PrintAVLTree(avl.GetRoot(), 1)

	 fmt.Println("\n根结点:", avl.GetRoot())
}
