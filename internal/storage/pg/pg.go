package pg

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func New() *sqlx.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, database)
	db, err := sqlx.Open("pgx", conn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	return db
}
