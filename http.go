package main

import (
	"net/http"
)

type myHandler struct{}

func (myHandler *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func main() {
	http.ListenAndServe("127.0.0.1:8888", &myHandler{})
}
