package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/williamcardozo/goexpert-client-server-api/pkg/models"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase() (*Database, error) {
	conn, err := sql.Open("sqlite3", "./cotacao.db")
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com banco: %w", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS ExchangeRate (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT,
		codein TEXT,
		name TEXT,
		high TEXT,
		low TEXT,
		varBid TEXT,
		pctChange TEXT,
		bid TEXT,
		ask TEXT,
		timestamp TEXT,
		create_date TEXT
	);`
	if _, err := conn.Exec(createTableSQL); err != nil {
		return nil, fmt.Errorf("erro ao criar tabela: %w", err)
	}

	return &Database{Conn: conn}, nil
}

func (db *Database) SaveExchangeRate(ctx context.Context, exchageRate *models.ExchangeRate) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	stmt, err := db.Conn.Prepare(`
		INSERT INTO ExchangeRate (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("erro ao preparar statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, exchageRate.USDBRL.Code, exchageRate.USDBRL.Codein, exchageRate.USDBRL.Name, exchageRate.USDBRL.High, exchageRate.USDBRL.Low, exchageRate.USDBRL.VarBid, exchageRate.USDBRL.PctChange, exchageRate.USDBRL.Bid, exchageRate.USDBRL.Ask, exchageRate.USDBRL.Timestamp, exchageRate.USDBRL.CreateDate)
	if err != nil {
		return fmt.Errorf("erro ao executar insert: %w", err)
	}

	log.Println("cotação salva no banco com sucesso:", exchageRate.USDBRL.Bid)
	return nil
}
