package main

import (
	"data-structure/tree/huffman/components"
	"fmt"
)

func main() {
	Nodes := components.Nodes
	Nodes = append(Nodes, components.Init(1))
	Nodes = append(Nodes, components.Init(3))
	Nodes = append(Nodes, components.Init(2))
	Nodes = append(Nodes, components.Init(2))
	Nodes = append(Nodes, components.Init(7))

	fmt.Println("初始序列为: ")
	fmt.Println(Nodes)
	fmt.Println("哈夫曼序列为: ")
	huffmanNodes := components.Huffman(Nodes)
	components.PrintHuffman(huffmanNodes)
}
