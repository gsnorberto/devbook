package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	r := router.Generate()

	port := fmt.Sprintf("127.0.0.1:%d", config.Port)

	fmt.Printf("Servidor rodando na porta %s...", port)

	log.Fatal(http.ListenAndServe(port, r))
}
