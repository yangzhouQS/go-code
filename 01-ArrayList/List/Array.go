package main

import (
	"errors"
	"fmt"
)

type List interface {
	Size() int                            // 数组大小
	Get(index int) interface{}            // 获取指定索引的值
	Set(index int, data interface{})      // 修改值
	Insert(index int, data interface{})   // 在指定位置插入
	Append(value interface{}) interface{} // 在末尾追加
	Clear() bool                          // 清空
	Delete(index int) interface{}         // 删除指定索引的元素
	String() string                       // 序列化输出
}
type ArrayList struct {
	dataStore []interface{} // 存储数组元素的容器
	TheSize   int           // 数组带下
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.TheSize
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index >= list.Size() || index < 0 {
		return nil, errors.New("index error")
	} else {
		return list.dataStore[index], nil
	}
}

func (list *ArrayList) Set(index int, value interface{}) error {
	panic("implement me")
}

func (list *ArrayList) Insert(index int, value interface{}) error {
	panic("implement me")
}

func (list *ArrayList) Append(value interface{}) int {
	list.dataStore = append(list.dataStore, value)
	list.TheSize++
	return list.TheSize
}

func (list *ArrayList) Clear() {
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
}

func (list *ArrayList) Delete(index int) error {
	panic("implement me")
}

func (list *ArrayList) String() string {
	return fmt.Sprint(list)
}

func main() {

}
