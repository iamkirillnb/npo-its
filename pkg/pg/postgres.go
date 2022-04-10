package pg

import (
	"github.com/iamkirillnb/Packages/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"npo-its/internal"
)

type DB struct {
	*sqlx.DB

	cfg *internal.DbConfig
	log *logger.Logger
}

func NewDB(config *internal.DbConfig, logger *logger.Logger) *DB {
	connStr := config.DbUrlConnection()
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		logger.Fatalf("connection to DB fail: %s", err)
	}

	err = db.Ping()
	if err != nil {
		logger.Fatalf("ping to DB fail: %s", err)
	}
	return &DB{
		DB:  db,
		cfg: config,
		log: logger,
	}
}
