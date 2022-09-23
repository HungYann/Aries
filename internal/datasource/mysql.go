package datasource

import (
	"database/sql"
	"fmt"
	"server/internal/config"
)

var (
	db *sql.DB
)

func setupDB() error {
	datasource := config.GetDataSource()
	var err error
	db, err = sql.Open("mysql", datasource.MySQL)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	return nil
}
