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

	if config.Porta == 0 {
		log.Fatal("A porta da API não foi encontrada nas variáveis de ambiente.")
	}

	fmt.Printf("Rodando API na porta %d...\n", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
