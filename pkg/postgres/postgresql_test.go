package postgres_test

import (
	"github.com/ngoctrng/calendarium/pkg/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Info struct {
	CurrentUser string `db:"current_user"`
}

func TestConnection(t *testing.T) {
	dbName, dbUser, dbPass := "test1", "test1", "123456"
	db := testutil.CreateConnection(t, dbName, dbUser, dbPass)
	testutil.MigrateTestDatabase(t, db)

	var info Info
	err := db.Get(&info, "SELECT current_user")
	assert.NoError(t, err)
	assert.Equal(t, dbUser, info.CurrentUser)
}
