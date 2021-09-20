# data-structure
## 目录
一、线性表  <br>
1.顺序表  <br>
2.单链表  <br>
二、树  <br>
1.哈夫曼树  <br>
三、排序  <br>
1.插入排序  <br>
2.选择排序  <br>
## 顺序表
### 规定: 
1.顺序表从0开始，次序从1开始  <br>
2.元素的值均 >= 0, 次序均 >= 1  <br>
### 结构
```
type SqLists struct {
	maxSize int
	data    []int
	length  int
}
```
### 方法
```
func Init(maxSize int) SqLists

func PrintSqList(sqList SqLists)

func PrintLength(sqList SqLists)

func IsEmpty(sqList SqLists) bool

func IsFull(sqList SqLists) bool 

func Add(sqList *SqLists, data int) 

func AddList(sqList *SqLists, data []int) 

func Delete(sqList *SqLists, locat int) 

func Insert(sqList *SqLists, locat int, data int) 

func GetElem(sqList SqLists, locat int) int 

func LocateElem(sqList SqLists, value int) int 
```
## 单链表
### 规定:
1.带头结点（头结点次序为0，值为-1） <br>
2.每个结点除头结点外，次序和值均 >= 0  <br>
3.增删改查按次序，1代表第一个有真正值的结点  <br>
4.返回-1，代表没有此项值或结点次序（头结点除外）  <br>
### 结构
```
type LNode struct {
	Data      int
	NextlNode *LNode
}
```
### 方法
```
func Init() LNode

func IsEmpty(lNode LNode) bool

func IsIllegal(lNode *LNode) bool

func Length(lNode *LNode) int

func PrintlNode(lNode LNode)

func AddInHead(lNode *LNode, newlNode *LNode)

func AddInEnd(lNode *LNode, newlNode *LNode)

func Delete(lNode *LNode, locat int) (elem int)

func DeleteNode(lNode *LNode, deleteNode *LNode)

func Insert(lNode *LNode, locat int, newlNode *LNode)

func InsertNext(lNode *LNode, newlNode *LNode)

func InsertPrior(lNode *LNode, newlNode *LNode)

func GetElem(lNode *LNode, locat int) *LNode

func LocatElem(lNode *LNode, value int)
```
## 哈夫曼树
### 规定
1.各结点的值均 >= 0  <br>
### 结构
```
type Node struct {
	Data         int
	LNode, RNode *Node
}
```
### 方法
```
func Init(data int) Node {

func SortNodes(Nodes []Node) {

func Huffman(Nodes []Node) []Node {

func PrintHuffman(huffmanNodes []Node) {
```
## 插入排序
### 规定
1.升序排列  <br>
2.插入排序: 从第2个元素开始, 向前寻找插入位置  <br>
  直接排序: 从后向前依次寻找  <br>
  折半排序: 从后向前，折半寻找(mid 向下取整)  <br>
  希尔排序：分组后进行直接排序，分组由多变少 <br>
3.选择排序：  <br>
  冒泡排序：从后往前选出最小元素  <br>
  快速排序：  <br>
### 插入方法
```
func DirectInsertSort(src []int)

func HarfInsertSort(src []int)

func ShellSort(src []int, increment int)
```
### 选择方法
```
func BubbleSort(src []int)
```