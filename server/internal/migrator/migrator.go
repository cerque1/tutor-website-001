package migrator

import (
	"log"
	"os"
	"time"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_attempts = 5
	_timeout  = time.Second
)

func Apply(path string, dsn string) {
	migrateLogger := log.New(os.Stdout, "[migrator] ", log.LstdFlags|log.Lshortfile)

	var m *migrate.Migrate
	var err error
	attempts := 0

	for attempts < _attempts {
		m, err = migrate.New(path, dsn)
		if err != nil {
			break
		}
		migrateLogger.Printf("trying to connect, attempt %d", attempts+1)
		time.Sleep(_timeout)
		attempts++
	}

	if err != nil {
		migrateLogger.Fatalf("postgres connect error: %s, with conn str: %s", err, dsn)
	}
	defer m.Close()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		migrateLogger.Fatalf("error apply migration: %v", err)
	}
	migrateLogger.Println("all migrations applied")
}
