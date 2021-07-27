package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "BnB app")
		fmt.Println("Bytes written:", n)
		if err != nil {
			fmt.Println(err)
		}
	})

	_ = http.ListenAndServe(":8080", nil)
}
