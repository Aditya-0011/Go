package routers

import (
	"database/sql"
	"net/http"
)

func RootRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/todo/", TodoRouter(db))
	return mux
}
