package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hi, this is home page")
	if err != nil {
		return
	}
}

// body 只能读取一次
func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "read body failed: %v", err)
		// 记住要返回，不然就还会执行后面的代码
		return
	}
	// 类型转换，将 []byte 转换为 string
	_, _ = fmt.Fprintf(w, "read the data: %s \n", string(body))

	// 尝试再次读取，啥也读不到，但是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		_, _ = fmt.Fprintf(w, "read the data one more time got error: %v", err)
		return
	}
	println(body)

	if body == nil {
		_, _ = fmt.Fprintf(w, "read the data one body : %v", body)
		return
	}
	_, _ = fmt.Fprintf(w, "read the data one more time: [%s] and read data length %d \n", string(body), len(body))
}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		_, _ = fmt.Fprint(w, "GetBody is nil \n")
	} else {
		_, _ = fmt.Fprintf(w, "GetBody not nil \n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	_, _ = fmt.Fprintf(w, "query is %v\n", values)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	_, _ = fmt.Fprintf(w, string(data))
}

func header(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "header is %v\n", r.Header)
}

func form(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		_, _ = fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	_, _ = fmt.Fprintf(w, "before parse form %v\n", r.Form)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/body/multi", getBodyIsNil)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/header", header)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/form", form)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
