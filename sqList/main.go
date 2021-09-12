package main

import "data-structure/sqList/components"

func main() {
	var alist components.SqLists
	alist = components.Init(5)
	data := []int{0, 1, 2}
	alist.AddList(data)
	alist.PrintLength()
	alist.PrintSqList()
}
