package main

import (
	"data-structure/sort/select/components"
	"fmt"
)

func main() {
	src := []int{2, 6, 9, 4, 1}
	fmt.Println("初始元素：")
	fmt.Println(src)
	components.BubbleSort(src)
	fmt.Println("排序后元素：")
	fmt.Println(src)
}
