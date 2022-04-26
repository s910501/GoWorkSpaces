package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:9000/xxx/?name=rekkles&age=18")
	// if err != nil {
	// 	fmt.Printf("Error message:%v", err)
	// 	return
	// }

	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:9000/xxx/")
	data.Set("name", "止上")
	data.Set("age", "18")
	queryStr := data.Encode()
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, _ := http.NewRequest("GET", urlObj.String(), nil)
	// resp, err := http.DefaultClient.Do(req)
	// if err != nil {
	// 	fmt.Printf("read resp.Body failed,err %v", err)
	// 	return
	// }
	// disable keepalive 单请求很频繁时，需要禁用
	tr := &http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("read resp.Body failed,err %v", err)
		return
	}
	//必须关闭
	defer resp.Body.Close()
	// read response
	// var data []byte
	// resp.Body.Read()
	// resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp.Body failed,err %v", err)
		return
	}

	fmt.Println(string(b))
}
