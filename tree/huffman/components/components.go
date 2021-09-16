package components

import "fmt"

type Node struct {
	Data         int
	LNode, RNode *Node
}

var Nodes []Node

func Init(data int) Node {
	return Node{
		Data:  data,
		LNode: nil,
		RNode: nil,
	}
}

func SortNodes(Nodes []Node) {
	for i := 0; i < len(Nodes)-1; i++ {
		flag := false
		for j := len(Nodes) - 1; j > i; j-- {
			if Nodes[j-1].Data > Nodes[j].Data {
				temp := Nodes[j]
				Nodes[j] = Nodes[j-1]
				Nodes[j-1] = temp
				flag = true
			}
		}
		if flag == false {
			return
		}
	}
}

func Huffman(Nodes []Node) []Node {
	SortNodes(Nodes)

	var desNodes []Node
	times := len(Nodes)

	for i := 0; i < times; i++ {
		if len(Nodes) > 2 || len(Nodes) == 2 {
			desNodes = append(desNodes, Nodes[0], Nodes[1])
			compoundNode := Node{
				Data:  Nodes[0].Data + Nodes[1].Data,
				LNode: &desNodes[i*2],
				RNode: &desNodes[i*2+1],
			}
			Nodes = append(Nodes[:0], Nodes[2:]...)
			Nodes = append(Nodes, compoundNode)
			SortNodes(Nodes)
		} else if len(Nodes) == 1 {
			desNodes = append(desNodes, Nodes[0])
			Nodes = append(Nodes[:0], Nodes[1:]...)
		} else if len(Nodes) == 0 {
			return desNodes
		}
	}
	return desNodes
}

func PrintHuffman(huffmanNodes []Node) {
	for i := 0; i < len(huffmanNodes); i++ {
		fmt.Printf("元素的值为: %v, ", huffmanNodes[i].Data)
		if huffmanNodes[i].LNode == nil {
			fmt.Printf("无左子树, ")
		} else {
			fmt.Printf("左子树的值为: %v,", huffmanNodes[i].LNode.Data)
		}
		if huffmanNodes[i].RNode == nil {
			fmt.Println("无右子树")
		} else {
			fmt.Printf("右子树的值为: %v\n", huffmanNodes[i].RNode.Data)
		}
	}
}
