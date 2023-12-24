package helpers

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"net/http"
)

const dbKey = "db"

func GetDBFromContext(r *http.Request) (*pgdb.DB, bool) {
	db, ok := r.Context().Value(dbKey).(*pgdb.DB)
	return db, ok
}
