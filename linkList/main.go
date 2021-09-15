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
	components.AddInEnd(&headlnode, &a)
	b := components.LNode{
		Data:      2,
		NextlNode: nil,
	}
	components.AddInEnd(&headlnode, &b)
	components.AddInEnd(&headlnode, &components.LNode{
		Data:      3,
		NextlNode: nil,
	})
	fmt.Println("初始链表:")
	components.PrintlNode(headlnode)
}
