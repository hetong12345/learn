package main

import (
	"./test"
	"fmt"
	"time"
)

type Ret interface {
	Post(url string) string
}
type Poster interface {
	Get(url string) string
}

func down(r Ret) string {
	return r.Post("http://www.renketong.com")
}
func main() {
	var r Ret
	r = test.Ret{"Mozilla/5.0", time.Minute}
	fmt.Println(down(r))
}
