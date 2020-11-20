package main

import (
	"01-ArrayList/Queue"
	"fmt"
	"io/ioutil"
)

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
