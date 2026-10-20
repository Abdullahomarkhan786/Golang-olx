package main

import (
	"Abdullahomarkhan786/Golang-olx/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		log.Fatalf("usage: make migrate <up|down> ")

	}
	cfg := config.MustLoad()
	m, err := migrate.New(
		"file://migrations",
		cfg.DatabaseUrl,
	)
	if err != nil {
		log.Fatalf("migration.new:%v", err)
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

	fmt.Println("running migration")

}
