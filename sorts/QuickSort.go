package main

import "fmt"

func main() {
	arr := []int{17, 9, 14, 10, 21, 5, 10}
	fmt.Println("before Quick-sort:", arr)
	QuickSort(arr, 0, len(arr) - 1)
	fmt.Println("after Quick-sort:", arr)
}

func QuickSort(arr[]int, left, right int) {
	if left < right {
		key := arr[left]
		low, high := left, right
		for low < high {
			for low < high && arr[high] > key {
				high--
			}
			arr[low] = arr[high]
			for low < high && arr[low] < key {
				low++
			}
			arr[high] = arr[low]
		}
		arr[low] = key
		QuickSort(arr, left, low - 1)
		QuickSort(arr, low + 1, right)
	}
}
