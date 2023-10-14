package main

import (
	"alura-go-store/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
