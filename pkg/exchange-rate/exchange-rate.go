package exchangerate

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/williamcardozo/goexpert-client-server-api/pkg/db"
	"github.com/williamcardozo/goexpert-client-server-api/pkg/models"
)

func getExchangeRate(ctx context.Context) (*models.ExchangeRate, error) {
	url := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return &models.ExchangeRate{}, fmt.Errorf("erro ao criar requisição para a API de câmbio: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &models.ExchangeRate{}, fmt.Errorf("erro ao enviar requisição para a API de câmbio: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &models.ExchangeRate{}, fmt.Errorf("erro ao ler resposta da API de câmbio: %w", err)
	}

	var exchange *models.ExchangeRate
	if err := json.Unmarshal(body, &exchange); err != nil {
		return &models.ExchangeRate{}, fmt.Errorf("erro ao fazer parse do JSON da API de câmbio: %w", err)
	}

	return exchange, nil
}

func GetExchangeRateBID() (string, error) {
	database, err := db.NewDatabase()
	if err != nil {
		return "", fmt.Errorf("erro ao inicializar banco de dados: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200 * time.Millisecond)
	defer cancel()

	exchange, err := getExchangeRate(ctx)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return "", fmt.Errorf("tempo limite excedido ao obter taxa de câmbio: %w", err)
		} else {
			return "", fmt.Errorf("erro ao obter taxa de câmbio: %w", err)
		}
	}

	if exchange.USDBRL.Bid == "" {
		return "", fmt.Errorf("taxa de câmbio não encontrada")
	}

	err = database.SaveExchangeRate(ctx, exchange)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return "", fmt.Errorf("tempo limite excedido ao salvar taxa de câmbio: %w", err)
		} else {
			return "", fmt.Errorf("erro ao salvar taxa de câmbio: %w", err)
		}
	}

	return exchange.USDBRL.Bid, nil
}
