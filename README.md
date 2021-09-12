# data-structure
## sqList 线性表
### 结构体
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

func (sqList SqLists) PrintSqList()

func (sqList SqLists) PrintLength()

func (sqList SqLists) IsEmpty() bool

func (sqList SqLists) IsFull() bool 

func (sqList *SqLists) Add(data int)

func (sqList *SqLists) AddList(data []int)

func (sqList *SqLists) Insert(locat int, data int) 

func (sqList *SqLists) Delete(locat int)

func (sqList SqLists) GetElem(locat int) int

func (sqList SqLists) LocateElem(value int) int
```
