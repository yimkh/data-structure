package main

import (
	"data-structure/linkList/components"
	"fmt"
	"testing"
)

func TestLength(t *testing.T) {
	headlnode := components.Init()
	components.IsEmpty(headlnode)
	fmt.Printf("the length is: %v\n", components.Length(&headlnode))
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

	fmt.Printf("the length is: %v\n", components.Length(&headlnode))
}

func TestAddInHead(t *testing.T) {
	headlnode := components.Init()
	components.IsEmpty(headlnode)
	fmt.Println("第一次头插")
	a := components.LNode{
		Data:      1,
		NextlNode: nil,
	}
	components.AddInHead(&headlnode, &a)
	components.PrintlNode(headlnode)
	fmt.Println("第二次头插")
	b := components.LNode{
		Data:      2,
		NextlNode: nil,
	}
	components.AddInHead(&headlnode, &b)
	components.PrintlNode(headlnode)
	fmt.Println("第三次头插")
	components.AddInHead(&headlnode, &components.LNode{
		Data:      3,
		NextlNode: nil,
	})
	components.PrintlNode(headlnode)
}

func TestAddInEnd(t *testing.T) {
	headlnode := components.Init()
	components.IsEmpty(headlnode)
	fmt.Println("第一次尾插")
	a := components.LNode{
		Data:      1,
		NextlNode: nil,
	}
	components.AddInEnd(&headlnode, &a)
	components.PrintlNode(headlnode)
	fmt.Println("第二次尾插")
	b := components.LNode{
		Data:      2,
		NextlNode: nil,
	}
	components.AddInEnd(&headlnode, &b)
	components.PrintlNode(headlnode)
	fmt.Println("第三次尾插")
	components.AddInEnd(&headlnode, &components.LNode{
		Data:      3,
		NextlNode: nil,
	})
	components.PrintlNode(headlnode)
}

func TestDelete(t *testing.T) {
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

	fmt.Println("----------")

	deleteElem2 := components.Delete(&headlnode, 2)
	fmt.Println("删除第2个元素,删除后链表为:")
	components.PrintlNode(headlnode)
	fmt.Printf("删除元素的值为: %v\n", deleteElem2)

	fmt.Println("----------")

	deleteElem5 := components.Delete(&headlnode, 5)
	fmt.Println("删除第5个元素,删除后链表为:")
	components.PrintlNode(headlnode)
	fmt.Printf("删除元素的值为: %v\n", deleteElem5)

	fmt.Println("--------------------")
}

func TestDeleteNode(t *testing.T) {
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

	fmt.Println("删除值为2的元素,删除后链表为:")
	components.DeleteNode(&headlnode, &b)
	components.PrintlNode(headlnode)
}

func TestInsert(t *testing.T) {
	headlnode := components.Init()
	components.IsEmpty(headlnode)
	fmt.Println("第一次插入在第一个位置")
	components.Insert(&headlnode, 1, &components.LNode{
		Data:      2,
		NextlNode: nil,
	})
	fmt.Println("第一次头插")
	a := components.LNode{
		Data:      1,
		NextlNode: nil,
	}
	components.AddInHead(&headlnode, &a)
	components.PrintlNode(headlnode)
	fmt.Println("第二次插入在第一个位置")
	components.Insert(&headlnode, 1, &components.LNode{
		Data:      2,
		NextlNode: nil,
	})
	components.PrintlNode(headlnode)
	fmt.Println("第二次头插")
	components.AddInHead(&headlnode, &components.LNode{
		Data:      3,
		NextlNode: nil,
	})
	components.PrintlNode(headlnode)
	fmt.Println("第三次插入在第五个位置")
	components.Insert(&headlnode, 5, &components.LNode{
		Data:      5,
		NextlNode: nil,
	})
	components.PrintlNode(headlnode)
}

func TestInsertNest(t *testing.T) {
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

	fmt.Printf("在元素%v之后插入%v\n", 2, 4)
	components.InsertNext(&b, &components.LNode{
		Data:      4,
		NextlNode: nil,
	})
	components.PrintlNode(headlnode)
}

func TestInsertPrior(t *testing.T) {
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

	fmt.Println("在值等于2的结点前插入一个等于4的结点")
	d := components.LNode{
		Data:      4,
		NextlNode: nil,
	}
	components.InsertPrior(&b, &d)
	components.PrintlNode(headlnode)
}

func TestGetElem(t *testing.T) {
	headlnode := components.Init()
	components.IsEmpty(headlnode)
	fmt.Println("第一次头插")
	a := components.LNode{
		Data:      1,
		NextlNode: nil,
	}
	components.AddInHead(&headlnode, &a)
	components.PrintlNode(headlnode)
	fmt.Println("第二次头插")
	b := components.LNode{
		Data:      2,
		NextlNode: nil,
	}
	components.AddInHead(&headlnode, &b)
	components.PrintlNode(headlnode)
	fmt.Println("第三次头插")
	components.AddInHead(&headlnode, &components.LNode{
		Data:      3,
		NextlNode: nil,
	})
	components.PrintlNode(headlnode)
	fmt.Printf("第%v个位置的值为: %v\n", 1, components.GetElem(&headlnode, 1).Data)
	fmt.Printf("第%v个位置的值为: %v\n", 3, components.GetElem(&headlnode, 3).Data)
}

func TestLocatElem(t *testing.T) {
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

	_, locat1 := components.LocatElem(&headlnode, 2)
	_, locat2 := components.LocatElem(&headlnode, 5)
	fmt.Printf("the location of %v: %v\n", 2, locat1)
	fmt.Printf("the location of %v: %v\n", 5, locat2)
}
