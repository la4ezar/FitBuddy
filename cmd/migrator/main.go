package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/FitBuddy/pkg/persistence"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current working directory:", cwd)

	databaseCfg := persistence.DatabaseConfig{
		User:               "postgres",
		Password:           "pgsql@12345",
		Host:               "127.0.0.1",
		Port:               "5432",
		Name:               "fitbuddy",
		SSLMode:            "disable",
		MaxOpenConnections: 10,
		MaxIdleConnections: 10,
		ConnMaxLifetime:    30 * time.Second,
	}

	db, closeFunc, err := persistence.Configure(ctx, databaseCfg)
	exitOnError(err, "Error while establishing the connection to the database")
	defer func() {
		err := closeFunc()
		exitOnError(err, "Error while closing the connection to the database")
	}()

	err = runMigrations(db, "file://sql")
	exitOnError(err, "Error while running migrations")

	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan

		err := runDownMigrations(db, "file://sql")
		exitOnError(err, "Error while running down migrations")

		cancel()
	}()

	<-ctx.Done()
}

func runMigrations(db *sql.DB, migrationPath string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	fmt.Println("Migrations applied successfully")
	return nil
}

func runDownMigrations(db *sql.DB, migrationPath string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	fmt.Println("Migrations rolled back successfully")
	return nil
}

func exitOnError(err error, context string) {
	if err != nil {
		wrappedError := errors.Wrap(err, context)
		fmt.Println(wrappedError)
		os.Exit(1)
	}
}
