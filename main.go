package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %v\n", name)
	fmt.Fprintf(w, "address = %v\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello!")
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current working directory: %s\n", cwd)

	staticDir := "./static"
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Printf("Warning: Static directory '%s' does not exist, creating it...\n", staticDir)
		if err := os.MkdirAll(staticDir, 0755); err != nil {
			log.Fatalf("Failed to create static directory: %v", err)
		}
	}

	files, err := filepath.Glob(filepath.Join(staticDir, "*"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Files in static directory: %v\n", files)

	fileServer := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
