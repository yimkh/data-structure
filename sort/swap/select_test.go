package main

import (
	"data-structure/sort/swap/components"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	src := []int{2, 6, 9, 4, 1}
	fmt.Println("初始元素：")
	fmt.Println(src)
	components.BubbleSort(src)
	fmt.Println("排序后元素：")
	fmt.Println(src)
}

func TestQuickSort(t *testing.T) {
	src := []int{2, 6, 9, 4, 1}
	fmt.Println("初始元素：")
	fmt.Println(src)
	components.QuickSort(src)
	fmt.Println("排序后元素：")
	fmt.Println(src)
}

func TestPartition(t *testing.T) {
	src := []int{2, 6, 9, 4, 1}
	fmt.Println("初始元素：")
	fmt.Println(src)
	pivotpos := components.Partition(src, 0, len(src)-1)
	fmt.Printf("分割元素的位置： 第%v个\n", pivotpos+1)
	fmt.Println("划分后元素：")
	fmt.Println(src)
}
