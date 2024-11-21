package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:qwerty@localhost:5432/myTube?sslmode=disable")
	if err!= nil {
          return nil, fmt.Errorf("failed to connect to the database: %w", err)
     }
	return conn , nil
}
