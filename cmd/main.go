package main

import (
	"fmt"
	"log"

	"github.com/williamcardozo/goexpert-client-server-api/pkg/client"
	"github.com/williamcardozo/goexpert-client-server-api/pkg/server"
)

func main() {
	go func() {
		server.InitServer()
		log.Println("Servidor iniciado com sucesso!")
	}()
	
	err := client.FetchExchangeRate()
	if err != nil {
		panic(fmt.Errorf("erro ao buscar taxa de câmbio: %w", err))
	}
	log.Println("Taxa de câmbio buscada com sucesso!")
}
