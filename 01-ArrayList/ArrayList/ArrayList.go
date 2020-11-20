package ArrayList

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int                                 // 数组大小
	Get(index int) (interface{}, error)        // 根据所有获取指定位置的值
	Set(index int, value interface{}) error    // 在指定索引位置插入元素
	Insert(index int, value interface{}) error // 在指定索引位置插入元素
	Append(value interface{}) int              // 追加
	Clear()                                    //清空
	Delete(index int) error                    //删除
	String() string                            // 返回字符串
	Iterator() Iterator                        // 继承接口
}

// 数据结构, 字符串,整数,实数
type ArrayList struct {
	dataStore []interface{} // 数组存储
	size      int           // 数组大小
}

// 实例化数组的方法
func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10) // 容量为10
	list.size = 0
	return list
}

// 数组大小
func (list *ArrayList) Size() int {
	return list.size
}

// 获取指定位置的值
func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= list.size {
		return nil, errors.New("获取的所有越界")
	}
	return list.dataStore[index], nil
}

func (list *ArrayList) Set(index int, value interface{}) error {
	if index < 0 || index >= list.size {
		return errors.New("获取的所有越界")
	}
	list.dataStore[index] = value // 修改值
	return nil
}

// 长度检测和扩容
func (list *ArrayList) checkIsFull() {
	if list.size == cap(list.dataStore) { // 判断大小和容量
		//newStore := make([]interface{}, 0, 2*list.size)
		// 内存开辟
		newStore := make([]interface{}, 2*list.size, 2*list.size)
		copy(newStore, list.dataStore)
		list.dataStore = newStore
	}
}

func (list *ArrayList) Insert(index int, value interface{}) error {
	if index < 0 || index >= list.size {
		return errors.New("获取的所有越界")
	}
	list.checkIsFull()
	list.dataStore = list.dataStore[:list.size+1] // 长度增加一位

	// 从index开始向后移动一位
	for i := list.size; i > index; i-- {
		list.dataStore[i] = list.dataStore[i-1]
	}
	// 指定位置插入
	list.dataStore[index] = value
	list.size++
	return nil
}

func (list *ArrayList) Append(value interface{}) int {
	list.dataStore = append(list.dataStore, value)
	list.size++
	return list.size
}

// 清空数组数据
func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.size = 0
}

// 删除指定所有处的元素
func (list *ArrayList) Delete(index int) error {
	if index < 0 || index >= list.size {
		return errors.New("获取的所有越界")
	}
	list.dataStore = append(list.dataStore[:index], list.dataStore[index+1:]...)
	list.size--
	return nil
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list.dataStore)
}
