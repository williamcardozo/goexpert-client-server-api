package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	exchangerate "github.com/williamcardozo/goexpert-client-server-api/pkg/exchange-rate"
)

func getExchangeRateHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	bid, err := exchangerate.GetExchangeRateBID(ctx)
	if err != nil {
		log.Printf("Erro ao buscar cotação: %v", err)
		http.Error(w, "Erro ao buscar cotação", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": bid})
}

func InitServer() {
	http.HandleFunc("/cotacao", getExchangeRateHandler)

	log.Println("Servidor iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	return
}
