package main

import "net/http"
import "fmt"

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}