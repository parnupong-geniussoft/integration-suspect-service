package databases

import (
	"integration-suspect-service/configs"
	"integration-suspect-service/pkg/utils"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgreSqlDbConnection(cfg *configs.Configs) (*sqlx.DB, error) {
	postgresUrl, err := utils.ConnectionUrlBuilder("postgresql", cfg)
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("pgx", postgresUrl)
	if err != nil {
		log.Printf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	log.Println("postgreSQL database has been connected 🐘")
	return db, nil
}
