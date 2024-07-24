package main

import (
	"fmt"
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
	mux.HandleFunc("/upload", uploadOneFile)
	mux.HandleFunc("/upload-one", uploadFile)
	return mux
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	t := toolkit.Tools{
		MaxFileSize: 1024 * 1024 * 2,
		AllowedFileTypes: []string{
			"image/jpeg",
			"image/png",
			"image/gif",
		},
	}
	uploadedFiles, err := t.UploadFiles(r, "./tmp", true)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	out := ""
	for _, f := range uploadedFiles {
		out += fmt.Sprintf("Uploaded file: %s (%d bytes)\n", f.OriginalFileName, f.FileSize)
	}
	_, _ = w.Write([]byte(out))
}

func uploadOneFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	t := toolkit.Tools{
		MaxFileSize: 1024 * 1024 * 2,
		AllowedFileTypes: []string{
			"image/jpeg",
			"image/png",
			"image/gif",
		},
	}
	uploadedFile, err := t.UploadOneFile(r, "./tmp", true)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	out := fmt.Sprintf("Uploaded file: %s (%d bytes)\n", uploadedFile.OriginalFileName, uploadedFile.FileSize)
	_, _ = w.Write([]byte(out))
}
