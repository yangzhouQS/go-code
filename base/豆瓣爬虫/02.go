package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	url := "https://search.douban.com/book/subject_search?search_text=JavaScript&cat=1001"
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	//增加header选项
	//reqest.Header.Add("Cookie", "xxxxxx")
	reqest.Header.Add("User-Agent", "mozilla/5.0 (windows nt 10.0; win64; x64) applewebkit/537.36 (khtml, like gecko) chrome/81.0.4044.138 safari/537.36")
	//reqest.Header.Add("X-Requested-With", "xxxx")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, resErr := client.Do(reqest)
	defer response.Body.Close()
	if resErr != nil {
		panic(resErr)
	}
	result, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("result %s\n", result)
}
