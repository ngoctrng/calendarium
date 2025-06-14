package testutil

import (
	"context"
	"github.com/ngoctrng/calendarium/pkg/migration"
	"github.com/ngoctrng/calendarium/pkg/postgres"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func MigrateTestDatabase(t testing.TB, db *sqlx.DB) {
	t.Helper()

	_, err := migration.Run(db.DB)
	assert.NoError(t, err)
}

func CreateConnection(t testing.TB, dbName string, dbUser string, dbPass string) *sqlx.DB {
	cont := SetupPostgresContainer(t, dbName, dbUser, dbPass)
	host, _ := cont.Host(context.Background())
	port, _ := cont.MappedPort(context.Background(), "5432")

	db, err := postgres.NewConnection(postgres.Options{
		DBName:   dbName,
		DBUser:   dbUser,
		Password: dbPass,
		Host:     host,
		Port:     port.Port(),
	})
	assert.NoError(t, err)

	return db
}
