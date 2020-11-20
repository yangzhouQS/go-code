package main

import (
	"01-ArrayList/Queue"
	"01-ArrayList/StackArray"
	"errors"
	"fmt"
	"io/ioutil"
)

func GetAll(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path)
	if err != nil {
		return files, errors.New("文件夹不可以读取")
	}
	for _, file := range read {
		if file.IsDir() { // 文件夹
			fullDir := path + "\\" + file.Name()
			files = append(files, fullDir)
			files, _ = GetAll(fullDir, files)
		} else {
			fullDir := path + "\\" + file.Name() // 构建新的路径
			files = append(files, fullDir)       // 追加路径
		}
	}
	return files, nil
}

// 递归实现文件夹遍历
func main01() {
	path := "F:\\Rust"
	files := []string{}
	files, _ = GetAll(path, files)
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
}

// 栈实现文件夹遍历
func main02() {
	path := "C:\\Users\\xianjs\\Desktop\\小狗\\2019年4月28日222236"
	files := []string{}
	stack := StackArray.NewStack()
	stack.Push(path)
	for !stack.IsEmpty() {
		newPath := stack.Pop().(string)
		//files = append(files, newPath)
		readDir, err := ioutil.ReadDir(newPath)
		if err != nil {
			panic("没有文件夹读取权限")
		}
		for _, file := range readDir {
			if file.IsDir() {
				fullPath := newPath + "\\" + file.Name()
				files = append(files, fullPath)
				stack.Push(fullPath)
			} else {
				files = append(files, fmt.Sprint(newPath+"\\"+file.Name()))
			}
		}
	}

	for k, fileName := range files {
		fmt.Println(k, fileName)
	}
}

func main03() {
	path := "C:\\Users\\xianjs\\Desktop\\小狗\\2019年4月28日222236"
	files := []string{}
	queue := Queue.NewQueue()
	queue.EnQueue(path)
	for {
		if !queue.IsEmpty() {
			filePath := queue.DeQueue().(string)
			readDir, err := ioutil.ReadDir(filePath)
			if err != nil {
				panic(err)
			}
			for _, dir := range readDir {
				if dir.IsDir() {
					newPath := filePath + "\\" + dir.Name()
					queue.EnQueue(newPath)
					files = append(files, newPath)
				} else {
					files = append(files, filePath+"\\"+dir.Name())
				}
			}
			//fmt.Println(readDir)
		} else {
			break
		}
	}
	fmt.Println("end")
	for k, v := range files {
		fmt.Println(k, v)
	}
}

func main() {
	path := "E:\\2017年毕业-实习i图片"
	queue := Queue.NewQueue()
	queue.EnQueue(path)
	fiels := []string{}
	for {
		if !queue.IsEmpty() {
			newPath := queue.DeQueue().(string)
			readDir, err := ioutil.ReadDir(newPath)
			if err != nil {
				panic(err)
			}
			for _, fileInfo := range readDir {
				if fileInfo.IsDir() {
					newPath2 := newPath + "\\" + fileInfo.Name()
					queue.EnQueue(newPath2)
					fiels = append(fiels, newPath2)
				} else {
					fiels = append(fiels, newPath+"\\"+fileInfo.Name())
				}
			}
		} else {
			break
		}
	}
	for index, value := range fiels {
		fmt.Println(index, value)
	}
}
