package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("04_Pacotes_importantes/08_FileServer/public/"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from blog"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
