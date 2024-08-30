package main

import (
	"fmt"
	"goHTML/vdom"
	"log"
	"net/http"
)

var (
	Tag        = vdom.Tag
	GenerateID = vdom.GenerateID
	Text       = vdom.Text
)

func main() {
	// Set up HTTP server
	http.HandleFunc("/", serverHomePage)

	// Start the server
	fmt.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// Note: This code will never be reached
	// fmt.Println("\n\n\nState Example:")
	// vdom.ExampleUsage4()
}

func serverHomePage(w http.ResponseWriter, r *http.Request) {
	page := vdom.HomePage()
	// type text/html
	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write([]byte(page))
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
