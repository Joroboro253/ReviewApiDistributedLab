package helpers

import (
	"github.com/jmoiron/sqlx"
	"net/http"
)

const dbKey = "db"

func GetDBFromContext(r *http.Request) (*sqlx.DB, bool) {
	db, ok := r.Context().Value(dbKey).(*sqlx.DB)
	return db, ok
}
