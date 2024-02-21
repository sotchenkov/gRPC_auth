package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func main() {
	var storagePath, migratpionsPath, migrationsTable string

	flag.StringVar(&storagePath, "storage-path", "", "")
	flag.StringVar(&migratpionsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-path", "migrations", "name of the migrations")
	flag.Parse()

	if storagePath == "" {
		panic("storage path is required")
	}
	if migratpionsPath == "" {
		panic("migratpions path is required")
	}

	m, err := migrate.New(
		"file://" + migratpionsPath,
		fmt.Sprint("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
