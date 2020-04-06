package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8888/")
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(content))
	}
}
