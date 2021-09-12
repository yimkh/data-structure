package main

import (
	"data-structure/sqList/components"
	"fmt"
)

func main() {
	var alist components.SqLists
	alist = components.Init(5)
	data := []int{0, 1, 2}
	alist.AddList(data)
	alist.PrintLength()
	alist.PrintSqList()

	v := 2
	fmt.Printf("the element location of %v : %v", v, alist.LocateElem(v))
}
