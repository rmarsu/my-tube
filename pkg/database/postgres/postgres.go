package database

import (
	"context"
	"os"

	"github.com/VandiKond/vanerrors"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, vanerrors.NewWrap("failed to connect to the database", err, vanerrors.EmptyHandler)
	}
	return dbpool, nil
}
