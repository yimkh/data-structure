package main

import (
	"data-structure/linkList/components"
	"fmt"
)

func main() {
	headlnode := components.Init()
	a := components.LNode{
		Data:      1,
		NextlNode: nil,
	}
	headlnode.AddInEnd(&a)
	b := components.LNode{
		Data:      2,
		NextlNode: nil,
	}
	headlnode.AddInEnd(&b)
	headlnode.AddInEnd(&components.LNode{
		Data:      3,
		NextlNode: nil,
	})
	fmt.Println("初始链表:")
	headlnode.PrintlNode()
}
