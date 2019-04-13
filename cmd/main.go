package main

import (
	"fmt"
	"net/http"
)

func Print(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	for key, value := range query {
		fmt.Printf("%s=%+v\n", key, value)
	}
}

func main() {
	http.HandleFunc("/test", Print)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}
