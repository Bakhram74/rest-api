package sql_store_test

import (
	"os"
	"testing"
)

var databaseUrl string

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "host=localhost port=5432 user=postgres dbname=restapi_test password=1234 sslmode=disable"
	}
	os.Exit(m.Run())
}
