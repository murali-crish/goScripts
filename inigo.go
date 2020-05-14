package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprint(res, "Hello, my name is Inigo Montoya")
}

func main() {
	http.HandleFunc("/", hello)
	_ = http.ListenAndServe("localhost:4000", nil)
}
