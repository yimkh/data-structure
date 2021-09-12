package main

import "data-structure/sqList/components"

func main() {
	components.Init(10)
	components.PrintSqList()
	for i := 0; i < 9; i++ {
		components.Add(i)
	}
	components.Insert(5, 100)
	components.PrintLength()
	components.Delete(2)
	components.PrintLength()
	components.Delete(10)
}
