package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	// Se fosse utilizado o Millisecond iria dar um erro context de timeout
	client := http.Client{Timeout: time.Millisecond}
	response, err := client.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
