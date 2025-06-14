package store_test

import (
	"github.com/ngoctrng/calendarium/internal/book"
	"github.com/ngoctrng/calendarium/internal/book/store"
	"github.com/ngoctrng/calendarium/pkg/testutil"
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/stretchr/testify/assert"
)

func TestBookStore(t *testing.T) {
	dbName, dbUser, dbPass := "test2", "test2", "123456"
	db := testutil.CreateConnection(t, dbName, dbUser, dbPass)
	testutil.MigrateTestDatabase(t, db)

	store := store.NewBookStore(db)

	t.Run("Save a book", func(t *testing.T) {
		want := book.NewBook("9781680500745", "Clojure Applied")
		err := store.Save(&want)

		assert.NoError(t, err)
		verifyInsertedBook(t, db, want.ISBN)
	})

	t.Run("Read existed book", func(t *testing.T) {
		want := book.NewBook("9781680507607", "Distributed Services with Go")
		err := store.Save(&want)
		assert.NoError(t, err)

		got, err := store.FindByISBN(want.ISBN)

		assert.NoError(t, err)
		assertFoundBook(t, got, want)
	})
}

func assertFoundBook(t *testing.T, got *book.Book, want book.Book) {
	t.Helper()

	assert.NotNil(t, got)
	assert.Equal(t, *got, want)
}

func verifyInsertedBook(t testing.TB, db *sqlx.DB, isbn string) {
	t.Helper()

	var got store.BookQuerySchema
	err := db.Get(&got, "SELECT isbn,name FROM books WHERE isbn=$1", isbn)
	assert.NoError(t, err)
	assert.NotNil(t, got)
	assert.Equal(t, isbn, got.ISBN)
}
