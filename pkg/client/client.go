package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/williamcardozo/goexpert-client-server-api/pkg/models"
)

func FetchExchangeRate() error {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	url := "http://localhost:8080/cotacao"
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return fmt.Errorf("tempo limite excedido ao criar requisição: %w", err)
		} else {
			return fmt.Errorf("erro ao criar requisição: %w", err)
		}
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("erro ao enviar requisição: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("erro ao ler resposta: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		err := string(body)
		return fmt.Errorf("erro ao buscar cotação: %s", err)
	}

	exchangeRateResponse := models.ExchangeRateResponse{}
	if err := json.Unmarshal(body, &exchangeRateResponse); err != nil {
		return fmt.Errorf("erro ao formatar response: %w", err)
	}
	formattedBid := fmt.Sprintf("Dólar: %s", exchangeRateResponse.Bid)

	os.WriteFile("cotacao.txt", []byte(formattedBid), 0644)
	return nil
}
