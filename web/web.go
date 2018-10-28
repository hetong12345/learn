package main

import (
	"github.com/hetong12345/learn/web/router"
	"net/http"
)

func main() {
	err := router.StartHandle()
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
