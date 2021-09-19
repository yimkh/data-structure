package main

import (
	"data-structure/sort/insert/components"
	"fmt"
)

func main() {
	src := []int{3, 5, 8, 4, 1}
	fmt.Println("初始元素")
	fmt.Println(src)
	components.ShellSort(src, 2)
	fmt.Println("排序后元素")
	fmt.Println(src)
}
