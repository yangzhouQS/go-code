package main

import (
	"fmt"
	"../github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "https://book.douban.com/"
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// 提交请求
	request, error := http.NewRequest("get", url, nil)
	if error != nil {
		panic(error)
	}
	request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	// 发起请求
	response, err := client.Do(request)
	// 关闭响应体
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("response status code :", response.StatusCode)
	}

	// 读取响应
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	//解决中文乱码
	bodystr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	fmt.Println(bodystr)

}
