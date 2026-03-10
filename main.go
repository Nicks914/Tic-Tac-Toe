package main

import (
	"fmt"
	"net/http"

	"tic-tac-toe/web"
)

func main() {

	fmt.Println("Server started at http://localhost:8080")

	http.HandleFunc("/", web.HomeHandler)
	http.HandleFunc("/move", web.MoveHandler)
	http.HandleFunc("/reset", web.ResetHandler)
	http.HandleFunc("/mode", web.ModeHandler)

	http.ListenAndServe(":8080", nil)
}
