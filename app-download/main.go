package main

import (
	"github.com/CristianC4/toolkit"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Printf("Starting server port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/download", downloadFile)
	return mux
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	tooler := toolkit.Tools{}
	tooler.DownloadStaticFile(w, r, "./files", "pic-dog.jpg", "white-wolf.png")
}
