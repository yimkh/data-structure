package main

import "data-structure/sqList/components"

func main() {
	var alist components.SqLists
	alist = components.Init(10)
	// alist.PrintSqList()
	// for i := 0; i < 9; i++ {
	// 	alist.Add(i)
	// }
	// alist.Insert(5, 100)
	// alist.PrintLength()
	// alist.Delete(2)
	// alist.PrintLength()
	// alist.Delete(10)

	data := []int{0, 1, 2}
	alist.AddList(data)
	alist.PrintLength()
	alist.PrintSqList()
}
