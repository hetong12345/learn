package main

import (
	"./retriever"
	"net/http"
	"net/http/httputil"
	"time"
)

type Ret struct {
	UserAgent string
	TimeOut   time.Duration
}

func download(r Ret) string {
	return r.Post("http://www.renketong.com")
}
func (r *Ret) Post(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	res, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(res)
}

type TreeNode struct {
	value       int
	left, right *TreeNode
}

func main() {
	var r Ret
	r = retriever.Ret{}
	println(download(r))
}