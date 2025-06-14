package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/ngoctrng/calendarium/internal/book/rest"
	"github.com/ngoctrng/calendarium/internal/book/store"
	"github.com/ngoctrng/calendarium/pkg/config"
	"github.com/ngoctrng/calendarium/pkg/postgres"

	sentrygo "github.com/getsentry/sentry-go"
	_ "github.com/lib/pq"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = sentrygo.Init(sentrygo.ClientOptions{
		Dsn:              cfg.SentryDSN,
		Environment:      cfg.AppEnv,
		AttachStacktrace: true,
	})
	if err != nil {
		log.Fatalf("cannot init sentry: %v", err)
	}
	defer sentrygo.Flush(5 * time.Second)

	db, err := postgres.NewConnection(postgres.ParseFromConfig(cfg))
	if err != nil {
		log.Fatal(err)
	}

	server, err := rest.New(rest.WithConfig(cfg))
	if err != nil {
		log.Fatal(err)
	}

	server.BookStore = store.NewBookStore(db)

	addr := fmt.Sprintf(":%d", cfg.Port)
	slog.Info("server started!", "port", cfg.Port)
	log.Fatal(http.ListenAndServe(addr, server))
}
