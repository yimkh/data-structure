package main

import (
	"data-structure/sort/insert/components"
	"fmt"
	"testing"
)

func TestDirectInsertSort(t *testing.T) {
	src := []int{3, 5, 8, 4, 1}
	fmt.Print("初始元素: ")
	fmt.Println(src)
	components.DirectInsertSort(src)
	fmt.Print("排序后元素: ")
	fmt.Println(src)
}

func TestHarfInsertSort(t *testing.T) {
	src := []int{3, 5, 8, 4, 1}
	fmt.Println("初始元素")
	fmt.Println(src)
	components.HarfInsertSort(src)
	fmt.Println("排序后元素")
	fmt.Println(src)
}

func TestShellSort(t *testing.T) {
	src := []int{3, 5, 8, 4, 1}
	fmt.Println("初始元素")
	fmt.Println(src)
	components.ShellSort(src, 2)
	fmt.Println("排序后元素")
	fmt.Println(src)
}
