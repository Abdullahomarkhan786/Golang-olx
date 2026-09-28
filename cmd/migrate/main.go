package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Abdullahomarkhan786/olx-api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//• package main: Tells the Go compiler that this file or package is an executable program rather than a shared library.

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		log.Fatal("usage:migrate <up|down>")
	}
	cfg := config.MustLoad()
	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseUrl,
	)
	if err != nil {
		log.Fatalf("migration.new %v", err)
	}
	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "down":
		if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command %s", os.Args[1])
	}
	fmt.Println("Running migration")
}
