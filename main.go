package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe("http://localhost:8888", nil)
}

func handle(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello world")
}
