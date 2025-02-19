package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/pedroafonso18/Maturador-GO/internal/config"
)

func GetConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), config.DBURL)
	if err != nil {
		fmt.Printf("Impossível se conectar ao banco")
		return nil, err
	}
	return conn, nil
}
