package apiserver

import (
	"database/sql"
	"github.com/Bakhram74/rest-api.git/internal/app/store/sql_store"
	"github.com/gorilla/sessions"
	"net/http"
)

func Start(config Config) error {
	db, err := newDb(config.DatabaseUrl)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sql_store.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	s := NewServer(store, sessionStore)
	return http.ListenAndServe(config.BindAddr, s)
}

func newDb(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
