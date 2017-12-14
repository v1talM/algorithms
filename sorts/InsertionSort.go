package main

import (
	"algorithm/Model"
	"fmt"
)

func main() {
	fmt.Println("before Insertion-sort:", Model.Arr)
	fmt.Println("after Insertion-sort:", InsertionSort(Model.Arr))
}

func InsertionSort(arr[]int) []int {
	for i := 1; i < Model.Len; i++ {
		k := arr[i]
		j := i - 1
		for j >=0 && arr[j] > k {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = k
	}
	return arr
}
