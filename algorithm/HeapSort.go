package main

import "fmt"

func parentNode(i int)  int{
	return (i - 1)/2
}

//左节点
func leftNode(i int) int{
	return 2*i + 1
}
//右节点
func rightNode(i int) int{
	return 2*i + 2
}
//创建heap
func buildHeap(heap []int) {
	length := len(heap)
	for i := length/2 - 1; i >= 0; i-- {
		maxHeap(heap, i, length)
	}
}

func maxHeap(heap []int, i int, length int) {
	left := leftNode(i)
	right := rightNode(i)
	largest := 0
	if left < length && heap[left] > heap[i] {
		largest = left
	}else {
		largest = i
	}
	if right < length && heap[right] > heap[largest] {
		largest = right
	}
	if largest != i {
		heap[i], heap[largest] = heap[largest], heap[i]
		//需要继续比较其父节点
		maxHeap(heap, largest, length)
	}
}

func main() {
	a := []int{1, 24, 35, 343, 463, 46, 34, 35, 12, 123, 245, 413, 5, 132}
	buildHeap(a)
	fmt.Println(a)
}
