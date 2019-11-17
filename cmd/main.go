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
	fmt.Printf("%v", string([]byte{0x49, 0x20, 0x4c, 0x6f, 0x76, 0x65, 0x20, 0x59, 0x6f, 0x75}))
}
