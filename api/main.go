package main

import (
	router "api/src/Router"
	"api/src/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Println((config.ConnectionString))

	fmt.Print("Rodando Api")
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprint(":%d", config.Port), r))
}
