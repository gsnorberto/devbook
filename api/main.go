package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("")

	r := router.Generate()

	log.Println("Servidor rodando na porta 5000...")
	log.Fatal(http.ListenAndServe("127.0.0.1:5000", r))
}
