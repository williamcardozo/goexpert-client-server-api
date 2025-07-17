package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	exchangerate "github.com/williamcardozo/goexpert-client-server-api/pkg/exchange-rate"
	"github.com/williamcardozo/goexpert-client-server-api/pkg/models"
)

func getExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	bid, err := exchangerate.GetExchangeRateBID()
	if err != nil {
		log.Printf("Erro ao buscar cotação: %v", err)
		http.Error(w, "Erro ao buscar cotação", http.StatusInternalServerError)
		return
	}

	resp := models.ExchangeRateResponse{
		Bid: bid,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func InitServer(ready chan<- struct{}) error {
	http.HandleFunc("/cotacao", getExchangeRateHandler)

	log.Println("Servidor iniciado na porta 8080...")
	ready <- struct{}{}

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Erro ao iniciar servidor: %v", err)
		ready <- struct{}{}
		return fmt.Errorf("erro ao iniciar servidor: %w", err)
	}
	ready <- struct{}{}
	return nil
}
