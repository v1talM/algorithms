package main

import (
	"algorithm/Model"
	"fmt"
)

func main() {
	fmt.Println("before Bubble-sort:", Model.Arr)
	fmt.Println("after Bubble-sort:", BubbleSort(Model.Arr))
}

func BubbleSort(arr[]int) []int {
	for i := 0; i < Model.Len; i++ {
		for j := i; j < Model.Len; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
