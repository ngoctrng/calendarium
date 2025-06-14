package migration_test

import (
	"testing"

	"github.com/ngoctrng/calendarium/pkg/migration"
	"github.com/ngoctrng/calendarium/pkg/testutil"

	"github.com/stretchr/testify/assert"
)

func TestMigration(t *testing.T) {
	dbName, dbUser, dbPass := "test1", "test1", "123456"
	db := testutil.CreateConnection(t, dbName, dbUser, dbPass)

	_, err := migration.Run(db.DB)
	assert.NoError(t, err)
}
