package store

type BookQuerySchema struct {
	ISBN string `db:"isbn"`
	Name string `db:"name"`
}
