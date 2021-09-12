package main

import (
	"data-structure/linkList/components"
	"fmt"
)

func main() {
	headlnode := components.Init()
	headlnode.IsEmpty()
	fmt.Println("第一次插入在第一个位置")
	headlnode.Insert(1, &components.LNode{
		Data:      2,
		NextlNode: nil,
	})
	fmt.Println("第一次头插")
	a := components.LNode{
		Data:      1,
		NextlNode: nil,
	}
	headlnode.AddInHead(&a)
	headlnode.PrintlNode()
	fmt.Println("第二次插入在第一个位置")
	headlnode.Insert(1, &components.LNode{
		Data:      2,
		NextlNode: nil,
	})
	headlnode.PrintlNode()
	fmt.Println("第二次头插")
	headlnode.AddInHead(&components.LNode{
		Data:      3,
		NextlNode: nil,
	})
	headlnode.PrintlNode()
}
