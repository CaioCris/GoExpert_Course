package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	client := http.Client{}
	json := bytes.NewBuffer([]byte(`{"name": "Caio Cris"}`))
	response, err := client.Post("https://google.com", "application/json", json)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	io.CopyBuffer(os.Stdout, response.Body, nil)
}
