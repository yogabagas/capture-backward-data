package postgre

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"my-github/capture-backward-data/domain/model"
	"my-github/capture-backward-data/domain/repository"
)

type sqlReadWriter struct {
	db *sql.DB
}

func NewPostgreDriver(host, name, user, password string, port int) (repository.PostgreRepository, error) {
	postgreConn := fmt.Sprintf("host=%s port=%d user%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, name)

	db, err := sql.Open("postgres", postgreConn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &sqlReadWriter{sql: db}, nil
}

func (db *sqlReadWriter) ReadDataInterval(ctx context.Context, from, to string) ([]model.AWBDetailPartner, error) {
	return nil, nil
}
