package basic

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get() {
	r, error := http.Get("https://news.baidu.com/sports")
	if error != nil {
		panic(error)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		fmt.Println("fetch error code: = ", r.StatusCode)
	}
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", string(content))
}

/**
Post请求
*/
func Post() {
	r, error := http.Get("https://news.baidu.com/sports")
	if error != nil {
		panic(error)
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		fmt.Println("fetch error code: = ", r.StatusCode)
	}
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", string(content))
}

var url = "https://news.baidu.com/sports"

/**
Http PUT请求
*/
func Put() {
	request, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer func() { _ = response.Body.Close() }()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func Delete() {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)
	defer func() { _ = response.Body.Close() }()
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

/**
打印响应体内容
*/
func PrintBody(r *http.Response) {
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

var url2 = "http://httpbin.org/get"

/**
Get参数传递测试
*/
func RequestGetParams() {
	request, err := http.NewRequest(http.MethodGet, url2, nil)
	if err != nil {
		panic(err)
	}
	params := request.URL.Query()
	params.Add("name", "sam")
	params.Add("age", "26")
	request.URL.RawQuery = params.Encode()

	fmt.Println("URL.String", request.URL.String()) // URL.String http://httpbin.org/get?age=26&name=sam
	fmt.Println("query params: ", params.Encode())  // age=26&name=sam
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	PrintBody(response)
}
