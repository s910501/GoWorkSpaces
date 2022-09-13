package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	// str := "<h1 style='color:red'>Hello Http!!</h1>"
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		// fmt.Println("open file err:", err)
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	name := queryParam.Get("name")
	age := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/xxx/", f2)
	http.ListenAndServe("127.0.0.1:9000", nil)
}
