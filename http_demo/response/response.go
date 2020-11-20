package httpResponse

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
)

var url = "http://httpbin.org/get"

func ResponseBody(r *http.Response) {
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}
func ResponseHeader(r *http.Response) {
	fmt.Println(r.Header.Get("content-type")) // application/json
	fmt.Printf("%s \n", r.Header)
}
func ResponseStatus(r *http.Response) {
	fmt.Println(r.StatusCode) // 200
	fmt.Println(r.Status)     // 200 OK
}
func ResponseEncoding(r *http.Response) {
	// 	content-type 和 <meta charset="UTF-8"> 会提供编码
	reader := bufio.NewReader(r.Body)
	bytes, _ := reader.Peek(1024)
	encoding, _, _ := charset.DetermineEncoding(bytes, r.Header.Get("content-type"))
	fmt.Println(encoding)

	// 编码转换
	newReader := transform.NewReader(reader, encoding.NewDecoder())
	fmt.Println(ioutil.ReadAll(newReader))
}
func PrintBody(r *http.Response) {
	defer r.Body.Close()
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}
