package main

import (
	"data-structure/sqList/components"
	"fmt"
)

func main() {
	var alist components.SqLists
	alist = components.Init(5)
	data := []int{0, 1, 2}
	components.AddList(&alist, data)
	components.PrintLength(alist)
	components.PrintSqList(alist)

	v := 2
	fmt.Printf("the element location of %v : %v", v, components.LocateElem(alist, v))
}
