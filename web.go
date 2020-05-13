package main

import (
	"fmt"
	"io"
	"net/http"
)

type webWriter struct{}

func (webWriter) Write(w []byte) (int, error) {
	fmt.Println(string(w))
	return len(w), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	e := webWriter{}
	io.Copy(e, resp.Body)
}
