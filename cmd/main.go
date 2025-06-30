package main

import (
	"log"

	"github.com/williamcardozo/goexpert-client-server-api/pkg/client"
	"github.com/williamcardozo/goexpert-client-server-api/pkg/server"
)

func main() {
	ready := make(chan struct{})

	go func() {
		if err := server.InitServer(ready); err != nil {
			log.Fatalf("Erro ao iniciar servidor: %v", err)
		}
	}()

	// Wait until server is ready
	<-ready

	// Now call client
	if err := client.FetchExchangeRate(); err != nil {
		log.Fatalf("Erro ao buscar taxa de câmbio: %v", err)
	}

	log.Println("Taxa de câmbio buscada com sucesso!")
}
