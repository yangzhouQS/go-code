package ArrayList

type StackArray interface {
	Size() int        // 大小
	Clear()           // 清空
	Pop() interface{} // 弹出
	Push()            // 压栈
	IsFull() bool     // 是否满了
	IsEmpty() bool    // 是否为空
}
type Stack struct {
	myArray *ArrayList // 保存数据容器,使用自己实现的数组保存
	capSize int        // 最大容量
}

func NewArrayListStack() *Stack {
	stack := new(Stack)
	stack.myArray = NewArrayList()
	stack.capSize = 10
	return stack
}

// 栈大小
func (stack *Stack) Size() int {
	return stack.myArray.size
}

// 清空栈
func (stack *Stack) Clear() {
	stack.myArray.Clear()
	stack.capSize = 10
}

// 出栈
func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty() {
		lastVal := stack.myArray.dataStore[stack.myArray.size-1]
		stack.myArray.Delete(stack.myArray.size-1)
		return lastVal // 返回最后一个数据
	}
	return nil
}

// 压栈
func (stack *Stack) Push(value interface{}) {
	// 栈满了
	if stack.IsFull() {

	} else { // 栈没有满
		stack.myArray.Append(value)
	}
}

// 栈是否满了
func (stack *Stack) IsFull() bool {
	return stack.myArray.size >= stack.capSize
}

// 是否为空
func (stack *Stack) IsEmpty() bool {
	return stack.myArray.size == 0
}
