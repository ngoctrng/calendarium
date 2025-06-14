package main

import (
	"fmt"
	"github.com/ngoctrng/calendarium/pkg/config"
	"github.com/ngoctrng/calendarium/pkg/migration"
	"github.com/ngoctrng/calendarium/pkg/postgres"
	"log"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("cannot load config: %v\n", err)
	}

	db, err := postgres.NewConnection(postgres.ParseFromConfig(cfg))
	if err != nil {
		log.Fatal(err)
	}

	total, err := migration.Run(db.DB)
	if err != nil {
		log.Fatalf("cannot execute migration: %v\n", err)
	}

	slog.Info(fmt.Sprintf("applied %d migrations\n", total))
}
