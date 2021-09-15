package main

import (
	"data-structure/sqList/components"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var alist components.SqLists

	fmt.Println("初始顺序表为:")
	alist = components.Init(5)
	for i := 0; i < 4; i++ {
		components.Add(&alist, i)
	}
	components.PrintSqList(alist)
	components.PrintLength(alist)

	fmt.Println("在第四个位置插入值为100的元素:")
	components.Insert(&alist, 3, 100)
	components.PrintSqList(alist)
	components.PrintLength(alist)

	fmt.Println("删除第3个位置上的元素:")
	components.Delete(&alist, 2)
	components.PrintSqList(alist)
	components.PrintLength(alist)

	fmt.Println("删除第6个位置上的元素:")
	components.Delete(&alist, 5)
	components.PrintSqList(alist)
	components.PrintLength(alist)
	fmt.Println("--------------------")
}

func TestAddList(t *testing.T) {
	var alist components.SqLists

	fmt.Println("初始顺序表为:")
	alist = components.Init(5)
	data := []int{0, 1, 2}
	components.AddList(&alist, data)
	components.PrintSqList(alist)
	components.PrintLength(alist)
}

func TestGetElem(t *testing.T) {
	var alist components.SqLists
	alist = components.Init(5)
	data := []int{0, 1, 2}
	components.AddList(&alist, data)
	components.PrintLength(alist)
	components.PrintSqList(alist)

	i := 2
	fmt.Printf("第%v个位置上的元素的值为: %v\n", i, components.GetElem(alist, i))
}

func TestLocateElem(t *testing.T) {
	var alist components.SqLists
	alist = components.Init(5)
	data := []int{0, 1, 2}
	components.AddList(&alist, data)
	components.PrintSqList(alist)
	components.PrintLength(alist)

	v := 2
	fmt.Printf("值为%v的元素位于: 第%v个\n", v, components.LocateElem(alist, v))
}
