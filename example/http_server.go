package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		writer.Write([]byte("Hello world\n"))
	})

	http.ListenAndServe(":8888", nil)
}

