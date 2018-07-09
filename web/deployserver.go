package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)
}
func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello World</h1>")
}
func reLaunch() {

}
