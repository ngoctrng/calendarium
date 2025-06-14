package rest_test

import (
	"github.com/ngoctrng/calendarium/internal/book"
	"github.com/ngoctrng/calendarium/internal/book/store"
	"github.com/ngoctrng/calendarium/pkg/testutil"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookAPI(t *testing.T) {
	dbName, dbUser, dbPass := "server", "server", "123456"
	db := testutil.CreateConnection(t, dbName, dbUser, dbPass)
	testutil.MigrateTestDatabase(t, db)
	store := store.NewBookStore(db)

	server := createBookServer(t, store)

	t.Run("create book", func(t *testing.T) {
		data := book.NewBook("9781680500745", "Clojure Applied")
		response := httptest.NewRecorder()
		request := newCreateBookRequest(data)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
		assert.Empty(t, response.Body.String())
	})

	t.Run("get book", func(t *testing.T) {
		data := book.NewBook("9781680500745", "Clojure Applied")
		err := store.Save(&data)
		assert.NoError(t, err)

		response := httptest.NewRecorder()
		request := newGetBookRequest(data.ISBN)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
		assertBookResponse(t, response.Body, data)
	})
}
