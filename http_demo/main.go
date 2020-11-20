package main

import (
	httpResponse "http_demo/response"
	"net/http"
)

func main() {
	// basic.Get()
	// basic.Post()
	// basic.Put()
	// basic.Delete()

	// get请求参数传递
	// basic.RequestGetParams()
	// var url = "http://httpbin.org/get"
	// var url2 = "https://studygolang.com/articles/2201"
	var url2 = "http://www.xiancn.com/quality/xawb_lianxi.html"
	r, _ := http.Get(url2)
	defer r.Body.Close()

	// 响应信息 测试
	httpResponse.ResponseBody(r)
	httpResponse.ResponseHeader(r)
	httpResponse.ResponseStatus(r)
	httpResponse.ResponseEncoding(r)
}
