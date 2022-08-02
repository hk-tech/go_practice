package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Entry struct {
	Initial string
	Fruit   string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/upload", uploadFile).Methods("POST")
	http.ListenAndServe(":8080", r)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		output := strings.Join(line[:], " for ") + "\n"
		fmt.Fprintf(w, output)
	}
}
