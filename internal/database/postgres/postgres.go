package postgres

import (
	"database/sql"
	"fmt"
	"testWallet/internal/config"
)


type PostgresRepository struct{
	db *sql.DB
}

func NewPGConnection()(*PostgresRepository, error){
	const op = "internal.database.postgres"

	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable",config)
	db, err := sql.Open("postgres", )
}