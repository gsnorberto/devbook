package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// environment variables
	config.Load()

	// configure the routes
	r := router.Generate()

	port := config.Port
	host := fmt.Sprintf("127.0.0.1:%d", port)

	fmt.Printf("Servidor rodando na porta %d...", port)

	log.Fatal(http.ListenAndServe(host, r))
}
