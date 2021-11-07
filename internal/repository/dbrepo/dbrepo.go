package dbrepo

import (
	"database/sql"

	"github.com/stasxgaming/bookings/internal/config"
	"github.com/stasxgaming/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(app *config.AppConfig, db *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: app,
		DB: db,
	}
}
