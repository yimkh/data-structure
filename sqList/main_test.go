package main

import (
	"data-structure/sqList/components"
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var alist components.SqLists

	fmt.Println("-----init-----")
	alist = components.Init(5)
	alist.PrintSqList()
	for i := 0; i < 4; i++ {
		alist.Add(i)
	}
	alist.PrintSqList()
	alist.PrintLength()

	fmt.Println("-----insert-----")
	alist.Insert(3, 100)
	alist.PrintSqList()
	alist.PrintLength()

	fmt.Println("-----delete1-----")
	alist.Delete(2)
	alist.PrintSqList()
	alist.PrintLength()

	fmt.Println("-----delete2-----")
	alist.Delete(5)
	alist.PrintSqList()
	alist.PrintLength()
}

func TestAddList(t *testing.T) {
	var alist components.SqLists

	fmt.Println("-----init-----")
	alist = components.Init(5)
	data := []int{0, 1, 2}
	alist.AddList(data)
	alist.PrintSqList()
	alist.PrintLength()
}

func TestGetElem(t *testing.T) {
	var alist components.SqLists
	alist = components.Init(5)
	data := []int{0, 1, 2}
	alist.AddList(data)
	alist.PrintLength()
	alist.PrintSqList()

	fmt.Println("-----getElem-----")
	i := 2
	fmt.Printf("the element of %v : %v", i, alist.GetElem(i))
}
