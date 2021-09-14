package main

import (
	"data-structure/linkList/components"
	"fmt"
	"testing"
)

func TestAddInHead(t *testing.T) {
	headlnode := components.Init()
	headlnode.IsEmpty()
	fmt.Println("第一次头插")
	a := components.LNode{
		Data:      1,
		NextlNode: nil,
	}
	headlnode.AddInHead(&a)
	headlnode.PrintlNode()
	fmt.Println("第二次头插")
	b := components.LNode{
		Data:      2,
		NextlNode: nil,
	}
	headlnode.AddInHead(&b)
	headlnode.PrintlNode()
	fmt.Println("第三次头插")
	headlnode.AddInHead(&components.LNode{
		Data:      3,
		NextlNode: nil,
	})
	headlnode.PrintlNode()
}

func TestAddInEnd(t *testing.T) {
	headlnode := components.Init()
	headlnode.IsEmpty()
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
}

func TestInsert(t *testing.T) {
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
	fmt.Println("第三次插入在第五个位置")
	headlnode.Insert(5, &components.LNode{
		Data:      5,
		NextlNode: nil,
	})
	headlnode.PrintlNode()
}

func TestInsertNest(t *testing.T) {
	headlnode := components.Init()
	headlnode.IsEmpty()
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
	fmt.Printf("在元素%v之后插入%v\n", 2, 4)
	b.InsertNext(&components.LNode{
		Data:      4,
		NextlNode: nil,
	})
	headlnode.PrintlNode()
}

func TestGetElem(t *testing.T) {
	headlnode := components.Init()
	headlnode.IsEmpty()
	fmt.Println("第一次头插")
	a := components.LNode{
		Data:      1,
		NextlNode: nil,
	}
	headlnode.AddInHead(&a)
	headlnode.PrintlNode()
	fmt.Println("第二次头插")
	b := components.LNode{
		Data:      2,
		NextlNode: nil,
	}
	headlnode.AddInHead(&b)
	headlnode.PrintlNode()
	fmt.Println("第三次头插")
	headlnode.AddInHead(&components.LNode{
		Data:      3,
		NextlNode: nil,
	})
	headlnode.PrintlNode()
	fmt.Printf("第%v个位置的值为: %v\n", 1, headlnode.GetElem(1).Data)
	fmt.Printf("第%v个位置的值为: %v\n", 3, headlnode.GetElem(3).Data)
}

func TestLocatElem(t *testing.T) {
	headlnode := components.Init()
	headlnode.IsEmpty()
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

	_, locat1 := headlnode.LocatElem(2)
	_, locat2 := headlnode.LocatElem(5)
	fmt.Printf("the location of %v: %v\n", 2, locat1)
	fmt.Printf("the location of %v: %v\n", 5, locat2)
}

func TestLength(t *testing.T) {
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
