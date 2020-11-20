package ArrayList

type StackArrayX interface {
	Size() int        // 大小
	Clear()           // 清空
	Pop() interface{} // 弹出
	Push()            // 压栈
	IsFull() bool     // 是否满了
	IsEmpty() bool    // 是否为空
}
type StackX struct {
	myArray *ArrayList // 保存数据容器,使用自己实现的数组保存
	Iterator    Iterator   //
}

func NewArrayListStackX() *StackX {
	stack := new(StackX)
	stack.myArray = NewArrayList()        // 数组
	stack.Iterator = stack.myArray.Iterator() // 迭代
	return stack
}

// 栈大小
func (stack *StackX) Size() int {
	return stack.myArray.size
}

// 清空栈
func (stack *StackX) Clear() {
	stack.myArray.Clear()
	stack.myArray.size = 0
}

// 出栈
func (stack *StackX) Pop() interface{} {
	if !stack.IsEmpty() {
		lastVal := stack.myArray.dataStore[stack.myArray.size-1]
		stack.myArray.Delete(stack.myArray.size - 1)
		return lastVal // 返回最后一个数据
	}
	return nil
}

// 压栈
func (stack *StackX) Push(value interface{}) {
	// 栈满了
	if stack.IsFull() {

	} else { // 栈没有满
		stack.myArray.Append(value)
	}
}

// 栈是否满了
func (stack *StackX) IsFull() bool {
	return stack.myArray.size >= 10
}

// 是否为空
func (stack *StackX) IsEmpty() bool {
	return stack.myArray.size == 0
}
