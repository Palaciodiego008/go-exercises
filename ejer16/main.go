package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3001", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving index.html")
	http.ServeFile(w, r, "./index.html")

}
