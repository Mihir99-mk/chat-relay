package config

import (
	"fmt"
	"log"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

func Migrate() error {
	env := NewEnv()
	url, err := GetConnectionString(env)
	if err != nil {
		return err
	}

	sqlDB, err := sql.Open("mysql", url)
	if err != nil {
		return err
	}
	gooseErr := goose.SetDialect("mysql")
	if gooseErr != nil {
		return gooseErr
	}

	if err := goose.Up(sqlDB.DB(), "./migration"); err != nil {
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	log.Println("migration up successfully!!")

	return nil
}
