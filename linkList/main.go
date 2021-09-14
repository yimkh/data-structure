package main

import (
	"data-structure/linkList/components"
	"fmt"
)

func main() {
	headlnode := components.Init()
	headlnode.IsEmpty()
	fmt.Printf("the length is: %v\n", headlnode.Length())
	fmt.Println("第一次尾插")
	a := components.LNode{
		Data:      1,
		NextlNode: nil,
	}
	headlnode.AddInEnd(&a)
	headlnode.PrintlNode()
	fmt.Println("第二次尾插")
	b := components.LNode{
		Data:      2,
		NextlNode: nil,
	}
	headlnode.AddInEnd(&b)
	headlnode.PrintlNode()
	fmt.Println("第三次尾插")
	headlnode.AddInEnd(&components.LNode{
		Data:      3,
		NextlNode: nil,
	})
	headlnode.PrintlNode()

	fmt.Printf("the length is: %v\n", headlnode.Length())
}
